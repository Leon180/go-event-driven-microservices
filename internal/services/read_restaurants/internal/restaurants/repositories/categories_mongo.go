package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
)

type ListCategoriesMongo interface {
	ListCategories(ctx context.Context) (aggregates.Categories, error)
}
