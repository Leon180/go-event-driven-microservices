package gincontrollers

import (
	customizeginresponse "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/response"
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/dtos"
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
	routerGroup.GET("", endpoint.Handle)
}

func (handle *getAccountImpl) Handle(c *gin.Context) {
	var getAccountRequest featuredtos.GetAccountRequest
	if err := c.ShouldBindQuery(&getAccountRequest); err != nil {
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
