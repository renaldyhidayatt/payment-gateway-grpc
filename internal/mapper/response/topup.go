package responsemapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
)

type topupResponseMapper struct {
}

func NewTopupResponseMapper() *topupResponseMapper {
	return &topupResponseMapper{}
}

func (s *topupResponseMapper) ToTopupResponse(topup *record.TopupRecord) *response.TopupResponse {
	return &response.TopupResponse{
		ID:          topup.ID,
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: topup.TopupAmount,
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

func (s *topupResponseMapper) ToTopupResponses(topups []*record.TopupRecord) []*response.TopupResponse {
	var responses []*response.TopupResponse

	for _, response := range topups {
		responses = append(responses, s.ToTopupResponse(response))
	}

	return responses
}

func (s *topupResponseMapper) ToTopupResponseDeleteAt(topup *record.TopupRecord) *response.TopupResponseDeleteAt {
	return &response.TopupResponseDeleteAt{
		ID:          topup.ID,
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: topup.TopupAmount,
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
		DeletedAt:   *topup.DeletedAt,
	}
}

func (s *topupResponseMapper) ToTopupResponsesDeleteAt(topups []*record.TopupRecord) []*response.TopupResponseDeleteAt {
	var responses []*response.TopupResponseDeleteAt

	for _, response := range topups {
		responses = append(responses, s.ToTopupResponseDeleteAt(response))
	}

	return responses
}
