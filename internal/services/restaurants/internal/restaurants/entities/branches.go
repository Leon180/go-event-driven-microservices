package entities

type Branch struct {
	ID           string `gorm:"primaryKey;type:uuid" comment:"ID"`
	RestaurantID string `gorm:"not null;type:uuid" comment:"Restaurant ID"`
	Name         string `gorm:"not null;type:varchar(255)" comment:"Name"`
	Description  string `gorm:"not null;type:text" comment:"Description"`
	CommonCQRSHistoryModel
}

type Address struct {
	ID         string `gorm:"primaryKey;type:uuid" comment:"ID"`
	BranchID   string `gorm:"not null;type:uuid" comment:"Branch ID"`
	Street     string `gorm:"not null;type:varchar(255)" comment:"Street"`
	City       string `gorm:"not null;type:varchar(255)" comment:"City or District"`
	PostalCode string `gorm:"not null;type:varchar(255)" comment:"Postal Code"`
	Country    string `gorm:"not null;type:varchar(255)" comment:"Country"`
	CommonCQRSHistoryModel
}

type PriceRange struct {
	ID       string `gorm:"primaryKey;type:uuid" comment:"ID"`
	BranchID string `gorm:"not null;type:uuid" comment:"Branch ID"`
	MinPrice int    `gorm:"not null;type:int" comment:"Min Price"`
	MaxPrice int    `gorm:"not null;type:int" comment:"Max Price"`
	CommonCQRSHistoryModel
}

type BranchCategoryRelation struct {
	ID         string `gorm:"primaryKey;type:uuid" comment:"ID"`
	BranchID   string `gorm:"not null;type:uuid" comment:"Branch ID"`
	CategoryID string `gorm:"not null;type:uuid" comment:"Category ID"`
	CommonCQRSHistoryModel
}

type Category struct {
	ID   string `gorm:"primaryKey;type:uuid" comment:"ID"`
	Name string `gorm:"not null;type:varchar(255)" comment:"Name"`
	CommonCQRSHistoryModel
}
