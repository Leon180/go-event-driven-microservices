package aggregates

import (
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
	validatesdtos "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/validates/dtos"
)

type BookDTOAggregateBuilder interface {
	// save book and related entities(table, available) to the aggregate
	SaveBook(book *dtos.Book) error

	// set all edit type code to none
	SetAllEditTypeCodeToNone()

	// get book aggregates
	GetAggregates() []Book

	// get pending edit entities during update
	GetEditEntities() []BookEditEntities
}

func NewBookDTOAggregateBuilder(uuidGenerator uuid.UUIDGenerator) BookDTOAggregateBuilder {
	return &bookDTOAggregateImpl{
		books:         make([]Book, 0),
		uuidGenerator: uuidGenerator,
	}
}

type bookDTOAggregateImpl struct {
	books []Book

	// dependencies
	uuidGenerator uuid.UUIDGenerator
}

type BookEditEntities struct {
	CreateEntities *BookCreateEntities
	UpdateEntities *BookUpdateEntities
	DeleteEntities *BookDeleteEntities
}

func newBookEditEntities() BookEditEntities {
	return BookEditEntities{
		CreateEntities: &BookCreateEntities{},
		UpdateEntities: &BookUpdateEntities{},
		DeleteEntities: &BookDeleteEntities{},
	}
}

type BookCreateEntities struct {
	Books []entities.Book
}

type BookUpdateEntities struct {
	Books []entities.UpdateBook
}

type BookDeleteEntities struct {
	Books []entities.Book
}

func (b *bookDTOAggregateImpl) SaveBook(book *dtos.Book) error {
	if book == nil {
		return nil
	}
	if err := validatesdtos.ValidateBook(book); err != nil {
		return err
	}
	for i := range b.books {
		if book.ID != nil && b.books[i].ID == *book.ID {
			if book.TableID == b.books[i].TableID && book.AvailableID == b.books[i].AvailableID {
				return b.updateBook(&b.books[i], book)
			}
			return customizeerrors.BookTableAndAvailableNotMatchError
		}
		// same table id and available id, duplicate book
		if b.books[i].TableID == book.TableID && b.books[i].AvailableID == book.AvailableID {
			return customizeerrors.BookAlreadyExistsError
		}
	}
	return b.addBook(book)
}

func (b *bookDTOAggregateImpl) addBook(book *dtos.Book) error {
	if book == nil {
		return nil
	}
	var bookID string
	if book.ID == nil {
		bookID = b.uuidGenerator.GenerateUUID()
	} else {
		bookID = *book.ID
	}
	timeNow := time.Now()
	b.books = append(b.books, Book{
		Book: entities.Book{
			ID:           bookID,
			TableID:      book.TableID,
			AvailableID:  book.AvailableID,
			Amount:       book.Amount,
			MobileNumber: book.MobileNumber,
			CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
				ActiveStatus: true,
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
			},
		},
		editTypeCode: enums.EditTypeCodeCreate,
	})
	return nil
}

func (b *bookDTOAggregateImpl) updateBook(book *Book, update *dtos.Book) error {
	if book == nil || update == nil {
		return nil
	}
	book.TableID = update.TableID
	book.AvailableID = update.AvailableID
	book.Amount = update.Amount
	book.MobileNumber = update.MobileNumber
	book.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	book.update = update.ToUpdateBook().RemoveUnchangedFields(book.Book)
	if book.update != nil {
		book.editTypeCode = enums.EditTypeCodeUpdate
	}
	return nil
}

func (b *bookDTOAggregateImpl) SetAllEditTypeCodeToNone() {
	for i := range b.books {
		b.books[i].editTypeCode = enums.EditTypeCodeNone
	}
}

func (impl *bookDTOAggregateImpl) GetAggregates() []Book {
	return impl.books
}

func (impl *bookDTOAggregateImpl) GetEditEntities() []BookEditEntities {
	res := make([]BookEditEntities, len(impl.books))
	for i, book := range impl.books {
		// Initialize container
		res[i] = newBookEditEntities()
		createEntities, updateEntities, deleteEntities := res[i].CreateEntities, res[i].UpdateEntities, res[i].DeleteEntities

		// Handle restaurant
		appendEntityByType(
			&createEntities.Books,
			&updateEntities.Books,
			&deleteEntities.Books,
			book.Book,
			book.update,
			book.editTypeCode,
		)
	}
	return res
}
