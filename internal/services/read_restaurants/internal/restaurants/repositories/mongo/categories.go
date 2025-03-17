package repositoriesmongo

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/mongodb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewListCategoriesMongo(
	db *mongo.Client,
	collections *mongodb.Collections,
	contextLogger contextloggers.ContextLogger,
) repositories.ListCategoriesMongo {
	return &listCategoriesMongoImpl{
		db:            db,
		collections:   collections,
		contextLogger: contextLogger,
	}
}

type listCategoriesMongoImpl struct {
	db            *mongo.Client
	collections   *mongodb.Collections
	contextLogger contextloggers.ContextLogger
}

func (impl *listCategoriesMongoImpl) ListCategories(ctx context.Context) (aggregates.Categories, error) {
	collection := impl.collections.Category
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to list categories", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories documents.Categories
	if err := cursor.All(ctx, &categories); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to list categories", err)
		return nil, err
	}
	return aggregates.CategoryDocuments(categories).ToAggregate(), nil
}
