package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/features/search_restaurants/queries"
	"github.com/gin-gonic/gin"
)

type searchRestaurantsImpl struct {
	searchRestaurantsQuery queries.SearchRestaurantsHandler
}

func NewSearchRestaurants(
	searchRestaurantsQuery queries.SearchRestaurantsHandler,
) customizegin.Endpoint {
	return &searchRestaurantsImpl{
		searchRestaurantsQuery: searchRestaurantsQuery,
	}
}

func (endpoint *searchRestaurantsImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/restaurants/search", endpoint.Handle)
}

// @Summary Search restaurants
// @Description Search restaurants
// @Tags restaurants
// @Produce json
// @Param restaurant body dtos.SearchRestaurants true "Restaurant"
// @Success 200 {object} customizegin.JSONResponse{data=[]dtos.Restaurant} "restaurant retrieved successfully"
// @Router /restaurants/search [post]
func (handle *searchRestaurantsImpl) Handle(c *gin.Context) {
	var req dtos.SearchRestaurants
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	restaurants, err := handle.searchRestaurantsQuery.SearchRestaurants(c.Request.Context(), &queries.SearchRestaurants{
		NameFilter:              req.NameFilter,
		NamePreciseSearch:       req.NamePreciseSearch,
		BranchID:                req.BranchID,
		DescriptionFilter:       req.DescriptionFilter,
		CityFilter:              req.CityFilter,
		CountryFilter:           req.CountryFilter,
		MaxPriceFilter:          req.MaxPriceFilter,
		MinPriceFilter:          req.MinPriceFilter,
		CategoryFilter:          req.CategoryFilter,
		TableAvailableWeek:      req.TableAvailableWeek,
		TableAvailableStartTime: req.TableAvailableStartTime,
		TableAvailableEndTime:   req.TableAvailableEndTime,
		OrderBy:                 req.OrderBy,
		Pagination:              req.Pagination,
	})
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
