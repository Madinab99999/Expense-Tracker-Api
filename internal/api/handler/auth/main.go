package auth_handler

import (
	"log/slog"

	auth_serv "github.com/Madinab99999/Expense-Tracker-Api/internal/api/service/auth"
)

type AuthHandler struct {
	logger  *slog.Logger
	service *auth_serv.AuthService
}

func NewAuthHandler(logger *slog.Logger, service *auth_serv.AuthService) *AuthHandler {
	return &AuthHandler{logger: logger, service: service}
}
