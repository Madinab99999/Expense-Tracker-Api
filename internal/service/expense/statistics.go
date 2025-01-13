package expense_service

import (
	"context"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *ExpenseService) GetExpenseStats(ctx context.Context, userID int64, startDate, endDate string) (*models.ExpenseStats, error) {
	log := s.logger.With("method", "GetExpenseStats")

	if err := s.validateDateFormat(&startDate); err != nil {
		log.ErrorContext(ctx, "invalid start date", "error", err, "user_id", userID)
		return nil, ErrInvalidOptions
	}

	if err := s.validateDateFormat(&endDate); err != nil {
		log.ErrorContext(ctx, "invalid end date", "error", err, "user_id", userID)
		return nil, ErrInvalidOptions
	}

	expenseStats, err := s.repo.GetExpenseStats(ctx, userID, startDate, endDate)
	if err != nil {
		log.ErrorContext(ctx, "failed to get expense statistics", "error", err, "user_id", userID)
		return nil, err
	}

	log.InfoContext(ctx, "success get expense statistics", "user_id", userID)

	return &expenseStats, nil
}
