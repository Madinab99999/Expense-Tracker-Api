package auth_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (r *AuthRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	log := r.logger.With("method", "GetUserByID")
	query := `
        SELECT 
            id,
            email,
            password_hash,
            salt,
            created_at,
            updated_at
        FROM users_
        WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Salt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.WarnContext(ctx, "no user found", "error", err, "user_id", user.ID)
			return nil, fmt.Errorf("get user by id: %w", ErrNotFound)
		}
		log.ErrorContext(ctx, "fail to scan user", "error", err, "user_id", user.ID)
		return nil, fmt.Errorf("get user by id: %w: %v", ErrDatabase, err)
	}
	log.InfoContext(ctx, "success get user by id", "user_id", user.ID)
	return user, nil
}
