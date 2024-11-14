package types

type ExpenseStore interface {
	GetExpenses() ([]Expense, error)
	GetExpense(ID int) (*Expense, error)
	CreateExpense(expense ExpenseCreated) (*Expense, error)
	UpdateExpense(ID int, expense ExpenseCreated) error
	DeleteExpense(ID int) error
}

type ExpenseCreated struct {
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
}

type Expense struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
}
