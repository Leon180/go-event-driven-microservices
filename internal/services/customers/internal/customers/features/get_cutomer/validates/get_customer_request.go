package validates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/get_cutomer/dtos"
)

func ValidateGetCustomerRequest(req *featuresdtos.GetCustomerRequest) error {
	return validates.ValidateMobileNumber(req.MobileNumber)
}
