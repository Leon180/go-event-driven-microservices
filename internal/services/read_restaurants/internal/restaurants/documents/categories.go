package documents

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

type Category struct {
	ID           string             `bson:"_id"`
	CategoryCode enums.CategoryCode `bson:"category_code"`
	CommonCQRSHistoryModel
}

type Categories []Category
