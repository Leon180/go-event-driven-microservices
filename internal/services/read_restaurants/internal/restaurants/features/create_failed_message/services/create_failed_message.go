package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type CreateFailedMessageHandler interface {
	CreateFailedMessage(ctx context.Context, aggregate *aggregates.FailedMessage) error
}

func NewCreateFailedMessageHandler(
	readFailedMessageMongo repositories.ReadFailedMessageMongo,
	createFailedMessageMongo repositories.CreateFailedMessageMongo,
) CreateFailedMessageHandler {
	return &createFailedMessageImpl{
		readFailedMessageMongo:   readFailedMessageMongo,
		createFailedMessageMongo: createFailedMessageMongo,
	}
}

type createFailedMessageImpl struct {
	readFailedMessageMongo   repositories.ReadFailedMessageMongo
	createFailedMessageMongo repositories.CreateFailedMessageMongo
}

func (handle *createFailedMessageImpl) CreateFailedMessage(
	ctx context.Context,
	aggregate *aggregates.FailedMessage,
) error {
	if aggregate == nil {
		return nil
	}
	existingFailedMessage, err := handle.readFailedMessageMongo.ReadFailedMessage(ctx, aggregate.MessageID)
	if err != nil {
		return err
	}
	if existingFailedMessage != nil {
		return customizeerrors.FailedMessageAlreadyExistsError
	}
	if err := handle.createFailedMessageMongo.Create(ctx, aggregate); err != nil {
		return err
	}
	return nil
}
