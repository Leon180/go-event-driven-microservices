package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/release_book/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/release_book/services"
	"github.com/gin-gonic/gin"
)

type releaseBookImpl struct {
	releaseBookService services.ReleaseBook
}

func NewReleaseBook(
	releaseBookService services.ReleaseBook,
) customizegin.Endpoint {
	return &releaseBookImpl{
		releaseBookService: releaseBookService,
	}
}

func (endpoint *releaseBookImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/book/release", endpoint.Handle)
}

// @Summary Release book
// @Description Release book
// @Tags books
// @Accept json
// @Produce json
// @Param book body featuresdtos.ReleaseBookRequest true "Book"
// @Success 200 {object} customizegin.JSONResponse "Book released successfully"
// @Router /book/release [post]
func (endpoint *releaseBookImpl) Handle(c *gin.Context) {
	var releaseBookRequest featuresdtos.ReleaseBookRequest
	if err := c.ShouldBindJSON(&releaseBookRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	err := endpoint.releaseBookService.ReleaseBook(c.Request.Context(), releaseBookRequest.BookID)
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "book released successfully")
}
