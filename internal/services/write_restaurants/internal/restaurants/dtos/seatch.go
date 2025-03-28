package dtos

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
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
	OrderBy                 []customizegorm.OrderBy
	Pagination              *customizegorm.Pagination
}

type SearchBooks struct {
	MobileNumber *string
	TableID      *string
	AvailableID  *string

	NameFilter        *string
	NamePreciseSearch bool

	TableAvailableWeek      []time.Weekday
	TableAvailableStartTime *string
	TableAvailableEndTime   *string

	OrderBy    []customizegorm.OrderBy
	Pagination *customizegorm.Pagination
}
