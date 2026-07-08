package expenses_transport

import (
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"

	"net/http"
	"time"
)

func (h *ExpensesHTTPHandler) CreateExpense(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	var expenseRequest CreateExpenseRequest
	if err := request.DecodeAndValidateRequest(r, &expenseRequest); err != nil {
		rh.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	var expenseDate time.Time
	if expenseRequest.ExpenseDate == nil || *expenseRequest.ExpenseDate == "" {
		expenseDate = time.Now()
	} else {
		parsed, err := time.Parse("2006-01-02", *expenseRequest.ExpenseDate)
		if err != nil {
			rh.ErrorResponse(core_errors.ErrInvalidArgument, "invalid expense_date format, user YYYY-MM-DD")
			return
		}
		expenseDate = parsed
	}

	newExpense := domain.NewExpense{
		CategoryID:  expenseRequest.CategoryID,
		Amount:      expenseRequest.Amount,
		Description: expenseRequest.Description,
		ExpenseDate: expenseDate,
	}

	expense, err := h.expensesService.CreateExpense(
		ctx,
		userID,
		newExpense,
	)

	if err != nil {
		rh.ErrorResponse(err, "failed to create expense")
		return
	}

	expenseResponse := expenseDTO(expense)

	rh.JSONResponse(
		http.StatusCreated,
		expenseResponse,
	)
}
