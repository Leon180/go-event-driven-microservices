package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/repositories"
	"gorm.io/gorm"
)

type createCreditCardImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewCreateCreditCard(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.CreateCreditCard {
	return &createCreditCardImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *createCreditCardImpl) CreateCreditCard(ctx context.Context, card *entities.CreditCard) error {
	if card == nil {
		return nil
	}
	if err := handle.db.WithContext(ctx).Create(card).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error creating credit card")
		return err
	}
	return nil
}

type createCreditCardsImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewCreateCreditCards(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.CreateCreditCards {
	return &createCreditCardsImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *createCreditCardsImpl) CreateCreditCards(ctx context.Context, cards entities.CreditCards) error {
	if cards == nil {
		return nil
	}
	if err := handle.db.WithContext(ctx).Create(cards).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error creating credit cards")
		return err
	}
	return nil
}
