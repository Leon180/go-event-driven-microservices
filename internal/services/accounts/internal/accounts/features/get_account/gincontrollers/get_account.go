package gincontrollers

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/get_account/services"
	"github.com/gin-gonic/gin"
)

type getAccountImpl struct {
	getAccountService services.GetAccount
}

func NewGetAccount(getAccountService services.GetAccount) *getAccountImpl {
	return &getAccountImpl{getAccountService: getAccountService}
}

func (handle *getAccountImpl) GetAccount(c *gin.Context) {
	var getAccountRequest featuredtos.GetAccountRequest
	if err := c.ShouldBindQuery(&getAccountRequest); err != nil {
		utilities.ResponseError(c, nil, "", err)
		return
	}
	account, err := handle.getAccountService.GetAccount(c.Request.Context(), &getAccountRequest)
	if err != nil {
		utilities.ResponseError(c, nil, "", err)
		return
	}
	utilities.ResponseSuccess(c, account, "account retrieved successfully")
}
