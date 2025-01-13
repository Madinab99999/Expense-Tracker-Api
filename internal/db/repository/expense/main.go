package expense_repository

import (
	"database/sql"
	"log/slog"
)

type ExpenseRepository struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewExpenseRepository(logger *slog.Logger, db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{logger: logger, db: db}
}
