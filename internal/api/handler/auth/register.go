package auth_handler

import (
	"errors"
	"net/http"

	auth_service "github.com/Madinab99999/Expense-Tracker-Api/internal/api/service/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/request"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/response"
)

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "Register")

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

	tokenPair, err := h.service.Register(ctx, req.Data.Email, req.Data.Password)
	if err != nil {
		switch {
		case errors.Is(err, auth_service.ErrValidation):
			log.ErrorContext(ctx, "validation failed",
				"error", err,
				"email", req.Data.Email,
			)
			http.Error(w, "invalid email or password format", http.StatusBadRequest)
		case errors.Is(err, auth_service.ErrUserExists):
			log.ErrorContext(ctx, "user already exists",
				"error", err,
				"email", req.Data.Email,
			)
			http.Error(w, "user already exists", http.StatusConflict)
		case errors.Is(err, auth_service.ErrInvalidCredentials):
			log.ErrorContext(ctx, "invalid credentials",
				"email", req.Data.Email,
			)
			http.Error(w, "registration failed", http.StatusBadRequest)
		default:
			log.ErrorContext(ctx, "registration failed", "error", err, "email", req.Data.Email)
			http.Error(w, "registration failed", http.StatusInternalServerError)
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
			"email", req.Data.Email,
		)
		return
	}

	log.InfoContext(
		ctx,
		"success register user",
		"email", req.Data.Email,
	)
}
