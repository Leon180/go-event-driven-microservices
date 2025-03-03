package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/features/create_account/services"
	"github.com/gin-gonic/gin"
)

type createAccountImpl struct {
	createAccountService services.CreateAccount
}

func NewCreateAccount(
	createAccountService services.CreateAccount,
) customizegin.Endpoint {
	return &createAccountImpl{
		createAccountService: createAccountService,
	}
}

func (endpoint *createAccountImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/account/create", endpoint.Handle)
}

// @Summary Create a new account
// @Description Create a new account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body featuresdtos.CreateAccountRequest true "Account"
// @Success 200 {object} customizegin.JSONResponse "Account created successfully"
// @Router /account/create [post]
func (endpoint *createAccountImpl) Handle(c *gin.Context) {
	var createAccountRequest featuresdtos.CreateAccountRequest
	if err := c.ShouldBindJSON(&createAccountRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.createAccountService.CreateAccount(c.Request.Context(), &createAccountRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "account created successfully")
}
