package expense_handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	expense_service "github.com/Madinab99999/Expense-Tracker-Api/internal/api/service/expense"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/response"
)

func (h *ExpenseHandler) GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetAllExpenses")

	user, ok := ctx.Value("user").(*auth.UserData)
	if !ok {
		log.ErrorContext(
			ctx,
			"failed to type cast user data",
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		log.ErrorContext(ctx, "invalid userID", "error", err)
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	params := models.ExpenseParametrs{
		Filter: &models.ExpenseFilter{
			Category:  extractCategory(r),
			TimeRange: extractTimeRange(r),
			StartDate: extractString(r, "start_date"),
			EndDate:   extractString(r, "end_date"),
			MinAmount: extractInt(r, "min_amount"),
			MaxAmount: extractInt(r, "max_amount"),
		},
		Sort: &models.SortOptions{
			SortBy: models.SortField(r.URL.Query().Get("sort_by")),
			Order:  models.SortOrder(r.URL.Query().Get("order")),
		},
		Pagination: models.Pagination{
			Cursor: extractInt(r, "cursor"),
			Limit:  extractLimit(r),
		},
	}

	expenses, err := h.service.GetAllExpenses(ctx, userID, params)
	if err != nil {
		if errors.Is(err, expense_service.ErrInvalidOptions) {
			log.ErrorContext(ctx, "invalid options", "user_id", userID)
			http.Error(w, "invalid options", http.StatusBadRequest)
			return
		}
		if err == sql.ErrNoRows {
			log.ErrorContext(ctx, "expenses not found", "user_id", userID)
			http.Error(w, "Expenses not found", http.StatusNotFound)
			return
		}
		log.ErrorContext(ctx, "failed to get all expenses", "error", err, "user_id", userID)
		http.Error(w, fmt.Sprintf("Failed to get all expenses: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if expenses == nil {
		log.InfoContext(ctx, "expenses not found", "user_id", userID)
		http.Error(w, "expenses not found", http.StatusNotFound)
		return
	}

	resp := models.ExpensesResponse{
		Data: &models.AllExpenses{
			Expenses: expenses,
		},
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		resp,
	); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err, "user_id", userID)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx, "success get all expenses",
		"user_id", userID)
}

func extractString(r *http.Request, key string) *string {
	value := r.URL.Query().Get(key)
	if value != "" {
		return &value
	}
	return nil
}

func extractInt(r *http.Request, key string) *int {
	value := r.URL.Query().Get(key)
	if value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return &intValue
		}
	}
	return nil
}

func extractTimeRange(r *http.Request) *models.TimeRange {
	value := r.URL.Query().Get("time_range")
	if value != "" {
		tr := models.TimeRange(value)
		return &tr
	}
	return nil
}

func extractCategory(r *http.Request) *models.Category {
	value := r.URL.Query().Get("category")
	if value != "" {
		tr := models.Category(value)
		return &tr
	}
	return nil
}

func extractLimit(r *http.Request) int {
	limit := 10
	value := r.URL.Query().Get("limit")
	if value != "" {
		if parsedLimit, err := strconv.Atoi(value); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}
	return limit
}
