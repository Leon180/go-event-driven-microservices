package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant_branch/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/get_restaurant_branch/queries"
	"github.com/gin-gonic/gin"
)

type getRestaurantBranchImpl struct {
	getRestaurantBranchQuery queries.GetRestaurantBranchHandler
}

func NewGetRestaurantBranch(
	getRestaurantBranchQuery queries.GetRestaurantBranchHandler,
) customizegin.Endpoint {
	return &getRestaurantBranchImpl{
		getRestaurantBranchQuery: getRestaurantBranchQuery,
	}
}

func (endpoint *getRestaurantBranchImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/restaurant/branch/get", endpoint.Handle)
}

// @Summary Get restaurant branch by id
// @Description Get restaurant branch by id
// @Tags restaurants
// @Produce json
// @Param restaurant body featuresdtos.GetRestaurantBranchRequest true "Restaurant Branch"
// @Success 200 {object} customizegin.JSONResponse "restaurant retrieved successfully"
// @Router /restaurant/branch/get [post]
func (handle *getRestaurantBranchImpl) Handle(c *gin.Context) {
	var req featuresdtos.GetRestaurantBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	restaurant, err := handle.getRestaurantBranchQuery.GetRestaurantBranch(
		c.Request.Context(),
		&queries.GetRestaurantBranch{BranchID: req.BranchID},
	)
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(
		c,
		restaurant.ToDTO(),
		"restaurant branch retrieved successfully",
	)
}
