package db

import (
	"context"
	"fmt"
)

func createTables(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS loan_records (
		id SERIAL PRIMARY KEY,
		borrower_name VARCHAR(100) NOT NULL,
		monthly_income NUMERIC NOT NULL,
		monthly_debts NUMERIC NOT NULL,
		loan_amount NUMERIC NOT NULL,
		property_value NUMERIC NOT NULL,
		credit_score INT NOT NULL,
		occupancy VARCHAR(50) NOT NULL,
		decision VARCHAR(10) NOT NULL,
		dti NUMERIC NOT NULL,
		ltv NUMERIC NOT NULL,
		reason TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Failed to create tables: %w", err)
	}
	return nil
}

func InsertLoanRecord(record LoanRecord) error {
	query := `INSERT INTO loan_records (borrower_name, monthly_income, monthly_debts, loan_amount, property_value, credit_score, occupancy, decision, dti, ltv, reason) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := db.ExecContext(context.Background(), query,
		record.Name,
		record.MonthlyIncome,
		record.MonthlyDebts,
		record.LoanAmount,
		record.PropertyValue,
		record.CreditScore,
		record.Occupancy,
		record.Decision,
		record.DTI,
		record.LTV,
		record.Reason,
	)

	if err != nil {
		return fmt.Errorf("Failed to insert loan record: %w", err)
	}

	return nil
}

func GetAllLoanRecords(ctx context.Context) ([]LoanRecord, error) {
	query := `SELECT id, borrower_name, monthly_income, monthly_debts, loan_amount, property_value, credit_score, occupancy, decision, dti, ltv, reason, created_at 
			  FROM loan_records ORDER BY created_at DESC`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to query loan records: %w", err)
	}
	defer rows.Close()

	records := []LoanRecord{}
	for rows.Next() {
		var record LoanRecord
		err := rows.Scan(
			&record.ID,
			&record.Name,
			&record.MonthlyIncome,
			&record.MonthlyDebts,
			&record.LoanAmount,
			&record.PropertyValue,
			&record.CreditScore,
			&record.Occupancy,
			&record.Decision,
			&record.DTI,
			&record.LTV,
			&record.Reason,
			&record.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan loan record: %w", err)
		}
		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Row iteration error: %w", err)
	}
	return records, nil
}
