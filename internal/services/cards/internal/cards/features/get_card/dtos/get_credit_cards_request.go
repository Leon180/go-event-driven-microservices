package featuresdtos

type GetCreditCardsRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required"`
	ActiveSwitch *bool  `json:"active_switch"`
}
