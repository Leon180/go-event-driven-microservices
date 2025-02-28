package validates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/create_customer/dtos"
)

func ValidateCreateCustomerRequest(req *featuresdtos.CreateCustomerRequest) error {
	if err := validates.ValidateMobileNumber(req.MobileNumber); err != nil {
		return err
	}
	if err := validates.ValidateEmail(req.Email); err != nil {
		return err
	}
	if err := validates.ValidateName(req.FirstName); err != nil {
		return err
	}
	if err := validates.ValidateName(req.LastName); err != nil {
		return err
	}
	return nil
}
