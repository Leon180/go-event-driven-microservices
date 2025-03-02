package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories"
	"gorm.io/gorm"
)

type readCreditCardByMobileNumberAndActiveSwitchImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadCreditCardByMobileNumberAndActiveSwitch(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadCreditCardByMobileNumberAndActiveSwitch {
	return &readCreditCardByMobileNumberAndActiveSwitchImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readCreditCardByMobileNumberAndActiveSwitchImpl) ReadCreditCardByMobileNumberAndActiveSwitch(
	ctx context.Context,
	mobileNumber string,
	activeSwitch *bool,
) (entities.CreditCards, error) {
	var cards entities.CreditCards
	sql := handle.db.WithContext(ctx)
	if activeSwitch != nil {
		sql = sql.Where("mobile_number = ? AND active_switch = ?", mobileNumber, *activeSwitch)
	} else {
		sql = sql.Where("mobile_number = ?", mobileNumber)
	}
	if err := sql.Find(&cards).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("error reading credit card by mobile number and active switch")
		return nil, err
	}
	return cards, nil
}

type readCreditCardImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadCreditCard(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadCreditCard {
	return &readCreditCardImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readCreditCardImpl) ReadCreditCard(ctx context.Context, id string) (*entities.CreditCard, error) {
	var card entities.CreditCard
	if err := handle.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&card).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error reading credit card by id")
		return nil, err
	}
	if card.ID == "" {
		return nil, nil
	}
	return &card, nil
}
