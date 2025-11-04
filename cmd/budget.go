package cmd

import "github.com/mesh-dell/expense-tracker/internal/expense"

func BudgetCommand(month int, amount float64) error {
	return expense.BudgetMonth(month, amount)
}
