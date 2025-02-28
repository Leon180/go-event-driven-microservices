package featuresdtos

type UpdateCustomerRequest struct {
	ID           string  `json:"id" binding:"required"`
	MobileNumber *string `json:"mobile_number"`
	Email        *string `json:"email"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
}
