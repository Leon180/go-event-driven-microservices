package dtos

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

type Restaurant struct {
	ID          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`

	Branches []Branch `json:"branches"`
}

type Branch struct {
	ID          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`

	Address    *Address    `json:"address"`
	PriceRange *PriceRange `json:"price_range"`
	Categories []Category  `json:"categories"`
	Availables []Available `json:"availables"`
	Tables     []Table     `json:"tables"`
}

type Address struct {
	ID         *string       `json:"id,omitempty"`
	Street     string        `json:"street"`
	City       enums.City    `json:"city"`
	PostalCode string        `json:"postal_code"`
	Country    enums.Country `json:"country"`
}

type PriceRange struct {
	ID       *string `json:"id,omitempty"`
	MinPrice int     `json:"min_price"`
	MaxPrice int     `json:"max_price"`
}

type Category struct {
	ID       string         `json:"id"`
	Category enums.Category `json:"category"`
	CommonCQRSHistoryModel
}

type Available struct {
	ID        *string      `json:"id,omitempty"`
	Weekday   time.Weekday `json:"weekday"`
	StartTime string       `json:"start_time"` // HH:MM
	EndTime   string       `json:"end_time"`   // HH:MM
}

type Table struct {
	ID       *string `json:"id,omitempty"`
	Capacity int     `json:"capacity"`
}
