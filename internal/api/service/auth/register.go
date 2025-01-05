package auth_service

import (
	"context"
	"fmt"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *AuthService) Register(ctx context.Context, email, password string) (*models.AuthTokenPair, error) {

	log := s.logger.With("method", "Register")

	authUser := &models.AuthUser{
		Email:    email,
		Password: password,
	}

	if err := s.validateAuthUser(authUser); err != nil {
		log.ErrorContext(ctx, "validation failed", "error", err)
		return nil, ErrValidation
	}

	tokenPepper := s.config.TokenPepper
	tokenSecret := s.config.TokenSecret
	passwordHash, salt, err := auth.HashPassword(password, tokenPepper)
	if err != nil {
		log.ErrorContext(ctx, "failed to hash password", "error", err)
		return nil, fmt.Errorf("password hashing failed: %w", err)
	}

	user, err := s.repo.CreateUser(ctx, &models.User{
		Email:        email,
		PasswordHash: passwordHash,
		Salt:         salt,
	})
	if err != nil {
		log.ErrorContext(ctx, "failed to create user", "error", err, "email", email)
		return nil, fmt.Errorf("user creation failed: %w", err)
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

	log.InfoContext(ctx, "user registered successfully", "email", email)
	return tokenPair, nil
}
