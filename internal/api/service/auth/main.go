package auth_service

import (
	"errors"
	"log/slog"

	auth_repo "github.com/Madinab99999/Expense-Tracker-Api/internal/api/repository/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/configs"

	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	config    *configs.Config
	logger    *slog.Logger
	repo      *auth_repo.AuthRepository
	validator *validator.Validate
}

func NewAuthService(logger *slog.Logger, config *configs.Config, repo *auth_repo.AuthRepository) *AuthService {
	v := validator.New()
	v.RegisterValidation("passwordComplexity", validatePasswordComplexity)
	return &AuthService{
		logger:    logger,
		config:    config,
		repo:      repo,
		validator: v,
	}
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidToken       = errors.New("invalid token")
	ErrValidation         = errors.New("validation error")
)
