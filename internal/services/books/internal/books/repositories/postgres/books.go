package repostgresespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/repositories"
	"gorm.io/gorm"
)

func NewSearchBooksFullInfo(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.SearchBooksFullInfo {
	return &searchBooksFullInfoImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type searchBooksFullInfoImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *searchBooksFullInfoImpl) SearchBooksFullInfo(
	ctx context.Context,
	searchBooks *dtos.SearchBooks,
) (entities.Books, error) {
	if searchBooks == nil {
		return nil, nil
	}
	sql := impl.buildBaseQuery(ctx).
		Scopes(
			impl.applyBranchIDFilter(searchBooks.BranchID),
			impl.applyDateFilter(searchBooks.StartDate, searchBooks.EndDate),
			impl.applyTimeFilter(searchBooks.StartTime, searchBooks.EndTime),
			impl.applyCapacityFilter(searchBooks.Capacity),
			impl.applyBookedFilter(searchBooks.Booked),
			impl.applyMobileNumberFilter(searchBooks.MobileNumber),
			impl.applyNoteFilter(searchBooks.Note),
			customizegorm.ApplyPagination(searchBooks.Pagination),
			customizegorm.ApplyOrdering(searchBooks.OrderBy),
		)
	var books entities.Books
	if err := sql.Find(&books).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to search books", err)
		return nil, err
	}
	return books, nil
}

func (impl *searchBooksFullInfoImpl) buildBaseQuery(ctx context.Context) *gorm.DB {
	return impl.db.WithContext(ctx)
}

func (impl *searchBooksFullInfoImpl) applyBranchIDFilter(branchID *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if branchID != nil {
			return db.Where("branch_id = ?", *branchID)
		}
		return db
	}
}

func (impl *searchBooksFullInfoImpl) applyDateFilter(startDate *string, endDate *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startDate != nil && endDate != nil {
			return db.Where("date >= ? AND date <= ?", *startDate, *endDate)
		}
		if startDate != nil {
			return db.Where("date >= ?", *startDate)
		}
		if endDate != nil {
			return db.Where("date <= ?", *endDate)
		}
		return db
	}
}

func (impl *searchBooksFullInfoImpl) applyTimeFilter(startTime *string, endTime *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startTime != nil && endTime != nil {
			return db.Where("start_time >= ? AND end_time <= ?", *startTime, *endTime)
		}
		if startTime != nil {
			return db.Where("start_time >= ?", *startTime)
		}
		if endTime != nil {
			return db.Where("end_time <= ?", *endTime)
		}
		return db
	}
}

func (impl *searchBooksFullInfoImpl) applyCapacityFilter(capacity *int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if capacity != nil {
			return db.Where("capacity >= ?", *capacity)
		}
		return db
	}
}

func (impl *searchBooksFullInfoImpl) applyBookedFilter(booked *bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if booked != nil {
			return db.Where("booked = ?", *booked)
		}
		return db
	}
}

func (impl *searchBooksFullInfoImpl) applyMobileNumberFilter(mobileNumber *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if mobileNumber != nil {
			return db.Where("mobile_number LIKE ?", "%"+*mobileNumber+"%")
		}
		return db
	}
}

func (impl *searchBooksFullInfoImpl) applyNoteFilter(note *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if note != nil {
			return db.Where("note LIKE ?", "%"+*note+"%")
		}
		return db
	}
}

func NewReadBook(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadBook {
	return &readBookImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type readBookImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *readBookImpl) ReadBook(ctx context.Context, id string) (*entities.Book, error) {
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
	return &updateBooksImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type updateBooksImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *updateBooksImpl) CreateBooks(ctx context.Context, books entities.Books) error {
	if len(books) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&books).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create books", err)
		return err
	}
	return nil
}

func (impl *updateBooksImpl) UpdateBook(ctx context.Context, updateBook *entities.UpdateBook) error {
	if updateBook == nil || updateBook.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.Book{}).Where("id = ?", updateBook.ID).Updates(updateBook.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update book", err)
		return err
	}
	return nil
}

func (impl *updateBooksImpl) DeleteBooks(ctx context.Context, ids []string) error {
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
