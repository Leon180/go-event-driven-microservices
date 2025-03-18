package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/list_categories/services"
	"github.com/gin-gonic/gin"
)

type listCategoriesImpl struct {
	listCategories services.ListCategories
}

func NewListCategories(listCategories services.ListCategories) customizegin.Endpoint {
	return &listCategoriesImpl{
		listCategories: listCategories,
	}
}

func (impl *listCategoriesImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/categories", impl.Handle)
}

// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Produce json
// @Success 200 {object} customizegin.JSONResponse "categories retrieved successfully"
// @Router /categories [get]
func (handle *listCategoriesImpl) Handle(c *gin.Context) {
	categories, err := handle.listCategories.ListCategories(c.Request.Context())
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(
		c,
		dtos.CategoriesEntity(categories).ToDTO(),
		"categories retrieved successfully",
	)
}
