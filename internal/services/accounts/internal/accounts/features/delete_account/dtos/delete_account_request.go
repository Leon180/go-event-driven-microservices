package dtos

type DeleteAccountRequest struct {
	ID string `json:"id" binding:"required"`
}
