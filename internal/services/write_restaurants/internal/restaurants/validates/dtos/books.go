package validatesdtos

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/dtos"
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
	if book.Table != nil {
		if err := ValidateTable(book.Table); err != nil {
			return err
		}
	}
	if book.Available != nil {
		if err := ValidateAvailable(book.Available); err != nil {
			return err
		}
	}
	return nil
}
