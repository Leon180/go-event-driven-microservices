package validates

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/validates"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/get_loans/dtos"
)

func ValidateGetLoansRequest(req *featuresdtos.GetLoansRequest) error {
	return validates.ValidateMobileNumber(req.MobileNumber)
}
