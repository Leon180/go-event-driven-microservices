package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
)

//go:generate mockgen -source=customers_create.go -destination=mocks/customers_create_mock.go -package=mocks

type CreateCustomer interface {
	CreateCustomer(ctx context.Context, customer *entities.Customer) error
}

type CreateCustomers interface {
	CreateCustomers(ctx context.Context, customers entities.Customers) error
}
