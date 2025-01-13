package auth_handler

import (
	"errors"
	"net/http"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
	auth_service "github.com/Madinab99999/Expense-Tracker-Api/internal/service/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/request"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/response"
)

func (h *AuthHandler) AccessToken(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log := h.logger.With("method", "AccessToken")

	var req models.AccessTokenRequest
	if err := request.JSON(w, r, &req); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	tokenPair, err := h.service.AccessToken(ctx, req.Data.RefreshToken)
	if err != nil {
		switch {
		case errors.Is(err, auth_service.ErrInvalidToken):
			http.Error(w, "token refresh failed", http.StatusUnauthorized)
		case errors.Is(err, auth_service.ErrUserNotFound):
			http.Error(w, "user not found", http.StatusUnauthorized)
		default:
			log.ErrorContext(ctx, "token refresh failed", "error", err)
			http.Error(w, "token refresh failed", http.StatusInternalServerError)
		}
		return
	}

	resp := &models.AuthResponse{Data: tokenPair}
	if err := response.JSON(
		w,
		http.StatusCreated,
		resp,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}
	log.InfoContext(
		ctx,
		"success generate access token",
	)
}
