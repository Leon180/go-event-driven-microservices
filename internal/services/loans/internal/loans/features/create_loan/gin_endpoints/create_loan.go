package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/create_loan/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/create_loan/services"
	"github.com/gin-gonic/gin"
)

type createLoanImpl struct {
	createLoanService services.CreateLoan
}

func NewCreateLoan(
	createLoanService services.CreateLoan,
) customizegin.Endpoint {
	return &createLoanImpl{
		createLoanService: createLoanService,
	}
}

func (endpoint *createLoanImpl) MapEndpoint(router *gin.RouterGroup) {
	router.POST("/loan/create", endpoint.Handle)
}

// @Summary Create a new loan
// @Description Create a new loan
// @Tags loans
// @Accept json
// @Produce json
// @Param loan body featuresdtos.CreateLoanRequest true "Loan"
// @Success 200 {object} customizegin.JSONResponse "Loan created successfully"
// @Router /loan/create [post]
func (endpoint *createLoanImpl) Handle(c *gin.Context) {
	var createLoanRequest featuresdtos.CreateLoanRequest
	if err := c.ShouldBindJSON(&createLoanRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	if err := endpoint.createLoanService.CreateLoan(c.Request.Context(), &createLoanRequest); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(c, nil, "loan created successfully")
}
