package aggregates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/documents"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
	"github.com/samber/lo"
)

type Books []Book

func (b Books) ToDTO() []dtos.Book {
	return lo.Map(b, func(book Book, _ int) dtos.Book {
		return *book.ToDTO()
	})
}

func (b Books) ToDocuments() []documents.Book {
	return lo.Map(b, func(book Book, _ int) documents.Book {
		return *book.ToDocument()
	})
}

type Book struct {
	entities.Book
	editTypeCode enums.EditTypeCode   `gorm:"-"`
	update       *entities.UpdateBook `gorm:"-"`
	Table        *Table               `gorm:"foreignKey:ID;references:TableID"     comment:"Table"`
	Available    *Available           `gorm:"foreignKey:ID;references:AvailableID" comment:"Available"`
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

func (b *Book) ToDocument() *documents.Book {
	return &documents.Book{
		ID:                     b.ID,
		TableID:                b.TableID,
		AvailableID:            b.AvailableID,
		Amount:                 b.Amount,
		MobileNumber:           b.MobileNumber,
		CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(b.CommonCQRSHistoryModel).ToDocument(),
		Table: func() documents.Table {
			if b.Table == nil {
				return documents.Table{}
			}
			return documents.Table{
				ID:                     b.Table.ID,
				BranchID:               b.Table.BranchID,
				Capacity:               b.Table.Capacity,
				CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(b.Table.CommonCQRSHistoryModel).ToDocument(),
			}
		}(),
		Available: func() documents.Available {
			if b.Available == nil {
				return documents.Available{}
			}
			return documents.Available{
				ID:                     b.Available.ID,
				BranchID:               b.Available.BranchID,
				Weekday:                b.Available.Weekday,
				StartTime:              b.Available.StartTime,
				EndTime:                b.Available.EndTime,
				CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(b.Available.CommonCQRSHistoryModel).ToDocument(),
			}
		}(),
	}
}

type BookDocuments documents.Books

func (b BookDocuments) ToAggregate() Books {
	return lo.Map(b, func(book documents.Book, _ int) Book {
		bookDocument := BookDocument(book)
		return *bookDocument.ToAggregate()
	})
}

type BookDocument documents.Book

func (b *BookDocument) ToAggregate() *Book {
	return &Book{
		Book: entities.Book{
			ID:                     b.ID,
			TableID:                b.TableID,
			AvailableID:            b.AvailableID,
			Amount:                 b.Amount,
			MobileNumber:           b.MobileNumber,
			CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(b.CommonCQRSHistoryModel).ToEntity(),
		},
		Table: &Table{
			Table: entities.Table{
				ID:                     b.Table.ID,
				BranchID:               b.Table.BranchID,
				Capacity:               b.Table.Capacity,
				CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(b.Table.CommonCQRSHistoryModel).ToEntity(),
			},
		},
		Available: &Available{
			Available: entities.Available{
				ID:                     b.Available.ID,
				BranchID:               b.Available.BranchID,
				Weekday:                b.Available.Weekday,
				StartTime:              b.Available.StartTime,
				EndTime:                b.Available.EndTime,
				CommonCQRSHistoryModel: CommonCQRSHistoryModelDocument(b.Available.CommonCQRSHistoryModel).ToEntity(),
			},
		},
	}
}
