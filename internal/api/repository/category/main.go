package category_repository

import (
	"database/sql"
	"log/slog"
)

type CategoryRepository struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewCategoryRepository(logger *slog.Logger, db *sql.DB) *CategoryRepository {
	return &CategoryRepository{logger: logger, db: db}
}
