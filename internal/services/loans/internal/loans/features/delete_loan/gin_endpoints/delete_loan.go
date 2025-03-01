package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/delete_loan/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/delete_loan/services"
	"github.com/gin-gonic/gin"
)

type deleteLoanImpl struct {
	deleteLoanService services.DeleteLoan
}

func NewDeleteLoan(
	deleteLoanService services.DeleteLoan,
) customizegin.Endpoint {
	return &deleteLoanImpl{
		deleteLoanService: deleteLoanService,
	}
}

func (endpoint *deleteLoanImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/loan/delete", endpoint.Handle)
}

// @Summary Delete a loan
// @Description Delete a loan
// @Tags loans
// @Produce json
// @Param loan body featuresdtos.DeleteLoanRequest true "Loan"
// @Success 200 {object} customizegin.JSONResponse "Loan deleted successfully"
// @Router /loan/delete [post]
func (endpoint *deleteLoanImpl) Handle(c *gin.Context) {
	var deleteLoanRequest featuresdtos.DeleteLoanRequest
	if err := c.ShouldBindJSON(&deleteLoanRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.deleteLoanService.DeleteLoan(c.Request.Context(), &deleteLoanRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "loan deleted successfully")
}
