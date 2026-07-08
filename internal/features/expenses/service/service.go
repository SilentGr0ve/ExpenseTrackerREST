package expenses_service

import (
	"context"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

type ExpensesService struct {
	expensesRepository ExpensesRepository
}

type ExpensesRepository interface {
	CreateExpense(ctx context.Context, expense domain.Expense) (domain.Expense, error)
	GetExpense(ctx context.Context, userID uuid.UUID, expenseID uuid.UUID) (domain.Expense, error)
	DeleteExpense(ctx context.Context, userID uuid.UUID, expenseID uuid.UUID) error
	GetExpenses(ctx context.Context, userID uuid.UUID, expenseQuery domain.ExpensesQuery) ([]domain.Expense, error)
	UpdateExpense(ctx context.Context, userID uuid.UUID, expense domain.Expense) (domain.Expense, error)
}

func NewExpensesService(expensesRepository ExpensesRepository) *ExpensesService {
	return &ExpensesService{
		expensesRepository: expensesRepository,
	}
}
