package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/customize_grpc"
	protobufsconvert "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/customize_grpc/convert/protobufs"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/customize_grpc/protobuf"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/repositories"
)

type CreateBooks interface {
	CreateBooks(ctx context.Context, req *featuresdtos.CreateBooksRequest) (entities.Books, error)
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

func (handle *createBooksImpl) CreateBooks(ctx context.Context, req *featuresdtos.CreateBooksRequest) (entities.Books, error) {
	if req == nil {
		return nil, nil
	}
	// check if book already exists
	books, err := handle.searchBooksFullInfoRepository.SearchBooksFullInfo(ctx, &dtos.SearchBooks{
		BranchID:  &req.BranchID,
		StartDate: &req.StartDate,
		EndDate:   &req.EndDate,
	})
	if err != nil {
		return nil, err
	}
	if len(books) != 0 {
		return books, customizeerrors.BookAlreadyExistsError
	}

	// get restaurant info from grpc
	restaurants, err := handle.grpcBookService.SearchRestaurants(ctx, &protobuf.SearchRestaurantsReq{
		BranchID: &req.BranchID,
	})
	if err != nil {
		return nil, err
	}
	if restaurants == nil || len(restaurants.Restaurants) == 0 {
		return nil, customizeerrors.RestaurantNotFoundError
	}

	// convert restaurants available and tables to books
	restaurant := protobufsconvert.RestaurantProtobufToAggregate(restaurants.Restaurants[0])

	books = entities.Books{}

	// create books
	tx, err := handle.updateBooksWithTransactionRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if err := tx.CreateBooks(ctx, books); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return books, nil
}
