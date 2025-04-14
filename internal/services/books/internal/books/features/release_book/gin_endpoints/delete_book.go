package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/delete_book/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/delete_book/services"
	"github.com/gin-gonic/gin"
)

type deleteBookImpl struct {
	deleteBookService services.DeleteBook
}

func NewDeleteBook(
	deleteBookService services.DeleteBook,
) customizegin.Endpoint {
	return &deleteBookImpl{
		deleteBookService: deleteBookService,
	}
}

func (endpoint *deleteBookImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/book/delete", endpoint.Handle)
}

// @Summary Delete a book
// @Description Delete a book
// @Tags books
// @Produce json
// @Param book body featuresdtos.DeleteBookRequest true "Book"
// @Success 200 {object} customizegin.JSONResponse "book deleted successfully"
// @Router /book/delete [post]
func (endpoint *deleteBookImpl) Handle(c *gin.Context) {
	var deleteBookRequest featuresdtos.DeleteBookRequest
	if err := c.ShouldBindJSON(&deleteBookRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.deleteBookService.DeleteBook(c.Request.Context(), &deleteBookRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "book deleted successfully")
}
