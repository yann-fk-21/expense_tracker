package api

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yann-fk-21/expense_tracker/services/expense"
	"net/http"
)

type Server struct {
	Addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *Server {
	return &Server{
		Addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {

	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	expenseStore := expense.NewStore(s.db)

	expenseHandlers := expense.NewHandler(expenseStore)

	expenseHandlers.RegisterRoutesHandler(subRouter)

	fmt.Println("server is running...")
	return http.ListenAndServe(s.Addr, router)
}
