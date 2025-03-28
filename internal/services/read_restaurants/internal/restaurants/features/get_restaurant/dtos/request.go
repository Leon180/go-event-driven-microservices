package featuresdtos

type GetRestaurantRequest struct {
	ID string `json:"id" binding:"required"`
}
