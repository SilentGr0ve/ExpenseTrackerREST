package expenses_transport

import (
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateExpenseRequest struct {
	CategoryID  uuid.UUID       `json:"category_id" validate:"required,uuid"`
	Amount      decimal.Decimal `json:"amount" validate:"required"`
	Description *string         `json:"description" validate:"omitempty"`
	ExpenseDate *string         `json:"expense_date" validate:"omitempty"`
}

type ExpenseResponse struct {
	ID          uuid.UUID       `json:"id"`
	UserID      uuid.UUID       `json:"user_id"`
	CategoryID  uuid.UUID       `json:"category_id"`
	Version     int             `json:"version"`
	Amount      decimal.Decimal `json:"amount"`
	Description *string         `json:"description"`
	ExpenseDate string          `json:"expense_date"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at"`
}

type ExpensePatchRequest struct {
	CategoryID  *string `json:"category_id" validate:"omitempty,uuid"`
	Amount      *string `json:"amount" validate:"omitempty"`
	ExpenseDate *string `json:"expense_date" validate:"omitempty"`
	Description *string `json:"description" validate:"omitempty"`
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
