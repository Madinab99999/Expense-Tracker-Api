package expense_service

import (
	"context"
	"fmt"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *ExpenseService) InsertExpense(ctx context.Context, expense *models.Expense, userID int64) (*models.Expense, error) {
	log := s.logger.With("method", "InsertExpense")

	if err := s.validateExpense(expense); err != nil {
		log.ErrorContext(ctx, "validation failed",
			"error", err,
			"user_id", userID,
		)
		return nil, err
	}

	if err := s.ValidateCategory(expense.Category); err != nil {
		log.ErrorContext(ctx, "failed to insert expense",
			"error", err,
			"user_id", userID)
		return nil, err
	}
	newExpense, err := s.repo.CreateExpense(ctx, expense, userID)
	if err != nil {
		log.ErrorContext(ctx, "failed to create expense",
			"error", err,
			"user_id", userID)
		return nil, fmt.Errorf("user creation failed: %w", err)
	}
	log.InfoContext(
		ctx, "success insert expense",
		"user_id", userID)
	return newExpense, nil
}
