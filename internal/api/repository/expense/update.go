package expense_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (r *ExpenseRepository) UpdateExpense(ctx context.Context, userID int64, expenseID int64, expense *models.Expense) (*models.Expense, error) {
	log := r.logger.With("method", "UpdateExpense")

	query := `
        UPDATE expense SET amount = $1, date_expense = $2, category = $3, description = $4, updated_at = NOW()
        WHERE id = $5 and user_id = $6  
        RETURNING id,user_id, amount, date_expense, category, description, created_at, updated_at `

	row := r.db.QueryRowContext(ctx, query,
		expense.Amount,
		expense.Date_expense,
		expense.Category,
		expense.Description,
		expenseID,
		userID,
	)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "failed to update expense", "error", err, "user_id", userID)
		return nil, err
	}

	var updatedExpense models.Expense
	err := row.Scan(
		&updatedExpense.ID,
		&updatedExpense.UserID,
		&updatedExpense.Amount,
		&updatedExpense.Date_expense,
		&updatedExpense.Category,
		&updatedExpense.Description,
		&updatedExpense.CreatedAt,
		&updatedExpense.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.WarnContext(ctx, "expense not found", "error", err, "user_id", userID)
			return nil, err
		}
		log.ErrorContext(ctx, "fail to scan expense", "error", err, "user_id", userID)
		return nil, err
	}
	log.InfoContext(ctx, "success update expense", "user_id", userID)
	return &updatedExpense, nil
}
