package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/list_categories/queries"
	"github.com/gin-gonic/gin"
)

type listCategoriesImpl struct {
	listCategoriesQuery queries.ListCategoriesHandler
}

func NewListCategories(listCategoriesQuery queries.ListCategoriesHandler) customizegin.Endpoint {
	return &listCategoriesImpl{
		listCategoriesQuery: listCategoriesQuery,
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
	categories, err := handle.listCategoriesQuery.ListCategories(c.Request.Context())
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(
		c,
		categories.ToDTO(),
		"categories retrieved successfully",
	)
}
