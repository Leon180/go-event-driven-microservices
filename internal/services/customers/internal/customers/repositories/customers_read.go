package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
)

//go:generate mockgen -source=customers_read.go -destination=mocks/customers_read_mock.go -package=mocks

type ReadCustomerByMobileNumberAndActiveSwitch interface {
	ReadCustomerByMobileNumberAndActiveSwitch(ctx context.Context, mobileNumber string, activeSwitch *bool) (entities.Customers, error)
}

type ReadCustomer interface {
	ReadCustomer(ctx context.Context, id string) (*entities.Customer, error)
}
