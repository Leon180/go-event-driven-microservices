package aggregates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	"github.com/samber/lo"
)

type Category entities.Category

func (c *Category) ToEntity() *entities.Category {
	e := entities.Category(*c)
	return &e
}

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

type Categories []Category

func (c Categories) ToEntities() []entities.Category {
	return lo.Map(c, func(category Category, _ int) entities.Category {
		return *category.ToEntity()
	})
}

func (c Categories) ToDTO() []dtos.Category {
	return lo.Map(c, func(category Category, _ int) dtos.Category {
		return *category.ToDTO()
	})
}

func (c Categories) GetUpdateCategories(original Categories) UpdateCategories {
	createCategories := []Category{}
	deleteCategories := []Category{}
	m := map[enums.CategoryCode]struct{}{}
	d := map[enums.CategoryCode]struct{}{}
	for _, category := range original {
		m[category.CategoryCode] = struct{}{}
	}
	for _, category := range c {
		if _, ok := m[category.CategoryCode]; !ok {
			createCategories = append(createCategories, category)
		}
		d[category.CategoryCode] = struct{}{}
	}
	for _, category := range original {
		if _, ok := d[category.CategoryCode]; !ok {
			deleteCategories = append(deleteCategories, category)
		}
	}
	return UpdateCategories{
		CreateCategories: createCategories,
		DeleteCategories: deleteCategories,
	}
}

type UpdateCategories struct {
	CreateCategories []Category
	DeleteCategories []Category
}
