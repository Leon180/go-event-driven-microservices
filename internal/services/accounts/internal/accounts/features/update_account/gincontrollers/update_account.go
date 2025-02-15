package gincontrollers

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/update_account/services"
	"github.com/gin-gonic/gin"
)

type updateAccountImpl struct {
	updateAccountService services.UpdateAccount
}

func NewUpdateAccount(updateAccountService services.UpdateAccount) *updateAccountImpl {
	return &updateAccountImpl{updateAccountService: updateAccountService}
}

func (handle *updateAccountImpl) UpdateAccount(c *gin.Context) {
	var updateAccountRequest featuredtos.UpdateAccountRequest
	if err := c.ShouldBindJSON(&updateAccountRequest); err != nil {
		utilities.ResponseError(c, nil, "", err)
		return
	}
	if err := handle.updateAccountService.UpdateAccount(c.Request.Context(), &updateAccountRequest); err != nil {
		utilities.ResponseError(c, nil, "", err)
		return
	}
	utilities.ResponseSuccess(c, nil, "account updated successfully")
}
