package ginendpoints

import (
	customizeginresponse "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/response"
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_accounts/services"
	"github.com/gin-gonic/gin"
)

type getAccountsByMobileNumberImpl struct {
	getAccountsByMobileNumberService services.GetAccountsByMobileNumber
}

func NewGetAccountsByMobileNumber(
	getAccountsByMobileNumberService services.GetAccountsByMobileNumber,
) customizeginendpoints.Endpoint {
	return &getAccountsByMobileNumberImpl{
		getAccountsByMobileNumberService: getAccountsByMobileNumberService,
	}
}

func (endpoint *getAccountsByMobileNumberImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/accounts/:mobile_number", endpoint.Handle)
}

// @Summary Get accounts by mobile number
// @Description Get accounts by mobile number
// @Tags accounts
// @Produce json
// @Param mobile_number path string true "Mobile Number"
// @Success 200 {object} customizeginresponse.JSONResponse "accounts retrieved successfully"
// @Router /accounts/{mobile_number} [get]
func (handle *getAccountsByMobileNumberImpl) Handle(c *gin.Context) {
	var req featuresdtos.GetAccountsByMobileNumberRequest
	if err := c.ShouldBindUri(&req); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	accounts, err := handle.getAccountsByMobileNumberService.GetAccountsByMobileNumber(c.Request.Context(), &req)
	if err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	customizeginresponse.ResponseSuccess(c, featuresdtos.AccountsWithHistory(accounts).ToGetAccountsResponse(), "account retrieved successfully")
}
