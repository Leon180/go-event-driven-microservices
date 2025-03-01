package featuresdtos

type GetLoansRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required"`
	ActiveSwitch *bool  `json:"active_switch"`
}
