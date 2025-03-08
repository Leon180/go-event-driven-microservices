package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/delete_restaurant/services"
	"github.com/gin-gonic/gin"
)

type deleteRestaurantImpl struct {
	deleteRestaurantService services.DeleteRestaurant
}

func NewDeleteRestaurant(
	deleteRestaurantService services.DeleteRestaurant,
) customizegin.Endpoint {
	return &deleteRestaurantImpl{
		deleteRestaurantService: deleteRestaurantService,
	}
}

func (endpoint *deleteRestaurantImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/restaurant/delete", endpoint.Handle)
}

// @Summary Delete a restaurant
// @Description Delete a restaurant
// @Tags restaurants
// @Produce json
// @Param restaurant body featuresdtos.DeleteRestaurantRequest true "Restaurant"
// @Success 200 {object} customizegin.JSONResponse "restaurant deleted successfully"
// @Router /restaurant/delete [post]
func (endpoint *deleteRestaurantImpl) Handle(c *gin.Context) {
	var deleteRestaurantRequest featuresdtos.DeleteRestaurantRequest
	if err := c.ShouldBindJSON(&deleteRestaurantRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.deleteRestaurantService.DeleteRestaurant(c.Request.Context(), &deleteRestaurantRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "restaurant deleted successfully")
}
