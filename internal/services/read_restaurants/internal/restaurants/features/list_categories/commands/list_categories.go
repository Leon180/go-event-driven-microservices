package commands

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type ListCategoriesHandler interface {
	ListCategories(ctx context.Context) (aggregates.Categories, error)
}

func NewListCategories(
	listCategoriesMongo repositories.ListCategoriesMongo,
	listCategoriesRedis repositories.ListCategoriesRedis,
) ListCategoriesHandler {
	return &listCategoriesImpl{
		listCategoriesMongo: listCategoriesMongo,
		listCategoriesRedis: listCategoriesRedis,
	}
}

type listCategoriesImpl struct {
	listCategoriesMongo repositories.ListCategoriesMongo
	listCategoriesRedis repositories.ListCategoriesRedis
}

func (impl *listCategoriesImpl) ListCategories(ctx context.Context) (aggregates.Categories, error) {
	categories, err := impl.listCategoriesRedis.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	if len(categories) > 0 {
		return categories, nil
	}
	categories, err = impl.listCategoriesMongo.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
