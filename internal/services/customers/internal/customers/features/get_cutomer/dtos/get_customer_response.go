package featuresdtos

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	"github.com/samber/lo"
)

type GetCustomerResponse struct {
	ID           string                            `json:"id"`
	MobileNumber string                            `json:"mobile_number"`
	Email        string                            `json:"email"`
	FirstName    string                            `json:"first_name"`
	LastName     string                            `json:"last_name"`
	ActiveSwitch bool                              `json:"active_switch"`
	History      dtos.CommonHistoryModelWithUpdate `json:"history"`
}

type CustomerEntity entities.Customer

func (c *CustomerEntity) ToGetCustomerResponse() *GetCustomerResponse {
	return &GetCustomerResponse{
		ID:           c.ID,
		MobileNumber: c.MobileNumber,
		Email:        c.Email,
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		ActiveSwitch: c.ActiveSwitch,
		History: dtos.CommonHistoryModelWithUpdate{
			CommonHistoryModel: dtos.CommonHistoryModel{
				CreatedAt: c.CreatedAt,
				DeletedAt: c.DeletedAt,
			},
			UpdatedAt: c.UpdatedAt,
		},
	}
}

type CustomerEntities entities.Customers

func (c CustomerEntities) ToGetCustomerResponse() []GetCustomerResponse {
	return lo.Map(c, func(customer entities.Customer, _ int) GetCustomerResponse {
		customerEntity := CustomerEntity(customer)
		return *customerEntity.ToGetCustomerResponse()
	})
}
