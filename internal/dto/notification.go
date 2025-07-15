package dto

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateNotificationDTO struct {
	NotificationID uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	Message        string    `json:"message"`
}

type NotificationResponseDTO struct {
	NotificationID uuid.UUID        `json:"id"`
	UserID         uuid.UUID        `json:"user_id"`
	CreatedAt      time.Time        `json:"created_at"`
	Message        string           `json:"message"`
	Status         string           `json:"status"`
	ReadAt         pgtype.Timestamp `json:"read_at"`
}
