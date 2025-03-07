package entities

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

type Category struct {
	ID           string             `gorm:"primaryKey;type:varchar(255)" comment:"ID"`
	CategoryCode enums.CategoryCode `gorm:"not null;type:int" comment:"Category Code"`
	CommonCQRSHistoryModel
}

type Categories []Category
