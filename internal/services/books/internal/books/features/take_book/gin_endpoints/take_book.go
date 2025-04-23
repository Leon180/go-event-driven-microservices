package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/take_book/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/take_book/services"
	"github.com/gin-gonic/gin"
)

type takeBookImpl struct {
	takeBookService services.TakeBook
}

func NewTakeBook(
	takeBookService services.TakeBook,
) customizegin.Endpoint {
	return &takeBookImpl{
		takeBookService: takeBookService,
	}
}

func (endpoint *takeBookImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/book/take", endpoint.Handle)
}

// @Summary Take book
// @Description Take book
// @Tags books
// @Accept json
// @Produce json
// @Param book body featuresdtos.TakeBookRequest true "Book"
// @Success 200 {object} customizegin.JSONResponse "Book taken successfully"
// @Router /book/take [post]
func (endpoint *takeBookImpl) Handle(c *gin.Context) {
	var takeBookRequest featuresdtos.TakeBookRequest
	if err := c.ShouldBindJSON(&takeBookRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	err := endpoint.takeBookService.TakeBook(c.Request.Context(), &takeBookRequest)
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "book taken successfully")
}
