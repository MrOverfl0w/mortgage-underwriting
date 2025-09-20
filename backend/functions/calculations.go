package functions

import "errors"

func CalculateDTI(monthlyDebts, monthlyIncome float64) (float64, error) {
	if monthlyIncome == 0 {
		return 0, errors.New("monthly income cannot be zero")
	}
	return monthlyDebts / monthlyIncome, nil
}

func CalculateLTV(loanAmount, propertyValue float64) (float64, error) {
	if propertyValue == 0 {
		return 0, errors.New("property value cannot be zero")
	}
	return loanAmount / propertyValue, nil
}
