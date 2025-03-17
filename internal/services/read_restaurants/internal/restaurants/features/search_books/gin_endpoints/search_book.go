package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_books/queries"
	"github.com/gin-gonic/gin"
)

type searchBooksImpl struct {
	searchBooksQuery queries.SearchBooksHandler
}

func NewSearchBooks(
	searchBooksQuery queries.SearchBooksHandler,
) customizegin.Endpoint {
	return &searchBooksImpl{
		searchBooksQuery: searchBooksQuery,
	}
}

func (endpoint *searchBooksImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/books/search", endpoint.Handle)
}

// @Summary Search books
// @Description Search books
// @Tags books
// @Produce json
// @Param book body dtos.SearchBooks true "Book"
// @Success 200 {object} customizegin.JSONResponse{data=[]dtos.Book} "book retrieved successfully"
// @Router /books/search [post]
func (handle *searchBooksImpl) Handle(c *gin.Context) {
	var req dtos.SearchBooks
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	books, err := handle.searchBooksQuery.SearchBooks(c.Request.Context(), &queries.SearchBooks{
		MobileNumber:            req.MobileNumber,
		TableID:                 req.TableID,
		AvailableID:             req.AvailableID,
		NameFilter:              req.NameFilter,
		NamePreciseSearch:       req.NamePreciseSearch,
		TableAvailableWeek:      req.TableAvailableWeek,
		TableAvailableStartTime: req.TableAvailableStartTime,
		TableAvailableEndTime:   req.TableAvailableEndTime,
		OrderBy:                 req.OrderBy,
		Pagination:              req.Pagination,
	})
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(
		c,
		books.ToDTO(),
		"books retrieved successfully",
	)
}
