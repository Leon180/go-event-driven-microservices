package featuresdtos

type CreateCustomerRequest struct {
	// @example John
	FirstName string `json:"first_name"    binding:"required"`
	// @example Doe
	LastName string `json:"last_name"     binding:"required"`
	// @example john.doe@example.com
	Email string `json:"email"         binding:"required"`
	// @example 0958308280
	MobileNumber string `json:"mobile_number" binding:"required"`
}
