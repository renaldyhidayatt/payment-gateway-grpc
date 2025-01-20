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
		WithdrawNo:     withdraw.WithdrawNo,
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
		WithdrawNo:     withdraw.WithdrawNo,
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

func (t *withdrawProtoMapper) ToResponseWithdrawMonthStatusSuccess(s *response.WithdrawResponseMonthStatusSuccess) *pb.WithdrawMonthStatusSuccessResponse {
	return &pb.WithdrawMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *withdrawProtoMapper) ToResponsesWithdrawMonthStatusSuccess(Withdraws []*response.WithdrawResponseMonthStatusSuccess) []*pb.WithdrawMonthStatusSuccessResponse {
	var WithdrawRecords []*pb.WithdrawMonthStatusSuccessResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToResponseWithdrawMonthStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

func (t *withdrawProtoMapper) ToWithdrawResponseYearStatusSuccess(s *response.WithdrawResponseYearStatusSuccess) *pb.WithdrawYearStatusSuccessResponse {
	return &pb.WithdrawYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *withdrawProtoMapper) ToWithdrawResponsesYearStatusSuccess(Withdraws []*response.WithdrawResponseYearStatusSuccess) []*pb.WithdrawYearStatusSuccessResponse {
	var WithdrawRecords []*pb.WithdrawYearStatusSuccessResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseYearStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

func (t *withdrawProtoMapper) ToResponseWithdrawMonthStatusFailed(s *response.WithdrawResponseMonthStatusFailed) *pb.WithdrawMonthStatusFailedResponse {
	return &pb.WithdrawMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *withdrawProtoMapper) ToResponsesWithdrawMonthStatusFailed(Withdraws []*response.WithdrawResponseMonthStatusFailed) []*pb.WithdrawMonthStatusFailedResponse {
	var WithdrawRecords []*pb.WithdrawMonthStatusFailedResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToResponseWithdrawMonthStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

func (t *withdrawProtoMapper) ToWithdrawResponseYearStatusFailed(s *response.WithdrawResponseYearStatusFailed) *pb.WithdrawYearStatusFailedResponse {
	return &pb.WithdrawYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *withdrawProtoMapper) ToWithdrawResponsesYearStatusFailed(Withdraws []*response.WithdrawResponseYearStatusFailed) []*pb.WithdrawYearStatusFailedResponse {
	var WithdrawRecords []*pb.WithdrawYearStatusFailedResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseYearStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

func (m *withdrawProtoMapper) ToResponseWithdrawMonthlyAmount(s *response.WithdrawMonthlyAmountResponse) *pb.WithdrawMonthlyAmountResponse {
	return &pb.WithdrawMonthlyAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (m *withdrawProtoMapper) ToResponseWithdrawMonthlyAmounts(s []*response.WithdrawMonthlyAmountResponse) []*pb.WithdrawMonthlyAmountResponse {
	var protoResponses []*pb.WithdrawMonthlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.ToResponseWithdrawMonthlyAmount(withdraw))
	}
	return protoResponses
}

func (m *withdrawProtoMapper) ToResponseWithdrawYearlyAmount(s *response.WithdrawYearlyAmountResponse) *pb.WithdrawYearlyAmountResponse {
	return &pb.WithdrawYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (m *withdrawProtoMapper) ToResponseWithdrawYearlyAmounts(s []*response.WithdrawYearlyAmountResponse) []*pb.WithdrawYearlyAmountResponse {
	var protoResponses []*pb.WithdrawYearlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.ToResponseWithdrawYearlyAmount(withdraw))
	}
	return protoResponses
}
