package expense

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yann-fk-21/expense_tracker/types"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Normal test
func TestExpenseHandlers(t *testing.T) {

	expenseStore := &mockExpenseStore{}
	handler := NewHandler(expenseStore)

	t.Run("should return get expense", func(t *testing.T) {
		statusCode := getExpenseHandler("/expenses", "/expenses", http.MethodGet, handler.getExpenses)

		if statusCode != http.StatusOK {
			t.Errorf("there error, code is %v, got %v ", statusCode, http.StatusOK)
		}
	})

	t.Run("should return status code 404 if expense not exist", func(t *testing.T) {
		statusCode := getExpenseHandler("/expenses/6", "/expenses/{id}", http.MethodGet, handler.getExpense)

		if statusCode != http.StatusNotFound {
			t.Errorf("there error, code is %v, got %v ", statusCode, http.StatusNotFound)
		}
	})

	t.Run("should create expense", func(t *testing.T) {
		payload := types.ExpenseCreated{
			Title: "Buy",
			Cost:  90.80,
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/expenses", bytes.NewBuffer(marshalled))
		if err != nil {
			log.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/expenses", handler.createExpense).Methods("POST")
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			t.Errorf("there error, code is %v, got %v", rr.Code, http.StatusCreated)
		}
	})

}

// Benchmark test
func BenchmarkExpenseHandlers(b *testing.B) {
	expenseStore := &mockExpenseStore{}
	handler := NewHandler(expenseStore)

	// benchmark test of getExpenses
	b.Run("performance of get expenses handler", func(b *testing.B) {
		statusCode := getExpenseHandler("/expenses", "/expenses", http.MethodGet, handler.getExpenses)

		if statusCode != http.StatusOK {
			b.Errorf("there error, code is %v, got %v ", statusCode, http.StatusOK)
		}
	})
}

type mockExpenseStore struct {
}

func (m *mockExpenseStore) GetExpenses() ([]types.Expense, error) {
	return make([]types.Expense, 0), nil
}
func (m *mockExpenseStore) GetExpense(ID int) (*types.Expense, error) {
	return &types.Expense{}, nil
}
func (m *mockExpenseStore) CreateExpense(expense types.ExpenseCreated) (*types.Expense, error) {
	return &types.Expense{}, nil
}
func (m *mockExpenseStore) UpdateExpense(ID int, expense types.ExpenseCreated) error {
	return nil
}

func getExpenseHandler(path, url, method string, handler func(w http.ResponseWriter, r *http.Request)) int {

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc(url, handler).Methods(method)
	router.ServeHTTP(rr, req)
	return rr.Code
}
