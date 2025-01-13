package expense_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (r *ExpenseRepository) GetInformationOfExpense(ctx context.Context, userID int64, expenseID int64) (*models.Expense, error) {
	log := r.logger.With("method", "GetInformationOfExpense")

	var expense models.Expense

	query := `
	SELECT id,user_id, amount, date_expense, category, description, created_at, updated_at 
	FROM expense WHERE id = $1 and user_id = $2 `
	row := r.db.QueryRowContext(ctx, query, expenseID, userID)
	err := row.Scan(
		&expense.ID,
		&expense.UserID,
		&expense.Amount,
		&expense.Date_expense,
		&expense.Category,
		&expense.Description,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.WarnContext(ctx, "expense not found", "error", err, "user_id", userID)
			return nil, err
		}
		log.ErrorContext(ctx, "fail to scan expense", "error", err, "user_id", userID)
		return nil, err
	}
	log.InfoContext(ctx, "success get information of expense", "user_id", userID)
	return &expense, nil
}
