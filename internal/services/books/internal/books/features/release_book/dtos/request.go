package featuresdtos

type DeleteBookRequest struct {
	ID string `json:"id" binding:"required"`
}
