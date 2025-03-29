package featuresdtos

type RestoreRestaurantRequest struct {
	ID string `json:"id" binding:"required"`
}
