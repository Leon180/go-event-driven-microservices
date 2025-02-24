package featuresdtos

type GetAccountRequest struct {
	MobileNumber string `uri:"mobile_number" binding:"required"`
}
