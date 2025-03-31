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
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/mongodb"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSearchFailedMessagesMongo(
	db *mongo.Client,
	collections *mongodb.Collections,
	contextLogger contextloggers.ContextLogger,
) repositories.SearchFailedMessagesMongo {
	return &searchFailedMessagesMongoImpl{
		db:            db,
		collections:   collections,
		contextLogger: contextLogger,
	}
}

type searchFailedMessagesMongoImpl struct {
	db            *mongo.Client
	collections   *mongodb.Collections
	contextLogger contextloggers.ContextLogger
}

func (impl *searchFailedMessagesMongoImpl) SearchFailedMessages(
	ctx context.Context,
	search *dtos.SearchFailedMessages,
) (aggregates.FailedMessages, error) {
	if search == nil {
		return nil, nil
	}

	pipeline := mongo.Pipeline{}

	matchStage := impl.buildMatchStage(search)
	if matchStage != nil {
		pipeline = append(pipeline, matchStage)
	}

	// Sort
	if len(search.OrderBy) > 0 {
		sortStage := impl.buildSortStage(search.OrderBy)
		pipeline = append(pipeline, sortStage)
	}

	// Pagination
	if search.Pagination != nil {
		pipeline = append(pipeline,
			bson.D{{Key: "$skip", Value: (search.Pagination.Page - 1) * search.Pagination.PageSize}},
			bson.D{{Key: "$limit", Value: search.Pagination.PageSize}},
		)
	}

	collection := impl.collections.Restaurant

	// Execute pipeline
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to search failed messages", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var failedMessages documents.FailedMessages
	if err := cursor.All(ctx, &failedMessages); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to search failed messages", err)
		return nil, err
	}

	return aggregates.FailedMessageDocuments(failedMessages).ToAggregate(), nil
}

func (impl *searchFailedMessagesMongoImpl) buildMatchStage(search *dtos.SearchFailedMessages) bson.D {
	match := bson.D{}

	// Name filter
	if search.Exchange != nil {
		match = append(match, bson.E{Key: "exchange", Value: primitive.Regex{
			Pattern: *search.Exchange,
			Options: "i",
		}})
	}

	// Description filter
	if search.RoutingKey != nil {
		match = append(match, bson.E{Key: "routing_key", Value: primitive.Regex{
			Pattern: *search.RoutingKey,
			Options: "i",
		}})
	}

	// Price range filter
	if search.Queue != nil {
		match = append(match, bson.E{Key: "queue", Value: primitive.Regex{
			Pattern: *search.Queue,
			Options: "i",
		}})
	}

	if search.MessageType != nil {
		match = append(match, bson.E{Key: "message_type", Value: primitive.Regex{
			Pattern: *search.MessageType,
			Options: "i",
		}})
	}

	if search.ContentType != nil {
		match = append(match, bson.E{Key: "content_type", Value: primitive.Regex{
			Pattern: *search.ContentType,
			Options: "i",
		}})
	}

	if search.DeliveryMode != nil {
		match = append(match, bson.E{Key: "delivery_mode", Value: *search.DeliveryMode})
	}

	if search.TimeStart != nil || search.TimeEnd != nil {
		timeMatch := bson.D{}
		if search.TimeStart != nil {
			st, _ := time.Parse(time.DateOnly, *search.TimeStart)
			timeMatch = append(
				timeMatch,
				bson.E{
					Key:   "time_stamp",
					Value: bson.D{{Key: "$gte", Value: st}},
				},
			)
		}
		if search.TimeEnd != nil {
			et, _ := time.Parse(time.DateOnly, *search.TimeEnd)
			timeMatch = append(
				timeMatch,
				bson.E{
					Key:   "time_stamp",
					Value: bson.D{{Key: "$lte", Value: et}},
				},
			)
		}
		match = append(match, bson.E{Key: "$and", Value: timeMatch})
	}
	if len(match) == 0 {
		return nil
	}
	return bson.D{{Key: "$match", Value: match}}
}

func (impl *searchFailedMessagesMongoImpl) buildSortStage(orderBy []utilitiesdb.OrderBy) bson.D {
	sort := bson.D{}
	for _, order := range orderBy {
		sort = append(sort, bson.E{Key: order.Field, Value: order.Direction.GetSortDirectionBsonValue()})
	}
	return bson.D{{Key: "$sort", Value: sort}}
}

func NewReadFailedMessageMongo(
	db *mongo.Client,
	collections *mongodb.Collections,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadFailedMessageMongo {
	return &readFailedMessageMongoImpl{
		db:            db,
		collections:   collections,
		contextLogger: contextLogger,
	}
}

type readFailedMessageMongoImpl struct {
	db            *mongo.Client
	collections   *mongodb.Collections
	contextLogger contextloggers.ContextLogger
}

func (impl *readFailedMessageMongoImpl) ReadFailedMessage(
	ctx context.Context,
	messageID string,
) (*aggregates.FailedMessage, error) {
	if messageID == "" {
		return nil, nil
	}
	collection := impl.collections.FailedMessage
	var failedMessage documents.FailedMessage
	if err := collection.FindOne(ctx, bson.D{{Key: "message_id", Value: messageID}}).Decode(&failedMessage); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to read failed message", err)
		return nil, err
	}
	aggregate := aggregates.FailedMessageDocument(failedMessage)
	return aggregate.ToAggregate(), nil
}

func NewCreateFailedMessageMongo(
	db *mongo.Client,
	collections *mongodb.Collections,
	contextLogger contextloggers.ContextLogger,
) repositories.CreateFailedMessageMongo {
	return &createFailedMessageMongoImpl{
		db:            db,
		collections:   collections,
		contextLogger: contextLogger,
	}
}

type createFailedMessageMongoImpl struct {
	db            *mongo.Client
	collections   *mongodb.Collections
	contextLogger contextloggers.ContextLogger
}

func (impl *createFailedMessageMongoImpl) Create(
	ctx context.Context,
	failedMessage *aggregates.FailedMessage,
) error {
	if failedMessage == nil {
		return nil
	}
	doc := failedMessage.ToDocument()
	collection := impl.collections.FailedMessage
	if _, err := collection.InsertOne(ctx, doc); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create failed message", err)
		return err
	}
	return nil
}
