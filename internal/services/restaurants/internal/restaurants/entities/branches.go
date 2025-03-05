package entities

import "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"

type Branch struct {
	ID           string `gorm:"primaryKey;type:uuid" comment:"ID"`
	RestaurantID string `gorm:"not null;type:uuid" comment:"Restaurant ID"`
	Name         string `gorm:"not null;type:varchar(255)" comment:"Name"`
	Description  string `gorm:"not null;type:text" comment:"Description"`
	CommonCQRSHistoryModel
}

type Branches []Branch

type UpdateBranch struct {
	ID           string
	Name         *string
	Description  *string
	ActiveStatus *bool
}

func (u *UpdateBranch) RemoveUnchangedFields(branch Branch) {
	if u.ID != branch.ID {
		return
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

type UpdateBranches []UpdateBranch

type Address struct {
	ID          string            `gorm:"primaryKey;type:uuid" comment:"ID"`
	BranchID    string            `gorm:"not null;type:uuid" comment:"Branch ID"`
	Street      string            `gorm:"not null;type:varchar(255)" comment:"Street"`
	CityCode    enums.CityCode    `gorm:"not null;type:int" comment:"City Code"`
	PostalCode  string            `gorm:"not null;type:varchar(255)" comment:"Postal Code"`
	CountryCode enums.CountryCode `gorm:"not null;type:varchar(255)" comment:"Country"`
	CommonCQRSHistoryModel
}

type Addresses []Address

type UpdateAddress struct {
	ID          string
	Street      *string
	CityCode    *enums.CityCode
	PostalCode  *string
	CountryCode *enums.CountryCode
}

func (u *UpdateAddress) RemoveUnchangedFields(address Address) {
	if u.ID != address.ID {
		return
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
	return updateMap
}

type UpdateAddresses []UpdateAddress

type PriceRange struct {
	ID       string `gorm:"primaryKey;type:uuid" comment:"ID"`
	BranchID string `gorm:"not null;type:uuid" comment:"Branch ID"`
	MinPrice int    `gorm:"not null;type:int" comment:"Min Price"`
	MaxPrice int    `gorm:"not null;type:int" comment:"Max Price"`
	CommonCQRSHistoryModel
}

type PriceRanges []PriceRange

type UpdatePriceRange struct {
	ID       string
	MinPrice *int
	MaxPrice *int
}

func (u *UpdatePriceRange) RemoveUnchangedFields(priceRange PriceRange) {
	if u.ID != priceRange.ID {
		return
	}
	if u.MinPrice != nil && *u.MinPrice == priceRange.MinPrice {
		u.MinPrice = nil
	}
	if u.MaxPrice != nil && *u.MaxPrice == priceRange.MaxPrice {
		u.MaxPrice = nil
	}
}

func (u *UpdatePriceRange) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.MinPrice != nil {
		updateMap["min_price"] = *u.MinPrice
	}
	if u.MaxPrice != nil {
		updateMap["max_price"] = *u.MaxPrice
	}
	return updateMap
}

type UpdatePriceRanges []UpdatePriceRange

type BranchCategoryRelation struct {
	ID         string `gorm:"primaryKey;type:uuid" comment:"ID"`
	BranchID   string `gorm:"not null;type:uuid" comment:"Branch ID"`
	CategoryID string `gorm:"not null;type:uuid" comment:"Category ID"`
	CommonCQRSHistoryModel
}

type BranchCategoryRelations []BranchCategoryRelation

type UpdateBranchCategoryRelation struct {
	ID         string
	CategoryID *string
}

func (u *UpdateBranchCategoryRelation) RemoveUnchangedFields(branchCategoryRelation BranchCategoryRelation) {
	if u.ID != branchCategoryRelation.ID {
		return
	}
	if u.CategoryID != nil && *u.CategoryID == branchCategoryRelation.CategoryID {
		u.CategoryID = nil
	}
}

func (u *UpdateBranchCategoryRelation) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.CategoryID != nil {
		updateMap["category_id"] = *u.CategoryID
	}
	return updateMap
}

type UpdateBranchCategoryRelations []UpdateBranchCategoryRelation

type Category struct {
	ID           string             `gorm:"primaryKey;type:uuid" comment:"ID"`
	CategoryCode enums.CategoryCode `gorm:"not null;type:int" comment:"Category Code"`
	CommonCQRSHistoryModel
}

type Categories []Category
