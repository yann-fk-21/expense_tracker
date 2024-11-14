package expense

import (
	"database/sql"
	"github.com/yann-fk-21/expense_tracker/types"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetExpenses() ([]types.Expense, error) {
	rows, err := s.db.Query("SELECT * FROM expenses")
	if err != nil {
		return nil, err
	}

	expenses := make([]types.Expense, 0)

	for rows.Next() {
		expense := new(types.Expense)
		err := rows.Scan(&expense.ID, &expense.Title, &expense.Cost)
		if err != nil {
			log.Fatal(err)
		}
		expenses = append(expenses, *expense)
	}

	return expenses, nil
}

func (s *Store) CreateExpense(expense types.ExpenseCreated) (*types.Expense, error) {
	result, err := s.db.Exec("INSERT INTO expenses(title, cost) VALUES (?, ?)", expense.Title, expense.Cost)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	newExpense := &types.Expense{
		ID:    int(id),
		Title: expense.Title,
		Cost:  expense.Cost,
	}

	return newExpense, nil
}

func (s *Store) GetExpense(id int) (*types.Expense, error) {
	rows, err := s.db.Query("SELECT * FROM expenses WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	expense := new(types.Expense)
	for rows.Next() {
		err = rows.Scan(&expense.ID, &expense.Title, &expense.Cost)
		if err != nil {
			return nil, err
		}
	}
	return expense, nil
}

func (s *Store) UpdateExpense(id int, expenseCreated types.ExpenseCreated) error {
	_, err := s.db.Exec("UPDATE expenses SET title = ?, cost = ? WHERE id = ?", expenseCreated.Title, expenseCreated.Cost, id)
	return err
}

func (s *Store) DeleteExpense(id int) error {
	_, err := s.db.Exec("DELETE FROM expenses WHERE id = ?", id)
	return err
}
