package featuresdtos

import (
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	"github.com/samber/lo"
)

type GetCreditCardsResponse struct {
	ID           string                            `json:"id"`
	MobileNumber string                            `json:"mobile_number"`
	TotalLimit   string                            `json:"total_limit"`
	AmountUsed   string                            `json:"amount_used"`
	ActiveSwitch bool                              `json:"active_switch"`
	History      dtos.CommonHistoryModelWithUpdate `json:"history"`
}

type CreditCardEntity entities.CreditCard

func (c *CreditCardEntity) ToGetCreditCardsResponse() *GetCreditCardsResponse {
	return &GetCreditCardsResponse{
		ID:           c.ID,
		MobileNumber: c.MobileNumber,
		TotalLimit:   c.TotalLimit.String(),
		AmountUsed:   c.AmountUsed.String(),
		ActiveSwitch: c.ActiveSwitch,
		History: dtos.CommonHistoryModelWithUpdate{
			CommonHistoryModel: dtos.CommonHistoryModel{
				CreatedAt: c.CreatedAt,
				DeletedAt: c.DeletedAt,
			},
			UpdatedAt: c.UpdatedAt,
		},
	}
}

type CreditCardEntities entities.CreditCards

func (c CreditCardEntities) ToGetCreditCardsResponse() []GetCreditCardsResponse {
	return lo.Map(c, func(creditCard entities.CreditCard, _ int) GetCreditCardsResponse {
		creditCardEntity := CreditCardEntity(creditCard)
		return *creditCardEntity.ToGetCreditCardsResponse()
	})
}
