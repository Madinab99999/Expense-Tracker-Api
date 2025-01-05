package expense_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
)

func (h *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("handler", "DeleteExpense")

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

	idString := r.URL.Path[len("/expenses/"):]
	expenseID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.ErrorContext(ctx, "invalid expense ID", "error", err)
		http.Error(w, "Invalid expense ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteExpense(ctx, userID, expenseID); err != nil {
		if err.Error() == fmt.Sprintf("expense with id %d not found", expenseID) {
			log.ErrorContext(ctx, "expense not found", "user_id", userID)
			http.Error(w, "404 Not Found", http.StatusNotFound)
		} else {
			log.ErrorContext(ctx, "failed to delete expense", "error", err, "user_id", userID)
			http.Error(w, fmt.Sprintf("Failed to delete expense: %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}

	log.InfoContext(
		ctx, "success delete expense",
		"user_id", userID)
}
