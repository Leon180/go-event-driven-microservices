package services

import (
	"context"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/get_cutomer/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/get_cutomer/validates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories"
)

type GetCustomerByMobileNumberAndActiveSwitch interface {
	GetCustomerByMobileNumberAndActiveSwitch(
		ctx context.Context,
		req *featuresdtos.GetCustomerRequest,
	) (entities.Customers, error)
}

func NewGetCustomerByMobileNumberAndActiveSwitch(
	readCustomerByMobileNumberAndActiveSwitchRepository repositories.ReadCustomerByMobileNumberAndActiveSwitch,
) GetCustomerByMobileNumberAndActiveSwitch {
	return &getCustomerByMobileNumberAndActiveSwitchImpl{
		readCustomerByMobileNumberAndActiveSwitchRepository: readCustomerByMobileNumberAndActiveSwitchRepository,
	}
}

type getCustomerByMobileNumberAndActiveSwitchImpl struct {
	readCustomerByMobileNumberAndActiveSwitchRepository repositories.ReadCustomerByMobileNumberAndActiveSwitch
}

func (handle *getCustomerByMobileNumberAndActiveSwitchImpl) GetCustomerByMobileNumberAndActiveSwitch(
	ctx context.Context,
	req *featuresdtos.GetCustomerRequest,
) (entities.Customers, error) {
	if req == nil {
		return nil, nil
	}
	if err := validates.ValidateGetCustomerRequest(req); err != nil {
		return nil, err
	}
	customers, err := handle.readCustomerByMobileNumberAndActiveSwitchRepository.ReadCustomerByMobileNumberAndActiveSwitch(
		ctx,
		req.MobileNumber,
		req.ActiveSwitch,
	)
	if err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, customizeerrors.CustomerNotFoundError
	}
	return customers, nil
}
