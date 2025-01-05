package expense_service

import (
	"context"
)

func (s *ExpenseService) DeleteExpense(ctx context.Context, userID int64, expenseID int64) error {
	log := s.logger.With("method", "DeleteExpense")

	if err := s.repo.DeleteExpense(ctx, userID, expenseID); err != nil {
		log.ErrorContext(ctx, "failed to delete expense", "error", err, "user_id", userID)
		return err
	}

	log.InfoContext(ctx, "success delete expense", "user_id", userID)

	return nil
}
