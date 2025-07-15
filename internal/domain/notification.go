package domain

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Notification struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Message   string
	Status    string
	CreatedAt time.Time
	ReadAt    pgtype.Timestamp
}
