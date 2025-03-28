package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/create_book/services"
	"github.com/gin-gonic/gin"
)

type createBookImpl struct {
	createBookService services.CreateBook
}

func NewCreateBook(
	createBookService services.CreateBook,
) customizegin.Endpoint {
	return &createBookImpl{
		createBookService: createBookService,
	}
}

func (endpoint *createBookImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/book/create", endpoint.Handle)
}

// @Summary Create a new book
// @Description Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body dtos.Book true "Book"
// @Success 200 {object} customizegin.JSONResponse "Book created successfully"
// @Router /book/create [post]
func (endpoint *createBookImpl) Handle(c *gin.Context) {
	var createBookRequest dtos.Book
	if err := c.ShouldBindJSON(&createBookRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.createBookService.CreateBook(c.Request.Context(), &createBookRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "book created successfully")
}
