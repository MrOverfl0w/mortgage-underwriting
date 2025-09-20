package db

type LoanRecord struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	MonthlyIncome float64 `json:"monthly_income"`
	MonthlyDebts  float64 `json:"monthly_debts"`
	LoanAmount    float64 `json:"loan_amount"`
	PropertyValue float64 `json:"property_value"`
	CreditScore   int     `json:"credit_score"`
	Occupancy     string  `json:"occupancy"`
	Decision      string  `json:"decision"`
	DTI           float64 `json:"dti"`
	LTV           float64 `json:"ltv"`
	Reason        string  `json:"reason,omitempty"`
	CreatedAt     string  `json:"created_at"`
}
