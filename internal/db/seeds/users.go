package seeds

import (
	"database/sql"
	"fmt"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
)

func (s *seeder) users(tx *sql.Tx) error {
	users := []struct {
		email    string
		password string
	}{
		{
			email:    "test@example.com",
			password: "Test123!@#",
		},
		{
			email:    "john.doe@example.com",
			password: "JohnDoe456!@#",
		},
		{
			email:    "jane.smith@example.com",
			password: "JaneSmith789!@#",
		},
	}

	sqlQuery := `
		INSERT INTO users_ (
			email,
			password_hash,
			salt
		)
		VALUES ($1, $2, $3)
		ON CONFLICT (email) DO NOTHING
		RETURNING id;
	`
	sqlStmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return fmt.Errorf("failed to prepare user seed statement: %w", err)
	}

	s.cleanups = append(
		s.cleanups, func() error {
			return sqlStmt.Close()
		},
	)

	for _, user := range users {
		passwordHash, salt, err := auth.HashPassword(user.password)
		if err != nil {
			return fmt.Errorf("failed to hash password for user %s: %w", user.email, err)
		}
		var id int64
		err = sqlStmt.QueryRow(
			user.email,
			passwordHash,
			salt,
		).Scan(&id)
		if err != nil {
			return err
		}
	}

	return nil
}
