package expense_handler

import (
	"log/slog"

	expense_service "github.com/Madinab99999/Expense-Tracker-Api/internal/api/service/expense"
)

type ExpenseHandler struct {
	logger  *slog.Logger
	service *expense_service.ExpenseService
}

func NewExpenseHandler(logger *slog.Logger, service *expense_service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{logger: logger, service: service}
}
