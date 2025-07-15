package service

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"testAPI/internal/dto"
	"testAPI/internal/notification"
	"testAPI/internal/repository"
	"testAPI/pkg/kafkaNP"
	"testAPI/pkg/logger"
	"testAPI/pkg/redisEM"
	"time"
)

type NotificaitonService struct {
	repo repository.NotificationRepo
	kfk  *kafkaNP.NotificaitonProducer
	rds  *redisEM.NotificaitonCache
}

func NewNotificationService(r repository.NotificationRepo, kafka *kafkaNP.NotificaitonProducer, rds *redisEM.NotificaitonCache) *NotificaitonService {
	return &NotificaitonService{
		repo: r,
		kfk:  kafka,
		rds:  rds,
	}
}

func (s *NotificaitonService) CreateNotification(ctx context.Context, in dto.CreateNotificationDTO) (uuid.UUID, error) {
	uuid, err := s.repo.Create(ctx, notification.CreateDtoToEntity(&in))
	if err != nil {
		logger.Log.Error("fail in service CreateNotification", zap.String("notification_uuid", uuid.String()), zap.Error(err))
		return uuid, err
	}

	go func() {
		err = s.kfk.SendUUID(uuid.String())
		if err != nil {
			logger.Log.Error("fail in service CreateNotification sent to kafka", zap.String("notification_uuid", uuid.String()), zap.Error(err))
		}
	}()

	return uuid, nil
}

func (s *NotificaitonService) ListNotifications(ctx context.Context) ([]dto.NotificationResponseDTO, error) {
	entities, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.NotificationResponseDTO, 0, len(entities))
	for _, entity := range entities {
		res = append(res, notification.EntityToDTO(&entity))
	}
	return res, nil
}

func (s *NotificaitonService) GetNotification(ctx context.Context, uid uuid.UUID) (dto.NotificationResponseDTO, error) {
	start := time.Now()
	notifRDS, err := s.rds.Get(ctx, uid.String())
	if err == nil && notifRDS != nil {
		logger.Log.Info("Cache hit")
		finish := time.Since(start)
		logger.Log.Info("Redis response time", zap.Duration("duration", finish))
		return notification.EntityToDTO(notifRDS), nil
	}
	logger.Log.Warn("Cache miss")
	start = time.Now()
	notificationDB, err := s.repo.GetByUUID(ctx, uid)
	if err != nil {
		return dto.NotificationResponseDTO{}, err
	}
	finish := time.Since(start)
	logger.Log.Info("DB response time", zap.Any("duration", finish))
	res := notification.EntityToDTO(&notificationDB)

	go func() {
		_ = s.rds.Set(context.Background(), uid.String(), &notificationDB, 10*time.Minute)
	}()
	return res, nil
}

func (s *NotificaitonService) UpdateStatus(ctx context.Context, uid uuid.UUID) (map[string]string, error) {
	res, err := s.repo.StatusUpdate(ctx, uid)
	if err != nil {
		return nil, err
	}
	_ = s.rds.Delete(ctx, uid.String())
	logger.Log.Info("Cache cleaned")

	return res, nil
}
