package gincontrollers

import (
	customizeginresponse "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/response"
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/services"
	"github.com/gin-gonic/gin"
)

type getAccountImpl struct {
	getAccountService services.GetAccount
}

func NewGetAccount(
	getAccountService services.GetAccount,
) customizeginendpoints.Endpoint {
	return &getAccountImpl{
		getAccountService: getAccountService,
	}
}

func (endpoint *getAccountImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/account/:mobile_number", endpoint.Handle)
}

// @Summary Get an account
// @Description Get an account
// @Tags accounts
// @Produce json
// @Param mobile_number path string true "Mobile Number"
// @Success 200 {object} customizeginresponse.JSONResponse "account retrieved successfully"
// @Router /account/{mobile_number} [get]
func (handle *getAccountImpl) Handle(c *gin.Context) {
	var getAccountRequest featuresdtos.GetAccountRequest
	if err := c.ShouldBindUri(&getAccountRequest); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	account, err := handle.getAccountService.GetAccount(c.Request.Context(), &getAccountRequest)
	if err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	customizeginresponse.ResponseSuccess(c, account, "account retrieved successfully")
}
