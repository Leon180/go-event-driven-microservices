package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/create_restaurant/services"
	"github.com/gin-gonic/gin"
)

type createRestaurantImpl struct {
	createRestaurantService services.CreateRestaurant
}

func NewCreateRestaurant(
	createRestaurantService services.CreateRestaurant,
) customizegin.Endpoint {
	return &createRestaurantImpl{
		createRestaurantService: createRestaurantService,
	}
}

func (endpoint *createRestaurantImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/restaurant/create", endpoint.Handle)
}

// @Summary Create a new restaurant
// @Description Create a new restaurant
// @Tags restaurants
// @Accept json
// @Produce json
// @Param restaurant body dtos.Restaurant true "Restaurant"
// @Success 200 {object} customizegin.JSONResponse "Restaurant created successfully"
// @Router /restaurant/create [post]
func (endpoint *createRestaurantImpl) Handle(c *gin.Context) {
	var createRestaurantRequest dtos.Restaurant
	if err := c.ShouldBindJSON(&createRestaurantRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.createRestaurantService.CreateRestaurant(c.Request.Context(), &createRestaurantRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "restaurant created successfully")
}
