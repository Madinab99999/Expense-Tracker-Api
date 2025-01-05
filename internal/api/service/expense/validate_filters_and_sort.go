package expense_service

import (
	"fmt"
	"time"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *ExpenseService) validateSortOptions(sort *models.SortOptions) error {
	if sort == nil {
		return nil
	}

	validSortFields := map[models.SortField]bool{
		models.SortByDateExpense: true,
		models.SortByAmount:      true,
		models.SortByCategory:    true,
		models.SortByDateCreate:  true,
	}

	validSortOrders := map[models.SortOrder]bool{
		models.Ascending:  true,
		models.Descending: true,
	}

	if _, ok := validSortFields[sort.SortBy]; !ok && sort.SortBy != "" {
		return fmt.Errorf("invalid sort field: %s", sort.SortBy)
	}

	if _, ok := validSortOrders[sort.Order]; !ok && sort.Order != "" {
		return fmt.Errorf("invalid sort order: %s", sort.Order)
	}

	return nil
}

func (s *ExpenseService) validateFilter(filter *models.ExpenseFilter) error {
	if filter == nil {
		return nil
	}

	if filter.TimeRange != nil {
		validTimeRanges := map[models.TimeRange]bool{
			models.PastWeek:    true,
			models.PastMonth:   true,
			models.Past3Months: true,
			models.PastYear:    true,
			models.NowDate:     true,
			models.CustomDate:  true,
		}

		if !validTimeRanges[*filter.TimeRange] {
			return fmt.Errorf("invalid time range: %s", *filter.TimeRange)
		}
	}

	if filter.Category != nil {
		err := s.ValidateCategory(*filter.Category)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ExpenseService) validateDateFormat(date *string) error {
	if date == nil {
		return nil
	}

	_, err := time.Parse("2006-01-02", *date)
	if err != nil {
		return fmt.Errorf("invalid date format: %s", *date)
	}

	return nil
}
