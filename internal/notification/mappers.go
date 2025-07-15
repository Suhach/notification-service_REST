package notification

import (
	"testAPI/internal/domain"
	"testAPI/internal/dto"
	"testAPI/internal/entity"
)

func EntityToDomain(e *entity.NotificationEntity) *domain.Notification {
	return &domain.Notification{
		ID:        e.ID,
		UserID:    e.UserID,
		Message:   e.Message,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
		ReadAt:    e.ReadAt,
	}
}

func CreateDtoToEntity(e *dto.CreateNotificationDTO) *entity.NotificationEntity {
	return &entity.NotificationEntity{
		UserID:  e.UserID,
		Message: e.Message,
		Status:  "sent",
	}
}

func EntityToDTO(e *entity.NotificationEntity) dto.NotificationResponseDTO {
	return dto.NotificationResponseDTO{
		NotificationID: e.ID,
		UserID:         e.UserID,
		Message:        e.Message,
		Status:         e.Status,
		CreatedAt:      e.CreatedAt,
		ReadAt:         e.ReadAt,
	}
}
