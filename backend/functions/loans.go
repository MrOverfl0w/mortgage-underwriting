package functions

const (
	loanApproved = "Approve"
	loanReferred = "Refer"
	loanDeclined = "Decline"
)

func GenerateLoanDecision(score int, dti, ltv float64, occupancy string, loanAmount, propertyValue float64) (decision string, reason string) {
	// Minimum property and loan amounts
	if propertyValue < 75000 {
		return loanDeclined, "Property value below minimum ($75,000)"
	}
	if loanAmount < 50000 {
		return loanDeclined, "Loan amount below minimum ($50,000)"
	}

	// Special rule: high credit, high DTI, higher LTV allowed
	if score >= 740 && dti <= 0.45 && ltv <= 0.95 {
		return loanApproved, "Excellent credit allows higher LTV"
	}

	// Minimum credit score by occupancy
	switch occupancy {
	case "primary":
		if score < 620 {
			return loanDeclined, "Credit score too low for primary residence"
		}
	case "secondary", "investment":
		if score < 680 {
			return loanDeclined, "Credit score too low for secondary/investment property"
		}
	}

	// Maximum LTV by occupancy
	switch occupancy {
	case "primary":
		if ltv > 0.90 {
			return loanDeclined, "LTV exceeds 90% for primary residence"
		}
	case "secondary", "investment":
		if ltv > 0.80 {
			return loanDeclined, "LTV exceeds 80% for secondary/investment property"
		}
	}

	// Maximum DTI by credit score
	if score < 700 && dti > 0.36 {
		return loanDeclined, "DTI exceeds 36% for credit score below 700"
	}
	if dti > 0.5 {
		return loanDeclined, "DTI exceeds 50%"
	}

	// Standard approval
	if score >= 700 && dti <= 0.43 && ((occupancy == "primary" && ltv <= 0.90) || (ltv <= 0.80)) {
		return loanApproved, "Meets standard approval criteria"
	}

	// Otherwise, refer for manual review
	return loanReferred, "Requires further review"
}
