package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/update_loan/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/update_loan/services"
	"github.com/gin-gonic/gin"
)

type updateLoanImpl struct {
	updateLoanService services.UpdateLoan
}

func NewUpdateLoan(
	updateLoanService services.UpdateLoan,
) customizegin.Endpoint {
	return &updateLoanImpl{
		updateLoanService: updateLoanService,
	}
}

func (endpoint *updateLoanImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.PUT("/loan/update", endpoint.Handle)
}

// @Summary Update an loan
// @Description Update an loan
// @Tags loans
// @Accept json
// @Produce json
// @Param loan body featuresdtos.UpdateLoanRequest true "Loan"
// @Success 200 {object} customizegin.JSONResponse "loan updated successfully"
// @Router /loan/update [put]
func (endpoint *updateLoanImpl) Handle(c *gin.Context) {
	var req featuresdtos.UpdateLoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.updateLoanService.UpdateLoan(c.Request.Context(), &req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "loan updated successfully")
}
