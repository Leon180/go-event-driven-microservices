package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/services"
	"github.com/gin-gonic/gin"
)

type deleteAccountImpl struct {
	deleteAccountService services.DeleteAccount
}

func NewDeleteAccount(
	deleteAccountService services.DeleteAccount,
) customizegin.Endpoint {
	return &deleteAccountImpl{
		deleteAccountService: deleteAccountService,
	}
}

func (endpoint *deleteAccountImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/account/delete", endpoint.Handle)
}

// @Summary Delete an account
// @Description Delete an account
// @Tags accounts
// @Produce json
// @Param account body featuresdtos.DeleteAccountRequest true "Account"
// @Success 200 {object} customizegin.JSONResponse "account deleted successfully"
// @Router /account/delete [post]
func (endpoint *deleteAccountImpl) Handle(c *gin.Context) {
	var deleteAccountRequest featuresdtos.DeleteAccountRequest
	if err := c.ShouldBindJSON(&deleteAccountRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.deleteAccountService.DeleteAccount(c.Request.Context(), &deleteAccountRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "account deleted successfully")
}
