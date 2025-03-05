package dtos

type Branch struct {
	ID          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`

	Address    *Address    `json:"address"`
	PriceRange *PriceRange `json:"priceRange"`
	Categories []Category  `json:"categories"`

	Tables []Table `json:"tables"`
}

type Address struct {
	ID         *string `json:"id,omitempty"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	PostalCode string  `json:"postalCode"`
	Country    string  `json:"country"`
}

type PriceRange struct {
	ID       *string `json:"id,omitempty"`
	MinPrice int     `json:"minPrice"`
	MaxPrice int     `json:"maxPrice"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	CommonCQRSHistoryModel
}
