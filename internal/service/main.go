package service

import (
	"log/slog"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/configs"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/db/repository"
	auth_service "github.com/Madinab99999/Expense-Tracker-Api/internal/service/auth"
	category_service "github.com/Madinab99999/Expense-Tracker-Api/internal/service/category"
	expense_service "github.com/Madinab99999/Expense-Tracker-Api/internal/service/expense"
)

type Service struct {
	logger          *slog.Logger
	config          *configs.Config
	AuthService     *auth_service.AuthService
	CategoryService *category_service.CategoryService
	ExpenseService  *expense_service.ExpenseService
}

func New(logger *slog.Logger, config *configs.Config, repo *repository.Repository) *Service {
	return &Service{
		logger:          logger,
		config:          config,
		AuthService:     auth_service.NewAuthService(logger, config, repo.AuthRepository),
		CategoryService: category_service.NewCategoryService(logger, repo.CategoryRepository),
		ExpenseService:  expense_service.NewExpenseService(logger, repo.ExpenseRepository),
	}
}
