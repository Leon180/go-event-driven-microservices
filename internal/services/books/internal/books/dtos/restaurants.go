package dtos

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
)

type Restaurant struct {
	ID          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`

	ActiveStatus bool `json:"active_status"`

	Branches []Branch `json:"branches"`
}

func (r *Restaurant) ToUpdateRestaurant() *entities.UpdateRestaurant {
	if r.ID == nil {
		return nil
	}
	return &entities.UpdateRestaurant{
		ID:          *r.ID,
		Name:        &r.Name,
		Description: &r.Description,
	}
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

func (b *Branch) ToUpdateBranch() *entities.UpdateBranch {
	if b.ID == nil {
		return nil
	}
	return &entities.UpdateBranch{
		ID:          *b.ID,
		Name:        &b.Name,
		Description: &b.Description,
	}
}

type Address struct {
	ID         *string       `json:"id,omitempty"`
	Street     string        `json:"street"`
	City       enums.City    `json:"city"`
	PostalCode string        `json:"postal_code"`
	Country    enums.Country `json:"country"`
}

func (a *Address) ToUpdateAddress() *entities.UpdateAddress {
	if a.ID == nil {
		return nil
	}
	cityCode := a.City.ToCityCode()
	countryCode := a.Country.ToCountryCode()
	return &entities.UpdateAddress{
		ID:          *a.ID,
		Street:      &a.Street,
		CityCode:    &cityCode,
		PostalCode:  &a.PostalCode,
		CountryCode: &countryCode,
	}
}

type PriceRange struct {
	ID       *string `json:"id,omitempty"`
	MinPrice int     `json:"min_price"`
	MaxPrice int     `json:"max_price"`
}

func (p *PriceRange) ToUpdatePriceRange() *entities.UpdatePriceRange {
	if p.ID == nil {
		return nil
	}
	return &entities.UpdatePriceRange{
		ID:       *p.ID,
		MinPrice: &p.MinPrice,
		MaxPrice: &p.MaxPrice,
	}
}

type Available struct {
	ID        *string      `json:"id,omitempty"`
	Weekday   time.Weekday `json:"weekday"`
	StartTime string       `json:"start_time"` // HH:MM
	EndTime   string       `json:"end_time"`   // HH:MM
}

func (a *Available) ToUpdateAvailable() *entities.UpdateAvailable {
	if a.ID == nil {
		return nil
	}
	return &entities.UpdateAvailable{
		ID:        *a.ID,
		Weekday:   &a.Weekday,
		StartTime: &a.StartTime,
		EndTime:   &a.EndTime,
	}
}

type Table struct {
	ID       *string `json:"id,omitempty"`
	Capacity int     `json:"capacity"`
}

func (t *Table) ToUpdateTable() *entities.UpdateTable {
	if t.ID == nil {
		return nil
	}
	return &entities.UpdateTable{
		ID:       *t.ID,
		Capacity: &t.Capacity,
	}
}
