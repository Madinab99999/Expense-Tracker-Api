package auth_repository

import (
	"database/sql"
	"errors"
	"log/slog"
)

type AuthRepository struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewAuthRepository(logger *slog.Logger, db *sql.DB) *AuthRepository {
	return &AuthRepository{logger: logger, db: db}
}

var (
	ErrNotFound  = errors.New("record not found")
	ErrDuplicate = errors.New("duplicate record")
	ErrDatabase  = errors.New("database error")
)
