package expense_service

import (
	"context"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *ExpenseService) UpdateExpense(ctx context.Context, userID int64, expenseID int64, expense *models.Expense) (*models.Expense, error) {
	log := s.logger.With("method", "UpdateExpense")

	if err := s.validateExpense(expense); err != nil {
		log.ErrorContext(ctx, "validation failed",
			"error", err,
			"user_id", userID)
		return nil, err
	}

	if err := s.ValidateCategory(expense.Category); err != nil {
		log.ErrorContext(ctx, "failed to update expense", "error", err, "user_id", userID)
		return nil, err
	}
	updateExpense, err := s.repo.UpdateExpense(ctx, userID, expenseID, expense)
	if err != nil {
		log.ErrorContext(ctx, "failed to update expense", "error", err, "user_id", userID)
		return nil, err
	}
	log.InfoContext(ctx, "success update expense", "user_id", userID)
	return updateExpense, nil
}
