package expenses

import "time"

type Expense struct {
	id          uint8
	date        time.Time
	description string
	amount      float64
}

func NewExpenses(id uint8, date time.Time, description string, amount float64) *Expense {
	return &Expense{
		id:          id,
		date:        date,
		description: description,
		amount:      amount,
	}
}
