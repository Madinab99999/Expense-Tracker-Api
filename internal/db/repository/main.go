package repository

import (
	"database/sql"
	"log/slog"

	auth_repo "github.com/Madinab99999/Expense-Tracker-Api/internal/db/repository/auth"
	cat_repo "github.com/Madinab99999/Expense-Tracker-Api/internal/db/repository/category"
	expense_repo "github.com/Madinab99999/Expense-Tracker-Api/internal/db/repository/expense"
)

type Repository struct {
	logger             *slog.Logger
	AuthRepository     *auth_repo.AuthRepository
	CategoryRepository *cat_repo.CategoryRepository
	ExpenseRepository  *expense_repo.ExpenseRepository
}

func New(logger *slog.Logger, db *sql.DB) *Repository {
	return &Repository{
		logger:             logger,
		AuthRepository:     auth_repo.NewAuthRepository(logger, db),
		CategoryRepository: cat_repo.NewCategoryRepository(logger, db),
		ExpenseRepository:  expense_repo.NewExpenseRepository(logger, db),
	}
}
