package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Expense struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	CategoryID uuid.UUID
	Version    int

	Amount      decimal.Decimal
	Description *string
	CreatedAt   time.Time
	ExpenseDate time.Time
	UpdatedAt   *time.Time
}
