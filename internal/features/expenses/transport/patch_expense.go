package expenses_transport

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/request"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (h *ExpensesHTTPHandler) PatchExpense(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		rh.ErrorResponse(
			core_errors.ErrUnauthorized,
			"unauthorized",
		)
		return
	}

	idStr := r.PathValue("id")
	expenseID, err := uuid.Parse(idStr)
	if err != nil {
		rh.ErrorResponse(
			core_errors.ErrInvalidArgument,
			"failed to parse id",
		)
		return
	}

	var patchRequest ExpensePatchRequest

	if err := request.DecodeAndValidateRequest(r, &patchRequest); err != nil {
		rh.ErrorResponse(
			err,
			"failed to decode and validate request")
		return
	}

	patch, err := parsePatchFields(patchRequest)
	if err != nil {
		rh.ErrorResponse(
			err,
			"invalid request field",
		)
		return
	}

	expense, err := h.expensesService.PatchExpense(ctx, userID, expenseID, patch)
	if err != nil {
		rh.ErrorResponse(
			err,
			"failed to update user",
		)
		return
	}

	rh.JSONResponse(
		http.StatusOK,
		expenseDTO(expense),
	)
}

func parsePatchFields(patchRequest ExpensePatchRequest) (domain.ExpensePatch, error) {

	var expense domain.ExpensePatch

	if patchRequest.CategoryID != nil {
		categoryID, err := uuid.Parse(*patchRequest.CategoryID)
		if err != nil {
			return domain.ExpensePatch{}, fmt.Errorf("invalid category_id: %w", core_errors.ErrInvalidArgument)
		}
		expense.CategoryID = &categoryID
	}

	if patchRequest.Amount != nil {
		amount, err := decimal.NewFromString(*patchRequest.Amount)
		if err != nil {
			return domain.ExpensePatch{}, fmt.Errorf("invalid amount: %w", core_errors.ErrInvalidArgument)
		}
		expense.Amount = &amount
	}

	if patchRequest.ExpenseDate != nil {
		expenseDate, err := time.Parse("2006-01-02", *patchRequest.ExpenseDate)
		if err != nil {
			return domain.ExpensePatch{}, fmt.Errorf("invalid expense_date: %w", core_errors.ErrInvalidArgument)
		}
		expense.ExpenseDate = &expenseDate
	}

	expense.Description = patchRequest.Description

	return expense, nil
}
