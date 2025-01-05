package expense_repository

import (
	"context"
	"fmt"
)

func (r *ExpenseRepository) DeleteExpense(ctx context.Context, userID int64, expenseID int64) error {
	log := r.logger.With("method", "DeleteExpense")

	query := `DELETE FROM expense WHERE id = $1 and user_id = $2`

	deleteExpense, err := r.db.ExecContext(ctx, query, expenseID, userID)
	if err != nil {
		log.ErrorContext(ctx, "failed to delete expense", "error", err, "user_id", userID)
		return err
	}

	countExpense, err := deleteExpense.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "failed to get rows affected", "error", err, "user_id", userID)
		return err
	}

	if countExpense == 0 {
		log.ErrorContext(ctx, "expense not found", "user_id", userID)
		return fmt.Errorf("expense with id %d not found", expenseID)
	}

	log.InfoContext(ctx, "success delete expense", "user_id", userID)
	return nil
}
