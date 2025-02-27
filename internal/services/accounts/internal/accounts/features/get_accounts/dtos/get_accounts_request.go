package featuresdtos

type GetAccountsByMobileNumberRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required"`
}
