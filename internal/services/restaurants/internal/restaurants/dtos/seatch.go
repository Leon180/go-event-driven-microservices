package dtos

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

type SearchRestaurants struct {
	NameFilter        *string
	NamePreciseSearch bool

	DescriptionFilter       *string
	CityFilter              []enums.City
	CountryFilter           []enums.Country
	MaxPriceFilter          *int
	MinPriceFilter          *int
	CategoryFilter          []enums.Category
	TableAvailableWeek      []time.Weekday
	TableAvailableStartTime *string
	TableAvailableEndTime   *string
	OrderBy                 []OrderBy
	Pagination              *Pagination
}
