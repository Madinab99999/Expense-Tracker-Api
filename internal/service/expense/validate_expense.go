package expense_service

import (
	"errors"
	"fmt"

	m "github.com/Madinab99999/Expense-Tracker-Api/internal/models"

	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidCategory    = errors.New("invalid category")
	ErrInvalidAmount      = errors.New("invalid amount")
	ErrInvalidDate        = errors.New("invalid date format")
	ErrInvalidDescription = errors.New("invalid description")
)

func (s *ExpenseService) validateExpense(expense *m.Expense) error {
	if err := s.validator.Struct(expense); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				switch e.Field() {
				case "Amount":
					return fmt.Errorf("%w: amount must be greater than 0", ErrInvalidAmount)
				case "Date_expense":
					return fmt.Errorf("%w: must be in format YYYY-MM-DD", ErrInvalidDate)
				case "Category":
					return fmt.Errorf("%w: category is required", ErrInvalidCategory)
				case "Description":
					return fmt.Errorf("%w: must be between 3 and 500 characters", ErrInvalidDescription)
				}
			}
		}
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}

func (s *ExpenseService) ValidateCategory(c m.Category) error {
	switch c {
	case m.Groceries, m.Leisure, m.Electronics, m.Utilities, m.Clothing, m.Health, m.Transport, m.Education, m.Credits, m.Others:
		return nil
	default:
		return ErrInvalidCategory
	}
}
