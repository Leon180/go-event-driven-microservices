package dtos

type Restaurant struct {
	ID          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`

	Branches []Branch `json:"branches"`
}
