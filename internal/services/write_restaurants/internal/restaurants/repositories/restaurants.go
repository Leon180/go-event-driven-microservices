package repositories

import (
	"context"

	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
)

//go:generate mockgen -source=restaurants.go -destination=./mocks/restaurants_mock.go -package=mocks

type SearchRestaurantsFullInfo interface {
	SearchRestaurantsFullInfo(
		ctx context.Context,
		searchRestaurants *dtos.SearchRestaurants,
	) (aggregates.Restaurants, error)
}

type ReadRestaurants interface {
	ReadRestaurantFullInfo(ctx context.Context, id string) (*aggregates.Restaurant, error)
	ReadRestaurant(ctx context.Context, id string) (*entities.Restaurant, error)
}

type UpdateRestaurants interface {
	CreateRestaurants(ctx context.Context, restaurants entities.Restaurants) error
	UpdateRestaurant(ctx context.Context, updateRestaurant *entities.UpdateRestaurant) error
	DeleteRestaurants(ctx context.Context, ids []string) error
	CreateBranches(ctx context.Context, branches entities.Branches) error
	UpdateBranch(ctx context.Context, updateBranch *entities.UpdateBranch) error
	DeleteBranches(ctx context.Context, ids []string) error
	CreateAddresses(ctx context.Context, addresses entities.Addresses) error
	UpdateAddress(ctx context.Context, updateAddress *entities.UpdateAddress) error
	DeleteAddresses(ctx context.Context, ids []string) error
	CreatePriceRanges(ctx context.Context, priceRanges entities.PriceRanges) error
	UpdatePriceRange(ctx context.Context, updatePriceRange *entities.UpdatePriceRange) error
	DeletePriceRanges(ctx context.Context, ids []string) error
	CreateBranchCategoryRelations(ctx context.Context, branchCategoryRelations entities.BranchCategoryRelations) error
	UpdateBranchCategoryRelation(
		ctx context.Context,
		updateBranchCategoryRelation *entities.UpdateBranchCategoryRelation,
	) error
	DeleteBranchCategoryRelations(ctx context.Context, ids []string) error
	CreateTables(ctx context.Context, tables entities.Tables) error
	UpdateTable(ctx context.Context, updateTable *entities.UpdateTable) error
	DeleteTables(ctx context.Context, ids []string) error
	CreateAvailables(ctx context.Context, availables entities.Availables) error
	UpdateAvailable(ctx context.Context, updateAvailable *entities.UpdateAvailable) error
	DeleteAvailables(ctx context.Context, ids []string) error
}

type UpdateRestaurantsWithTransaction interface {
	customizegorm.Transaction
	UpdateRestaurants
	UpdateOutboxMessages
}
