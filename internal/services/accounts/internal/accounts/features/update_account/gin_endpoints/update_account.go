package gincontrollers

import (
	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	customizeginresponse "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/response"
	customizeginendpoints "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/server/endpoints"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
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
	routerGroup.PUT("/account/:mobile_number", endpoint.Handle)
}

// @Summary Update an account
// @Description Update an account
// @Tags accounts
// @Accept json
// @Produce json
// @Param mobile_number path string true "Mobile Number"
// @Param account body featuresdtos.UpdateAccountRequest true "Account"
// @Success 200 {object} customizeginresponse.JSONResponse "account updated successfully"
// @Router /account/{mobile_number} [put]
func (endpoint *updateAccountImpl) Handle(c *gin.Context) {
	var updateAccountRequest featuresdtos.UpdateAccountRequest
	if err := c.ShouldBindJSON(&updateAccountRequest); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	updateAccountRequest.MobileNumber = c.Param("mobile_number")
	if updateAccountRequest.MobileNumber == "" {
		customizeginresponse.ResponseError(c, nil, "", customizeerrors.HTTPBadRequestError)
		return
	}
	if err := endpoint.updateAccountService.UpdateAccount(c.Request.Context(), &updateAccountRequest); err != nil {
		customizeginresponse.ResponseError(c, nil, "", err)
		return
	}
	customizeginresponse.ResponseSuccess(c, nil, "account updated successfully")
}
