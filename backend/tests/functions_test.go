package main

import (
	"errors"
	"testing"

	"MrOverflow.github.io/mortgage-underwriting/backend/functions"
)

func TestGenerateLoanDecision(t *testing.T) {
	tests := []struct {
		name          string
		score         int
		dti           float64
		ltv           float64
		occupancy     string
		loanAmount    float64
		propertyValue float64
		wantDecision  string
		wantReason    string
	}{
		{
			name:          "Declined due to low property value",
			score:         700,
			dti:           0.3,
			ltv:           0.8,
			occupancy:     "primary",
			loanAmount:    100000,
			propertyValue: 70000,
			wantDecision:  "Decline",
			wantReason:    "Property value below minimum ($75,000)",
		},
		{
			name:          "Declined due to low loan amount",
			score:         700,
			dti:           0.3,
			ltv:           0.8,
			occupancy:     "primary",
			loanAmount:    40000,
			propertyValue: 100000,
			wantDecision:  "Decline",
			wantReason:    "Loan amount below minimum ($50,000)",
		},
		{
			name:          "Declined due to low credit score for primary",
			score:         600,
			dti:           0.3,
			ltv:           0.8,
			occupancy:     "primary",
			loanAmount:    100000,
			propertyValue: 100000,
			wantDecision:  "Decline",
			wantReason:    "Credit score too low for primary residence",
		},
		{
			name:          "Declined due to low credit score for investment",
			score:         670,
			dti:           0.3,
			ltv:           0.8,
			occupancy:     "investment",
			loanAmount:    100000,
			propertyValue: 100000,
			wantDecision:  "Decline",
			wantReason:    "Credit score too low for secondary/investment property",
		},
		{
			name:          "Declined due to high LTV for primary",
			score:         700,
			dti:           0.3,
			ltv:           0.91,
			occupancy:     "primary",
			loanAmount:    91000,
			propertyValue: 100000,
			wantDecision:  "Decline",
			wantReason:    "LTV exceeds 90% for primary residence",
		},
		{
			name:          "Declined due to high LTV for investment",
			score:         700,
			dti:           0.3,
			ltv:           0.81,
			occupancy:     "investment",
			loanAmount:    81000,
			propertyValue: 100000,
			wantDecision:  "Decline",
			wantReason:    "LTV exceeds 80% for secondary/investment property",
		},
		{
			name:          "Declined due to high DTI for low credit",
			score:         650,
			dti:           0.41,
			ltv:           0.8,
			occupancy:     "primary",
			loanAmount:    80000,
			propertyValue: 100000,
			wantDecision:  "Decline",
			wantReason:    "DTI exceeds 36% for credit score below 700",
		},
		{
			name:          "Declined due to very high DTI",
			score:         750,
			dti:           0.51,
			ltv:           0.8,
			occupancy:     "primary",
			loanAmount:    80000,
			propertyValue: 100000,
			wantDecision:  "Decline",
			wantReason:    "DTI exceeds 43%",
		},
		{
			name:          "Approved: excellent credit, high LTV allowed",
			score:         750,
			dti:           0.44,
			ltv:           0.95,
			occupancy:     "primary",
			loanAmount:    95000,
			propertyValue: 100000,
			wantDecision:  "Approve",
			wantReason:    "Excellent credit allows higher LTV",
		},
		{
			name:          "Approved: standard approval",
			score:         720,
			dti:           0.42,
			ltv:           0.85,
			occupancy:     "primary",
			loanAmount:    85000,
			propertyValue: 100000,
			wantDecision:  "Approve",
			wantReason:    "Meets standard approval criteria",
		},
		{
			name:          "Referred: does not meet any approval/decline rule",
			score:         700,
			dti:           0.44,
			ltv:           0.85,
			occupancy:     "primary",
			loanAmount:    85000,
			propertyValue: 100000,
			wantDecision:  "Refer",
			wantReason:    "Requires further review",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotDecision, gotReason := functions.GenerateLoanDecision(
				tc.score, tc.dti, tc.ltv, tc.occupancy, tc.loanAmount, tc.propertyValue,
			)
			if gotDecision != tc.wantDecision {
				t.Errorf("Decision: got %v, want %v", gotDecision, tc.wantDecision)
			}
			if gotReason != tc.wantReason {
				t.Errorf("Reason: got %v, want %v", gotReason, tc.wantReason)
			}
		})
	}
}

func TestCalculateDTI(t *testing.T) {
	tests := []struct {
		name          string
		monthlyDebts  float64
		monthlyIncome float64
		wantDTI       float64
		wantErr       error
	}{
		{
			name:          "Normal case",
			monthlyDebts:  1500,
			monthlyIncome: 5000,
			wantDTI:       0.3,
			wantErr:       nil,
		},
		{
			name:          "Zero income",
			monthlyDebts:  1500,
			monthlyIncome: 0,
			wantDTI:       0,
			wantErr:       errors.New("monthly income cannot be zero"),
		},
		{
			name:          "Zero debts",
			monthlyDebts:  0,
			monthlyIncome: 5000,
			wantDTI:       0,
			wantErr:       nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotDTI, gotErr := functions.CalculateDTI(tc.monthlyDebts, tc.monthlyIncome)
			if gotDTI != tc.wantDTI && gotErr != tc.wantErr {
				t.Errorf("DTI: got %v, want %v", gotDTI, tc.wantDTI)
				t.Errorf("Error: got %v, want %v", gotErr, tc.wantErr)
			}
		})
	}
}

func TestCalculateLTV(t *testing.T) {
	tests := []struct {
		name          string
		loanAmount    float64
		propertyValue float64
		wantLTV       float64
		wantErr       error
	}{
		{
			name:          "Normal case",
			loanAmount:    180000,
			propertyValue: 200000,
			wantLTV:       0.9,
			wantErr:       nil,
		},
		{
			name:          "Zero property value",
			loanAmount:    180000,
			propertyValue: 0,
			wantLTV:       0,
			wantErr:       errors.New("property value cannot be zero"),
		},
		{
			name:          "Zero loan amount",
			loanAmount:    0,
			propertyValue: 200000,
			wantLTV:       0,
			wantErr:       nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotLTV, gotErr := functions.CalculateLTV(tc.loanAmount, tc.propertyValue)
			if gotLTV != tc.wantLTV && gotErr != tc.wantErr {
				t.Errorf("LTV: got %v, want %v", gotLTV, tc.wantLTV)
				t.Errorf("Error: got %v, want %v", gotErr, tc.wantErr)
			}
		})
	}
}
