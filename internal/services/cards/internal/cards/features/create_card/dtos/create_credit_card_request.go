package featuresdtos

type CreateCreditCardRequest struct {
	// @example 0958308280
	MobileNumber string `json:"mobile_number" binding:"required"`
	// @example 100000
	TotalLimit string `json:"total_limit"   binding:"required"`
}
