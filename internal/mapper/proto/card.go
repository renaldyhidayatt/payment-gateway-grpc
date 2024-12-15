package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type cardProtoMapper struct {
}

func NewCardProtoMapper() *cardProtoMapper {
	return &cardProtoMapper{}
}

func (s *cardProtoMapper) ToResponseCard(card *response.CardResponse) *pb.CardResponse {
	return &pb.CardResponse{
		Id:           int32(card.ID),
		UserId:       int32(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		Cvv:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}

func (s *cardProtoMapper) ToResponsesCard(cards []*response.CardResponse) []*pb.CardResponse {
	responses := make([]*pb.CardResponse, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseCard(card))
	}
	return responses
}
