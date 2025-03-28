package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
)

type ListCategories interface {
	ListCategories(ctx context.Context) (entities.Categories, error)
}
