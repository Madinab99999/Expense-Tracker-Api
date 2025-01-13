package auth_handler

import (
	"errors"
	"net/http"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
	auth_service "github.com/Madinab99999/Expense-Tracker-Api/internal/service/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/request"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/response"
)

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "Login")

	var req models.AuthRequest
	if err := request.JSON(w, r, &req); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	tokenPair, err := h.service.Login(ctx, req.Data.Email, req.Data.Password)
	if err != nil {
		switch {
		case errors.Is(err, auth_service.ErrValidation):
			log.ErrorContext(ctx, "validation failed",
				"error", err,
				"email", req.Data.Email,
			)
			http.Error(w, "invalid email or password format", http.StatusBadRequest)
		case errors.Is(err, auth_service.ErrInvalidCredentials):
			log.ErrorContext(ctx, "invalid credentials",
				"email", req.Data.Email,
			)
			http.Error(w, "login failed", http.StatusUnauthorized)
		default:
			log.ErrorContext(ctx, "login failed", "error", err, "email", req.Data.Email)
			http.Error(w, "login failed", http.StatusInternalServerError)
		}
		return
	}

	resp := &models.AuthResponse{Data: tokenPair}
	if err := response.JSON(
		w,
		http.StatusOK,
		resp,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
			"email", req.Data.Email,
		)
		return
	}

	log.InfoContext(
		ctx,
		"success login user",
		"email", req.Data.Email,
	)
}
