package featuresdtos

type DeleteAccountRequest struct {
	ID string `json:"id" binding:"required"`
}
