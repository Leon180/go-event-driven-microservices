package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/update_taked_book/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/update_taked_book/services"
	"github.com/gin-gonic/gin"
)

type updateTakedBookImpl struct {
	updateTakedBookService services.UpdateTakedBook
}

func NewUpdateTakedBook(
	updateTakedBookService services.UpdateTakedBook,
) customizegin.Endpoint {
	return &updateTakedBookImpl{
		updateTakedBookService: updateTakedBookService,
	}
}

func (endpoint *updateTakedBookImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/book/update-taked", endpoint.Handle)
}

// @Summary Update taked book
// @Description Update taked book
// @Tags books
// @Accept json
// @Produce json
// @Param book body featuresdtos.UpdateTakedBookRequest true "Book"
// @Success 200 {object} customizegin.JSONResponse "Book updated successfully"
// @Router /book/update-taked [post]
func (endpoint *updateTakedBookImpl) Handle(c *gin.Context) {
	var updateTakedBookRequest featuresdtos.UpdateTakedBookRequest
	if err := c.ShouldBindJSON(&updateTakedBookRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	err := endpoint.updateTakedBookService.UpdateTakedBook(c.Request.Context(), &updateTakedBookRequest)
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "book taken successfully")
}
