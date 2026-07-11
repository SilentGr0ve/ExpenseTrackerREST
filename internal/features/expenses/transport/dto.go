package expenses_transport

import (
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateExpenseRequest struct {
	CategoryID  uuid.UUID       `json:"category_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Amount      decimal.Decimal `json:"amount" validate:"required" example:"599.50"`
	Description *string         `json:"description" validate:"omitempty" example:"Description for a new expense"`
	ExpenseDate *string         `json:"expense_date" validate:"omitempty" example:"2026-12-31"`
}

type ExpenseResponse struct {
	ID          uuid.UUID       `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserID      uuid.UUID       `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	CategoryID  uuid.UUID       `json:"category_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Version     int             `json:"version" example:"1"`
	Amount      decimal.Decimal `json:"amount" example:"599.50"`
	Description *string         `json:"description" example:"Description for an expense"`
	ExpenseDate string          `json:"expense_date" example:"2026-12-31"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at"`
}

type ExpensePatchRequest struct {
	CategoryID  *string `json:"category_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Amount      *string `json:"amount" validate:"omitempty" example:"599.50"`
	ExpenseDate *string `json:"expense_date" validate:"omitempty" example:"2026-12-31"`
	Description *string `json:"description" validate:"omitempty" example:"Patched description for an expense"`
}

func expenseDTO(expense domain.Expense) ExpenseResponse {
	return ExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		CategoryID:  expense.CategoryID,
		Version:     expense.Version,
		Amount:      expense.Amount,
		Description: expense.Description,
		ExpenseDate: expense.ExpenseDate.Format("2006-01-02"),
		CreatedAt:   expense.CreatedAt,
		UpdatedAt:   expense.UpdatedAt,
	}
}

func ToDTOFromDomain(expenses []domain.Expense) []ExpenseResponse {
	expensesResponse := make([]ExpenseResponse, 0, len(expenses))

	for _, expense := range expenses {
		expensesResponse = append(expensesResponse, expenseDTO(expense))
	}

	return expensesResponse
}
