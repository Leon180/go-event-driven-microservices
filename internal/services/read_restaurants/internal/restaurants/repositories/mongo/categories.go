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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewSyncCategoriesMongo(
	db *mongo.Client,
	collections *mongodb.Collections,
	contextLogger contextloggers.ContextLogger,
) repositories.SyncCategoriesMongo {
	return &syncCategoriesMongoImpl{
		db:            db,
		collections:   collections,
		contextLogger: contextLogger,
	}
}

type syncCategoriesMongoImpl struct {
	db            *mongo.Client
	collections   *mongodb.Collections
	contextLogger contextloggers.ContextLogger
}

func (impl *syncCategoriesMongoImpl) SyncCategories(ctx context.Context, newCategories aggregates.Categories) error {
	collection := impl.collections.Category

	// Convert new categories to map for easy lookup
	newCategoryMap := make(map[enums.CategoryCode]*documents.Category)
	for _, category := range newCategories {
		doc := category.ToDocument()
		newCategoryMap[category.CategoryCode] = doc
	}

	// Get existing categories
	var existingCategories []documents.Category
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to get existing categories", err)
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &existingCategories); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to get existing categories", err)
		return err
	}

	// Prepare bulk operations
	var operations []mongo.WriteModel

	// Handle updates and identifies deletes
	existingIDs := make(map[enums.CategoryCode]bool)
	for _, existing := range existingCategories {
		existingIDs[existing.CategoryCode] = true

		if _, exists := newCategoryMap[existing.CategoryCode]; !exists {
			// Delete if not in new set
			operation := mongo.NewDeleteOneModel().
				SetFilter(bson.M{"_id": existing.ID})
			operations = append(operations, operation)
		}
	}

	// Handle inserts for new categories
	for _, category := range newCategoryMap {
		if !existingIDs[category.CategoryCode] {
			operation := mongo.NewInsertOneModel().
				SetDocument(category)
			operations = append(operations, operation)
		}
	}

	// Execute bulk write if there are operations
	if len(operations) > 0 {
		opts := options.BulkWrite().SetOrdered(false)
		_, err := collection.BulkWrite(ctx, operations, opts)
		if err != nil {
			impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
				Error("failed to sync categories", err)
			return err
		}
	}

	return nil
}

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
