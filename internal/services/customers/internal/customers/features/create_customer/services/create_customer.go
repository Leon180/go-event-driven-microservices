package services

import (
	"context"
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/create_customer/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/create_customer/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories"
)

type CreateCustomer interface {
	CreateCustomer(ctx context.Context, req *featuresdtos.CreateCustomerRequest) error
}

func NewCreateCustomer(
	createCustomerRepository repositories.CreateCustomer,
	readCustomerByMobileNumberAndActiveSwitch repositories.ReadCustomerByMobileNumberAndActiveSwitch,
	uuidGenerator uuid.UUIDGenerator,
) CreateCustomer {
	return &createCustomerImpl{
		createCustomerRepository:                  createCustomerRepository,
		readCustomerByMobileNumberAndActiveSwitch: readCustomerByMobileNumberAndActiveSwitch,
		uuidGenerator:                             uuidGenerator,
	}
}

type createCustomerImpl struct {
	createCustomerRepository                  repositories.CreateCustomer
	readCustomerByMobileNumberAndActiveSwitch repositories.ReadCustomerByMobileNumberAndActiveSwitch
	uuidGenerator                             uuid.UUIDGenerator
}

func (handle *createCustomerImpl) CreateCustomer(ctx context.Context, req *featuresdtos.CreateCustomerRequest) error {
	if req == nil {
		return nil
	}
	if err := validates.ValidateCreateCustomerRequest(req); err != nil {
		return err
	}

	customers, err := handle.readCustomerByMobileNumberAndActiveSwitch.ReadCustomerByMobileNumberAndActiveSwitch(
		ctx,
		req.MobileNumber,
		nil,
	)
	if err != nil {
		return err
	}
	if len(customers) > 0 {
		for _, customer := range customers {
			if customer.IsActive() {
				return customizeerrors.CustomerAlreadyExistsError
			}
		}
		return customizeerrors.CustomerExistsButInactiveError
	}

	systemTime := time.Now()
	customerEntity := entities.Customer{
		ID:           handle.uuidGenerator.GenerateUUID(),
		MobileNumber: req.MobileNumber,
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		ActiveSwitch: true,
		CommonHistoryModelWithUpdate: entities.CommonHistoryModelWithUpdate{
			CommonHistoryModel: entities.CommonHistoryModel{
				CreatedAt: systemTime,
			},
			UpdatedAt: systemTime,
		},
	}

	return handle.createCustomerRepository.CreateCustomer(ctx, &customerEntity)
}
