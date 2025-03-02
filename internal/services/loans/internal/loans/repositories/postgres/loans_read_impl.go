package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/repositories"
	"gorm.io/gorm"
)

type readLoanByMobileNumberAndActiveSwitchImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadLoanByMobileNumberAndActiveSwitch(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadLoanByMobileNumberAndActiveSwitch {
	return &readLoanByMobileNumberAndActiveSwitchImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readLoanByMobileNumberAndActiveSwitchImpl) ReadLoanByMobileNumberAndActiveSwitch(
	ctx context.Context,
	mobileNumber string,
	activeSwitch *bool,
) (entities.Loans, error) {
	var loans entities.Loans
	sql := handle.db.WithContext(ctx)
	if activeSwitch != nil {
		sql = sql.Where("mobile_number = ? AND active_switch = ?", mobileNumber, *activeSwitch)
	} else {
		sql = sql.Where("mobile_number = ?", mobileNumber)
	}
	if err := sql.Find(&loans).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("error reading loan by mobile number and active switch")
		return nil, err
	}
	return loans, nil
}

type readLoanImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadLoan(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadLoan {
	return &readLoanImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readLoanImpl) ReadLoan(ctx context.Context, id string) (*entities.Loan, error) {
	var loan entities.Loan
	if err := handle.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&loan).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error reading loan by id")
		return nil, err
	}
	if loan.ID == "" {
		return nil, nil
	}
	return &loan, nil
}
