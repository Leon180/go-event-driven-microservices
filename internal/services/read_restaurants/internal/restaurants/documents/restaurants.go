package documents

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
)

type Restaurant struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	CommonCQRSHistoryModel
	Branches []Branch `bson:"branches"`
}

type Restaurants []Restaurant

type Branch struct {
	ID           string `bson:"_id"`
	RestaurantID string `bson:"restaurant_id"`
	Name         string `bson:"name"`
	Description  string `bson:"description"`
	CommonCQRSHistoryModel
	Address          Address                  `bson:"address"`
	PriceRange       PriceRange               `bson:"price_range"`
	BranchCategories []BranchCategoryRelation `bson:"branch_categories"`
	Availables       []Available              `bson:"availables"`
	Tables           []Table                  `bson:"tables"`
}

type Branches []Branch

type Address struct {
	ID          string            `bson:"_id"`
	BranchID    string            `bson:"branch_id"`
	Street      string            `bson:"street"`
	CityCode    enums.CityCode    `bson:"city_code"`
	PostalCode  string            `bson:"postal_code"`
	CountryCode enums.CountryCode `bson:"country_code"`
	CommonCQRSHistoryModel
}

type Addresses []Address

type PriceRange struct {
	ID       string `bson:"_id"`
	BranchID string `bson:"branch_id"`
	MinPrice int    `bson:"min_price"`
	MaxPrice int    `bson:"max_price"`
	CommonCQRSHistoryModel
}

type PriceRanges []PriceRange

type BranchCategoryRelation struct {
	ID         string `bson:"_id"`
	BranchID   string `bson:"branch_id"`
	CategoryID string `bson:"category_id"`
	CommonCQRSHistoryModel
	Category Category `bson:"category"`
}

type BranchCategoryRelations []BranchCategoryRelation

type Available struct {
	ID        string       `bson:"_id"`
	BranchID  string       `bson:"branch_id"`
	Weekday   time.Weekday `bson:"weekday"`
	StartTime string       `bson:"start_time"`
	EndTime   string       `bson:"end_time"`
	CommonCQRSHistoryModel
}

type Availables []Available

type Table struct {
	ID       string `bson:"_id"`
	BranchID string `bson:"branch_id"`
	Capacity int    `bson:"capacity"`
	CommonCQRSHistoryModel
}

type Tables []Table
