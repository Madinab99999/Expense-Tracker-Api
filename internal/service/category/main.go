package category_service

import (
	"log/slog"

	cat_repo "github.com/Madinab99999/Expense-Tracker-Api/internal/db/repository/category"
)

type CategoryService struct {
	logger *slog.Logger
	repo   *cat_repo.CategoryRepository
}

func NewCategoryService(logger *slog.Logger, repo *cat_repo.CategoryRepository) *CategoryService {
	return &CategoryService{logger: logger, repo: repo}
}
