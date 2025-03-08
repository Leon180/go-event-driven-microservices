package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/restore_restaurant/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/restore_restaurant/services"
	"github.com/gin-gonic/gin"
)

type restoreRestaurantImpl struct {
	restoreRestaurantService services.RestoreRestaurant
}

func NewRestoreRestaurant(
	restoreRestaurantService services.RestoreRestaurant,
) customizegin.Endpoint {
	return &restoreRestaurantImpl{
		restoreRestaurantService: restoreRestaurantService,
	}
}

func (endpoint *restoreRestaurantImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/restaurant/restore", endpoint.Handle)
}

// @Summary Restore a restaurant
// @Description Restore a restaurant
// @Tags restaurants
// @Produce json
// @Param restaurant body featuresdtos.RestoreRestaurantRequest true "Restaurant"
// @Success 200 {object} customizegin.JSONResponse "restaurant restored successfully"
// @Router /restaurant/restore [post]
func (endpoint *restoreRestaurantImpl) Handle(c *gin.Context) {
	var restoreRestaurantRequest featuresdtos.RestoreRestaurantRequest
	if err := c.ShouldBindJSON(&restoreRestaurantRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.restoreRestaurantService.RestoreRestaurant(c.Request.Context(), &restoreRestaurantRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "restaurant restored successfully")
}
