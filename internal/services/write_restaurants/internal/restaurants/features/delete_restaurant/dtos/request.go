package featuresdtos

type DeleteRestaurantRequest struct {
	ID string `json:"id" binding:"required"`
}
