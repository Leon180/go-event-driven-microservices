package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/aggregates"
	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/customize_grpc"
	protobufsconvert "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/customize_grpc/convert/protobufs"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/customize_grpc/protobuf"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/repositories"
	"github.com/samber/lo"
)

type CreateBooks interface {
	CreateBooks(ctx context.Context, req *featuresdtos.CreateBooksRequest) (aggregates.Books, error)
}

func NewCreateBooks(
	uuidGenerator uuid.UUIDGenerator,
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction],
	searchBooksFullInfoRepository repositories.SearchBooksFullInfo,
	grpcBookService customizegrpc.GRPCBookService,
) CreateBooks {
	return &createBooksImpl{
		uuidGenerator:                        uuidGenerator,
		updateBooksWithTransactionRepository: updateBooksWithTransactionRepository,
		searchBooksFullInfoRepository:        searchBooksFullInfoRepository,
		grpcBookService:                      grpcBookService,
	}
}

type createBooksImpl struct {
	uuidGenerator                        uuid.UUIDGenerator
	updateBooksWithTransactionRepository customizegorm.Transactor[repositories.UpdateBooksWithTransaction]
	searchBooksFullInfoRepository        repositories.SearchBooksFullInfo
	grpcBookService                      customizegrpc.GRPCBookService
}

func (handle *createBooksImpl) CreateBooks(ctx context.Context, req *featuresdtos.CreateBooksRequest) (aggregates.Books, error) {
	if req == nil {
		return nil, nil
	}
	// check if book already exists
	bookEntities, err := handle.searchBooksFullInfoRepository.SearchBooksFullInfo(ctx, &dtos.SearchBooks{
		BranchID:  &req.BranchID,
		StartDate: &req.StartDate,
		EndDate:   &req.EndDate,
	})
	if err != nil {
		return nil, err
	}
	books := aggregates.BookEntities(bookEntities).ToAggregates()
	if books.InTimePeriod(req.StartDate, req.EndDate) {
		return books, customizeerrors.BookAlreadyExistsError
	}

	// get restaurant info from grpc
	restaurantProtobuf, err := handle.grpcBookService.GetRestaurantBranch(ctx, &protobuf.GetRestaurantBranchReq{
		BranchId: req.BranchID,
	})
	if err != nil {
		return nil, err
	}
	if restaurantProtobuf == nil {
		return nil, customizeerrors.RestaurantNotFoundError
	}

	// convert restaurants available and tables to books
	restaurant := protobufsconvert.RestaurantProtobufToAggregate(restaurantProtobuf)
	branch, ok := lo.Find(restaurant.Branches, func(branch aggregates.Branch) bool {
		return branch.ID == req.BranchID
	})
	if !ok {
		return nil, customizeerrors.BranchNotFoundError
	}
	startDate, _ := time.Parse(time.DateOnly, req.StartDate)
	endDate, _ := time.Parse(time.DateOnly, req.EndDate)
	dateBooks := aggregates.NewBranchBooksBuilder(handle.uuidGenerator).
		SetBranch(&branch).
		BuildBoooksByTimePeriod(startDate, endDate).
		GetDateBooks()

	bookEntities = entities.Books{}
	for _, dateBook := range dateBooks {
		bookEntities = append(bookEntities, dateBook.Books.ToEntities()...)
	}

	// create books
	tx, err := handle.updateBooksWithTransactionRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if err := tx.CreateBooks(ctx, bookEntities); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return books, nil
}
