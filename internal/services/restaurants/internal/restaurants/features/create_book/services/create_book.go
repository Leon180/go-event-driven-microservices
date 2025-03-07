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

type CreateBook interface {
	CreateBook(ctx context.Context, req *dtos.Book) error
}

func NewCreateBook(
	uuidGenerator uuid.UUIDGenerator,
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction],
	searchBooksFullInfoRepository repositories.SearchBooksFullInfo,
) CreateBook {
	return &createBookImpl{
		uuidGenerator:                        uuidGenerator,
		updateBooksWithTransactionRepository: updateBooksWithTransactionRepository,
		searchBooksFullInfoRepository:        searchBooksFullInfoRepository,
	}
}

type createBookImpl struct {
	uuidGenerator                        uuid.UUIDGenerator
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction]
	searchBooksFullInfoRepository        repositories.SearchBooksFullInfo
}

func (handle *createBookImpl) CreateBook(ctx context.Context, req *dtos.Book) error {
	if req == nil {
		return nil
	}
	// build restaurant create entities by aggregate
	bookDTOAggregateBuilder := aggregates.NewBookDTOAggregateBuilder(handle.uuidGenerator)
	err := bookDTOAggregateBuilder.SaveBook(req)
	if err != nil {
		return err
	}
	editEntities := bookDTOAggregateBuilder.GetEditEntities()
	if len(editEntities) == 0 || editEntities[0].CreateEntities == nil {
		return nil
	}
	createEntities := *editEntities[0].CreateEntities

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
