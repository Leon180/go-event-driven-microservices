package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/create_loan/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/create_loan/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories"
	loannumberutilities "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/utilities/loan_number"
	"github.com/shopspring/decimal"
)

type CreateLoan interface {
	CreateLoan(ctx context.Context, req *featuresdtos.CreateLoanRequest) error
}

func NewCreateLoan(
	createLoanRepository repositories.CreateLoan,
	readLoanByMobileNumberAndActiveSwitch repositories.ReadLoanByMobileNumberAndActiveSwitch,
	uuidGenerator uuid.UUIDGenerator,
	loanNumberGenerator loannumberutilities.LoanNumberGenerator,
) CreateLoan {
	return &createLoanImpl{
		createLoanRepository:                  createLoanRepository,
		readLoanByMobileNumberAndActiveSwitch: readLoanByMobileNumberAndActiveSwitch,
		uuidGenerator:                         uuidGenerator,
		loanNumberGenerator:                   loanNumberGenerator,
	}
}

type createLoanImpl struct {
	createLoanRepository                  repositories.CreateLoan
	readLoanByMobileNumberAndActiveSwitch repositories.ReadLoanByMobileNumberAndActiveSwitch
	uuidGenerator                         uuid.UUIDGenerator
	loanNumberGenerator                   loannumberutilities.LoanNumberGenerator
}

func (handle *createLoanImpl) CreateLoan(ctx context.Context, req *featuresdtos.CreateLoanRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateCreateLoanRequest(req); err != nil {
		return err
	}

	activeSwitch := true
	loans, err := handle.readLoanByMobileNumberAndActiveSwitch.ReadLoanByMobileNumberAndActiveSwitch(
		ctx,
		req.MobileNumber,
		&activeSwitch,
	)
	if err != nil {
		return err
	}
	if len(loans) > 0 {
		return customizeerrors.LoanAlreadyExistsError
	}

	systemTime := time.Now()
	loanEntity := entities.Loan{
		ID:           handle.uuidGenerator.GenerateUUID(),
		MobileNumber: req.MobileNumber,
		LoanNumber:   handle.loanNumberGenerator.GenerateLoanNumber(),
		LoanTypeCode: req.LoanType.ToLoanTypeCode(),
		TotalAmount: func() decimal.Decimal {
			totalAmount, _ := decimal.NewFromString(req.TotalAmount)
			return totalAmount
		}(),
		PaidAmount: decimal.Zero,
		InterestRate: func() decimal.Decimal {
			interestRate, _ := decimal.NewFromString(req.InterestRate)
			return interestRate
		}(),
		Term:         req.Term,
		ActiveSwitch: true,
		CommonHistoryModelWithUpdate: entities.CommonHistoryModelWithUpdate{
			CommonHistoryModel: entities.CommonHistoryModel{
				CreatedAt: systemTime,
			},
			UpdatedAt: systemTime,
		},
	}

	return handle.createLoanRepository.CreateLoan(ctx, &loanEntity)
}
