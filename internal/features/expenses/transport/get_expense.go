package expenses_transport

import (
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
	"github.com/google/uuid"
)

// GetExpense godoc
// @Summary Get an expense by ID
// @Description Returns a single expense by its ID for the authorized user
// @Tags expenses
// @Produce json
// @Security BearerAuth
// @Param id path string true "Expense ID"
// @Success 200 {object} ExpenseResponse "Expense found"
// @Failure 400 {object} response.ErrorResponse "Invalid expense ID"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /expenses/{id} [get]
func (h *ExpensesHTTPHandler) GetExpense(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	id := r.PathValue("id")
	expenseID, err := uuid.Parse(id)
	if err != nil {
		rh.ErrorResponse(core_errors.ErrInvalidArgument, "failed to parse id")
		return
	}

	expense, err := h.expensesService.GetExpense(ctx, userID, expenseID)
	if err != nil {
		rh.ErrorResponse(err, "failed to get expense")
		return
	}

	expenseResponse := ExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		CategoryID:  expense.CategoryID,
		Version:     expense.Version,
		Amount:      expense.Amount,
		Description: expense.Description,
		ExpenseDate: expense.ExpenseDate.Format("2026-01-02"),
		CreatedAt:   expense.CreatedAt,
		UpdatedAt:   expense.UpdatedAt,
	}

	rh.JSONResponse(
		http.StatusOK,
		expenseResponse,
	)
}
