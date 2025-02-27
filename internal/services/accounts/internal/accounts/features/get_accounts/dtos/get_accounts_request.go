package featuresdtos

type GetAccountsByMobileNumberRequest struct {
	MobileNumber string `uri:"mobile_number" binding:"required"`
}
