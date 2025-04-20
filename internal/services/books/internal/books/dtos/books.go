package dtos

import "time"

type Book struct {
	ID           string    `json:"id"`
	BranchID     string    `json:"branch_id"`
	Date         time.Time `json:"date"`
	StartTime    string    `json:"start_time"`
	EndTime      string    `json:"end_time"`
	Capacity     int       `json:"capacity"`
	Booked       bool      `json:"booked"`
	Amount       int       `json:"amount"`
	MobileNumber string    `json:"mobile_number"`
	Note         string    `json:"note"`
}
