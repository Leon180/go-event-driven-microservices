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

type Tables []Table

type UpdateTable struct {
	ID       string
	Capacity *int
}

func (u *UpdateTable) RemoveUnchangedFields(table Table) {
	if u.ID != table.ID {
		return
	}
	if u.Capacity != nil && *u.Capacity == table.Capacity {
		u.Capacity = nil
	}
}

func (u *UpdateTable) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.Capacity != nil {
		updateMap["capacity"] = *u.Capacity
	}
	return updateMap
}

type UpdateTables []UpdateTable

type TableAvailable struct {
	ID        string       `gorm:"primaryKey;type:uuid" comment:"ID"`
	TableID   string       `gorm:"not null;type:uuid" comment:"Table ID"`
	Weekday   time.Weekday `gorm:"not null;type:int" comment:"Weekday"`
	StartTime string       `gorm:"not null;type:varchar(5)" comment:"Start Time HH:MM"`
	EndTime   string       `gorm:"not null;type:varchar(5)" comment:"End Time HH:MM"`
	Booked    bool         `gorm:"not null;type:boolean" comment:"Booked"`
	CommonCQRSHistoryModel
}

type TableAvailables []TableAvailable

type UpdateTableAvailable struct {
	ID        string
	Weekday   *time.Weekday
	StartTime *string
	EndTime   *string
	Booked    *bool
}

func (u *UpdateTableAvailable) RemoveUnchangedFields(tableAvailable TableAvailable) {
	if u.ID != tableAvailable.ID {
		return
	}
	if u.Weekday != nil && *u.Weekday == tableAvailable.Weekday {
		u.Weekday = nil
	}
	if u.StartTime != nil && *u.StartTime == tableAvailable.StartTime {
		u.StartTime = nil
	}
	if u.EndTime != nil && *u.EndTime == tableAvailable.EndTime {
		u.EndTime = nil
	}
	if u.Booked != nil && *u.Booked == tableAvailable.Booked {
		u.Booked = nil
	}
}

func (u *UpdateTableAvailable) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.Weekday != nil {
		updateMap["weekday"] = *u.Weekday
	}
	if u.StartTime != nil {
		updateMap["start_time"] = *u.StartTime
	}
	if u.EndTime != nil {
		updateMap["end_time"] = *u.EndTime
	}
	if u.Booked != nil {
		updateMap["booked"] = *u.Booked
	}
	return updateMap
}

type UpdateTableAvailables []UpdateTableAvailable
