package expense_repository

import (
	"context"
	"fmt"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (r *ExpenseRepository) GetExpenseStats(ctx context.Context, userID int64, startDate, endDate string) (models.ExpenseStats, error) {
	log := r.logger.With("method", "GetExpenseStats")
	stats := models.ExpenseStats{
		Categories: make(map[models.Category]models.CategoryStats),
	}
	query := `
	SELECT 
    category,
	COUNT(*) AS count,
    COALESCE(SUM(amount), 0) AS total_amount,
    COALESCE(MAX(amount), 0) AS highest_amount,
    COALESCE(MIN(amount), 0) AS lowest_amount,
    COALESCE(AVG(amount), 0) AS average_amount
	FROM expense
	WHERE user_id = $1 AND date_expense BETWEEN $2 AND $3
	GROUP BY category
	ORDER BY total_amount DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID, startDate, endDate)
	if err != nil {
		log.ErrorContext(ctx, "failed to query expense stats", "error", err, "user_id", userID)
		return stats, fmt.Errorf("failed to query expense stats: %w", err)
	}
	defer rows.Close()

	var totalAmount, totalCount int
	isEmpty := true
	for rows.Next() {
		var category models.Category
		var stat models.CategoryStats

		err = rows.Scan(
			&category,
			&stat.TotalCount,
			&stat.TotalAmount,
			&stat.HighestAmount,
			&stat.LowestAmount,
			&stat.AverageAmount,
		)
		if err != nil {
			log.ErrorContext(ctx, "failed to scan expense stats", "error", err, "user_id", userID)
			return stats, fmt.Errorf("failed to scan expense stats: %w", err)
		}
		isEmpty = false
		totalAmount += stat.TotalAmount
		totalCount += stat.TotalCount
		stats.Categories[category] = stat
		if stat.HighestAmount > stats.HighestAmount {
			stats.HighestAmount = stat.HighestAmount
		}
		if stats.LowestAmount == 0 || stat.LowestAmount < stats.LowestAmount {
			stats.LowestAmount = stat.LowestAmount
		}
	}

	if isEmpty {
		log.InfoContext(ctx, "expenses not found", "user_id", userID)
		stats.LowestAmount = 0
		stats.HighestAmount = 0
		stats.AverageAmount = 0.0
		stats.TotalAmount = 0
		stats.TotalCount = 0
		return stats, nil
	}

	stats.TotalAmount = totalAmount
	stats.TotalCount = totalCount

	if totalCount > 0 {
		stats.AverageAmount = float64(totalAmount / totalCount)
	}
	log.InfoContext(ctx, "success get expense statistics", "user_id", userID)
	return stats, nil
}
