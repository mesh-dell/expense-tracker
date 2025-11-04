package cmd

import "github.com/mesh-dell/expense-tracker/internal/expense"

func DeleteCommand(id int) error {
	return expense.DeleteExpense(id)
}
