package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testAPI/internal/database/postgres"
	"testAPI/internal/handlers"
	openapi "testAPI/internal/ogenerated"
	"testAPI/internal/repository"
	"testAPI/internal/service"
	"testAPI/pkg/kafkaNP"
	"testAPI/pkg/logger"
	"testAPI/pkg/prometheus"
	"testAPI/pkg/redisEM"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logger.Log.Error("recover from panic", zap.Any("err", err))
		}
	}()
	r := gin.New()
	if err := logger.Init(); err != nil {
		log.Fatalf("Logger init error: %v", err)
	}
	defer logger.Log.Sync()
	logger.Log.Info("Logger init success", zap.String("app", "main"))

	prometheus.InitMetrics() //init metrics
	r.Use(prometheus.PrometheusMiddleware())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	//KAFKA
	producer, err := kafkaNP.NewNotificationProducer([]string{"kafka:29092"}, "notification_topic")
	if err != nil {
		logger.Log.Fatal("Kafka init failed:", zap.Error(err))
	}
	//REDIS
	redisCache := redisEM.NewNotificaitonCache("localhost:6379")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	err = postgres.Init() //DB init
	if err != nil {
		logger.Log.Fatal("postgres.Init error", zap.Error(err))
	}
	repo := repository.NewNotificationRepository(postgres.Pool)
	service := service.NewNotificationService(repo, producer, redisCache)
	handler := handlers.NewNotificationHandler(service)

	openapi.RegisterHandlers(r, handler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	if err := postgres.Init(); err != nil {
		logger.Log.Fatal("DB init error", zap.Error(err))
	}
	go func() {
		logger.Log.Info("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Error("Server forced to shutdown", zap.Error(err))
	} else {
		logger.Log.Info("Server exited gracefully")
	}

}
