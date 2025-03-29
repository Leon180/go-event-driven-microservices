package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/features/update_restaurant/services"
	"github.com/gin-gonic/gin"
)

type updateRestaurantImpl struct {
	updateRestaurantService services.UpdateRestaurant
}

func NewUpdateRestaurant(
	updateRestaurantService services.UpdateRestaurant,
) customizegin.Endpoint {
	return &updateRestaurantImpl{
		updateRestaurantService: updateRestaurantService,
	}
}

func (endpoint *updateRestaurantImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.PUT("/restaurant/update", endpoint.Handle)
}

// @Summary Update a restaurant
// @Description Update a restaurant
// @Tags restaurants
// @Accept json
// @Produce json
// @Param restaurant body dtos.Restaurant true "Restaurant"
// @Success 200 {object} customizegin.JSONResponse "restaurant updated successfully"
// @Router /restaurant/update [put]
func (endpoint *updateRestaurantImpl) Handle(c *gin.Context) {
	var req dtos.Restaurant
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.updateRestaurantService.UpdateRestaurant(c.Request.Context(), &req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "restaurant updated successfully")
}
