package expenses_transport

import (
	"context"
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	httpserver "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http"
	"github.com/google/uuid"
)

type ExpensesHTTPHandler struct {
	expensesService ExpensesService
}

type ExpensesService interface {
	CreateExpense(ctx context.Context, userID uuid.UUID, newExpense domain.NewExpense) (domain.Expense, error)
	GetExpense(ctx context.Context, userID uuid.UUID, expenseID uuid.UUID) (domain.Expense, error)
	DeleteExpense(ctx context.Context, userID uuid.UUID, expenseID uuid.UUID) error
	GetExpenses(ctx context.Context, userID uuid.UUID, query domain.ExpensesQuery) ([]domain.Expense, error)
	PatchExpense(ctx context.Context, userID uuid.UUID, expenseID uuid.UUID, patch domain.ExpensePatch) (domain.Expense, error)
}

func NewExpensesHTTPHandler(expensesService ExpensesService) *ExpensesHTTPHandler {
	return &ExpensesHTTPHandler{
		expensesService: expensesService,
	}
}

func (h *ExpensesHTTPHandler) ProtectedRoutes() []httpserver.Route {
	return []httpserver.Route{
		{Method: http.MethodPost, Path: "/expenses", Handler: h.CreateExpense},
		{Method: http.MethodGet, Path: "/expenses/{id}", Handler: h.GetExpense},
		{Method: http.MethodDelete, Path: "/expenses/{id}", Handler: h.DeleteExpense},
		{Method: http.MethodGet, Path: "/expenses", Handler: h.GetExpenses},
		{Method: http.MethodPatch, Path: "/expenses/{id}", Handler: h.PatchExpense},
	}
}
