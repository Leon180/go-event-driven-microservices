package featuresdtos

type UpdateCreditCardRequest struct {
	ID           string  `json:"id"            binding:"required"`
	MobileNumber *string `json:"mobile_number"`
	TotalLimit   *string `json:"total_limit"`
	AmountUsed   *string `json:"amount_used"`
}
