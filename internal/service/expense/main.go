package expense_service

import (
	"log/slog"

	expense_repo "github.com/Madinab99999/Expense-Tracker-Api/internal/db/repository/expense"

	"github.com/go-playground/validator/v10"
)

type ExpenseService struct {
	logger    *slog.Logger
	repo      *expense_repo.ExpenseRepository
	validator *validator.Validate
}

func NewExpenseService(logger *slog.Logger, repo *expense_repo.ExpenseRepository) *ExpenseService {
	v := validator.New()
	return &ExpenseService{
		logger:    logger,
		repo:      repo,
		validator: v,
	}
}
