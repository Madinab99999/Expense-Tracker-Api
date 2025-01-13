package auth_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
	auth_repository "github.com/Madinab99999/Expense-Tracker-Api/internal/db/repository/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *AuthService) AccessToken(ctx context.Context, refreshToken string) (*models.AuthTokenPair, error) {
	log := s.logger.With("method", "AccessToken")

	tokenSecret := s.config.TokenSecret
	userData, err := auth.ParseToken(refreshToken, tokenSecret)
	if err != nil {
		log.ErrorContext(ctx, "failed to parse refresh token", "error", err)
		return nil, ErrInvalidToken
	}

	user, err := s.repo.GetUserByID(ctx, userData.ID)
	if err != nil {
		if errors.Is(err, auth_repository.ErrNotFound) {
			log.ErrorContext(ctx, "failed to get user", "error", err)
			return nil, ErrUserNotFound
		}
		log.ErrorContext(ctx, "failed to get user", "error", err)
		return nil, fmt.Errorf("database error: %w", err)
	}

	tokenPair, err := auth.GenerateTokenPair(
		&auth.UserData{
			ID:    fmt.Sprint(user.ID),
			Email: user.Email,
		},
		tokenSecret,
	)
	if err != nil {
		log.ErrorContext(ctx, "failed to generate tokens", "error", err)
		return nil, fmt.Errorf("token generation failed: %w", err)
	}

	log.InfoContext(ctx, "success generate access token", "user_id", user.ID)
	return tokenPair, nil
}
