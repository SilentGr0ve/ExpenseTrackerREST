package expenses_service

import (
	"context"
	"fmt"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
)

func (s *ExpensesService) CreateExpense(ctx context.Context, userID uuid.UUID, newExpense domain.NewExpense) (domain.Expense, error) {
	expense := domain.Expense{
		ID:          uuid.New(),
		UserID:      userID,
		CategoryID:  newExpense.CategoryID,
		Version:     1,
		Amount:      newExpense.Amount,
		Description: newExpense.Description,
		CreatedAt:   time.Now(),
		ExpenseDate: newExpense.ExpenseDate,
		UpdatedAt:   nil,
	}

	created, err := s.expensesRepository.CreateExpense(ctx, expense)
	if err != nil {
		return domain.Expense{}, fmt.Errorf("create expense: %w", err)
	}

	return created, nil
}
