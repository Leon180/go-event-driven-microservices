package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
)

type SyncCategoriesMongo interface {
	SyncCategories(ctx context.Context, categories aggregates.Categories) error
}

type ListCategoriesMongo interface {
	ListCategories(ctx context.Context) (aggregates.Categories, error)
}
