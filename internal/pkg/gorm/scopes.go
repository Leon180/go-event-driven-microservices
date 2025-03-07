package postgresgorm

import (
	"gorm.io/gorm"
)

func ApplyPagination(pagination *Pagination) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination != nil {
			return db.Limit(pagination.PageSize).
				Offset((pagination.Page - 1) * pagination.PageSize)
		}
		return db
	}
}

func ApplyOrdering(orderBy []OrderBy) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, ob := range orderBy {
			db = db.Order(ob.ToSort())
		}
		return db
	}
}
