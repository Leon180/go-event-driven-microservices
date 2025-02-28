package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/delete_customer/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories"
)

type DeleteCustomer interface {
	DeleteCustomer(ctx context.Context, req *featuresdtos.DeleteCustomerRequest) error
}

func NewDeleteCustomer(
	readCustomerRepository repositories.ReadCustomer,
	updateCustomerByIDRepository repositories.UpdateCustomerByID,
) DeleteCustomer {
	return &deleteCustomerImpl{
		readCustomerRepository:       readCustomerRepository,
		updateCustomerByIDRepository: updateCustomerByIDRepository,
	}
}

type deleteCustomerImpl struct {
	readCustomerRepository       repositories.ReadCustomer
	updateCustomerByIDRepository repositories.UpdateCustomerByID
}

func (handle *deleteCustomerImpl) DeleteCustomer(ctx context.Context, req *featuresdtos.DeleteCustomerRequest) error {
	if req == nil {
		return nil
	}
	if req.ID == "" {
		return customizeerrors.InvalidIDError
	}
	customer, err := handle.readCustomerRepository.ReadCustomer(ctx, req.ID)
	if err != nil {
		return err
	}
	if customer == nil {
		return customizeerrors.CustomerNotFoundError
	}
	if !customer.IsActive() {
		return customizeerrors.CustomerAlreadyDeletedError
	}
	activeSwitch := false
	updateCustomer := entities.UpdateCustomer{
		ID:           customer.ID,
		ActiveSwitch: &activeSwitch,
	}
	return handle.updateCustomerByIDRepository.UpdateCustomerByID(ctx, updateCustomer)
}
