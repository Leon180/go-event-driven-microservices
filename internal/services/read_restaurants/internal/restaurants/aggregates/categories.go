package aggregates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
	"github.com/samber/lo"
)

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

func (c *Category) ToDocument() *documents.Category {
	return &documents.Category{
		ID:           c.ID,
		CategoryCode: c.CategoryCode,
		CommonCQRSHistoryModel: documents.CommonCQRSHistoryModel{
			ActiveStatus: c.ActiveStatus,
			CreatedAt:    c.CreatedAt,
			UpdatedAt:    c.UpdatedAt,
		},
	}
}

type Categories []Category

func (c Categories) ToDTO() []dtos.Category {
	return lo.Map(c, func(category Category, _ int) dtos.Category {
		return *category.ToDTO()
	})
}

func (c Categories) ToDocuments() []documents.Category {
	return lo.Map(c, func(category Category, _ int) documents.Category {
		return *category.ToDocument()
	})
}

type CategoryDocument documents.Category

func (c *CategoryDocument) ToAggregate() *Category {
	return &Category{
		ID:           c.ID,
		CategoryCode: c.CategoryCode,
		CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
			ActiveStatus: c.ActiveStatus,
			CreatedAt:    c.CreatedAt,
			UpdatedAt:    c.UpdatedAt,
		},
	}
}

type CategoryDocuments documents.Categories

func (c CategoryDocuments) ToAggregate() Categories {
	return lo.Map(c, func(category documents.Category, _ int) Category {
		categoryDocument := CategoryDocument(category)
		return *categoryDocument.ToAggregate()
	})
}
