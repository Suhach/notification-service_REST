package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"testAPI/internal/entity"
	"testAPI/pkg/logger"
)

type NotificationRepo interface {
	Create(context.Context, *entity.NotificationEntity) (uuid.UUID, error)
	GetAll(context.Context) ([]entity.NotificationEntity, error)
	GetByUUID(context.Context, uuid.UUID) (entity.NotificationEntity, error)
	StatusUpdate(context.Context, uuid.UUID) (map[string]string, error)
}

type NotificationRepository struct {
	db *pgxpool.Pool
}

func NewNotificationRepository(db *pgxpool.Pool) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) Create(ctx context.Context, notif *entity.NotificationEntity) (uuid.UUID, error) {
	query := `INSERT INTO notifications (user_id, message, status)
              VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, notif.UserID, notif.Message, notif.Status)
	if err != nil {
		logger.Log.Error("Fail to insert data into notifications table", zap.Error(err))
		return uuid.Nil, err
	}

	var notifID uuid.UUID
	nid := `SELECT id 
			FROM notifications 
			WHERE user_id = $1`
	err = r.db.QueryRow(ctx, nid, notif.UserID).Scan(&notifID)

	return notifID, nil
}

func (r *NotificationRepository) GetAll(ctx context.Context) ([]entity.NotificationEntity, error) {
	query := `SELECT id, user_id, message, status, created_at, read_at
				FROM notifications
				ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]entity.NotificationEntity, 0)
	for rows.Next() {
		var n entity.NotificationEntity
		if err := rows.Scan(&n.ID, &n.UserID, &n.Message, &n.Status, &n.CreatedAt, &n.ReadAt); err != nil {
			return nil, err
		}
		res = append(res, n)
	}
	return res, nil
}

func (r *NotificationRepository) GetByUUID(ctx context.Context, uid uuid.UUID) (entity.NotificationEntity, error) {
	query := `SELECT id, user_id, message, status, created_at, read_at
			  FROM notifications
			  WHERE id = $1`
	var res entity.NotificationEntity
	err := r.db.QueryRow(ctx, query, uid).Scan(
		&res.ID,
		&res.UserID,
		&res.Message,
		&res.Status,
		&res.CreatedAt,
		&res.ReadAt,
	)
	logger.Log.Info("Get notification by uuid", zap.String("uid", uid.String()))
	if err != nil {
		logger.Log.Error("failed to get notification by uuid", zap.String("uuid", uid.String()), zap.Error(err))
	}
	return res, nil
}

func (r *NotificationRepository) StatusUpdate(ctx context.Context, uid uuid.UUID) (map[string]string, error) {
	query := `UPDATE notifications
			  SET status = 'seen',
				  read_at = CURRENT_TIMESTAMP	
			  WHERE id = $1`
	_, err := r.db.Exec(ctx, query, uid)
	if err != nil {
		logger.Log.Error("fail to queryStatusUpdate", zap.Error(err))
		return nil, err
	}

	return map[string]string{"status": "seen"}, err
}
