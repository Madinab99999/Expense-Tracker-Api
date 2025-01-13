package category_repository

import (
	"context"

	m "github.com/Madinab99999/Expense-Tracker-Api/internal/models"
)

func (r *CategoryRepository) GetAllCategories(ctx context.Context) ([]m.Category, error) {
	log := r.logger.With("method", "GetAllCategories")
	log.InfoContext(ctx, "success get all categories")
	return []m.Category{
		m.Groceries,
		m.Leisure,
		m.Electronics,
		m.Utilities,
		m.Clothing,
		m.Health,
		m.Transport,
		m.Education,
		m.Credits,
		m.Others,
	}, nil
}
