package dtos

import (
	"time"
)

type Table struct {
	ID       *string `json:"id,omitempty"`
	Capacity int     `json:"capacity"`

	TableAvailables []TableAvailable `json:"tableAvailables"`
}

type TableAvailable struct {
	ID        *string      `json:"id,omitempty"`
	Weekday   time.Weekday `json:"weekday"`
	StartTime string       `json:"startTime"` // HH:MM
	EndTime   string       `json:"endTime"`   // HH:MM
	Booked    bool         `json:"booked"`
}
