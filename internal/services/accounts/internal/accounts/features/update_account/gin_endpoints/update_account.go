package gincontrollers

import (
	customizeginresponse "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/response"
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/services"
	"github.com/gin-gonic/gin"
)

type updateAccountImpl struct {
	updateAccountService services.UpdateAccount
}

func NewUpdateAccount(
	updateAccountService services.UpdateAccount,
) customizeginendpoints.Endpoint {
	return &updateAccountImpl{
		updateAccountService: updateAccountService,
	}
}

func (endpoint *updateAccountImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.PUT("", endpoint.Handle)
}

func (endpoint *updateAccountImpl) Handle(c *gin.Context) {
	var updateAccountRequest featuredtos.UpdateAccountRequest
	if err := c.ShouldBindJSON(&updateAccountRequest); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.updateAccountService.UpdateAccount(c.Request.Context(), &updateAccountRequest); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	customizeginresponse.ResponseSuccess(c, nil, "account updated successfully")
}
