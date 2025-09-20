package main

import (
	"MrOverflow.github.io/mortgage-underwriting/backend/handlers"
	"github.com/gorilla/mux"
)

func addRoutes(r *mux.Router) {
	r.HandleFunc("/api/request-loan", handlers.LoanSolicitationHandler).Methods("POST")
	r.HandleFunc("/api/loan-history", handlers.LoanHistoryHandler).Methods("GET")
}
