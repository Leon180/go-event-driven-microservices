package aggregates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
	"github.com/samber/lo"
)

type Books []Book

func (b Books) ToDTO() []dtos.Book {
	return lo.Map(b, func(book Book, _ int) dtos.Book {
		return *book.ToDTO()
	})
}

type Book struct {
	entities.Book

	// Relations
	Table     *Table     `gorm:"foreignKey:ID;references:TableID" comment:"Table"`
	Available *Available `gorm:"foreignKey:ID;references:AvailableID" comment:"Available"`
}

func (b *Book) TableName() string {
	return "book"
}

func (b *Book) ToDTO() *dtos.Book {
	return &dtos.Book{
		ID:           &b.ID,
		TableID:      b.TableID,
		AvailableID:  b.AvailableID,
		Amount:       b.Amount,
		MobileNumber: b.MobileNumber,
		CommonCQRSHistoryModel: dtos.CommonCQRSHistoryModel{
			ActiveStatus: b.ActiveStatus,
			CreatedAt:    b.CreatedAt,
			UpdatedAt:    b.UpdatedAt,
		},
		Table: func() *dtos.Table {
			if b.Table == nil {
				return nil
			}
			return b.Table.ToDTO()
		}(),
		Available: func() *dtos.Available {
			if b.Available == nil {
				return nil
			}
			return b.Available.ToDTO()
		}(),
	}
}
