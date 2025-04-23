package validatesdtos

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/features/take_book/dtos"
)

func ValidateTakeBookRequest(req *featuresdtos.TakeBookRequest) error {
	if req == nil {
		return nil
	}
	if req.BookID == "" {
		return customizeerrors.InvalidIDError
	}
	if !enums.MobileNumberFormat.ValidateFormat(req.MobileNumber) {
		return customizeerrors.InvalidMobileNumberError
	}
	if req.Amount <= 1 {
		return customizeerrors.BookAmountInvalidError
	}
	return nil
}
