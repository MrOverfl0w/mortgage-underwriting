package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"MrOverflow.github.io/mortgage-underwriting/backend/db"
	"MrOverflow.github.io/mortgage-underwriting/backend/functions"
)

func LoanSolicitationHandler(w http.ResponseWriter, r *http.Request) {
	var requestData LoanSolicitationRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		fmt.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request, please check your inputs", http.StatusBadRequest)
		return
	}
	dti, err := functions.CalculateDTI(requestData.MonthlyDebts, requestData.MonthlyIncome)
	if err != nil {
		http.Error(w, "Error calculating DTI: "+err.Error(), http.StatusBadRequest)
		return
	}
	ltv, err := functions.CalculateLTV(requestData.LoanAmount, requestData.PropertyValue)
	if err != nil {
		http.Error(w, "Error calculating LTV: "+err.Error(), http.StatusBadRequest)
		return
	}
	decision, reason := functions.GenerateLoanDecision(requestData.CreditScore, dti, ltv, requestData.Occupancy, requestData.LoanAmount, requestData.PropertyValue)

	go func() {
		err := db.InsertLoanRecord(db.LoanRecord{
			Name:          requestData.Name,
			MonthlyIncome: requestData.MonthlyIncome,
			MonthlyDebts:  requestData.MonthlyDebts,
			LoanAmount:    requestData.LoanAmount,
			PropertyValue: requestData.PropertyValue,
			CreditScore:   requestData.CreditScore,
			Occupancy:     requestData.Occupancy,
			Decision:      decision,
			DTI:           dti,
			LTV:           ltv,
			Reason:        reason,
		})

		if err != nil {
			log.Printf("Failed to insert loan record: %v", err)
		}
	}()

	responseData := LoanSolicitationResponse{
		Decision: decision,
		DTI:      dti,
		LTV:      ltv,
		Reason:   reason,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func LoanHistoryHandler(w http.ResponseWriter, r *http.Request) {
	records, err := db.GetAllLoanRecords(r.Context())
	if err != nil {
		http.Error(w, "Error retrieving loan records: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}
