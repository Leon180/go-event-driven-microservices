package featuresdtos

type GetCustomerRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required"`
	ActiveSwitch *bool  `json:"active_switch"`
}
