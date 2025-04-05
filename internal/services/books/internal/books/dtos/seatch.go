package dtos

import (
	"time"

	utilitiesdb "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/db"
)

type SearchBooks struct {
	MobileNumber *string
	TableID      *string
	AvailableID  *string

	NameFilter        *string
	NamePreciseSearch bool

	TableAvailableWeek      []time.Weekday
	TableAvailableStartTime *string
	TableAvailableEndTime   *string

	OrderBy    []utilitiesdb.OrderBy
	Pagination *utilitiesdb.Pagination
}
