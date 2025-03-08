package services

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/repositories"
)

type SearchBooks interface {
	SearchBooks(
		ctx context.Context,
		req *dtos.SearchBooks,
	) (aggregates.Books, error)
}

func NewSearchBooks(
	searchBooksFullInfoRepository repositories.SearchBooksFullInfo,
) SearchBooks {
	return &searchBooksImpl{searchBooksFullInfoRepository: searchBooksFullInfoRepository}
}

type searchBooksImpl struct {
	searchBooksFullInfoRepository repositories.SearchBooksFullInfo
}

func (handle *searchBooksImpl) SearchBooks(
	ctx context.Context,
	req *dtos.SearchBooks,
) (aggregates.Books, error) {
	if req == nil {
		return nil, nil
	}
	return handle.searchBooksFullInfoRepository.SearchBooksFullInfo(ctx, req)
}
