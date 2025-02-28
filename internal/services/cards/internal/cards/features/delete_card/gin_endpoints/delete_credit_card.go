package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/delete_card/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/features/delete_card/services"
	"github.com/gin-gonic/gin"
)

type deleteCreditCardImpl struct {
	deleteCreditCardService services.DeleteCreditCard
}

func NewDeleteCreditCard(
	deleteCreditCardService services.DeleteCreditCard,
) customizegin.Endpoint {
	return &deleteCreditCardImpl{
		deleteCreditCardService: deleteCreditCardService,
	}
}

func (endpoint *deleteCreditCardImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/credit-card/delete", endpoint.Handle)
}

// @Summary Delete a credit card
// @Description Delete a credit card
// @Tags cards
// @Produce json
// @Param credit-card body featuresdtos.DeleteCreditCardRequest true "Credit Card"
// @Success 200 {object} customizegin.JSONResponse "Credit Card deleted successfully"
// @Router /credit-card/delete [post]
func (endpoint *deleteCreditCardImpl) Handle(c *gin.Context) {
	var deleteCreditCardRequest featuresdtos.DeleteCreditCardRequest
	if err := c.ShouldBindJSON(&deleteCreditCardRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.deleteCreditCardService.DeleteCreditCard(c.Request.Context(), &deleteCreditCardRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "card deleted successfully")
}
