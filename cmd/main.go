package main

import (
	"expenses_tracker/internal/expenses"
	"fmt"
	"time"
)

func main() {
	chipotleExpense := expenses.NewExpenses(1, time.Now().Local(), "Chipotle", 11.56)
	fmt.Println(chipotleExpense)
}

// Adding an Expense

// Updating an Expense

// Delete an Expense

// View all Expenses

// View specific Expenses in a month (of current year)
