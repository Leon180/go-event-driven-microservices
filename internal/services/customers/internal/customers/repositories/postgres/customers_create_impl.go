package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories"
	"gorm.io/gorm"
)

type createCustomerImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewCreateCustomer(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.CreateCustomer {
	return &createCustomerImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *createCustomerImpl) CreateCustomer(ctx context.Context, customer *entities.Customer) error {
	if customer == nil {
		return nil
	}
	if err := handle.db.WithContext(ctx).Create(customer).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error creating customer")
		return err
	}
	return nil
}

type createCustomersImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewCreateCustomers(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.CreateCustomers {
	return &createCustomersImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *createCustomersImpl) CreateCustomers(ctx context.Context, customers entities.Customers) error {
	if customers == nil {
		return nil
	}
	if err := handle.db.WithContext(ctx).Create(customers).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error creating customers")
		return err
	}
	return nil
}
