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

func ValidateEmail(st any) error {
	if !enums.EmailFormat.ValidateFormat(st) {
		return customizeerrors.InvalidEmailError
	}
	return nil
}

func ValidateName(st string) error {
	if st == "" {
		return customizeerrors.InvalidNameError
	}
	return nil
}
