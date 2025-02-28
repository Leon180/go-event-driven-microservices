package repositoriespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/repositories"
	"gorm.io/gorm"
)

type readCustomerByMobileNumberAndActiveSwitchImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadCustomerByMobileNumberAndActiveSwitch(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadCustomerByMobileNumberAndActiveSwitch {
	return &readCustomerByMobileNumberAndActiveSwitchImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readCustomerByMobileNumberAndActiveSwitchImpl) ReadCustomerByMobileNumberAndActiveSwitch(ctx context.Context, mobileNumber string, activeSwitch *bool) (entities.Customers, error) {
	var customers entities.Customers
	sql := handle.db.WithContext(ctx)
	if activeSwitch != nil {
		sql = sql.Where("mobile_number = ? AND active_switch = ?", mobileNumber, *activeSwitch)
	} else {
		sql = sql.Where("mobile_number = ?", mobileNumber)
	}
	if err := sql.Find(&customers).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error reading customer by mobile number and active switch")
		return nil, err
	}
	return customers, nil
}

type readCustomerImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func NewReadCustomer(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadCustomer {
	return &readCustomerImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

func (handle *readCustomerImpl) ReadCustomer(ctx context.Context, id string) (*entities.Customer, error) {
	var customer entities.Customer
	if err := handle.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&customer).Error; err != nil {
		handle.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("error reading customer by id")
		return nil, err
	}
	if customer.ID == "" {
		return nil, nil
	}
	return &customer, nil
}
