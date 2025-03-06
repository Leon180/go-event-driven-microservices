package dtos

type Book struct {
	ID           string `json:"id"`
	TableID      string `json:"table_id"`
	AvailableID  string `json:"available_id"`
	Amount       int    `json:"amount"`
	MobileNumber string `json:"mobile_number"`

	Table     Table     `json:"table"`
	Available Available `json:"available"`
}
