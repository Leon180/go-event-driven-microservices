package dtos

import "github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"

type Book struct {
	ID           *string    `json:"id"`
	TableID      string     `json:"table_id"`
	AvailableID  string     `json:"available_id"`
	Amount       int        `json:"amount"`
	MobileNumber string     `json:"mobile_number"`
	Table        *Table     `json:"table"`
	Available    *Available `json:"available"`
	CommonCQRSHistoryModel
}

func (r *Book) ToUpdateBook() *entities.UpdateBook {
	if r.ID == nil {
		return nil
	}
	return &entities.UpdateBook{
		ID:           *r.ID,
		TableID:      r.TableID,
		AvailableID:  r.AvailableID,
		Amount:       &r.Amount,
		MobileNumber: &r.MobileNumber,
	}
}
