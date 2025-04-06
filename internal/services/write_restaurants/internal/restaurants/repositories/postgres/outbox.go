package repostgresespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"gorm.io/gorm"
)

func NewUpdateOutboxMessages(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateOutboxMessages {
	return &updateOutboxMessagesImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type updateOutboxMessagesImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *updateOutboxMessagesImpl) CreateOutboxMessages(
	ctx context.Context,
	entities entities.OutboxMessages,
) error {
	if len(entities) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&entities).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create outbox messages", err)
		return err
	}
	return nil
}

func (impl *updateOutboxMessagesImpl) UpdateOutboxMessage(
	ctx context.Context,
	update *entities.UpdateOutboxMessage,
) error {
	if update == nil || update.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.OutboxMessage{}).Where("id = ?", update.ID).Updates(update.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update outbox message", err)
		return err
	}
	return nil
}

func (impl *updateOutboxMessagesImpl) DeleteOutboxMessages(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.OutboxMessage{}).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete outbox messages", err)
		return err
	}
	return nil
}

func NewUpdateOutboxMessagesWithTransaction(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) customizegorm.Transactor[repositories.UpdateOutboxMessagesWithTransaction] {
	return &UpdateOutboxMessagesWithTransactionImpl{
		TransactorImpl: TransactorImpl{
			db:            db,
			contextLogger: contextLogger,
		},
	}
}

type UpdateOutboxMessagesWithTransactionImpl struct {
	TransactorImpl
}

func (impl *UpdateOutboxMessagesWithTransactionImpl) BeginTx(
	ctx context.Context,
) (repositories.UpdateOutboxMessagesWithTransaction, error) {
	tx := impl.db.WithContext(ctx).Begin()
	return &UpdateOutboxMessagesTransactionImpl{
		TransactionImpl: TransactionImpl{
			Db:            tx,
			ContextLogger: impl.contextLogger,
		},
	}, nil
}

type UpdateOutboxMessagesTransactionImpl struct {
	TransactionImpl
}

func (impl *UpdateOutboxMessagesTransactionImpl) CreateOutboxMessages(
	ctx context.Context,
	entities entities.OutboxMessages,
) error {
	return NewUpdateOutboxMessages(impl.Db, impl.ContextLogger).CreateOutboxMessages(ctx, entities)
}

func (impl *UpdateOutboxMessagesTransactionImpl) UpdateOutboxMessage(
	ctx context.Context,
	update *entities.UpdateOutboxMessage,
) error {
	return NewUpdateOutboxMessages(impl.Db, impl.ContextLogger).UpdateOutboxMessage(ctx, update)
}

func (impl *UpdateOutboxMessagesTransactionImpl) DeleteOutboxMessages(
	ctx context.Context,
	ids []string,
) error {
	return NewUpdateOutboxMessages(impl.Db, impl.ContextLogger).DeleteOutboxMessages(ctx, ids)
}

func NewListOutboxMessages(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ListOutboxMessages {
	return &listOutboxMessagesImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type listOutboxMessagesImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *listOutboxMessagesImpl) ListOutboxMessages(
	ctx context.Context,
	status []enums.OutboxStatus,
	retry *int,
) (entities.OutboxMessages, error) {
	var outboxMessages entities.OutboxMessages
	query := impl.db.WithContext(ctx)
	if len(status) > 0 {
		query = query.Where("status IN (?)", status)
	}
	if retry != nil {
		query = query.Where("retry_count <= ?", retry)
	}
	if err := query.Find(&outboxMessages).Order("updated_at ASC").Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to list outbox messages", err)
		return nil, err
	}
	return outboxMessages, nil
}
