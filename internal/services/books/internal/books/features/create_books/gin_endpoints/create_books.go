package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_books/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/create_books/services"
	"github.com/gin-gonic/gin"
)

type createBooksImpl struct {
	createBooksService services.CreateBooks
}

func NewCreateBooks(
	createBooksService services.CreateBooks,
) customizegin.Endpoint {
	return &createBooksImpl{
		createBooksService: createBooksService,
	}
}

func (endpoint *createBooksImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/books/create", endpoint.Handle)
}

// @Summary Create books by branch id, start date and end date
// @Description Create books by branch id, start date and end date
// @Tags books
// @Accept json
// @Produce json
// @Param book body featuresdtos.CreateBooksRequest true "Book"
// @Success 200 {object} customizegin.JSONResponse "Books created successfully"
// @Router /books/create [post]
func (endpoint *createBooksImpl) Handle(c *gin.Context) {
	var createBooksRequest featuresdtos.CreateBooksRequest
	if err := c.ShouldBindJSON(&createBooksRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	books, err := endpoint.createBooksService.CreateBooks(c.Request.Context(), &createBooksRequest)
	if err != nil {
		customizegin.ResponseError(c, books, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "books created successfully")
}
