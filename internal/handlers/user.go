package handlers

import (
	"elibrary/internal/models"
	"elibrary/internal/repository"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: repo}
}

func (h *UserHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var subReq models.SubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&subReq); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.userRepo.FindUserByEmail(subReq.Email)
	if err != nil {
		// Create new user if not exists
		user = &models.User{Email: subReq.Email}
		if err := h.userRepo.CreateUser(user); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create subscription")
			return
		}
	}

	if err := h.userRepo.UpdateSubscription(user.ID, true); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update subscription")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "subscribed"})
}

func (h *UserHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	// Similar implementation for unsubscribe
}
