package aggregates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
	"github.com/samber/lo"
)

type Restaurants []Restaurant

func (r Restaurants) ToDTO() []dtos.Restaurant {
	return lo.Map(r, func(restaurant Restaurant, _ int) dtos.Restaurant {
		return *restaurant.ToDTO()
	})
}

func (r Restaurants) ToDocuments() []documents.Restaurant {
	return lo.Map(r, func(restaurant Restaurant, _ int) documents.Restaurant {
		return *restaurant.ToDocument()
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
		ID:          &r.ID,
		Name:        r.Name,
		Description: r.Description,
		Branches: lo.Map(r.Branches, func(b Branch, _ int) dtos.Branch {
			return *b.ToDTO()
		}),
	}
}

func (r *Restaurant) ToDocument() *documents.Restaurant {
	return &documents.Restaurant{
		ID:                     r.ID,
		Name:                   r.Name,
		Description:            r.Description,
		CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(r.CommonCQRSHistoryModel).ToDocument(),
		Branches: lo.Map(r.Branches, func(b Branch, _ int) documents.Branch {
			return documents.Branch{
				ID:                     b.ID,
				RestaurantID:           r.ID,
				Name:                   b.Name,
				Description:            b.Description,
				CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(b.CommonCQRSHistoryModel).ToDocument(),
				Address: func() documents.Address {
					if b.Address == nil {
						return documents.Address{}
					}
					return documents.Address{
						ID:          b.Address.ID,
						BranchID:    b.ID,
						Street:      b.Address.Street,
						CityCode:    b.Address.CityCode,
						PostalCode:  b.Address.PostalCode,
						CountryCode: b.Address.CountryCode,
						CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(
							b.Address.CommonCQRSHistoryModel,
						).ToDocument(),
					}
				}(),
				PriceRange: func() documents.PriceRange {
					if b.PriceRange == nil {
						return documents.PriceRange{}
					}
					return documents.PriceRange{
						ID:       b.PriceRange.ID,
						BranchID: b.ID,
						MinPrice: b.PriceRange.MinPrice,
						MaxPrice: b.PriceRange.MaxPrice,
						CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(
							b.PriceRange.CommonCQRSHistoryModel,
						).ToDocument(),
					}
				}(),
				BranchCategories: lo.Map(
					b.BranchCategoryRelations,
					func(bcr BranchCategoryRelation, _ int) documents.BranchCategoryRelation {
						return documents.BranchCategoryRelation{
							ID:         bcr.ID,
							BranchID:   b.ID,
							CategoryID: bcr.CategoryID,
							CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(
								bcr.CommonCQRSHistoryModel,
							).ToDocument(),
							Category: documents.Category{
								ID:           bcr.Category.ID,
								CategoryCode: bcr.Category.CategoryCode,
								CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(
									bcr.Category.CommonCQRSHistoryModel,
								).ToDocument(),
							},
						}
					},
				),
				Tables: lo.Map(b.Tables, func(t Table, _ int) documents.Table {
					return documents.Table{
						ID:                     t.ID,
						BranchID:               b.ID,
						Capacity:               t.Capacity,
						CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(t.CommonCQRSHistoryModel).ToDocument(),
					}
				}),
				Availables: lo.Map(b.Availables, func(a Available, _ int) documents.Available {
					return documents.Available{
						ID:                     a.ID,
						BranchID:               b.ID,
						Weekday:                a.Weekday,
						StartTime:              a.StartTime,
						EndTime:                a.EndTime,
						CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(a.CommonCQRSHistoryModel).ToDocument(),
					}
				}),
			}
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

type Category entities.Category

func (c *Category) TableName() string {
	return "category"
}

func (c *Category) ToDTO() *dtos.Category {
	return &dtos.Category{
		ID:       c.ID,
		Category: c.CategoryCode.ToCategory(),
		CommonCQRSHistoryModel: dtos.CommonCQRSHistoryModel{
			ActiveStatus: c.ActiveStatus,
			CreatedAt:    c.CreatedAt,
			UpdatedAt:    c.UpdatedAt,
		},
	}
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

type RestaurantDocuments documents.Restaurants

func (r RestaurantDocuments) ToAggregate() Restaurants {
	return lo.Map(r, func(r documents.Restaurant, _ int) Restaurant {
		restaurant := RestaurantDocument(r)
		return *restaurant.ToAggregate()
	})
}

type RestaurantDocument documents.Restaurant

func (r *RestaurantDocument) ToAggregate() *Restaurant {
	if r == nil {
		return nil
	}

	return &Restaurant{
		Restaurant: entities.Restaurant{
			ID:                     r.ID,
			Name:                   r.Name,
			Description:            r.Description,
			CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(r.CommonCQRSHistoryModel).ToEntity(),
		},
		Branches: lo.Map(r.Branches, func(b documents.Branch, _ int) Branch {
			return Branch{
				Branch: entities.Branch{
					ID:                     b.ID,
					RestaurantID:           r.ID,
					Name:                   b.Name,
					Description:            b.Description,
					CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(b.CommonCQRSHistoryModel).ToEntity(),
				},
				Address: &Address{
					Address: entities.Address{
						ID:          b.Address.ID,
						BranchID:    b.ID,
						Street:      b.Address.Street,
						CityCode:    b.Address.CityCode,
						PostalCode:  b.Address.PostalCode,
						CountryCode: b.Address.CountryCode,
						CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(
							b.Address.CommonCQRSHistoryModel,
						).ToEntity(),
					},
				},
				PriceRange: &PriceRange{
					PriceRange: entities.PriceRange{
						ID:       b.PriceRange.ID,
						BranchID: b.ID,
						MinPrice: b.PriceRange.MinPrice,
						MaxPrice: b.PriceRange.MaxPrice,
						CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(
							b.PriceRange.CommonCQRSHistoryModel,
						).ToEntity(),
					},
				},
				BranchCategoryRelations: lo.Map(
					b.BranchCategories,
					func(bcr documents.BranchCategoryRelation, _ int) BranchCategoryRelation {
						return BranchCategoryRelation{
							BranchCategoryRelation: entities.BranchCategoryRelation{
								ID:         bcr.ID,
								BranchID:   b.ID,
								CategoryID: bcr.CategoryID,
								CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(
									bcr.CommonCQRSHistoryModel,
								).ToEntity(),
							},
							Category: Category{
								ID:           bcr.Category.ID,
								CategoryCode: bcr.Category.CategoryCode,
								CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(
									bcr.Category.CommonCQRSHistoryModel,
								).ToEntity(),
							},
						}
					},
				),
				Tables: lo.Map(b.Tables, func(t documents.Table, _ int) Table {
					return Table{
						Table: entities.Table{
							ID:                     t.ID,
							BranchID:               b.ID,
							Capacity:               t.Capacity,
							CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(t.CommonCQRSHistoryModel).ToEntity(),
						},
					}
				}),
				Availables: lo.Map(b.Availables, func(a documents.Available, _ int) Available {
					return Available{
						Available: entities.Available{
							ID:                     a.ID,
							BranchID:               b.ID,
							Weekday:                a.Weekday,
							StartTime:              a.StartTime,
							EndTime:                a.EndTime,
							CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(a.CommonCQRSHistoryModel).ToEntity(),
						},
					}
				}),
			}
		}),
	}
}

type CommonCQRSHistoryModelDocument documents.CommonCQRSHistoryModel

func (c CommonCQRSHistoryModelDocument) ToEntity() entities.CommonCQRSHistoryModel {
	return entities.CommonCQRSHistoryModel{
		ActiveStatus: c.ActiveStatus,
		CreatedAt:    c.CreatedAt,
		CreatedBy:    c.CreatedBy,
		UpdatedAt:    c.UpdatedAt,
		UpdatedBy:    c.UpdatedBy,
	}
}

func (c CommonCQRSHistoryModelDocument) ToDocument() documents.CommonCQRSHistoryModel {
	return documents.CommonCQRSHistoryModel{
		ActiveStatus: c.ActiveStatus,
		CreatedAt:    c.CreatedAt,
		CreatedBy:    c.CreatedBy,
		UpdatedAt:    c.UpdatedAt,
		UpdatedBy:    c.UpdatedBy,
	}
}
