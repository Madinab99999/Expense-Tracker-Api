package auth_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
	auth_repository "github.com/Madinab99999/Expense-Tracker-Api/internal/db/repository/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.AuthTokenPair, error) {
	log := s.logger.With("method", "Login")

	authUser := &models.AuthUser{
		Email:    email,
		Password: password,
	}

	if err := s.validateAuthUser(authUser); err != nil {
		log.ErrorContext(ctx, "validation failed", "error", err, "email", email)
		return nil, ErrValidation
	}

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, auth_repository.ErrNotFound) {
			log.ErrorContext(ctx, "failed to get user", "error", err, "email", email)
			return nil, ErrInvalidCredentials
		}
		log.ErrorContext(ctx, "failed to get user", "error", err, "email", email)
		return nil, fmt.Errorf("database error: %w", err)
	}

	tokenSecret := s.config.TokenSecret
	isValid, err := auth.VerifyPassword(password, user.PasswordHash, user.Salt)
	if err != nil {
		log.ErrorContext(ctx, "failed to verify password", "error", err)
		return nil, fmt.Errorf("password verification failed: %w", err)
	}
	if !isValid {
		log.ErrorContext(ctx, "failed to verify password", "error", err)
		return nil, ErrInvalidCredentials
	}

	tokenPair, err := auth.GenerateTokenPair(
		&auth.UserData{
			ID:    fmt.Sprint(user.ID),
			Email: user.Email,
		},
		tokenSecret,
	)
	if err != nil {
		log.ErrorContext(ctx, "failed to generate tokens", "error", err, "email", email)
		return nil, fmt.Errorf("token generation failed: %w", err)
	}

	log.InfoContext(ctx, "user logged in successfully", "email", email)
	return tokenPair, nil
}
