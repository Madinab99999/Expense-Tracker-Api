package auth_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	log := r.logger.With("method", "GetUserByEmail")
	query := `
        SELECT 
            id,
            email,
            password_hash,
            salt,
            created_at,
            updated_at
        FROM users_
        WHERE email = $1`

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Salt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.WarnContext(ctx, "no user found", "error", err, "email", email)
			return nil, fmt.Errorf("get user by email: %w", ErrNotFound)
		}
		log.ErrorContext(ctx, "fail to scan user", "error", err, "email", email)
		return nil, fmt.Errorf("get user by email: %w: %v", ErrDatabase, err)
	}
	log.InfoContext(ctx, "success get user by email", "email", email)
	return user, nil
}
