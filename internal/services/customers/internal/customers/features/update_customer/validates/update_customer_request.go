package validates

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/update_customer/dtos"
)

func ValidateUpdateCustomerRequest(req featuresdtos.UpdateCustomerRequest) error {
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	if req.MobileNumber != nil {
		if err := validates.ValidateMobileNumber(*req.MobileNumber); err != nil {
			return err
		}
	}
	if req.Email != nil {
		if err := validates.ValidateEmail(*req.Email); err != nil {
			return err
		}
	}
	if req.FirstName != nil {
		if err := validates.ValidateName(*req.FirstName); err != nil {
			return err
		}
	}
	if req.LastName != nil {
		if err := validates.ValidateName(*req.LastName); err != nil {
			return err
		}
	}
	return nil
}
