package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type CreateRestaurant interface {
	CreateRestaurant(ctx context.Context, req *dtos.Restaurant) error
}

func NewCreateRestaurant(
	uuidGenerator uuid.UUIDGenerator,
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction],
	searchRestaurantsFullInfoRepository repositories.SearchRestaurantsFullInfo,
) CreateRestaurant {
	return &createRestaurantImpl{
		uuidGenerator: uuidGenerator,
		updateRestaurantsWithTransactionRepository: updateRestaurantsWithTransactionRepository,
		searchRestaurantsFullInfoRepository:        searchRestaurantsFullInfoRepository,
	}
}

type createRestaurantImpl struct {
	uuidGenerator                              uuid.UUIDGenerator
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction]
	searchRestaurantsFullInfoRepository        repositories.SearchRestaurantsFullInfo
}

func (handle *createRestaurantImpl) CreateRestaurant(ctx context.Context, req *dtos.Restaurant) error {
	if req == nil {
		return nil
	}

	// check if restaurant already exists
	restaurants, err := handle.searchRestaurantsFullInfoRepository.SearchRestaurantsFullInfo(ctx, &dtos.SearchRestaurants{
		NameFilter:        &req.Name,
		NamePreciseSearch: true,
	})
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
	err = restaurantDTOAggregateBuilder.SaveRestaurant(req)
	if err != nil {
		return err
	}
	editEntities := restaurantDTOAggregateBuilder.GetEditEntities()
	if len(editEntities) == 0 || editEntities[0].CreateEntities == nil {
		return nil
	}
	createEntities := *editEntities[0].CreateEntities

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

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
