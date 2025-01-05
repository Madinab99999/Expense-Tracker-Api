package expense_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/auth"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/request"
	"github.com/Madinab99999/Expense-Tracker-Api/pkg/httputils/response"
)

func (h *ExpenseHandler) InsertExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "InsertExpense")

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
		log.ErrorContext(
			ctx, "invalid userID",
			"error", err,
			"user_id", userID,
		)
		http.Error(w, "Invalid userID", http.StatusBadRequest)
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

	expense, err := h.service.InsertExpense(ctx, req.Data, userID)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to insert expense",
			"error", err,
			"user_id", userID,
		)
		http.Error(w, "failed to insert expense", http.StatusInternalServerError)
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
		log.ErrorContext(ctx, "failed to encode response", "error", err, "user_id", userID)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx, "success insert expense",
		"user_id", userID,
	)
}
