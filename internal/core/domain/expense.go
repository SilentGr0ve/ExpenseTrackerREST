package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Expense struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	CategoryID  uuid.UUID
	Version     int
	Amount      decimal.Decimal
	Description *string
	ExpenseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type NewExpense struct {
	CategoryID  uuid.UUID
	Amount      decimal.Decimal
	Description *string
	ExpenseDate time.Time
}

type ExpensesQuery struct {
	CategoryID *uuid.UUID
	DateFrom   *time.Time
	DateTo     *time.Time
	Limit      *int
	Offset     *int
}

type ExpensePatch struct {
	CategoryID  *uuid.UUID
	Amount      *decimal.Decimal
	ExpenseDate *time.Time
	Description *string
}
