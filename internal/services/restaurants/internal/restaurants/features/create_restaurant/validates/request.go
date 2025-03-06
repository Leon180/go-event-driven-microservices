package validates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/features/create_restaurant/dtos"
	validatesdtos "github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/validates/dtos"
)

func ValidateCreateRestaurantRequest(req *featuresdtos.CreateRestaurantRequest) error {
	if req == nil {
		return nil
	}
	restaurant := dtos.Restaurant(*req)
	return validatesdtos.ValidateRestaurant(&restaurant)
}
