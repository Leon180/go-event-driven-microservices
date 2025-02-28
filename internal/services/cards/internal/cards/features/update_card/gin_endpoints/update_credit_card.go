package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/update_card/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/update_card/services"
	"github.com/gin-gonic/gin"
)

type updateCreditCardImpl struct {
	updateCreditCardService services.UpdateCreditCard
}

func NewUpdateCreditCard(
	updateCreditCardService services.UpdateCreditCard,
) customizegin.Endpoint {
	return &updateCreditCardImpl{
		updateCreditCardService: updateCreditCardService,
	}
}

func (endpoint *updateCreditCardImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.PUT("/credit-card/update", endpoint.Handle)
}

// @Summary Update an card
// @Description Update an card
// @Tags cards
// @Accept json
// @Produce json
// @Param card body featuresdtos.UpdateCreditCardRequest true "Card"
// @Success 200 {object} customizegin.JSONResponse "card updated successfully"
// @Router /credit-card/update [put]
func (endpoint *updateCreditCardImpl) Handle(c *gin.Context) {
	var req featuresdtos.UpdateCreditCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.updateCreditCardService.UpdateCreditCard(c.Request.Context(), &req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "card updated successfully")
}
