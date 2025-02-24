package featuresdtos

type DeleteAccountRequest struct {
	ID string `uri:"id" binding:"required"`
}
