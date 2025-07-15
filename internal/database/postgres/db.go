package postgres

import (
	"context"
	"testAPI/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var Pool *pgxpool.Pool

func Init() error {
	dbURL := "postgres://notif_user:notif_pass@host.docker.internal:5432/notifications_db?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		logger.Log.Error("fail to connect to database", zap.String("db", "postgres"))
		return err
	}
	Pool = pool
	return nil
}
