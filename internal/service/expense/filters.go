package expense_service

import (
	"time"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *ExpenseService) processTimeFilter(filter *models.ExpenseFilter) {
	now := time.Now()

	if filter.TimeRange == nil {
		return
	}

	switch *filter.TimeRange {
	case models.PastWeek:
		startDate := now.AddDate(0, 0, -7)
		filter.StartDate = stringPointer(startDate.Format("2006-01-02"))
		filter.EndDate = stringPointer(now.Format("2006-01-02"))

	case models.PastMonth:
		startDate := now.AddDate(0, -1, 0)
		filter.StartDate = stringPointer(startDate.Format("2006-01-02"))
		filter.EndDate = stringPointer(now.Format("2006-01-02"))

	case models.Past3Months:
		startDate := now.AddDate(0, -3, 0)
		filter.StartDate = stringPointer(startDate.Format("2006-01-02"))
		filter.EndDate = stringPointer(now.Format("2006-01-02"))

	case models.PastYear:
		startDate := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		filter.StartDate = stringPointer(startDate.Format("2006-01-02"))
		filter.EndDate = stringPointer(now.Format("2006-01-02"))

	case models.NowDate:
		filter.StartDate = stringPointer(now.Format("2006-01-02"))
		filter.EndDate = stringPointer(now.Format("2006-01-02"))

	case models.CustomDate:

		if filter.StartDate == nil {
			filter.StartDate = stringPointer(now.Format("2006-01-02"))
		}

		if filter.EndDate == nil {
			filter.EndDate = stringPointer(now.Format("2006-01-02"))
		}
	}
}

func stringPointer(s string) *string {
	return &s
}
