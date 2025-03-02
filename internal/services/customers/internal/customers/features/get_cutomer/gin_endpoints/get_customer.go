package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/get_cutomer/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/features/get_cutomer/services"
	"github.com/gin-gonic/gin"
)

type getCustomerByMobileNumberAndActiveSwitchImpl struct {
	getCustomerByMobileNumberAndActiveSwitchService services.GetCustomerByMobileNumberAndActiveSwitch
}

func NewGetCustomerByMobileNumberAndActiveSwitch(
	getCustomerByMobileNumberAndActiveSwitchService services.GetCustomerByMobileNumberAndActiveSwitch,
) customizegin.Endpoint {
	return &getCustomerByMobileNumberAndActiveSwitchImpl{
		getCustomerByMobileNumberAndActiveSwitchService: getCustomerByMobileNumberAndActiveSwitchService,
	}
}

func (endpoint *getCustomerByMobileNumberAndActiveSwitchImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("customer/get", endpoint.Handle)
}

// @Summary Get customer by mobile number and active switch
// @Description Get customer by mobile number and active switch
// @Tags customers
// @Produce json
// @Param customer body featuresdtos.GetCustomerRequest true "Customer"
// @Success 200 {object} customizegin.JSONResponse "customer retrieved successfully"
// @Router /customer/get [post]
func (handle *getCustomerByMobileNumberAndActiveSwitchImpl) Handle(c *gin.Context) {
	var req featuresdtos.GetCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customers, err := handle.getCustomerByMobileNumberAndActiveSwitchService.GetCustomerByMobileNumberAndActiveSwitch(
		c.Request.Context(),
		&req,
	)
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(
		c,
		featuresdtos.CustomerEntities(customers).ToGetCustomerResponse(),
		"customer retrieved successfully",
	)
}
