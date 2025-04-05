package validatesdtos

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/dtos"
)

func ValidateBook(book *dtos.Book) error {
	if book == nil {
		return nil
	}
	if book.TableID == "" || book.AvailableID == "" {
		return customizeerrors.InvalidIDError
	}
	if book.Amount <= 1 {
		return customizeerrors.BookAmountInvalidError
	}
	if !enums.MobileNumberFormat.ValidateFormat(book.MobileNumber) {
		return customizeerrors.InvalidMobileNumberError
	}
	return nil
}
