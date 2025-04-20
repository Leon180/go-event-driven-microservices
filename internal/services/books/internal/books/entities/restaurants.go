package entities

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

type Restaurant struct {
	ID          string `gorm:"primaryKey;type:varchar(255)" comment:"ID"`
	Name        string `gorm:"not null;type:varchar(255)"   comment:"Name"`
	Description string `gorm:"not null;type:text"           comment:"Description"`
	CommonCQRSHistoryModel
}

type Restaurants []Restaurant

type Branch struct {
	ID           string `gorm:"primaryKey;type:varchar(255)" comment:"ID"`
	RestaurantID string `gorm:"not null;type:varchar(255)"   comment:"Restaurant ID"`
	Name         string `gorm:"not null;type:varchar(255)"   comment:"Name"`
	Description  string `gorm:"not null;type:text"           comment:"Description"`
	CommonCQRSHistoryModel
}

type Branches []Branch

type Address struct {
	ID          string            `gorm:"primaryKey;type:varchar(255)" comment:"ID"`
	BranchID    string            `gorm:"not null;type:varchar(255)"   comment:"Branch ID"`
	Street      string            `gorm:"not null;type:varchar(255)"   comment:"Street"`
	CityCode    enums.CityCode    `gorm:"not null;type:int"            comment:"City Code"`
	PostalCode  string            `gorm:"not null;type:varchar(255)"   comment:"Postal Code"`
	CountryCode enums.CountryCode `gorm:"not null;type:int"            comment:"Country"`
	CommonCQRSHistoryModel
}

type Addresses []Address

type PriceRange struct {
	ID       string `gorm:"primaryKey;type:varchar(255)" comment:"ID"`
	BranchID string `gorm:"not null;type:varchar(255)"   comment:"Branch ID"`
	MinPrice int    `gorm:"not null;type:int"            comment:"Min Price"`
	MaxPrice int    `gorm:"not null;type:int"            comment:"Max Price"`
	CommonCQRSHistoryModel
}

type PriceRanges []PriceRange

type BranchCategoryRelation struct {
	ID         string `gorm:"primaryKey;type:varchar(255)" comment:"ID"`
	BranchID   string `gorm:"not null;type:varchar(255)"   comment:"Branch ID"`
	CategoryID string `gorm:"not null;type:varchar(255)"   comment:"Category ID"`
	CommonCQRSHistoryModel
}

type BranchCategoryRelations []BranchCategoryRelation

type Available struct {
	ID        string       `gorm:"primaryKey;type:varchar(255)" comment:"ID"`
	BranchID  string       `gorm:"not null;type:varchar(255)"   comment:"Branch ID"`
	Weekday   time.Weekday `gorm:"not null;type:int"            comment:"Weekday"`
	StartTime string       `gorm:"not null;type:varchar(5)"     comment:"Start Time HH:MM"`
	EndTime   string       `gorm:"not null;type:varchar(5)"     comment:"End Time HH:MM"`
	CommonCQRSHistoryModel
}

type Availables []Available

type Table struct {
	ID       string `gorm:"primaryKey;type:varchar(255)" comment:"ID"`
	BranchID string `gorm:"not null;type:varchar(255)"   comment:"Branch ID"`
	Capacity int    `gorm:"not null;type:int"            comment:"Capacity"`
	CommonCQRSHistoryModel
}

type Tables []Table

type UpdateRestaurant struct {
	ID           string
	Name         *string
	Description  *string
	ActiveStatus *bool
}

func (u *UpdateRestaurant) RemoveUnchangedFields(restaurant Restaurant) *UpdateRestaurant {
	if u == nil {
		return nil
	}
	if u.ID != restaurant.ID {
		return nil
	}
	if u.Name != nil && *u.Name == restaurant.Name {
		u.Name = nil
	}
	if u.Description != nil && *u.Description == restaurant.Description {
		u.Description = nil
	}
	if u.ActiveStatus != nil && *u.ActiveStatus == restaurant.ActiveStatus {
		u.ActiveStatus = nil
	}
	return u
}

func (u *UpdateRestaurant) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.Name != nil {
		updateMap["name"] = *u.Name
	}
	if u.Description != nil {
		updateMap["description"] = *u.Description
	}
	if u.ActiveStatus != nil {
		updateMap["active_status"] = *u.ActiveStatus
	}
	return updateMap
}

type UpdateBranch struct {
	ID           string
	Name         *string
	Description  *string
	ActiveStatus *bool
}

func (u *UpdateBranch) RemoveUnchangedFields(branch Branch) *UpdateBranch {
	if u == nil {
		return nil
	}
	if u.ID != branch.ID {
		return nil
	}
	if u.Name != nil && *u.Name == branch.Name {
		u.Name = nil
	}
	if u.Description != nil && *u.Description == branch.Description {
		u.Description = nil
	}
	if u.ActiveStatus != nil && *u.ActiveStatus == branch.ActiveStatus {
		u.ActiveStatus = nil
	}
	return u
}

func (u *UpdateBranch) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.Name != nil {
		updateMap["name"] = *u.Name
	}
	if u.Description != nil {
		updateMap["description"] = *u.Description
	}
	if u.ActiveStatus != nil {
		updateMap["active_status"] = *u.ActiveStatus
	}
	return updateMap
}

type UpdateAddress struct {
	ID           string
	Street       *string
	CityCode     *enums.CityCode
	PostalCode   *string
	CountryCode  *enums.CountryCode
	ActiveStatus *bool
}

func (u *UpdateAddress) RemoveUnchangedFields(address Address) *UpdateAddress {
	if u == nil {
		return nil
	}
	if u.ID != address.ID {
		return nil
	}
	if u.Street != nil && *u.Street == address.Street {
		u.Street = nil
	}
	if u.CityCode != nil && *u.CityCode == address.CityCode {
		u.CityCode = nil
	}
	if u.PostalCode != nil && *u.PostalCode == address.PostalCode {
		u.PostalCode = nil
	}
	if u.CountryCode != nil && *u.CountryCode == address.CountryCode {
		u.CountryCode = nil
	}
	if u.ActiveStatus != nil && *u.ActiveStatus == address.ActiveStatus {
		u.ActiveStatus = nil
	}
	return u
}

func (u *UpdateAddress) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.Street != nil {
		updateMap["street"] = *u.Street
	}
	if u.CityCode != nil {
		updateMap["city_code"] = *u.CityCode
	}
	if u.PostalCode != nil {
		updateMap["postal_code"] = *u.PostalCode
	}
	if u.CountryCode != nil {
		updateMap["country_code"] = *u.CountryCode
	}
	if u.ActiveStatus != nil {
		updateMap["active_status"] = *u.ActiveStatus
	}
	return updateMap
}

type UpdatePriceRange struct {
	ID           string
	MinPrice     *int
	MaxPrice     *int
	ActiveStatus *bool
}

func (u *UpdatePriceRange) RemoveUnchangedFields(priceRange PriceRange) *UpdatePriceRange {
	if u == nil {
		return nil
	}
	if u.ID != priceRange.ID {
		return nil
	}
	if u.MinPrice != nil && *u.MinPrice == priceRange.MinPrice {
		u.MinPrice = nil
	}
	if u.MaxPrice != nil && *u.MaxPrice == priceRange.MaxPrice {
		u.MaxPrice = nil
	}
	if u.ActiveStatus != nil && *u.ActiveStatus == priceRange.ActiveStatus {
		u.ActiveStatus = nil
	}
	return u
}

func (u *UpdatePriceRange) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.MinPrice != nil {
		updateMap["min_price"] = *u.MinPrice
	}
	if u.MaxPrice != nil {
		updateMap["max_price"] = *u.MaxPrice
	}
	if u.ActiveStatus != nil {
		updateMap["active_status"] = *u.ActiveStatus
	}
	return updateMap
}

type UpdateBranchCategoryRelation struct {
	ID           string
	CategoryID   *string
	ActiveStatus *bool
}

func (u *UpdateBranchCategoryRelation) RemoveUnchangedFields(
	branchCategoryRelation BranchCategoryRelation,
) *UpdateBranchCategoryRelation {
	if u == nil {
		return nil
	}
	if u.ID != branchCategoryRelation.ID {
		return nil
	}
	if u.CategoryID != nil && *u.CategoryID == branchCategoryRelation.CategoryID {
		u.CategoryID = nil
	}
	if u.ActiveStatus != nil && *u.ActiveStatus == branchCategoryRelation.ActiveStatus {
		u.ActiveStatus = nil
	}
	return u
}

func (u *UpdateBranchCategoryRelation) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.CategoryID != nil {
		updateMap["category_id"] = *u.CategoryID
	}
	if u.ActiveStatus != nil {
		updateMap["active_status"] = *u.ActiveStatus
	}
	return updateMap
}

type UpdateTable struct {
	ID           string
	Capacity     *int
	ActiveStatus *bool
}

func (u *UpdateTable) RemoveUnchangedFields(table Table) *UpdateTable {
	if u == nil {
		return nil
	}
	if u.ID != table.ID {
		return nil
	}
	if u.Capacity != nil && *u.Capacity == table.Capacity {
		u.Capacity = nil
	}
	if u.ActiveStatus != nil && *u.ActiveStatus == table.ActiveStatus {
		u.ActiveStatus = nil
	}
	return u
}

func (u *UpdateTable) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.Capacity != nil {
		updateMap["capacity"] = *u.Capacity
	}
	if u.ActiveStatus != nil {
		updateMap["active_status"] = *u.ActiveStatus
	}
	return updateMap
}

type UpdateAvailable struct {
	ID           string
	Weekday      *time.Weekday
	StartTime    *string
	EndTime      *string
	ActiveStatus *bool
}

func (u *UpdateAvailable) RemoveUnchangedFields(available Available) *UpdateAvailable {
	if u == nil {
		return nil
	}
	if u.ID != available.ID {
		return nil
	}
	if u.Weekday != nil && *u.Weekday == available.Weekday {
		u.Weekday = nil
	}
	if u.StartTime != nil && *u.StartTime == available.StartTime {
		u.StartTime = nil
	}
	if u.EndTime != nil && *u.EndTime == available.EndTime {
		u.EndTime = nil
	}
	if u.ActiveStatus != nil && *u.ActiveStatus == available.ActiveStatus {
		u.ActiveStatus = nil
	}
	return u
}

func (u *UpdateAvailable) ToUpdateMap() map[string]any {
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
	if u.ActiveStatus != nil {
		updateMap["active_status"] = *u.ActiveStatus
	}
	return updateMap
}
