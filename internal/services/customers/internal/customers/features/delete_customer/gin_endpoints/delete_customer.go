package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/delete_customer/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/delete_customer/services"
	"github.com/gin-gonic/gin"
)

type deleteCustomerImpl struct {
	deleteCustomerService services.DeleteCustomer
}

func NewDeleteCustomer(
	deleteCustomerService services.DeleteCustomer,
) customizegin.Endpoint {
	return &deleteCustomerImpl{
		deleteCustomerService: deleteCustomerService,
	}
}

func (endpoint *deleteCustomerImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/customer/delete", endpoint.Handle)
}

// @Summary Delete a customer
// @Description Delete a customer
// @Tags customers
// @Produce json
// @Param customer body featuresdtos.DeleteCustomerRequest true "Customer"
// @Success 200 {object} customizegin.JSONResponse "Customer deleted successfully"
// @Router /customer/delete [post]
func (endpoint *deleteCustomerImpl) Handle(c *gin.Context) {
	var deleteCustomerRequest featuresdtos.DeleteCustomerRequest
	if err := c.ShouldBindJSON(&deleteCustomerRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.deleteCustomerService.DeleteCustomer(c.Request.Context(), &deleteCustomerRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "card deleted successfully")
}
