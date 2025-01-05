package expense_handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/request"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/response"
)

func (h *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "UpdateExpense")

	user, ok := ctx.Value("user").(*auth.UserData)
	if !ok {
		log.ErrorContext(
			ctx, "failed to type cast user data",
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		log.ErrorContext(ctx, "invalid userID", "error", err, "user_id", userID)
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	idString := r.URL.Path[len("/expenses/"):]
	expenseID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.ErrorContext(ctx, "invalid expense ID", "error", err, "user_id", userID)
		http.Error(w, "Invalid expense ID", http.StatusBadRequest)
		return
	}

	var req models.ExpenseRequest
	if err := request.JSON(w, r, &req); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
			"user_id", userID,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	expense, err := h.service.UpdateExpense(ctx, userID, expenseID, req.Data)
	if err != nil {
		if err == sql.ErrNoRows {
			log.ErrorContext(ctx, "expense not found", "user_id", userID)
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		log.ErrorContext(ctx, "failed to update expense", "error", err, "user_id", userID)
		http.Error(w, fmt.Sprintf("Failed to update expense: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	resp := models.ExpenseResponse{
		Data: expense,
	}

	if err := response.JSON(
		w,
		http.StatusCreated,
		resp,
	); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx, "success update expense",
		"user_id", userID)
}
