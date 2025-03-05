package dtos

import "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"

type Branch struct {
	ID          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`

	Address    *Address    `json:"address"`
	PriceRange *PriceRange `json:"priceRange"`
	Categories []Category  `json:"categories"`

	Tables []Table `json:"tables"`
}

type Address struct {
	ID         *string       `json:"id,omitempty"`
	Street     string        `json:"street"`
	City       enums.City    `json:"city"`
	PostalCode string        `json:"postalCode"`
	Country    enums.Country `json:"country"`
}

type PriceRange struct {
	ID       *string `json:"id,omitempty"`
	MinPrice int     `json:"minPrice"`
	MaxPrice int     `json:"maxPrice"`
}

type Category struct {
	ID       string         `json:"id"`
	Category enums.Category `json:"category"`
	CommonCQRSHistoryModel
}
