package expense

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yann-fk-21/expense_tracker/types"
	"github.com/yann-fk-21/expense_tracker/utils"
	"net/http"
	"strconv"
)

type Handler struct {
	store types.ExpenseStore
}

func NewHandler(store types.ExpenseStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutesHandler(router *mux.Router) {
	router.HandleFunc("/expenses", h.getExpenses).Methods("GET")
	router.HandleFunc("/expenses/{id}", h.getExpense).Methods("GET")
	router.HandleFunc("/expenses", h.createExpense).Methods("POST")
	router.HandleFunc("/expenses/{id}", h.updateExpense).Methods("PUT")
	router.HandleFunc("/expenses/{id}", h.deleteExpense).Methods("DELETE")
}

func (h *Handler) createExpense(w http.ResponseWriter, r *http.Request) {
	var expense types.ExpenseCreated
	err := utils.ParseJson(r, &expense)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(expense); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid body, err=%v", errors))
		return
	}

	newExpense, err := h.store.CreateExpense(expense)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, newExpense)

}

func (h *Handler) getExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.store.GetExpenses()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, expenses)

}

func (h *Handler) getExpense(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idStr, ok := params["id"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("empty id"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	expense, err := h.store.GetExpense(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if expense.ID == 0 {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("expense %v not found", id))
		return
	}

	utils.WriteJson(w, http.StatusOK, expense)
}

func (h *Handler) updateExpense(w http.ResponseWriter, r *http.Request) {

	var expenseCreated types.ExpenseCreated
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing id expense"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ParseJson(r, &expenseCreated)

	expense, err := h.store.GetExpense(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	err = h.store.UpdateExpense(expense.ID, expenseCreated)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"response": fmt.Sprintf("expense %v is updated", expense.ID)})
}

func (h *Handler) deleteExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("missing id params"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	_, err = h.store.GetExpense(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("expense not found"))
		return
	}

	err = h.store.DeleteExpense(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": fmt.Sprintf("expense with ID %v is deleted", id)})
}
