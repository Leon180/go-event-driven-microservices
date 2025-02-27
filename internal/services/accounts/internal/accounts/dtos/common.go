package dtos

import (
	"time"

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
