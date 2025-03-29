package aggregates

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
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
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

func appendEntityByType[T any, U any](
	create *[]T,
	update *[]U,
	delete *[]T,
	entity T,
	updateEntity *U,
	editTypeCode enums.EditTypeCode,
) {
	switch editTypeCode {
	case enums.EditTypeCodeCreate, enums.EditTypeCodeDelete:
		appendEntityByTypeCD(create, delete, entity, editTypeCode)
	case enums.EditTypeCodeUpdate:
		appendEntityByTypeU(update, updateEntity)
	}
}

func appendEntityByTypeCD[T any](create *[]T, delete *[]T, entity T, editTypeCode enums.EditTypeCode) {
	switch editTypeCode {
	case enums.EditTypeCodeCreate:
		*create = append(*create, entity)
	case enums.EditTypeCodeDelete:
		*delete = append(*delete, entity)
	}
}

func appendEntityByTypeU[U any](update *[]U, updateEntity *U) {
	*update = append(*update, *updateEntity)
}
