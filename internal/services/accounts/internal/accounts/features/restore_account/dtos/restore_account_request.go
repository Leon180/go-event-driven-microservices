package featuresdtos

type RestoreAccountRequest struct {
	ID string `json:"id" binding:"required"`
}
