package gincontrollers

import (
	customizeginresponse "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/response"
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/services"
	"github.com/gin-gonic/gin"
)

type deleteAccountImpl struct {
	deleteAccountService services.DeleteAccount
}

func NewDeleteAccount(
	deleteAccountService services.DeleteAccount,
) customizeginendpoints.Endpoint {
	return &deleteAccountImpl{
		deleteAccountService: deleteAccountService,
	}
}

func (endpoint *deleteAccountImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.DELETE("/account/:id", endpoint.Handle)
}

// @Summary Delete an account
// @Description Delete an account
// @Tags accounts
// @Produce json
// @Param id path string true "Account ID to delete"
// @Success 200 {object} customizeginresponse.JSONResponse "account deleted successfully"
// @Router /account/{id} [delete]
func (endpoint *deleteAccountImpl) Handle(c *gin.Context) {
	var deleteAccountRequest featuresdtos.DeleteAccountRequest
	if err := c.ShouldBindUri(&deleteAccountRequest); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.deleteAccountService.DeleteAccount(c.Request.Context(), &deleteAccountRequest); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	customizeginresponse.ResponseSuccess(c, nil, "account deleted successfully")
}
