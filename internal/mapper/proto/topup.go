package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type topupProtoMapper struct {
}

func NewTopupProtoMapper() *topupProtoMapper {
	return &topupProtoMapper{}
}

func (t *topupProtoMapper) ToResponseTopup(topup *response.TopupResponse) *pb.TopupResponse {
	return &pb.TopupResponse{
		Id:          int32(topup.ID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

func (t *topupProtoMapper) ToResponsesTopup(topups []*response.TopupResponse) []*pb.TopupResponse {
	var responses []*pb.TopupResponse

	for _, response := range topups {
		responses = append(responses, t.ToResponseTopup(response))
	}

	return responses
}

func (t *topupProtoMapper) ToResponseTopupDeleteAt(topup *response.TopupResponseDeleteAt) *pb.TopupResponseDeleteAt {
	return &pb.TopupResponseDeleteAt{
		Id:          int32(topup.ID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
		DeletedAt:   topup.DeletedAt,
	}
}

func (t *topupProtoMapper) ToResponsesTopupDeleteAt(topups []*response.TopupResponseDeleteAt) []*pb.TopupResponseDeleteAt {
	var responses []*pb.TopupResponseDeleteAt

	for _, response := range topups {
		responses = append(responses, t.ToResponseTopupDeleteAt(response))
	}

	return responses
}
