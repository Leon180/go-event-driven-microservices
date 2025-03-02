package featuresdtos

type CreateAccountRequest struct {
	// @example 0958308280
	MobileNumber string `json:"mobile_number" binding:"required"`
	// @enum "checking", "savings", "currency", "salary", "business"
	AccountType string `json:"account_type"  binding:"required"`
	// @enum "台北市中山區", "台北市松山區", "台北市信義區", "台北市文山區", "台北市北投區", "台北市南港區", "台北市萬華區"
	Branch string `json:"branch"        binding:"required"`
}
