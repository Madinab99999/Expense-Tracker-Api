package expense_handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
	expense_service "github.com/Madinab99999/Expense-Tracker-Api/internal/service/expense"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/response"
)

func (h *ExpenseHandler) GetExpenseStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetExpenseStats")

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

	startDate := extractString(r, "start_date")
	endDate := extractString(r, "end_date")
	if startDate == nil || endDate == nil {
		log.ErrorContext(ctx, "start_date and end_date are required", "user_id", userID)
		http.Error(w, "start_date and end_date are required", http.StatusBadRequest)
		return
	}

	expenseStats, err := h.service.GetExpenseStats(ctx, userID, *startDate, *endDate)
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
		log.ErrorContext(ctx, "failed to get expense statistics", "error", err, "user_id", userID)
		http.Error(w, fmt.Sprintf("Failed to get expense statistics: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if expenseStats == nil {
		log.InfoContext(ctx, "expenses not found", "user_id", userID)
		http.Error(w, "expenses not found", http.StatusNotFound)
		return
	}

	resp := models.ExpenseStatsResponse{
		Data: expenseStats,
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

	log.InfoContext(ctx, "success get expense statistics", "user_id", userID)
}
