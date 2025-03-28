package dtos

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	"gorm.io/gorm"
)

type CommonHistoryModelWithUpdate struct {
	CommonHistoryModel
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type CommonHistoryModel struct {
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy string         `json:"created_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	DeletedBy string         `json:"deleted_by"`
}

type CommonCQRSHistoryModel struct {
	ActiveStatus bool      `json:"activeStatus"`
	CreatedAt    time.Time `json:"createdAt"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    string    `json:"updatedBy"`
}

type CommonCQRSHistoryModelEntity entities.CommonCQRSHistoryModel

func (c *CommonCQRSHistoryModelEntity) ToDTO() *CommonCQRSHistoryModel {
	return &CommonCQRSHistoryModel{
		ActiveStatus: c.ActiveStatus,
		CreatedAt:    c.CreatedAt,
		CreatedBy:    c.CreatedBy,
		UpdatedAt:    c.UpdatedAt,
		UpdatedBy:    c.UpdatedBy,
	}
}
