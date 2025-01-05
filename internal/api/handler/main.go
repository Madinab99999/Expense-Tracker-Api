package handler

import (
	"log/slog"

	auth_handler "github.com/Madinab99999/Expense-Tracker-Api/internal/api/handler/auth"
	category_handler "github.com/Madinab99999/Expense-Tracker-Api/internal/api/handler/category"
	expense_handler "github.com/Madinab99999/Expense-Tracker-Api/internal/api/handler/expense"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/api/service"
)

type Handler struct {
	logger          *slog.Logger
	AuthHandler     *auth_handler.AuthHandler
	CategoryHandler *category_handler.CategoryHandler
	ExpenseHandler  *expense_handler.ExpenseHandler
}

func New(logger *slog.Logger, service *service.Service) *Handler {
	return &Handler{
		logger:          logger,
		AuthHandler:     auth_handler.NewAuthHandler(logger, service.AuthService),
		CategoryHandler: category_handler.NewCategoryHandler(logger, service.CategoryService),
		ExpenseHandler:  expense_handler.NewExpenseHandler(logger, service.ExpenseService),
	}
}
