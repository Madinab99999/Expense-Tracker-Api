package category_handler

import (
	"log/slog"

	cat_service "github.com/Madinab99999/Expense-Tracker-Api/internal/api/service/category"
)

type CategoryHandler struct {
	logger  *slog.Logger
	service *cat_service.CategoryService
}

func NewCategoryHandler(logger *slog.Logger, service *cat_service.CategoryService) *CategoryHandler {
	return &CategoryHandler{logger: logger, service: service}
}
