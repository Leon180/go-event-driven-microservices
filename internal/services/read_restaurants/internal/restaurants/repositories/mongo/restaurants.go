package repositoriesmongo

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSearchRestaurantsMongo(
	db *mongo.Client,
	collection *mongo.Collection,
	contextLogger contextloggers.ContextLogger,
) repositories.SearchRestaurantsMongo {
	return &SearchRestaurantsMongoImpl{
		db:            db,
		collection:    collection,
		contextLogger: contextLogger,
	}
}

type SearchRestaurantsMongoImpl struct {
	db            *mongo.Client
	collection    *mongo.Collection
	contextLogger contextloggers.ContextLogger
}

func (impl *SearchRestaurantsMongoImpl) SearchRestaurants(
	ctx context.Context,
	search *dtos.SearchRestaurants,
) (aggregates.Restaurants, error) {
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

	// Execute pipeline
	cursor, err := impl.collection.Aggregate(ctx, pipeline)
	if err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to search restaurants", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var restaurants documents.Restaurants
	if err := cursor.All(ctx, &restaurants); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to search restaurants", err)
		return nil, err
	}

	return aggregates.RestaurantDocuments(restaurants).ToAggregate(), nil
}

func (impl *SearchRestaurantsMongoImpl) buildMatchStage(search *dtos.SearchRestaurants) bson.D {
	match := bson.D{}

	// Name filter
	if search.NameFilter != nil {
		if search.NamePreciseSearch {
			match = append(match, bson.E{Key: "name", Value: *search.NameFilter})
		} else {
			match = append(match, bson.E{Key: "name", Value: primitive.Regex{
				Pattern: *search.NameFilter,
				Options: "i",
			}})
		}
	}

	// Description filter
	if search.DescriptionFilter != nil {
		match = append(match, bson.E{Key: "description", Value: primitive.Regex{
			Pattern: *search.DescriptionFilter,
			Options: "i",
		}})
	}

	// Price range filter
	if search.MinPriceFilter != nil || search.MaxPriceFilter != nil {
		priceMatch := bson.D{}
		if search.MinPriceFilter != nil {
			priceMatch = append(
				priceMatch,
				bson.E{
					Key:   "branches.price_range.min_price",
					Value: bson.D{{Key: "$gte", Value: *search.MinPriceFilter}},
				},
			)
		}
		if search.MaxPriceFilter != nil {
			priceMatch = append(
				priceMatch,
				bson.E{
					Key:   "branches.price_range.max_price",
					Value: bson.D{{Key: "$lte", Value: *search.MaxPriceFilter}},
				},
			)
		}
		match = append(match, bson.E{Key: "$and", Value: priceMatch})
	}

	// City filter
	if len(search.CityFilter) > 0 {
		cityCodes := lo.Map(search.CityFilter, func(city enums.City, _ int) int {
			return int(city.ToCityCode())
		})
		match = append(match, bson.E{Key: "branches.address.city_code", Value: bson.D{{Key: "$in", Value: cityCodes}}})
	}

	// Country filter
	if len(search.CountryFilter) > 0 {
		countryCodes := lo.Map(search.CountryFilter, func(country enums.Country, _ int) int {
			return int(country.ToCountryCode())
		})
		match = append(
			match,
			bson.E{Key: "branches.address.country_code", Value: bson.D{{Key: "$in", Value: countryCodes}}},
		)
	}

	// Category filter
	if len(search.CategoryFilter) > 0 {
		categoryCodes := lo.Map(search.CategoryFilter, func(category enums.Category, _ int) int {
			return int(category.ToCategoryCode())
		})
		match = append(
			match,
			bson.E{Key: "branches.categories.category_code", Value: bson.D{{Key: "$in", Value: categoryCodes}}},
		)
	}

	// Availability filter
	if len(search.TableAvailableWeek) > 0 || search.TableAvailableStartTime != nil ||
		search.TableAvailableEndTime != nil {
		availMatch := bson.D{}

		if len(search.TableAvailableWeek) > 0 {
			weekdays := lo.Map(search.TableAvailableWeek, func(w time.Weekday, _ int) int {
				return int(w)
			})
			availMatch = append(
				availMatch,
				bson.E{Key: "branches.availables.weekday", Value: bson.D{{Key: "$in", Value: weekdays}}},
			)
		}

		if search.TableAvailableStartTime != nil {
			availMatch = append(
				availMatch,
				bson.E{
					Key:   "branches.availables.start_time",
					Value: bson.D{{Key: "$gte", Value: *search.TableAvailableStartTime}},
				},
			)
		}

		if search.TableAvailableEndTime != nil {
			availMatch = append(
				availMatch,
				bson.E{
					Key:   "branches.availables.end_time",
					Value: bson.D{{Key: "$lte", Value: *search.TableAvailableEndTime}},
				},
			)
		}

		match = append(match, bson.E{Key: "$and", Value: availMatch})
	}

	if len(match) == 0 {
		return nil
	}
	return bson.D{{Key: "$match", Value: match}}
}

func (impl *SearchRestaurantsMongoImpl) buildSortStage(orderBy []customizegorm.OrderBy) bson.D {
	sort := bson.D{}
	for _, order := range orderBy {
		sort = append(sort, bson.E{Key: order.Field, Value: order.Direction.GetSortDirectionBsonValue()})
	}
	return bson.D{{Key: "$sort", Value: sort}}
}

func NewReadRestaurantsMongo(
	db *mongo.Client,
	collection *mongo.Collection,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadRestaurantsMongo {
	return &ReadRestaurantsMongoImpl{
		db:            db,
		collection:    collection,
		contextLogger: contextLogger,
	}
}

type ReadRestaurantsMongoImpl struct {
	db            *mongo.Client
	collection    *mongo.Collection
	contextLogger contextloggers.ContextLogger
}

func (impl *ReadRestaurantsMongoImpl) ReadRestaurant(
	ctx context.Context,
	id string,
) (*aggregates.Restaurant, error) {
	if id == "" {
		return nil, nil
	}
	var restaurant documents.Restaurant
	if err := impl.collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&restaurant); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).
			Error("failed to read restaurant full info", err)
		return nil, err
	}
	aggregate := aggregates.RestaurantDocument(restaurant)
	return aggregate.ToAggregate(), nil
}

// Update Restaurants Impl
func NewUpdateRestaurantsMongo(
	db *mongo.Client,
	collection *mongo.Collection,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateRestaurantsMongo {
	return &UpdateRestaurantsMongoImpl{
		db:            db,
		collection:    collection,
		contextLogger: contextLogger,
	}
}

type UpdateRestaurantsMongoImpl struct {
	db            *mongo.Client
	collection    *mongo.Collection
	contextLogger contextloggers.ContextLogger
}

func (impl *UpdateRestaurantsMongoImpl) CreateRestaurants(
	ctx context.Context,
	restaurants aggregates.Restaurants,
) error {
	if len(restaurants) == 0 {
		return nil
	}
	restaurantsDocs := restaurants.ToDocuments()
	docs := make([]interface{}, len(restaurantsDocs))
	for i, doc := range restaurantsDocs {
		docs[i] = doc
	}
	if _, err := impl.collection.InsertMany(ctx, docs); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create restaurants", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsMongoImpl) UpdateRestaurant(
	ctx context.Context,
	restaurant *aggregates.Restaurant,
) error {
	if restaurant == nil || restaurant.ID == "" {
		return nil
	}
	doc := restaurant.ToDocument()
	if _, err := impl.collection.UpdateOne(ctx, bson.D{{Key: "_id", Value: restaurant.ID}}, bson.D{{Key: "$set", Value: *doc}}); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update restaurant", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsMongoImpl) DeleteRestaurants(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if _, err := impl.collection.DeleteMany(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}); err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete restaurants", err)
		return err
	}
	return nil
}
