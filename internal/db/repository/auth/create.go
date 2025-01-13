package auth_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (r *AuthRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	log := r.logger.With("method", "CreateUser")
	query := `
        INSERT INTO users_ (
            email,
            password_hash,
            salt
        ) VALUES ($1, $2, $3)
        RETURNING id`

	row := r.db.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.Salt,
	)
	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to create new user", "error", err, "email", user.Email)
		return nil, err
	}

	err := row.Scan(&user.ID)
	if err != nil {
		if isDuplicateKeyError(err) {
			log.ErrorContext(ctx, "fail to scan user", "error", err, "email", user.Email)
			return nil, fmt.Errorf("create user: %w: %v", ErrDuplicate, err)
		}
		log.ErrorContext(ctx, "fail to scan user", "error", err, "email", user.Email)
		return nil, fmt.Errorf("create user: %w: %v", ErrDatabase, err)
	}

	log.InfoContext(ctx, "success create user", "user_id", user.ID)

	return user, nil
}

func isDuplicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}
