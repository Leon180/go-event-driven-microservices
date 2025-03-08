package documents

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
	ActiveStatus bool      `bson:"active_status"`
	CreatedAt    time.Time `bson:"created_at"`
	CreatedBy    string    `bson:"created_by"`
	UpdatedAt    time.Time `bson:"updated_at"`
	UpdatedBy    string    `bson:"updated_by"`
}

func (c *CommonCQRSHistoryModel) IsActive() bool {
	return c.ActiveStatus
}
