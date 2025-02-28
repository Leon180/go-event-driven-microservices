package dtos

type Customer struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	ActiveSwitch bool   `json:"active_switch"`
	CommonHistoryModelWithUpdate
}
