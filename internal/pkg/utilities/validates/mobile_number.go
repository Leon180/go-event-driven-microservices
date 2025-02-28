package validates

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/shopspring/decimal"
)

func ValidateMobileNumber(st any) error {
	if !enums.MobileNumberFormat.ValidateFormat(st) {
		return customizeerrors.InvalidMobileNumberError
	}
	return nil
}

func ValidateDecimal(st any) error {
	var v decimal.Decimal
	if err := v.Scan(st); err != nil {
		return customizeerrors.InvalidDecimalError
	}
	return nil
}
