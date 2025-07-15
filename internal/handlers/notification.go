package handlers

import (
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"go.uber.org/zap"
	"net/http"
	"testAPI/internal/dto"
	"testAPI/internal/service"
	"testAPI/pkg/logger"
)

type NotificationHandler struct {
	service *service.NotificaitonService
}

func NewNotificationHandler(s *service.NotificaitonService) *NotificationHandler {
	return &NotificationHandler{
		service: s,
	}
}

func (h *NotificationHandler) PostNotifications(c *gin.Context) {
	var input dto.CreateNotificationDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Log.Error("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	uid, err := h.service.CreateNotification(c.Request.Context(), input)
	if err != nil {
		logger.Log.Error("Failed to create notification", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create notification"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"notification_uuid": uid})
	logger.Log.Info("Created notification", zap.Any("notification", input))
}

func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	notifs, err := h.service.ListNotifications(c.Request.Context())
	if err != nil {
		logger.Log.Error("Failed to get notifications", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get notifications"})
		return
	}
	c.JSON(http.StatusOK, notifs)
	logger.Log.Info("Retrieved notifications", zap.Any("notifications", notifs))
}

// TODO: исправить парсинг uuid из query запроса!
func (h *NotificationHandler) GetNotificationUuid(c *gin.Context, uuid openapi_types.UUID) {
	notification, err := h.service.GetNotification(c.Request.Context(), uuid)
	if err != nil {
		logger.Log.Error("Failed to get notification", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get notification"})
		return
	}
	c.JSON(http.StatusOK, notification)
	logger.Log.Info("Retrieved notification", zap.Any("notification", notification))
}

// TODO: исправить парсинг uuid из query запроса!
func (h *NotificationHandler) PatchNotificationUuid(c *gin.Context, uuid openapi_types.UUID) {
	updatedStatus, err := h.service.UpdateStatus(c.Request.Context(), uuid)
	if err != nil {
		logger.Log.Error("Failed to update status", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update status"})
		return
	}
	c.JSON(http.StatusOK, updatedStatus)
	logger.Log.Info("Updated status", zap.Any("status", updatedStatus))
}
