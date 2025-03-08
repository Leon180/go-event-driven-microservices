package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant/services"
	"github.com/gin-gonic/gin"
)

type getRestaurantImpl struct {
	getRestaurantService services.GetRestaurant
}

func NewGetRestaurant(
	getRestaurantService services.GetRestaurant,
) customizegin.Endpoint {
	return &getRestaurantImpl{
		getRestaurantService: getRestaurantService,
	}
}

func (endpoint *getRestaurantImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/restaurant/get", endpoint.Handle)
}

// @Summary Get restaurant by id
// @Description Get restaurant by id
// @Tags restaurants
// @Produce json
// @Param restaurant body featuresdtos.GetRestaurantRequest true "Restaurant"
// @Success 200 {object} customizegin.JSONResponse "restaurant retrieved successfully"
// @Router /restaurant/get [post]
func (handle *getRestaurantImpl) Handle(c *gin.Context) {
	var req featuresdtos.GetRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	restaurant, err := handle.getRestaurantService.GetRestaurant(c.Request.Context(), &req)
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(
		c,
		restaurant.ToDTO(),
		"restaurant retrieved successfully",
	)
}
