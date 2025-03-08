package services

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type ListCategories interface {
	ListCategories(ctx context.Context) (entities.Categories, error)
}

func NewListCategories(listCategories repositories.ListCategories) ListCategories {
	return &listCategoriesImpl{
		listCategories: listCategories,
	}
}

type listCategoriesImpl struct {
	listCategories repositories.ListCategories
}

func (impl *listCategoriesImpl) ListCategories(ctx context.Context) (entities.Categories, error) {
	return impl.listCategories.ListCategories(ctx)
}
