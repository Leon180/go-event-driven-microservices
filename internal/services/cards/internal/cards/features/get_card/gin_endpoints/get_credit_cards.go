package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/get_card/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/get_card/services"
	"github.com/gin-gonic/gin"
)

type getCreditCardsByMobileNumberAndActiveSwitchImpl struct {
	getCreditCardsByMobileNumberAndActiveSwitchService services.GetCreditCardsByMobileNumberAndActiveSwitch
}

func NewGetCreditCardsByMobileNumberAndActiveSwitch(
	getCreditCardsByMobileNumberAndActiveSwitchService services.GetCreditCardsByMobileNumberAndActiveSwitch,
) customizegin.Endpoint {
	return &getCreditCardsByMobileNumberAndActiveSwitchImpl{
		getCreditCardsByMobileNumberAndActiveSwitchService: getCreditCardsByMobileNumberAndActiveSwitchService,
	}
}

func (endpoint *getCreditCardsByMobileNumberAndActiveSwitchImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("credit-cards/get", endpoint.Handle)
}

// @Summary Get credit cards by mobile number and active switch
// @Description Get credit cards by mobile number and active switch
// @Tags cards
// @Produce json
// @Param account body featuresdtos.GetCreditCardsRequest true "Credit Card"
// @Success 200 {object} customizegin.JSONResponse "credit card retrieved successfully"
// @Router /credit-cards/get [post]
func (handle *getCreditCardsByMobileNumberAndActiveSwitchImpl) Handle(c *gin.Context) {
	var req featuresdtos.GetCreditCardsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	creditCards, err := handle.getCreditCardsByMobileNumberAndActiveSwitchService.GetCreditCardsByMobileNumberAndActiveSwitch(
		c.Request.Context(),
		&req,
	)
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(
		c,
		featuresdtos.CreditCardEntities(creditCards).ToGetCreditCardsResponse(),
		"credit card retrieved successfully",
	)
}
