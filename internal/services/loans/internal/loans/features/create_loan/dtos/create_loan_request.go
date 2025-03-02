package featuresdtos

import "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"

type CreateLoanRequest struct {
	// @example 0958308280
	MobileNumber string `json:"mobile_number" binding:"required"`
	// @example "home", "car"
	LoanType enums.LoanType `json:"loan_type"     binding:"required"`
	// @example 10000000
	TotalAmount string `json:"total_amount"  binding:"required"`
	// @example 0.03
	InterestRate string `json:"interest_rate" binding:"required"`
	// @example 84
	Term int `json:"term"          binding:"required"`
}
