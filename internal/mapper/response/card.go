package responsemapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
)

type cardResponseMapper struct {
}

func NewCardResponseMapper() *cardResponseMapper {
	return &cardResponseMapper{}
}

func (s *cardResponseMapper) ToCardResponse(card record.CardRecord) *response.CardResponse {
	return &response.CardResponse{
		ID:           card.ID,
		UserID:       card.UserID,
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}

func (s *cardResponseMapper) ToCardsResponse(cards []*record.CardRecord) []*response.CardResponse {
	var response []*response.CardResponse

	for _, card := range cards {
		response = append(response, s.ToCardResponse(*card))
	}

	return response
}
