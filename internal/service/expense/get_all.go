package expense_service

import (
	"context"
	"errors"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

var ErrInvalidOptions = errors.New("invalid options")

func (s *ExpenseService) GetAllExpenses(ctx context.Context, userID int64, params models.ExpenseParametrs) ([]models.Expense, error) {
	log := s.logger.With("method", "GetAllExpenses")

	if err := s.validateSortOptions(params.Sort); err != nil {
		log.ErrorContext(ctx, "invalid sort options", "error", err, "user_id", userID)
		return nil, ErrInvalidOptions
	}

	if err := s.validateFilter(params.Filter); err != nil {
		log.ErrorContext(ctx, "invalid filter options", "error", err, "user_id", userID)
		return nil, ErrInvalidOptions
	}

	if err := s.validateDateFormat(params.Filter.StartDate); err != nil {
		log.ErrorContext(ctx, "invalid start date", "error", err, "user_id", userID)
		return nil, ErrInvalidOptions
	}

	if err := s.validateDateFormat(params.Filter.EndDate); err != nil {
		log.ErrorContext(ctx, "invalid end date", "error", err, "user_id", userID)
		return nil, ErrInvalidOptions
	}

	if params.Filter != nil {
		s.processTimeFilter(params.Filter)
	}

	if params.Pagination.Limit == 0 {
		params.Pagination.Limit = 10
	}

	expenses, err := s.repo.GetAllExpenses(ctx, userID, params)
	if err != nil {
		log.ErrorContext(ctx, "failed to get all expenses", "error", err, "user_id", userID)
		return nil, err
	}

	log.InfoContext(ctx, "success get all expenses", "user_id", userID)

	return expenses, nil
}
