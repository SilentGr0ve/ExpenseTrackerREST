package statistics_transport

import (
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StatisticsResponse struct {
	Period     PeriodResponse            `json:"period"`
	Summary    SummaryResponse           `json:"summary"`
	ByCategory []CategorySummaryResponse `json:"by_category"`
}

type PeriodResponse struct {
	From string `json:"date_from"`
	To   string `json:"date_to"`
}

type SummaryResponse struct {
	TotalAmount   decimal.Decimal `json:"total_amount"`
	TotalExpenses int             `json:"total_expenses"`
	AverageAmount decimal.Decimal `json:"average_amount"`
}

type CategorySummaryResponse struct {
	CategoryID    uuid.UUID       `json:"category_id"`
	CategoryName  string          `json:"category_name"`
	TotalAmount   decimal.Decimal `json:"total_amount"`
	ExpensesCount int             `json:"expenses_count"`
}

func toResponse(statistics domain.Statistics) StatisticsResponse {

	var fromStr, toStr string
	if statistics.Period.From != nil {
		fromStr = statistics.Period.From.Format("2006-01-02")
	}
	if statistics.Period.To != nil {
		toStr = statistics.Period.To.Format("2006-01-02")
	}
	periodResponse := PeriodResponse{
		From: fromStr,
		To:   toStr,
	}

	summaryResponse := SummaryResponse{
		TotalAmount:   statistics.Summary.TotalAmount,
		TotalExpenses: statistics.Summary.TotalExpenses,
		AverageAmount: statistics.Summary.AverageAmount,
	}

	var byCategorySummary []CategorySummaryResponse

	for _, category := range statistics.ByCategory {

		categorySummaryResponse := CategorySummaryResponse{
			CategoryID:    category.CategoryID,
			CategoryName:  category.CategoryName,
			TotalAmount:   category.TotalAmount,
			ExpensesCount: category.ExpensesCount,
		}

		byCategorySummary = append(byCategorySummary, categorySummaryResponse)
	}

	statisticsResponse := StatisticsResponse{
		Period:     periodResponse,
		Summary:    summaryResponse,
		ByCategory: byCategorySummary,
	}

	return statisticsResponse
}
