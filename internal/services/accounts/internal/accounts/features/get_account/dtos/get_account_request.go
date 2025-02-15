package dtos

type GetAccountRequest struct {
	MobileNumber string `form:"mobile_number" binding:"required"`
}
