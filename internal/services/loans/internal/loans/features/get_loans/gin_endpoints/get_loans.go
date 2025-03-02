package ginendpoints

import (
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	featuresdtos "github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/get_loans/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/features/get_loans/services"
	"github.com/gin-gonic/gin"
)

type getLoansByMobileNumberAndActiveSwitchImpl struct {
	getLoansByMobileNumberAndActiveSwitchService services.GetLoansByMobileNumberAndActiveSwitch
}

func NewGetLoansByMobileNumberAndActiveSwitch(
	getLoansByMobileNumberAndActiveSwitchService services.GetLoansByMobileNumberAndActiveSwitch,
) customizegin.Endpoint {
	return &getLoansByMobileNumberAndActiveSwitchImpl{
		getLoansByMobileNumberAndActiveSwitchService: getLoansByMobileNumberAndActiveSwitchService,
	}
}

func (endpoint *getLoansByMobileNumberAndActiveSwitchImpl) MapEndpoint(routerGroup *gin.RouterGroup) {
	routerGroup.POST("loans/get", endpoint.Handle)
}

// @Summary Get loans by mobile number and active switch
// @Description Get loans by mobile number and active switch
// @Tags loans
// @Produce json
// @Param account body featuresdtos.GetLoansRequest true "Loan"
// @Success 200 {object} customizegin.JSONResponse "loan retrieved successfully"
// @Router /loans/get [post]
func (handle *getLoansByMobileNumberAndActiveSwitchImpl) Handle(c *gin.Context) {
	var req featuresdtos.GetLoansRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	loans, err := handle.getLoansByMobileNumberAndActiveSwitchService.GetLoansByMobileNumberAndActiveSwitch(
		c.Request.Context(),
		&req,
	)
	if err != nil {
		customizegin.ResponseError(c, nil, "", err)
		return
	}
	customizegin.ResponseSuccess(
		c,
		featuresdtos.LoanEntities(loans).ToGetLoansResponse(),
		"loan retrieved successfully",
	)
}
