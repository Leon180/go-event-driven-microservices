package aggregates

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	"github.com/samber/lo"
)

type Books []Book

func (b Books) ToEntities() entities.Books {
	return lo.Map(b, func(book Book, _ int) entities.Book {
		return *book.ToEntity()
	})
}

func (b Books) ToDTO() []dtos.Book {
	return lo.Map(b, func(book Book, _ int) dtos.Book {
		return *book.ToDTO()
	})
}

// InTimePeriod checks if the book is in the time period
// startDate/endDate format: YYYY-MM-DD
func (b Books) InTimePeriod(startDate string, endDate string) bool {
	sd, _ := time.Parse(time.DateOnly, startDate)
	ed, _ := time.Parse(time.DateOnly, endDate)
	for _, book := range b {
		bd, _ := time.Parse(time.DateOnly, book.Date)
		if (bd.After(sd) || bd.Equal(sd)) && (bd.Before(ed) || bd.Equal(ed)) {
			return true
		}
	}
	return false
}

type Book entities.Book

func (b *Book) ToEntity() *entities.Book {
	bs := *b
	eb := entities.Book(bs)
	return &eb
}

func (b *Book) ToDTO() *dtos.Book {
	return &dtos.Book{
		ID:       b.ID,
		BranchID: b.BranchID,
		Date: func() time.Time {
			date, _ := time.Parse(time.DateOnly, b.Date)
			return date
		}(),
		StartTime:    b.StartTime,
		EndTime:      b.EndTime,
		Capacity:     b.Capacity,
		Booked:       b.Booked,
		Amount:       b.Amount,
		MobileNumber: b.MobileNumber,
		Note:         b.Note,
	}
}

type BookEntities entities.Books

func (b BookEntities) ToAggregates() Books {
	return lo.Map(b, func(book entities.Book, _ int) Book {
		be := BookEntity(book)
		return *be.ToAggregate()
	})
}

type BookEntity entities.Book

func (b *BookEntity) ToAggregate() *Book {
	bs := *b
	ab := Book(bs)
	return &ab
}
