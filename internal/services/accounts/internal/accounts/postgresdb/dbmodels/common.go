package dbmodels

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/dtos"
	"gorm.io/gorm"
)

type CommonHistoryModelWithUpdate struct {
	CommonHistoryModel
	UpdatedAt time.Time
	UpdatedBy string
}

func (c *CommonHistoryModelWithUpdate) ToDTOs() dtos.CommonHistoryModelWithUpdate {
	return dtos.CommonHistoryModelWithUpdate{
		CommonHistoryModel: c.CommonHistoryModel.ToDTOs(),
		UpdatedAt:          c.UpdatedAt,
		UpdatedBy:          c.UpdatedBy,
	}
}

type CommonHistoryModel struct {
	CreatedAt time.Time
	CreatedBy string
	DeletedAt gorm.DeletedAt
	DeletedBy string
}

func (c *CommonHistoryModel) ToDTOs() dtos.CommonHistoryModel {
	return dtos.CommonHistoryModel{
		CreatedAt: c.CreatedAt,
		CreatedBy: c.CreatedBy,
		DeletedAt: c.DeletedAt,
		DeletedBy: c.DeletedBy,
	}
}
