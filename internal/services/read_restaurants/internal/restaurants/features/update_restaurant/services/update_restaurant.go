package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
)

type UpdateRestaurant interface {
	UpdateRestaurant(ctx context.Context, req *dtos.Restaurant) error
}

type updateRestaurantImpl struct {
	uuidGenerator                              uuid.UUIDGenerator
	readRestaurantsRepository                  repositories.ReadRestaurants
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction]
	contextlogger                              contextloggers.ContextLogger
}

func NewUpdateRestaurant(
	uuidGenerator uuid.UUIDGenerator,
	readRestaurantsRepository repositories.ReadRestaurants,
	updateRestaurantsWithTransactionRepository customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction],
	contextlogger contextloggers.ContextLogger,
) UpdateRestaurant {
	return &updateRestaurantImpl{
		uuidGenerator:                              uuidGenerator,
		readRestaurantsRepository:                  readRestaurantsRepository,
		updateRestaurantsWithTransactionRepository: updateRestaurantsWithTransactionRepository,
		contextlogger:                              contextlogger,
	}
}

func (handle *updateRestaurantImpl) UpdateRestaurant(ctx context.Context, req *dtos.Restaurant) error {
	if req == nil || req.ID == nil {
		return nil
	}
	// check if restaurant exists
	restaurant, err := handle.readRestaurantsRepository.ReadRestaurantFullInfo(ctx, *req.ID)
	if err != nil {
		return err
	}
	if restaurant == nil {
		return customizeerrors.RestaurantNotFoundError
	}

	// build restaurant update entities by aggregate
	restaurantDTOAggregateBuilder := aggregates.NewRestaurantDTOAggregateBuilder(handle.uuidGenerator)
	err = restaurantDTOAggregateBuilder.SaveRestaurant(restaurant.ToDTO())
	if err != nil {
		return err
	}
	restaurantDTOAggregateBuilder.SetAllEditTypeCodeToNone()
	err = restaurantDTOAggregateBuilder.SaveRestaurant(req)
	if err != nil {
		return err
	}
	editEntities := restaurantDTOAggregateBuilder.GetEditEntities()
	if len(editEntities) == 0 {
		return customizeerrors.NoChangesError
	}
	if len(editEntities) > 1 {
		handle.contextlogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("multiple edit entities error occur while updating restaurant, should only have one edit entity")
		return customizeerrors.HTTPInternalServerError
	}
	createEntities := editEntities[0].CreateEntities
	updateEntities := editEntities[0].UpdateEntities
	deleteEntities := editEntities[0].DeleteEntities

	tx, err := handle.updateRestaurantsWithTransactionRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = handle.createRestaurants(ctx, tx, createEntities)
	if err != nil {
		return err
	}

	err = handle.updateRestaurants(ctx, tx, updateEntities)
	if err != nil {
		return err
	}
	err = handle.deleteRestaurants(ctx, tx, deleteEntities)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (handle *updateRestaurantImpl) createRestaurants(
	ctx context.Context,
	tx repositories.UpdateRestaurantsWithTransaction,
	createEntities *aggregates.RestaurantCreateEntities,
) error {
	if createEntities == nil {
		return nil
	}
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
	return nil
}

func (handle *updateRestaurantImpl) updateRestaurants(
	ctx context.Context,
	tx repositories.UpdateRestaurantsWithTransaction,
	updateEntities *aggregates.RestaurantUpdateEntities,
) error {
	if updateEntities == nil {
		return nil
	}
	for _, restaurant := range updateEntities.Restaurants {
		if err := tx.UpdateRestaurant(ctx, &restaurant); err != nil {
			return err
		}
	}
	for _, branch := range updateEntities.Branches {
		if err := tx.UpdateBranch(ctx, &branch); err != nil {
			return err
		}
	}
	for _, address := range updateEntities.Addresses {
		if err := tx.UpdateAddress(ctx, &address); err != nil {
			return err
		}
	}
	for _, priceRange := range updateEntities.PriceRanges {
		if err := tx.UpdatePriceRange(ctx, &priceRange); err != nil {
			return err
		}
	}
	for _, branchCategoryRelation := range updateEntities.BranchCategoryRelations {
		if err := tx.UpdateBranchCategoryRelation(ctx, &branchCategoryRelation); err != nil {
			return err
		}
	}
	for _, table := range updateEntities.Tables {
		if err := tx.UpdateTable(ctx, &table); err != nil {
			return err
		}
	}
	for _, available := range updateEntities.Availables {
		if err := tx.UpdateAvailable(ctx, &available); err != nil {
			return err
		}
	}
	return nil
}

func (handle *updateRestaurantImpl) deleteRestaurants(
	ctx context.Context,
	tx repositories.UpdateRestaurantsWithTransaction,
	deleteEntities *aggregates.RestaurantDeleteEntities,
) error {
	if deleteEntities == nil {
		return nil
	}
	for _, restaurant := range deleteEntities.Restaurants {
		if err := tx.UpdateRestaurant(ctx, &entities.UpdateRestaurant{ID: restaurant.ID, ActiveStatus: lo.ToPtr(false)}); err != nil {
			return err
		}
	}
	for _, branch := range deleteEntities.Branches {
		if err := tx.UpdateBranch(ctx, &entities.UpdateBranch{ID: branch.ID, ActiveStatus: lo.ToPtr(false)}); err != nil {
			return err
		}
	}
	for _, address := range deleteEntities.Addresses {
		if err := tx.UpdateAddress(ctx, &entities.UpdateAddress{ID: address.ID, ActiveStatus: lo.ToPtr(false)}); err != nil {
			return err
		}
	}
	for _, priceRange := range deleteEntities.PriceRanges {
		if err := tx.UpdatePriceRange(ctx, &entities.UpdatePriceRange{ID: priceRange.ID, ActiveStatus: lo.ToPtr(false)}); err != nil {
			return err
		}
	}
	for _, branchCategoryRelation := range deleteEntities.BranchCategoryRelations {
		if err := tx.UpdateBranchCategoryRelation(ctx, &entities.UpdateBranchCategoryRelation{ID: branchCategoryRelation.ID, ActiveStatus: lo.ToPtr(false)}); err != nil {
			return err
		}
	}
	for _, table := range deleteEntities.Tables {
		if err := tx.UpdateTable(ctx, &entities.UpdateTable{ID: table.ID, ActiveStatus: lo.ToPtr(false)}); err != nil {
			return err
		}
	}
	for _, available := range deleteEntities.Availables {
		if err := tx.UpdateAvailable(ctx, &entities.UpdateAvailable{ID: available.ID, ActiveStatus: lo.ToPtr(false)}); err != nil {
			return err
		}
	}
	return nil
}
