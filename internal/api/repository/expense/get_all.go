package expense_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (r *ExpenseRepository) GetAllExpenses(ctx context.Context, userID int64, params models.ExpenseParametrs) ([]models.Expense, error) {
	log := r.logger.With("method", "GetAllExpenses")

	queryBuilder := strings.Builder{}
	queryBuilder.WriteString(`
	SELECT id,user_id, amount, date_expense, category, description, created_at, updated_at 
	FROM expense WHERE user_id = $1
	`)

	args := []interface{}{userID}
	paramCount := 1

	if params.Filter != nil {
		if params.Filter.Category != nil {
			queryBuilder.WriteString(fmt.Sprintf(" AND category = $%d", paramCount+1))
			args = append(args, *params.Filter.Category)
			paramCount++
		}
		if params.Filter.TimeRange != nil {
			queryBuilder.WriteString(fmt.Sprintf(" AND date_expense BETWEEN $%d AND $%d", paramCount+1, paramCount+2))
			args = append(args, *params.Filter.StartDate, *params.Filter.EndDate)
			paramCount += 2
		}
		if params.Filter.MinAmount != nil {
			queryBuilder.WriteString(fmt.Sprintf(" AND amount >= $%d", paramCount+1))
			args = append(args, *params.Filter.MinAmount)
			paramCount++
		}
		if params.Filter.MaxAmount != nil {
			queryBuilder.WriteString(fmt.Sprintf(" AND amount <= $%d", paramCount+1))
			args = append(args, *params.Filter.MaxAmount)
			paramCount++
		}
	}

	if params.Sort != nil && params.Sort.SortBy != "" && params.Sort.Order != "" {
		queryBuilder.WriteString(fmt.Sprintf(" ORDER BY %s %s", params.Sort.SortBy, params.Sort.Order))
	}

	if params.Pagination.Cursor != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND id > $%d", paramCount+1))
		args = append(args, *params.Pagination.Cursor)
		paramCount++
	}
	queryBuilder.WriteString(fmt.Sprintf(" LIMIT $%d", paramCount+1))
	args = append(args, params.Pagination.Limit)

	query := queryBuilder.String()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.ErrorContext(ctx, "failed to query expenses", "error", err, "user_id", userID)
		return nil, fmt.Errorf("failed to query expenses: %w", err)
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
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
			log.ErrorContext(ctx, "failed to scan expenses", "error", err, "user_id", userID)
			return nil, fmt.Errorf("failed to scan expenses: %w", err)
		}
		expenses = append(expenses, expense)
	}
	if err = rows.Err(); err != nil {
		log.ErrorContext(ctx, "failed to scan expense rows", "error", err, "user_id", userID)
		return nil, fmt.Errorf("failed to scan expense rows: %w", err)
	}
	log.InfoContext(ctx, "success get all expenses", "user_id", userID)
	return expenses, nil
}
