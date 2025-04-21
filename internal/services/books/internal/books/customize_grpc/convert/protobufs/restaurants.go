package protobufsconvert

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/customize_grpc/protobuf"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	"github.com/samber/lo"
)

func RestaurantProtobufToAggregate(r *protobuf.Restaurant) aggregates.Restaurant {
	return aggregates.Restaurant{
		Restaurant: entities.Restaurant{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
		},
		Branches: lo.Map(r.Branches, func(b *protobuf.Branch, _ int) aggregates.Branch {
			return aggregates.Branch{
				Branch: entities.Branch{
					ID:           b.ID,
					RestaurantID: r.ID,
					Name:         b.Name,
					Description:  b.Description,
				},
				Address: func() *aggregates.Address {
					if b.Address == nil {
						return nil
					}
					return &aggregates.Address{
						Address: entities.Address{
							ID:          b.Address.ID,
							BranchID:    b.ID,
							Street:      b.Address.Street,
							CityCode:    enums.City(b.Address.City).ToCityCode(),
							PostalCode:  b.Address.PostalCode,
							CountryCode: enums.Country(b.Address.Country).ToCountryCode(),
						},
					}
				}(),
				PriceRange: func() *aggregates.PriceRange {
					if b.PriceRange == nil {
						return nil
					}
					return &aggregates.PriceRange{
						PriceRange: entities.PriceRange{
							ID:       b.PriceRange.ID,
							BranchID: b.ID,
							MinPrice: int(b.PriceRange.MinPrice),
							MaxPrice: int(b.PriceRange.MaxPrice),
						},
					}
				}(),
				BranchCategoryRelations: lo.Map(b.Categories, func(c *protobuf.Category, _ int) aggregates.BranchCategoryRelation {
					return aggregates.BranchCategoryRelation{
						BranchCategoryRelation: entities.BranchCategoryRelation{
							BranchID:   b.ID,
							CategoryID: c.ID,
						},
						Category: func() aggregates.Category {
							if c == nil {
								return aggregates.Category{}
							}
							return aggregates.Category{
								ID:           c.ID,
								CategoryCode: enums.Category(c.Name).ToCategoryCode(),
							}
						}(),
					}
				}),
				Tables: lo.Map(b.Tables, func(t *protobuf.Table, _ int) aggregates.Table {
					return aggregates.Table{
						Table: entities.Table{
							ID:       t.ID,
							BranchID: b.ID,
							Capacity: int(t.Capacity),
						},
					}
				}),
				Availables: lo.Map(b.Availables, func(a *protobuf.Available, _ int) aggregates.Available {
					return aggregates.Available{
						Available: entities.Available{
							ID:        a.ID,
							BranchID:  b.ID,
							Weekday:   time.Weekday(a.Weekday),
							StartTime: a.StartTime,
							EndTime:   a.EndTime,
						},
					}
				}),
			}
		}),
	}
}
