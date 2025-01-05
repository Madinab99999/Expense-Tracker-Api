package seeds

import (
	"database/sql"
	"fmt"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *seeder) expenses(tx *sql.Tx) error {
	rows, err := tx.Query("SELECT id FROM users_ WHERE email IN ($1, $2, $3)",
		"test@example.com",
		"john.doe@example.com",
		"jane.smith@example.com",
	)
	if err != nil {
		return fmt.Errorf("failed to query user IDs: %w", err)
	}
	defer rows.Close()

	userIDs := make([]int64, 0)
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return fmt.Errorf("failed to scan user ID: %w", err)
		}
		userIDs = append(userIDs, id)
	}

	if len(userIDs) == 0 {
		return fmt.Errorf("no users found for seeding expenses")
	}

	expenses := []struct {
		userID       int64
		amount       int
		date_expense string
		category     models.Category
		description  string
	}{
		{
			userID:       userIDs[0],
			amount:       5000,
			date_expense: "2024-12-21",
			category:     models.Groceries,
			description:  "Weekly grocery shopping at Walmart",
		},
		{
			userID:       userIDs[0],
			amount:       12000,
			date_expense: "2024-12-05",
			category:     models.Electronics,
			description:  "New wireless headphones",
		},
		{
			userID:       userIDs[0],
			amount:       3500,
			date_expense: "2024-12-12",
			category:     models.Transport,
			description:  "Monthly bus pass",
		},
		{
			userID:       userIDs[1],
			amount:       150000,
			date_expense: "2024-12-09",
			category:     models.Education,
			description:  "University semester fee",
		},
		{
			userID:       userIDs[1],
			amount:       8000,
			date_expense: "2024-12-01",
			category:     models.Health,
			description:  "Monthly gym membership",
		},
		// Jane Smith's expenses
		{
			userID:       userIDs[2],
			amount:       25000,
			date_expense: "2024-12-08",
			category:     models.Utilities,
			description:  "Electricity and water bill",
		},
		{
			userID:       userIDs[2],
			amount:       15000,
			date_expense: "2024-12-03",
			category:     models.Leisure,
			description:  "Concert tickets",
		},
	}

	sqlQuery := `
		INSERT INTO expense (
			user_id,
			amount,
			date_expense,
			category,
			description
		)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT DO NOTHING
		RETURNING id;
	`
	sqlStmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return err
	}

	s.cleanups = append(
		s.cleanups, func() error {
			return sqlStmt.Close()
		},
	)

	for _, expense := range expenses {
		var id int64
		err = sqlStmt.QueryRow(
			expense.userID,
			expense.amount,
			expense.date_expense,
			expense.category,
			expense.description,
		).Scan(&id)
		if err != nil {
			return fmt.Errorf("failed to insert expense for user %d: %w", expense.userID, err)
		}
	}

	return nil
}
