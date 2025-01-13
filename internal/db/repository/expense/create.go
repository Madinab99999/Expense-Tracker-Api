package expense_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (r *ExpenseRepository) CreateExpense(ctx context.Context, expense *models.Expense, userID int64) (*models.Expense, error) {
	log := r.logger.With("method", "CreateExpense")

	query := `
        INSERT INTO expense (user_id, amount, date_expense, category, description)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id,user_id, amount, date_expense, category, description, created_at, updated_at `

	row := r.db.QueryRowContext(ctx, query,
		userID,
		expense.Amount,
		expense.Date_expense,
		expense.Category,
		expense.Description,
	)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "failed to create new  expense", "error", err, "user_id", userID)
		return nil, err
	}

	var createdExpense models.Expense
	err := row.Scan(
		&createdExpense.ID,
		&createdExpense.UserID,
		&createdExpense.Amount,
		&createdExpense.Date_expense,
		&createdExpense.Category,
		&createdExpense.Description,
		&createdExpense.CreatedAt,
		&createdExpense.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.WarnContext(ctx, "no values was found", "error", err, "user_id", userID)
			return nil, err
		}
		log.ErrorContext(ctx, "fail to scan expense", "error", err, "user_id", userID)
		return nil, err
	}
	log.InfoContext(ctx, "success create expense", "user_id", userID)
	return &createdExpense, nil
}
