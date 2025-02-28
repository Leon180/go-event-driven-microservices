package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/create_card/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/create_card/services"
	"github.com/gin-gonic/gin"
)

type createCreditCardImpl struct {
	createCreditCardService services.CreateCreditCard
}

func NewCreateCreditCard(
	createCreditCardService services.CreateCreditCard,
) customizegin.Endpoint {
	return &createCreditCardImpl{
		createCreditCardService: createCreditCardService,
	}
}

func (endpoint *createCreditCardImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/credit-card/create", endpoint.Handle)
}

// @Summary Create a new credit card
// @Description Create a new credit card
// @Tags cards
// @Accept json
// @Produce json
// @Param credit-card body featuresdtos.CreateCreditCardRequest true "Credit Card"
// @Success 200 {object} customizegin.JSONResponse "Credit Card created successfully"
// @Router /credit-card/create [post]
func (endpoint *createCreditCardImpl) Handle(c *gin.Context) {
	var createCreditCardRequest featuresdtos.CreateCreditCardRequest
	if err := c.ShouldBindJSON(&createCreditCardRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.createCreditCardService.CreateCreditCard(c.Request.Context(), &createCreditCardRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "card created successfully")
}
