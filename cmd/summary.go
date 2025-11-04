package cmd

import "github.com/mesh-dell/expense-tracker/internal/expense"

func SummaryCommand(month int) error {
	return expense.SummaryExpenses(month)
}
