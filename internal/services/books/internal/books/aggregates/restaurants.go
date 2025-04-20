package aggregates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	"github.com/samber/lo"
)

type Restaurants []Restaurant

func (r Restaurants) ToDTO() []dtos.Restaurant {
	return lo.Map(r, func(restaurant Restaurant, _ int) dtos.Restaurant {
		return *restaurant.ToDTO()
	})
}

type Restaurant struct {
	entities.Restaurant
	editTypeCode enums.EditTypeCode         `gorm:"-"`
	update       *entities.UpdateRestaurant `gorm:"-"`
	Branches     []Branch                   `gorm:"foreignKey:RestaurantID;references:ID" comment:"Branches"`
}

func (r *Restaurant) TableName() string {
	return "restaurant"
}

func (r *Restaurant) ToDTO() *dtos.Restaurant {
	return &dtos.Restaurant{
		ID:           &r.ID,
		Name:         r.Name,
		Description:  r.Description,
		ActiveStatus: r.ActiveStatus,
		Branches: lo.Map(r.Branches, func(b Branch, _ int) dtos.Branch {
			return *b.ToDTO()
		}),
	}
}

type Branch struct {
	entities.Branch
	editTypeCode            enums.EditTypeCode       `gorm:"-"`
	update                  *entities.UpdateBranch   `gorm:"-"`
	Address                 *Address                 `gorm:"foreignKey:BranchID;references:ID" comment:"Address"`
	PriceRange              *PriceRange              `gorm:"foreignKey:BranchID;references:ID" comment:"Price Range"`
	BranchCategoryRelations []BranchCategoryRelation `gorm:"foreignKey:BranchID;references:ID" comment:"Branch Category Relation"`
	Tables                  []Table                  `gorm:"foreignKey:BranchID;references:ID" comment:"Tables"`
	Availables              []Available              `gorm:"foreignKey:BranchID;references:ID" comment:"Availables"`
}

func (b *Branch) TableName() string {
	return "branch"
}

func (b *Branch) ToDTO() *dtos.Branch {
	return &dtos.Branch{
		ID:          &b.ID,
		Name:        b.Name,
		Description: b.Description,
		Address:     b.Address.ToDTO(),
		PriceRange:  b.PriceRange.ToDTO(),
		Categories: lo.Map(b.BranchCategoryRelations, func(bcr BranchCategoryRelation, _ int) dtos.Category {
			return *bcr.Category.ToDTO()
		}),
		Tables: lo.Map(b.Tables, func(t Table, _ int) dtos.Table {
			return *t.ToDTO()
		}),
		Availables: lo.Map(b.Availables, func(a Available, _ int) dtos.Available {
			return *a.ToDTO()
		}),
	}
}

type Address struct {
	entities.Address
	editTypeCode enums.EditTypeCode      `gorm:"-"`
	update       *entities.UpdateAddress `gorm:"-"`
}

func (a *Address) TableName() string {
	return "address"
}

func (a *Address) ToDTO() *dtos.Address {
	return &dtos.Address{
		ID:         &a.ID,
		Street:     a.Street,
		City:       a.CityCode.ToCity(),
		PostalCode: a.PostalCode,
		Country:    a.CountryCode.ToCountry(),
	}
}

type PriceRange struct {
	entities.PriceRange
	editTypeCode enums.EditTypeCode         `gorm:"-"`
	update       *entities.UpdatePriceRange `gorm:"-"`
}

func (p *PriceRange) TableName() string {
	return "price_range"
}

func (p *PriceRange) ToDTO() *dtos.PriceRange {
	return &dtos.PriceRange{
		ID:       &p.ID,
		MinPrice: p.MinPrice,
		MaxPrice: p.MaxPrice,
	}
}

type BranchCategoryRelation struct {
	entities.BranchCategoryRelation
	editTypeCode enums.EditTypeCode                     `gorm:"-"`
	update       *entities.UpdateBranchCategoryRelation `gorm:"-"`
	Category     Category                               `gorm:"foreignKey:ID;references:CategoryID" comment:"Category"`
}

func (bcr *BranchCategoryRelation) TableName() string {
	return "branch_category_relation"
}

type Table struct {
	entities.Table
	editTypeCode enums.EditTypeCode    `gorm:"-"`
	update       *entities.UpdateTable `gorm:"-"`
}

func (t *Table) TableName() string {
	return "table"
}

func (t *Table) ToDTO() *dtos.Table {
	return &dtos.Table{
		ID:       &t.ID,
		Capacity: t.Capacity,
	}
}

type Available struct {
	entities.Available
	editTypeCode enums.EditTypeCode        `gorm:"-"`
	update       *entities.UpdateAvailable `gorm:"-"`
}

func (a *Available) TableName() string {
	return "available"
}

func (a *Available) ToDTO() *dtos.Available {
	return &dtos.Available{
		ID:        &a.ID,
		Weekday:   a.Weekday,
		StartTime: a.StartTime,
		EndTime:   a.EndTime,
	}
}
