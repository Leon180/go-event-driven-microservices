package entities

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

type CommonCQRSHistoryModel struct {
	ActiveStatus bool `gorm:"not null;type:boolean;default:true" comment:"Active Status"`
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
}

func (c *CommonCQRSHistoryModel) IsActive() bool {
	return c.ActiveStatus
}
