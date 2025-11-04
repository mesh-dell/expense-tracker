package cmd

import (
	"flag"
	"fmt"
)

func Execute(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("expense-tracker No command provided")
	}

	command := args[1]

	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		amountFlag := addCmd.Float64("amount", 0, "amount of an expense")
		descriptionFlag := addCmd.String("description", "", "description of an expense")
		categoryFlag := addCmd.String("category", "", "category of the expense")

		addCmd.Parse(args[2:])

		if *amountFlag <= 0 {
			return fmt.Errorf("please provide a valid amount")
		}
		if *descriptionFlag == "" || *categoryFlag == "" {
			return fmt.Errorf("please provide a description and category")
		}
		return AddCommand(*descriptionFlag, *categoryFlag, *amountFlag)
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		categoryFlag := listCmd.String("category", "all", "category filter")
		listCmd.Parse(args[2:])
		return ListCommand(*categoryFlag)
	case "summary":
		summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		monthFlag := summaryCmd.Int("month", 0, "month filter")
		summaryCmd.Parse(args[2:])
		return SummaryCommand(*monthFlag)
	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		idFlag := deleteCmd.Int("id", 0, "expense id to delete")
		deleteCmd.Parse(args[2:])

		if *idFlag == 0 {
			return fmt.Errorf("please provide an id")
		}

		return DeleteCommand(*idFlag)
	case "budget":
		budgetFlag := flag.NewFlagSet("budget", flag.ExitOnError)
		monthFlag := budgetFlag.Int("month", 0, "month to budget")
		amountFlag := budgetFlag.Float64("amount", 0, "amount to allocate")
		budgetFlag.Parse(args[2:])

		if *monthFlag <= 0 {
			return fmt.Errorf("please enter a month")
		}

		if *amountFlag <= 0 {
			return fmt.Errorf("please enter a valid amount")
		}
		return BudgetCommand(*monthFlag, *amountFlag)
	case "update":
		updateFlag := flag.NewFlagSet("update", flag.ExitOnError)
		idFlag := updateFlag.Int("id", 0, "id of expense to update")
		amountFlag := updateFlag.Float64("amount", 0, "new amount")

		descriptionFlag := updateFlag.String("description", "", "new description")
		categoryFlag := updateFlag.String("category", "", "new category")
		updateFlag.Parse(args[2:])

		if *idFlag == 0 {
			return fmt.Errorf("please provide an id")
		}

		if *amountFlag <= 0 {
			return fmt.Errorf("please provide a valid amount")
		}
		if *descriptionFlag == "" {
			return fmt.Errorf("please provide a description")
		}
		if *categoryFlag == "" {
			return fmt.Errorf("please provide a category")
		}
		return UpdateCommand(*idFlag, *amountFlag, *descriptionFlag, *categoryFlag)
	case "help":
		PrintHelp()
		return nil
	default:
		return fmt.Errorf("expense-tracker unknown command %s", command)
	}
}

func PrintHelp() {
	fmt.Println(`
Expense Tracker CLI
-------------------
Manage and summarize your expenses.

Usage:
  expense-tracker [command] [options]

Commands:

  add
    Add a new expense.
    Options:
      --description <text>    Expense description
      --amount <number>       Expense amount

  update
    Update an existing expense.
    Options:
      --id <number>           Expense ID
      --description <text>    New description (optional)
      --amount <number>       New amount (optional)

  delete
    Delete an expense by ID.
    Options:
      --id <number>           Expense ID

  list
    View all expenses.
    Options:
	  --category <string>     Expense Category

  summary
    View total or monthly summary.
    Options:
      --month <1-12>          Filter by month (current year)

  budget
	Add a monthly budget
	Options
	  --month <1-12>          Month to set budget (current year)
	  --amount <decimal>	  Amount to allocate to budget	  
  help
    Show this message.

Examples:
  expense-tracker add --description "Lunch" --amount 20 --category "food"
  expense-tracker update --id 1 --amount 25 --description "Supper" --category "food"
  expense-tracker budget --month 12 --amount 250
  expense-tracker delete --id 2
  expense-tracker list --category "Luxuries"
  expense-tracker summary --month 8
	`)
}
