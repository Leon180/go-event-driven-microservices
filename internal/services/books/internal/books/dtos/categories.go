package dtos

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	"github.com/samber/lo"
)

type Category struct {
	ID       string         `json:"id"`
	Category enums.Category `json:"category"`
	CommonCQRSHistoryModel
}

type Categories []Category

type CategoryEntity entities.Category

func (c *CategoryEntity) ToDTO() *Category {
	return &Category{
		ID:       c.ID,
		Category: c.CategoryCode.ToCategory(),
		CommonCQRSHistoryModel: CommonCQRSHistoryModel{
			ActiveStatus: c.ActiveStatus,
			CreatedAt:    c.CreatedAt,
			CreatedBy:    c.CreatedBy,
			UpdatedAt:    c.UpdatedAt,
			UpdatedBy:    c.UpdatedBy,
		},
	}
}

type CategoriesEntity []entities.Category

func (c CategoriesEntity) ToDTO() Categories {
	return lo.Map(c, func(category entities.Category, _ int) Category {
		categoryEntity := CategoryEntity(category)
		return *categoryEntity.ToDTO()
	})
}
