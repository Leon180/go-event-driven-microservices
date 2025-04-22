package queries

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type GetRestaurantBranch struct {
	BranchID string
}

type GetRestaurantBranchHandler interface {
	GetRestaurantBranch(
		ctx context.Context,
		command *GetRestaurantBranch,
	) (*aggregates.Restaurant, error)
}

func NewGetRestaurantBranchHandler(
	readRestaurantBranchMongo repositories.ReadRestaurantBranchMongo,
) GetRestaurantBranchHandler {
	handle := &getRestaurantBranchImpl{
		readRestaurantBranchMongo: readRestaurantBranchMongo,
	}
	return handle
}

type getRestaurantBranchImpl struct {
	readRestaurantBranchMongo repositories.ReadRestaurantBranchMongo
}

func (handle *getRestaurantBranchImpl) GetRestaurantBranch(
	ctx context.Context,
	command *GetRestaurantBranch,
) (*aggregates.Restaurant, error) {
	if command == nil {
		return nil, nil
	}
	if command.BranchID == "" {
		return nil, customizeerrors.InvalidIDError
	}
	restaurant, err := handle.readRestaurantBranchMongo.ReadRestaurantBranch(ctx, command.BranchID)
	if err != nil {
		return nil, err
	}
	if restaurant == nil {
		return nil, customizeerrors.RestaurantNotFoundError
	}
	return restaurant, nil
}
