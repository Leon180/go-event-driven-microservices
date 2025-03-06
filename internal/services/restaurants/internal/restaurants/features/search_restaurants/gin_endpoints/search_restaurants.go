package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/search_restaurants/services"
	"github.com/gin-gonic/gin"
)

type searchRestaurantsImpl struct {
	searchRestaurantsService services.SearchRestaurants
}

func NewSearchRestaurants(
	searchRestaurantsService services.SearchRestaurants,
) customizegin.Endpoint {
	return &searchRestaurantsImpl{
		searchRestaurantsService: searchRestaurantsService,
	}
}

func (endpoint *searchRestaurantsImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/restaurants/search", endpoint.Handle)
}

// @Summary Search restaurants
// @Description Search restaurants
// @Tags restaurants
// @Produce json
// @Param restaurant body dtos.SearchRestaurantsRequest true "Restaurant"
// @Success 200 {object} customizegin.JSONResponse{data=[]dtos.Restaurant} "restaurant retrieved successfully"
// @Router /restaurants/search [post]
func (handle *searchRestaurantsImpl) Handle(c *gin.Context) {
	var req dtos.SearchRestaurants
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	restaurants, err := handle.searchRestaurantsService.SearchRestaurants(c.Request.Context(), &req)
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(
		c,
		restaurants.ToDTO(),
		"restaurants retrieved successfully",
	)
}
