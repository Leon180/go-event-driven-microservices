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
	routerGroup.POST("/accounts/get", endpoint.Handle)
}

// @Summary Get accounts by mobile number
// @Description Get accounts by mobile number
// @Tags accounts
// @Produce json
// @Param account body featuresdtos.GetAccountsByMobileNumberRequest true "Account"
// @Success 200 {object} customizeginresponse.JSONResponse "accounts retrieved successfully"
// @Router /accounts/get [post]
func (handle *getAccountsByMobileNumberImpl) Handle(c *gin.Context) {
	var req featuresdtos.GetAccountsByMobileNumberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
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
