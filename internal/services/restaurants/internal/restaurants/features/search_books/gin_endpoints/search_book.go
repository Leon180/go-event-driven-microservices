package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/search_books/services"
	"github.com/gin-gonic/gin"
)

type searchBooksImpl struct {
	searchBooksService services.SearchBooks
}

func NewSearchBooks(
	searchBooksService services.SearchBooks,
) customizegin.Endpoint {
	return &searchBooksImpl{
		searchBooksService: searchBooksService,
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
	books, err := handle.searchBooksService.SearchBooks(c.Request.Context(), &req)
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
