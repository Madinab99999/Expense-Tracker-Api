package expense_service

import (
	"context"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *ExpenseService) GetInformationOfExpense(ctx context.Context, userID int64, expenseID int64) (*models.Expense, error) {
	log := s.logger.With("method", "GetInformationOfExpense")

	expense, err := s.repo.GetInformationOfExpense(ctx, userID, expenseID)
	if err != nil {
		log.ErrorContext(ctx, "failed to get information of expense", "error", err, "user_id", userID)
		return nil, err
	}

	log.InfoContext(ctx, "success get information of expense", "user_id", userID)

	return expense, nil
}
