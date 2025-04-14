package aggregatesconvert

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/customize_grpc/protobuf"
	"github.com/samber/lo"
)

type RestaurantAggregate aggregates.Restaurant

func (r *RestaurantAggregate) ToProto() *protobuf.Restaurant {
	return &protobuf.Restaurant{
		ID:           r.ID,
		Name:         r.Name,
		Description:  r.Description,
		ActiveStatus: r.ActiveStatus,
		Branches: lo.Map(r.Branches, func(b aggregates.Branch, _ int) *protobuf.Branch {
			ba := BranchAggregate(b)
			return ba.ToProto()
		}),
	}
}

type BranchAggregate aggregates.Branch

func (b *BranchAggregate) ToProto() *protobuf.Branch {
	return &protobuf.Branch{
		ID:          b.ID,
		Name:        b.Name,
		Description: b.Description,
		Address: func() *protobuf.Address {
			if b.Address == nil {
				return nil
			}
			aa := AddressAggregate(*b.Address)
			return aa.ToProto()
		}(),
		PriceRange: func() *protobuf.PriceRange {
			if b.PriceRange == nil {
				return nil
			}
			pr := PriceRangeAggregate(*b.PriceRange)
			return pr.ToProto()
		}(),
		Categories: lo.Map(
			b.BranchCategoryRelations,
			func(c aggregates.BranchCategoryRelation, _ int) *protobuf.Category {
				ca := CategoryAggregate(c.Category)
				return ca.ToProto()
			},
		),
		Availables: lo.Map(b.Availables, func(a aggregates.Available, _ int) *protobuf.Available {
			aa := AvailableAggregate(a)
			return aa.ToProto()
		}),
		Tables: lo.Map(b.Tables, func(t aggregates.Table, _ int) *protobuf.Table {
			ta := TableAggregate(t)
			return ta.ToProto()
		}),
	}
}

type AddressAggregate aggregates.Address

func (a *AddressAggregate) ToProto() *protobuf.Address {
	return &protobuf.Address{
		ID:         a.ID,
		Street:     a.Street,
		City:       a.CityCode.ToCity().ToString(),
		PostalCode: a.PostalCode,
		Country:    a.CountryCode.ToCountry().ToString(),
	}
}

type PriceRangeAggregate aggregates.PriceRange

func (p *PriceRangeAggregate) ToProto() *protobuf.PriceRange {
	return &protobuf.PriceRange{
		ID:       p.ID,
		MinPrice: int32(p.MinPrice),
		MaxPrice: int32(p.MaxPrice),
	}
}

type CategoryAggregate aggregates.Category

func (c *CategoryAggregate) ToProto() *protobuf.Category {
	return &protobuf.Category{
		ID:   c.ID,
		Name: c.CategoryCode.ToCategory().String(),
	}
}

type AvailableAggregate aggregates.Available

func (a *AvailableAggregate) ToProto() *protobuf.Available {
	return &protobuf.Available{
		ID:        a.ID,
		Weekday:   int32(a.Weekday),
		StartTime: a.StartTime,
		EndTime:   a.EndTime,
	}
}

type TableAggregate aggregates.Table

func (t *TableAggregate) ToProto() *protobuf.Table {
	return &protobuf.Table{
		ID:       t.ID,
		Capacity: int32(t.Capacity),
	}
}
