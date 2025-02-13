package handlers

import (
	"elibrary/internal/repository"
	"elibrary/internal/services"
	"encoding/json"
	"net/http"
)

type NotificationHandler struct {
	notifyService *services.NotificationService
}

func NewNotificationHandler(db *repository.Database) *NotificationHandler {
	return &NotificationHandler{
		notifyService: services.NewNotificationService(db),
	}
}

func (h *NotificationHandler) NotifyAvailable(w http.ResponseWriter, r *http.Request) {
	// Implementation for manual notification trigger
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "notifications processed"})
}
