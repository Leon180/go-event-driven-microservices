package repositoriesmongo

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	utilitiesdb "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/db"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSearchBooksMongo(
	db *mongo.Client,
	collection *mongo.Collection,
	contextLogger contextloggers.ContextLogger,
) repositories.SearchBooksMongo {
	return &SearchBooksMongoImpl{
		db:            db,
		collection:    collection,
		contextLogger: contextLogger,
	}
}

type SearchBooksMongoImpl struct {
	db            *mongo.Client
	collection    *mongo.Collection
	contextLogger contextloggers.ContextLogger
}

func (impl *SearchBooksMongoImpl) SearchBooks(
	ctx context.Context,
	searchBooks *dtos.SearchBooks,
) (aggregates.Books, error) {
	if searchBooks == nil {
		return nil, nil
	}

	pipeline := mongo.Pipeline{}

	// Match stage
	matchStage := impl.buildMatchStage(searchBooks)
	if matchStage != nil {
		pipeline = append(pipeline, matchStage)
	}

	// Sort
	if len(searchBooks.OrderBy) > 0 {
		sortStage := impl.buildSortStage(searchBooks.OrderBy)
		pipeline = append(pipeline, sortStage)
	}

	// Pagination
	if searchBooks.Pagination != nil {
		pipeline = append(pipeline,
			bson.D{{Key: "$skip", Value: (searchBooks.Pagination.Page - 1) * searchBooks.Pagination.PageSize}},
			bson.D{{Key: "$limit", Value: searchBooks.Pagination.PageSize}},
		)
	}

	cursor, err := impl.collection.Aggregate(ctx, pipeline)
	if err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to search books", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var books documents.Books
	if err := cursor.All(ctx, &books); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to decode books", err)
		return nil, err
	}

	return aggregates.BookDocuments(books).ToAggregate(), nil
}

func (impl *SearchBooksMongoImpl) buildMatchStage(search *dtos.SearchBooks) bson.D {
	match := bson.D{}

	// Mobile number filter
	if search.MobileNumber != nil {
		match = append(match, bson.E{Key: "mobile_number", Value: *search.MobileNumber})
	}

	// Table ID filter
	if search.TableID != nil {
		match = append(match, bson.E{Key: "table_id", Value: *search.TableID})
	}

	// Available ID filter
	if search.AvailableID != nil {
		match = append(match, bson.E{Key: "available_id", Value: *search.AvailableID})
	}

	// Availability filter
	if len(search.TableAvailableWeek) > 0 || search.TableAvailableStartTime != nil ||
		search.TableAvailableEndTime != nil {
		availMatch := bson.D{}

		if len(search.TableAvailableWeek) > 0 {
			weekdays := lo.Map(search.TableAvailableWeek, func(w time.Weekday, _ int) int {
				return int(w)
			})
			availMatch = append(availMatch, bson.E{
				Key:   "available.weekday",
				Value: bson.D{{Key: "$in", Value: weekdays}},
			})
		}

		if search.TableAvailableStartTime != nil {
			availMatch = append(availMatch, bson.E{
				Key:   "available.start_time",
				Value: bson.D{{Key: "$gte", Value: *search.TableAvailableStartTime}},
			})
		}

		if search.TableAvailableEndTime != nil {
			availMatch = append(availMatch, bson.E{
				Key:   "available.end_time",
				Value: bson.D{{Key: "$lte", Value: *search.TableAvailableEndTime}},
			})
		}

		match = append(match, bson.E{Key: "$and", Value: availMatch})
	}

	if len(match) == 0 {
		return nil
	}
	return bson.D{{Key: "$match", Value: match}}
}

func (impl *SearchBooksMongoImpl) buildSortStage(orderBy []utilitiesdb.OrderBy) bson.D {
	sort := bson.D{}
	for _, order := range orderBy {
		sort = append(sort, bson.E{Key: order.Field, Value: order.Direction.GetSortDirectionBsonValue()})
	}
	return bson.D{{Key: "$sort", Value: sort}}
}

func NewReadBooksMongo(
	db *mongo.Client,
	collection *mongo.Collection,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadBooksMongo {
	return &ReadBooksMongoImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type ReadBooksMongoImpl struct {
	db            *mongo.Client
	collection    *mongo.Collection
	contextLogger contextloggers.ContextLogger
}

func (impl *ReadBooksMongoImpl) ReadBook(ctx context.Context, id string) (*aggregates.Book, error) {
	if id == "" {
		return nil, nil
	}
	var book documents.Book
	if err := impl.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to read book", err)
		return nil, err
	}
	bookDocument := aggregates.BookDocument(book)
	return bookDocument.ToAggregate(), nil
}

func NewUpdateBooksMongo(
	db *mongo.Client,
	collection *mongo.Collection,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateBooksMongo {
	return &UpdateBooksMongoImpl{
		db:            db,
		collection:    collection,
		contextLogger: contextLogger,
	}
}

type UpdateBooksMongoImpl struct {
	db            *mongo.Client
	collection    *mongo.Collection
	contextLogger contextloggers.ContextLogger
}

func (impl *UpdateBooksMongoImpl) CreateBooks(ctx context.Context, books aggregates.Books) error {
	if len(books) == 0 {
		return nil
	}
	bookDocuments := books.ToDocuments()
	docs := make([]interface{}, len(bookDocuments))
	for i, doc := range bookDocuments {
		docs[i] = doc
	}
	if _, err := impl.collection.InsertMany(ctx, docs); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create books", err)
		return err
	}
	return nil
}

func (impl *UpdateBooksMongoImpl) UpdateBook(ctx context.Context, updateBook *aggregates.Book) error {
	if updateBook == nil || updateBook.ID == "" {
		return nil
	}
	bookDocument := updateBook.ToDocument()
	if _, err := impl.collection.UpdateOne(ctx, bson.M{"_id": updateBook.ID}, bson.M{"$set": *bookDocument}); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update book", err)
		return err
	}
	return nil
}

func (impl *UpdateBooksMongoImpl) DeleteBooks(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if _, err := impl.collection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": ids}}); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete books", err)
		return err
	}
	return nil
}
