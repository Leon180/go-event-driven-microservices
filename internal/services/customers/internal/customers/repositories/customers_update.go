package repositories

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
)

//go:generate mockgen -source=customers_update.go -destination=./mocks/customers_update_mock.go -package=mocks

type UpdateCustomerByID interface {
	UpdateCustomerByID(ctx context.Context, update entities.UpdateCustomer) error
}
