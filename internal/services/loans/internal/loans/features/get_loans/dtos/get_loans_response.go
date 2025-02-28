package featuresdtos

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	"github.com/samber/lo"
)

type GetLoansResponse struct {
	ID           string                            `json:"id"`
	MobileNumber string                            `json:"mobile_number"`
	LoanType     enums.LoanType                    `json:"loan_type"`
	TotalAmount  string                            `json:"total_amount"`
	PaidAmount   string                            `json:"paid_amount"`
	InterestRate string                            `json:"interest_rate"`
	Term         int                               `json:"term"`
	ActiveSwitch bool                              `json:"active_switch"`
	History      dtos.CommonHistoryModelWithUpdate `json:"history"`
}

type LoanEntity entities.Loan

func (c *LoanEntity) ToGetLoansResponse() *GetLoansResponse {
	return &GetLoansResponse{
		ID:           c.ID,
		MobileNumber: c.MobileNumber,
		LoanType:     c.LoanTypeCode.ToLoanType(),
		TotalAmount:  c.TotalAmount.String(),
		PaidAmount:   c.PaidAmount.String(),
		InterestRate: c.InterestRate.String(),
		Term:         c.Term,
		ActiveSwitch: c.ActiveSwitch,
		History: dtos.CommonHistoryModelWithUpdate{
			CommonHistoryModel: dtos.CommonHistoryModel{
				CreatedAt: c.CreatedAt,
				DeletedAt: c.DeletedAt,
			},
			UpdatedAt: c.UpdatedAt,
		},
	}
}

type LoanEntities entities.Loans

func (c LoanEntities) ToGetLoansResponse() []GetLoansResponse {
	return lo.Map(c, func(loan entities.Loan, _ int) GetLoansResponse {
		loanEntity := LoanEntity(loan)
		return *loanEntity.ToGetLoansResponse()
	})
}
