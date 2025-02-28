package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/update_customer/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/update_customer/services"
	"github.com/gin-gonic/gin"
)

type updateCustomerImpl struct {
	updateCustomerService services.UpdateCustomer
}

func NewUpdateCustomer(
	updateCustomerService services.UpdateCustomer,
) customizegin.Endpoint {
	return &updateCustomerImpl{
		updateCustomerService: updateCustomerService,
	}
}

func (endpoint *updateCustomerImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.PUT("/customer/update", endpoint.Handle)
}

// @Summary Update an customer
// @Description Update an customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body featuresdtos.UpdateCustomerRequest true "Customer"
// @Success 200 {object} customizegin.JSONResponse "customer updated successfully"
// @Router /customer/update [put]
func (endpoint *updateCustomerImpl) Handle(c *gin.Context) {
	var req featuresdtos.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.updateCustomerService.UpdateCustomer(c.Request.Context(), &req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "customer updated successfully")
}
