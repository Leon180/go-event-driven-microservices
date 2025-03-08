package repostgresespostgres

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
	"gorm.io/gorm"
)

func NewSearchBooksFullInfo(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.SearchBooksFullInfo {
	return &SearchBooksFullInfoImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type SearchBooksFullInfoImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *SearchBooksFullInfoImpl) SearchBooksFullInfo(
	ctx context.Context,
	searchBooks *dtos.SearchBooks,
) (aggregates.Books, error) {
	if searchBooks == nil {
		return nil, nil
	}
	sql := impl.buildBaseQuery(ctx).
		Scopes(
			impl.applyMobileNumberFilter(searchBooks.MobileNumber),
			impl.applyTableIDFilter(searchBooks.TableID),
			impl.applyAvailableIDFilter(searchBooks.AvailableID),
			impl.applyTableAvailabilityFilter(searchBooks.TableAvailableWeek, searchBooks.TableAvailableStartTime, searchBooks.TableAvailableEndTime),
			customizegorm.ApplyPagination(searchBooks.Pagination),
			customizegorm.ApplyOrdering(searchBooks.OrderBy),
		)
	var books []aggregates.Book
	if err := sql.Find(&books).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to search books", err)
		return nil, err
	}
	return books, nil
}

func (impl *SearchBooksFullInfoImpl) buildBaseQuery(ctx context.Context) *gorm.DB {
	return impl.db.WithContext(ctx).
		Preload("Table").
		Preload("Available")
}

func (impl *SearchBooksFullInfoImpl) applyMobileNumberFilter(mobileNumber *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if mobileNumber != nil {
			return db.Where("mobile_number = ?", *mobileNumber)
		}
		return db
	}
}

func (impl *SearchBooksFullInfoImpl) applyTableIDFilter(tableID *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tableID != nil {
			return db.Where("table_id = ?", *tableID)
		}
		return db
	}
}

func (impl *SearchBooksFullInfoImpl) applyAvailableIDFilter(availableID *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if availableID != nil {
			return db.Where("available_id = ?", *availableID)
		}
		return db
	}
}

func (impl *SearchBooksFullInfoImpl) applyTableAvailabilityFilter(
	availableWeek []time.Weekday,
	availableStartTime *string,
	availableEndTime *string,
) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(availableWeek) == 0 && availableStartTime == nil && availableEndTime == nil {
			return db
		}

		query := db.Where(`book.available_id IN (
			SELECT id FROM available WHERE 1=1`)

		if len(availableWeek) > 0 {
			query = query.Where("AND available.weekday IN (?)", availableWeek)
		}

		if availableStartTime != nil {
			query = query.Where("AND available.start_time >= ?", *availableStartTime)
		}

		if availableEndTime != nil {
			query = query.Where("AND available.end_time <= ?", *availableEndTime)
		}

		return query.Where(")")
	}
}

func NewReadBooks(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadBooks {
	return &ReadBooksImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type ReadBooksImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *ReadBooksImpl) ReadBookFullInfo(ctx context.Context, id string) (*aggregates.Book, error) {
	if id == "" {
		return nil, nil
	}
	var book aggregates.Book
	if err := impl.db.WithContext(ctx).
		Preload("Table").
		Preload("Available").
		Where("id = ?", id).Limit(1).Find(&book).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to read book full info", err)
		return nil, err
	}
	return &book, nil
}

func (impl *ReadBooksImpl) ReadBook(ctx context.Context, id string) (*entities.Book, error) {
	if id == "" {
		return nil, nil
	}
	var book entities.Book
	if err := impl.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&book).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to read book", err)
		return nil, err
	}
	return &book, nil
}

func NewUpdateBooks(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateBooks {
	return &UpdateBooksImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type UpdateBooksImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *UpdateBooksImpl) CreateBooks(ctx context.Context, books entities.Books) error {
	if len(books) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&books).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create books", err)
		return err
	}
	return nil
}

func (impl *UpdateBooksImpl) UpdateBook(ctx context.Context, updateBook *entities.UpdateBook) error {
	if updateBook == nil || updateBook.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.Book{}).Where("id = ?", updateBook.ID).Updates(updateBook.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update book", err)
		return err
	}
	return nil
}

func (impl *UpdateBooksImpl) DeleteBooks(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Book{}).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete books", err)
		return err
	}
	return nil
}

func NewUpdateBooksWithTransaction(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) customizegorm.Transactor[repositories.UpdateBooksWithTransaction] {
	return &UpdateBooksWithTransactionImpl{
		TransactorImpl: TransactorImpl{
			db:            db,
			contextLogger: contextLogger,
		},
	}
}

type UpdateBooksWithTransactionImpl struct {
	TransactorImpl
}

func (impl *UpdateBooksWithTransactionImpl) BeginTx(
	ctx context.Context,
) (repositories.UpdateBooksWithTransaction, error) {
	tx := impl.db.WithContext(ctx).Begin()
	return &UpdateBooksTransactionImpl{
		TransactionImpl: TransactionImpl{
			Db:            tx,
			ContextLogger: impl.contextLogger,
		},
	}, nil
}

type UpdateBooksTransactionImpl struct {
	TransactionImpl
}

func (impl *UpdateBooksTransactionImpl) CreateBooks(ctx context.Context, books entities.Books) error {
	return NewUpdateBooks(impl.Db, impl.ContextLogger).CreateBooks(ctx, books)
}

func (impl *UpdateBooksTransactionImpl) UpdateBook(ctx context.Context, updateBook *entities.UpdateBook) error {
	return NewUpdateBooks(impl.Db, impl.ContextLogger).UpdateBook(ctx, updateBook)
}

func (impl *UpdateBooksTransactionImpl) DeleteBooks(ctx context.Context, ids []string) error {
	return NewUpdateBooks(impl.Db, impl.ContextLogger).DeleteBooks(ctx, ids)
}
