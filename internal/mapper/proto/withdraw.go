package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type withdrawProtoMapper struct {
}

func NewWithdrawProtoMapper() *withdrawProtoMapper {
	return &withdrawProtoMapper{}
}

func (w *withdrawProtoMapper) ToResponseWithdrawal(withdraw *response.WithdrawResponse) *pb.WithdrawResponse {
	return &pb.WithdrawResponse{
		WithdrawId:     int32(withdraw.ID),
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int32(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
	}
}

func (w *withdrawProtoMapper) ToResponsesWithdrawal(withdraws []*response.WithdrawResponse) []*pb.WithdrawResponse {
	var responseWithdraws []*pb.WithdrawResponse

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.ToResponseWithdrawal(withdraw))
	}

	return responseWithdraws
}

func (w *withdrawProtoMapper) ToResponseWithdrawalDeleteAt(withdraw *response.WithdrawResponseDeleteAt) *pb.WithdrawResponseDeleteAt {
	return &pb.WithdrawResponseDeleteAt{
		WithdrawId:     int32(withdraw.ID),
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int32(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
		DeletedAt:      withdraw.DeletedAt,
	}
}

func (w *withdrawProtoMapper) ToResponsesWithdrawalDeleteAt(withdraws []*response.WithdrawResponseDeleteAt) []*pb.WithdrawResponseDeleteAt {
	var responseWithdraws []*pb.WithdrawResponseDeleteAt

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.ToResponseWithdrawalDeleteAt(withdraw))
	}

	return responseWithdraws
}
