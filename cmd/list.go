package cmd

import "github.com/mesh-dell/expense-tracker/internal/expense"

func ListCommand(category string) error {
	return expense.ListExpenses(category)
}
