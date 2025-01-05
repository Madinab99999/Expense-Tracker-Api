package category_service

import (
	"context"

	m "github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (s *CategoryService) GetAllCategories(ctx context.Context) ([]m.Category, error) {
	log := s.logger.With("method", "GetAllCategories")

	cat, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		log.ErrorContext(ctx, "failed to get all categoriese", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success get all categories")

	return cat, nil
}
