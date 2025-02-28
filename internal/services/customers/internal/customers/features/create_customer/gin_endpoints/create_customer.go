package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/create_customer/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/create_customer/services"
	"github.com/gin-gonic/gin"
)

type createCustomerImpl struct {
	createCustomerService services.CreateCustomer
}

func NewCreateCustomer(
	createCustomerService services.CreateCustomer,
) customizegin.Endpoint {
	return &createCustomerImpl{
		createCustomerService: createCustomerService,
	}
}

func (endpoint *createCustomerImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/customer/create", endpoint.Handle)
}

// @Summary Create a new customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body featuresdtos.CreateCustomerRequest true "Customer"
// @Success 200 {object} customizegin.JSONResponse "Customer created successfully"
// @Router /customer/create [post]
func (endpoint *createCustomerImpl) Handle(c *gin.Context) {
	var createCustomerRequest featuresdtos.CreateCustomerRequest
	if err := c.ShouldBindJSON(&createCustomerRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.createCustomerService.CreateCustomer(c.Request.Context(), &createCustomerRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "customer created successfully")
}
