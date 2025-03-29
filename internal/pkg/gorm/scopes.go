package postgresgorm

import (
	utilitiesdb "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/db"
	"gorm.io/gorm"
)

func ApplyPagination(pagination *utilitiesdb.Pagination) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination != nil {
			return db.Limit(pagination.PageSize).
				Offset((pagination.Page - 1) * pagination.PageSize)
		}
		return db
	}
}

func ApplyOrdering(orderBy []utilitiesdb.OrderBy) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, ob := range orderBy {
			db = db.Order(ob.ToSort())
		}
		return db
	}
}
