package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/update_customer/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/update_customer/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories"
)

type UpdateCustomer interface {
	UpdateCustomer(ctx context.Context, req *featuresdtos.UpdateCustomerRequest) error
}

type updateCustomerImpl struct {
	readCustomer                 repositories.ReadCustomer
	updateCustomerByIDRepository repositories.UpdateCustomerByID
}

func NewUpdateCustomer(
	readCustomer repositories.ReadCustomer,
	updateCustomerByIDRepository repositories.UpdateCustomerByID,
) UpdateCustomer {
	return &updateCustomerImpl{
		readCustomer:                 readCustomer,
		updateCustomerByIDRepository: updateCustomerByIDRepository,
	}
}

func (handle *updateCustomerImpl) UpdateCustomer(ctx context.Context, req *featuresdtos.UpdateCustomerRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateUpdateCustomerRequest(*req); err != nil {
		return err
	}
	customer, err := handle.readCustomer.ReadCustomer(ctx, req.ID)
	if err != nil {
		return err
	}
	if customer == nil {
		return customizeerrors.CustomerNotFoundError
	}
	updateCustomer := entities.UpdateCustomer{
		ID:           req.ID,
		MobileNumber: req.MobileNumber,
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
	}
	updateCustomer.RemoveUnchangedFields(*customer)
	if updateCustomer.NoUpdates() {
		return customizeerrors.CustomerNoUpdatesError
	}
	return handle.updateCustomerByIDRepository.UpdateCustomerByID(ctx, updateCustomer)
}
