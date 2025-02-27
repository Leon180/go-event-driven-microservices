package ginendpoints

import (
	customizeginresponse "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/response"
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/restore_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/restore_account/services"
	"github.com/gin-gonic/gin"
)

type restoreAccountImpl struct {
	restoreAccountService services.RestoreAccount
}

func NewRestoreAccount(
	restoreAccountService services.RestoreAccount,
) customizeginendpoints.Endpoint {
	return &restoreAccountImpl{
		restoreAccountService: restoreAccountService,
	}
}

func (endpoint *restoreAccountImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/account/restore", endpoint.Handle)
}

// @Summary Restore an account
// @Description Restore an account
// @Tags accounts
// @Produce json
// @Param account body featuresdtos.RestoreAccountRequest true "Account"
// @Success 200 {object} customizeginresponse.JSONResponse "account restored successfully"
// @Router /account/restore [post]
func (endpoint *restoreAccountImpl) Handle(c *gin.Context) {
	var restoreAccountRequest featuresdtos.RestoreAccountRequest
	if err := c.ShouldBindJSON(&restoreAccountRequest); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.restoreAccountService.RestoreAccount(c.Request.Context(), &restoreAccountRequest); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	customizeginresponse.ResponseSuccess(c, nil, "account restored successfully")
}
