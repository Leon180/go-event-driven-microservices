package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
)

//go:generate mockgen -source=restaurants.go -destination=./mocks/restaurants_mock.go -package=mocks

type SearchRestaurantsFullInfo interface {
	SearchRestaurantsFullInfo(ctx context.Context, searchRestaurants dtos.SearchRestaurants) ([]aggregates.Restaurant, error)
}
type Restaurants interface {
	CreateRestaurants(ctx context.Context, restaurants entities.Restaurants) error
	ReadRestaurantFullInfo(ctx context.Context, id string) (aggregates.Restaurant, error)
	UpdateRestaurant(ctx context.Context, updateRestaurant entities.UpdateRestaurant) error
	DeleteRestaurants(ctx context.Context, ids []string) error
}
type Branches interface {
	CreateBranches(ctx context.Context, branches entities.Branches) error
	UpdateBranch(ctx context.Context, updateBranch entities.UpdateBranch) error
	DeleteBranches(ctx context.Context, ids []string) error
}

type Addresses interface {
	CreateAddresses(ctx context.Context, addresses entities.Addresses) error
	UpdateAddress(ctx context.Context, updateAddress entities.UpdateAddress) error
	DeleteAddresses(ctx context.Context, ids []string) error
}

type PriceRanges interface {
	CreatePriceRanges(ctx context.Context, priceRanges entities.PriceRanges) error
	UpdatePriceRange(ctx context.Context, updatePriceRange entities.UpdatePriceRange) error
	DeletePriceRanges(ctx context.Context, ids []string) error
}

type BranchCategoryRelations interface {
	CreateBranchCategoryRelations(ctx context.Context, branchCategoryRelations entities.BranchCategoryRelations) error
	UpdateBranchCategoryRelation(ctx context.Context, updateBranchCategoryRelation entities.UpdateBranchCategoryRelation) error
	DeleteBranchCategoryRelations(ctx context.Context, ids []string) error
}

type Tables interface {
	CreateTables(ctx context.Context, tables entities.Tables) error
	UpdateTable(ctx context.Context, updateTable entities.UpdateTable) error
	DeleteTables(ctx context.Context, ids []string) error
}

type TableAvailables interface {
	CreateTableAvailables(ctx context.Context, tableAvailables entities.TableAvailables) error
	UpdateTableAvailable(ctx context.Context, updateTableAvailable entities.UpdateTableAvailable) error
	DeleteTableAvailables(ctx context.Context, ids []string) error
}
