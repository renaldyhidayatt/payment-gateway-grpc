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

func (t *topupProtoMapper) ToResponseTopupMonthStatusSuccess(s *response.TopupResponseMonthStatusSuccess) *pb.TopupMonthStatusSuccessResponse {
	return &pb.TopupMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *topupProtoMapper) ToResponsesTopupMonthStatusSuccess(topups []*response.TopupResponseMonthStatusSuccess) []*pb.TopupMonthStatusSuccessResponse {
	var topupRecords []*pb.TopupMonthStatusSuccessResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToResponseTopupMonthStatusSuccess(topup))
	}

	return topupRecords
}

func (t *topupProtoMapper) ToTopupResponseYearStatusSuccess(s *response.TopupResponseYearStatusSuccess) *pb.TopupYearStatusSuccessResponse {
	return &pb.TopupYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *topupProtoMapper) ToTopupResponsesYearStatusSuccess(topups []*response.TopupResponseYearStatusSuccess) []*pb.TopupYearStatusSuccessResponse {
	var topupRecords []*pb.TopupYearStatusSuccessResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseYearStatusSuccess(topup))
	}

	return topupRecords
}

func (t *topupProtoMapper) ToResponseTopupMonthStatusFailed(s *response.TopupResponseMonthStatusFailed) *pb.TopupMonthStatusFailedResponse {
	return &pb.TopupMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *topupProtoMapper) ToResponsesTopupMonthStatusFailed(topups []*response.TopupResponseMonthStatusFailed) []*pb.TopupMonthStatusFailedResponse {
	var topupRecords []*pb.TopupMonthStatusFailedResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToResponseTopupMonthStatusFailed(topup))
	}

	return topupRecords
}

func (t *topupProtoMapper) ToTopupResponseYearStatusFailed(s *response.TopupResponseYearStatusFailed) *pb.TopupYearStatusFailedResponse {
	return &pb.TopupYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *topupProtoMapper) ToTopupResponsesYearStatusFailed(topups []*response.TopupResponseYearStatusFailed) []*pb.TopupYearStatusFailedResponse {
	var topupRecords []*pb.TopupYearStatusFailedResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseYearStatusFailed(topup))
	}

	return topupRecords
}

func (t *topupProtoMapper) ToResponseTopupMonthlyMethod(s *response.TopupMonthMethodResponse) *pb.TopupMonthMethodResponse {
	return &pb.TopupMonthMethodResponse{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int32(s.TotalTopups),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *topupProtoMapper) ToResponseTopupMonthlyMethods(s []*response.TopupMonthMethodResponse) []*pb.TopupMonthMethodResponse {
	var topupProtos []*pb.TopupMonthMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.ToResponseTopupMonthlyMethod(topup))
	}
	return topupProtos
}

func (t *topupProtoMapper) ToResponseTopupYearlyMethod(s *response.TopupYearlyMethodResponse) *pb.TopupYearlyMethodResponse {
	return &pb.TopupYearlyMethodResponse{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int32(s.TotalTopups),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *topupProtoMapper) ToResponseTopupYearlyMethods(s []*response.TopupYearlyMethodResponse) []*pb.TopupYearlyMethodResponse {
	var topupProtos []*pb.TopupYearlyMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.ToResponseTopupYearlyMethod(topup))
	}
	return topupProtos
}

func (t *topupProtoMapper) ToResponseTopupMonthlyAmount(s *response.TopupMonthAmountResponse) *pb.TopupMonthAmountResponse {
	return &pb.TopupMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *topupProtoMapper) ToResponseTopupMonthlyAmounts(s []*response.TopupMonthAmountResponse) []*pb.TopupMonthAmountResponse {
	var topupProtos []*pb.TopupMonthAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.ToResponseTopupMonthlyAmount(topup))
	}
	return topupProtos
}

func (t *topupProtoMapper) ToResponseTopupYearlyAmount(s *response.TopupYearlyAmountResponse) *pb.TopupYearlyAmountResponse {
	return &pb.TopupYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *topupProtoMapper) ToResponseTopupYearlyAmounts(s []*response.TopupYearlyAmountResponse) []*pb.TopupYearlyAmountResponse {
	var topupProtos []*pb.TopupYearlyAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.ToResponseTopupYearlyAmount(topup))
	}
	return topupProtos
}
