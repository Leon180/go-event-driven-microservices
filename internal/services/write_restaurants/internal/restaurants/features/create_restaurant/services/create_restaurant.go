package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/messaging/serializers"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	createRestaurantEvents "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/create_restaurant/events"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type CreateRestaurant interface {
	CreateRestaurant(ctx context.Context, req *dtos.Restaurant) error
}

func NewCreateRestaurant(
	uuidGenerator uuid.UUIDGenerator,
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction],
	searchRestaurantsFullInfoRepository repositories.SearchRestaurantsFullInfo,
	createRestaurantMessageBuilder createRestaurantEvents.CreateRestaurantMessageBuilder,
	messageSerializer serializers.MessageSerializer,
) CreateRestaurant {
	return &createRestaurantImpl{
		uuidGenerator: uuidGenerator,
		updateRestaurantsWithTransactionRepository: updateRestaurantsWithTransactionRepository,
		searchRestaurantsFullInfoRepository:        searchRestaurantsFullInfoRepository,
		createRestaurantMessageBuilder:             createRestaurantMessageBuilder,
		messageSerializer:                          messageSerializer,
	}
}

type createRestaurantImpl struct {
	uuidGenerator                              uuid.UUIDGenerator
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction]
	searchRestaurantsFullInfoRepository        repositories.SearchRestaurantsFullInfo
	createRestaurantMessageBuilder             createRestaurantEvents.CreateRestaurantMessageBuilder
	messageSerializer                          serializers.MessageSerializer
}

func (handle *createRestaurantImpl) CreateRestaurant(ctx context.Context, req *dtos.Restaurant) error {
	if req == nil {
		return nil
	}

	// check if restaurant already exists
	restaurants, err := handle.searchRestaurantsFullInfoRepository.SearchRestaurantsFullInfo(
		ctx,
		&dtos.SearchRestaurants{
			NameFilter:        &req.Name,
			NamePreciseSearch: true,
		},
	)
	if err != nil {
		return err
	}
	if lo.ContainsBy(restaurants, func(restaurant aggregates.Restaurant) bool {
		return restaurant.ActiveStatus
	}) {
		return customizeerrors.RestaurantAlreadyExistsError
	}

	if lo.ContainsBy(restaurants, func(restaurant aggregates.Restaurant) bool {
		return !restaurant.ActiveStatus
	}) {
		return customizeerrors.RestaurantAlreadyExistsButInactiveError
	}

	// build restaurant create entities by aggregate
	restaurantDTOAggregateBuilder := aggregates.NewRestaurantDTOAggregateBuilder(handle.uuidGenerator)
	if err := restaurantDTOAggregateBuilder.SaveRestaurant(req); err != nil {
		return err
	}
	editEntities := restaurantDTOAggregateBuilder.GetEditEntities()
	if len(editEntities) == 0 || editEntities[0].CreateEntities == nil {
		return nil
	}
	createEntities := *editEntities[0].CreateEntities

	aggregates := restaurantDTOAggregateBuilder.GetAggregates()
	outboxMessages := make(entities.OutboxMessages, len(aggregates))
	for i, restaurant := range aggregates {
		message := handle.createRestaurantMessageBuilder.Build(&restaurant)
		serializationResult, err := handle.messageSerializer.Serialize(message)
		if err != nil {
			return err
		}
		t := time.Now()
		outboxMessages[i] = entities.OutboxMessage{
			ID:        handle.uuidGenerator.GenerateUUID(),
			MessageID: message.ID(),
			Type:      message.Type(),
			Payload:   serializationResult.Data,
			Status:    enums.OutboxStatusPending,
			CreatedAt: t,
			UpdatedAt: t,
		}
	}

	// create
	tx, err := handle.updateRestaurantsWithTransactionRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := tx.CreateRestaurants(ctx, createEntities.Restaurants); err != nil {
		return err
	}
	if err := tx.CreateBranches(ctx, createEntities.Branches); err != nil {
		return err
	}
	if err := tx.CreateAddresses(ctx, createEntities.Addresses); err != nil {
		return err
	}
	if err := tx.CreatePriceRanges(ctx, createEntities.PriceRanges); err != nil {
		return err
	}
	if err := tx.CreateBranchCategoryRelations(ctx, createEntities.BranchCategoryRelations); err != nil {
		return err
	}
	if err := tx.CreateTables(ctx, createEntities.Tables); err != nil {
		return err
	}
	if err := tx.CreateAvailables(ctx, createEntities.Availables); err != nil {
		return err
	}

	// store outbox messages
	if err := tx.CreateOutboxMessages(ctx, outboxMessages); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
