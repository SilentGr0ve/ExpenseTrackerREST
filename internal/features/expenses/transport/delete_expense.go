package expenses_transport

import (
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
	"github.com/google/uuid"
)

// DeleteExpense godoc
// @Summary Delete an expense
// @Description Delete an expense by ID for the authorized user
// @Tags expenses
// @Security BearerAuth
// @Param id path string true "Expense ID"
// @Success 204 "Expense deleted"
// @Failure 400 {object} response.ErrorResponse "Invalid expense id"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Expense not found"
// @Router /expenses/{id} [delete]
func (h *ExpensesHTTPHandler) DeleteExpense(rw http.ResponseWriter, r *http.Request) {
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
		rh.ErrorResponse(err, "failed to parse id")
		return
	}

	err = h.expensesService.DeleteExpense(ctx, userID, expenseID)
	if err != nil {
		rh.ErrorResponse(err, "failed to delete expense")
		return
	}

	rh.NoContent()
}
