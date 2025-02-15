package gincontrollers

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	featuredtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/services"
	"github.com/gin-gonic/gin"
)

type createAccountImpl struct {
	createAccountService services.CreateAccount
}

func NewCreateAccount(createAccountService services.CreateAccount) *createAccountImpl {
	return &createAccountImpl{createAccountService: createAccountService}
}

func (handle *createAccountImpl) CreateAccount(c *gin.Context) {
	var createAccountRequest featuredtos.CreateAccountRequest
	if err := c.ShouldBindJSON(&createAccountRequest); err != nil {
		utilities.ResponseError(c, nil, "", err)
		return
	}
	if err := handle.createAccountService.CreateAccount(c.Request.Context(), &createAccountRequest); err != nil {
		utilities.ResponseError(c, nil, "", err)
		return
	}
	utilities.ResponseSuccess(c, nil, "account created successfully")
}
