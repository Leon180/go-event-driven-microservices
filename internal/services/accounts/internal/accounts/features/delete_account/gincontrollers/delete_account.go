package gincontrollers

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/delete_account/services"
	"github.com/gin-gonic/gin"
)

type deleteAccountImpl struct {
	deleteAccountService services.DeleteAccount
}

func NewDeleteAccount(deleteAccountService services.DeleteAccount) *deleteAccountImpl {
	return &deleteAccountImpl{deleteAccountService: deleteAccountService}
}

func (handle *deleteAccountImpl) DeleteAccount(c *gin.Context) {
	var deleteAccountRequest featuredtos.DeleteAccountRequest
	if err := c.ShouldBindJSON(&deleteAccountRequest); err != nil {
		utilities.ResponseError(c, nil, "", err)
		return
	}
	if err := handle.deleteAccountService.DeleteAccount(c.Request.Context(), &deleteAccountRequest); err != nil {
		utilities.ResponseError(c, nil, "", err)
		return
	}
	utilities.ResponseSuccess(c, nil, "account deleted successfully")
}
