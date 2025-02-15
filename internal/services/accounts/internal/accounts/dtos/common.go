package dtos

import (
	"time"

	"gorm.io/gorm"
)

type CommonHistoryModelWithUpdate struct {
	CommonHistoryModel
	UpdatedAt time.Time
	UpdatedBy string
}

type CommonHistoryModel struct {
	CreatedAt time.Time
	CreatedBy string
	DeletedAt gorm.DeletedAt
	DeletedBy string
}
