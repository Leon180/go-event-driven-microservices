package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
)

type ListCategoriesRedis interface {
	ListCategories(ctx context.Context) (documents.Categories, error)
}

type SetCategoriesRedis interface {
	SetCategory(ctx context.Context, category *documents.Category) error
}
