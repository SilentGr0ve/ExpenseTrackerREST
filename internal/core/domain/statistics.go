package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Statistics struct {
	Period     Period
	Summary    Summary
	ByCategory []CategorySummary
}

type Period struct {
	From *time.Time
	To   *time.Time
}

type Summary struct {
	TotalAmount   decimal.Decimal
	TotalExpenses int
	AverageAmount decimal.Decimal
}

type CategorySummary struct {
	CategoryID    uuid.UUID
	CategoryName  string
	TotalAmount   decimal.Decimal
	ExpensesCount int
}
