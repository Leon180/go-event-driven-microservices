package featuresdtos

type UpdateTakedBookRequest struct {
	BookID       string `json:"book_id"`
	MobileNumber string `json:"mobile_number"`

	Amount *int    `json:"amount"`
	Note   *string `json:"note"`
}
