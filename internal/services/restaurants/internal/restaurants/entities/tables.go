package entities

import (
	"time"
)

type Table struct {
	ID       string `gorm:"primaryKey;type:uuid" comment:"ID"`
	BranchID string `gorm:"not null;type:uuid" comment:"Branch ID"`
	Capacity int    `gorm:"not null;type:int" comment:"Capacity"`
	CommonCQRSHistoryModel
}

type TableAvailable struct {
	ID        string       `gorm:"primaryKey;type:uuid" comment:"ID"`
	TableID   string       `gorm:"not null;type:uuid" comment:"Table ID"`
	Weekday   time.Weekday `gorm:"not null;type:int" comment:"Weekday"`
	StartTime string       `gorm:"not null;type:varchar(5)" comment:"Start Time HH:MM"`
	EndTime   string       `gorm:"not null;type:varchar(5)" comment:"End Time HH:MM"`
	Booked    bool         `gorm:"not null;type:boolean" comment:"Booked"`
	CommonCQRSHistoryModel
}
