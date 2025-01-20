package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type transferProtoMapper struct{}

func NewTransferProtoMapper() *transferProtoMapper {
	return &transferProtoMapper{}
}

func (t *transferProtoMapper) ToResponseTransfer(transfer *response.TransferResponse) *pb.TransferResponse {
	return &pb.TransferResponse{
		Id:             int32(transfer.ID),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}

func (t *transferProtoMapper) ToResponsesTransfer(transfers []*response.TransferResponse) []*pb.TransferResponse {
	var responses []*pb.TransferResponse

	for _, response := range transfers {
		responses = append(responses, t.ToResponseTransfer(response))
	}

	return responses
}

func (t *transferProtoMapper) ToResponseTransferDeleteAt(transfer *response.TransferResponseDeleteAt) *pb.TransferResponseDeleteAt {
	return &pb.TransferResponseDeleteAt{
		Id:             int32(transfer.ID),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
		DeletedAt:      transfer.DeletedAt,
	}
}

func (t *transferProtoMapper) ToResponsesTransferDeleteAt(transfers []*response.TransferResponseDeleteAt) []*pb.TransferResponseDeleteAt {
	var responses []*pb.TransferResponseDeleteAt

	for _, response := range transfers {
		responses = append(responses, t.ToResponseTransferDeleteAt(response))
	}

	return responses
}

func (t *transferProtoMapper) ToResponseTransferMonthStatusSuccess(s *response.TransferResponseMonthStatusSuccess) *pb.TransferMonthStatusSuccessResponse {
	return &pb.TransferMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *transferProtoMapper) ToResponsesTransferMonthStatusSuccess(Transfers []*response.TransferResponseMonthStatusSuccess) []*pb.TransferMonthStatusSuccessResponse {
	var TransferRecords []*pb.TransferMonthStatusSuccessResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToResponseTransferMonthStatusSuccess(Transfer))
	}

	return TransferRecords
}

func (t *transferProtoMapper) ToTransferResponseYearStatusSuccess(s *response.TransferResponseYearStatusSuccess) *pb.TransferYearStatusSuccessResponse {
	return &pb.TransferYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *transferProtoMapper) ToTransferResponsesYearStatusSuccess(Transfers []*response.TransferResponseYearStatusSuccess) []*pb.TransferYearStatusSuccessResponse {
	var TransferRecords []*pb.TransferYearStatusSuccessResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseYearStatusSuccess(Transfer))
	}

	return TransferRecords
}

func (t *transferProtoMapper) ToResponseTransferMonthStatusFailed(s *response.TransferResponseMonthStatusFailed) *pb.TransferMonthStatusFailedResponse {
	return &pb.TransferMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *transferProtoMapper) ToResponsesTransferMonthStatusFailed(Transfers []*response.TransferResponseMonthStatusFailed) []*pb.TransferMonthStatusFailedResponse {
	var TransferRecords []*pb.TransferMonthStatusFailedResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToResponseTransferMonthStatusFailed(Transfer))
	}

	return TransferRecords
}

func (t *transferProtoMapper) ToTransferResponseYearStatusFailed(s *response.TransferResponseYearStatusFailed) *pb.TransferYearStatusFailedResponse {
	return &pb.TransferYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *transferProtoMapper) ToTransferResponsesYearStatusFailed(Transfers []*response.TransferResponseYearStatusFailed) []*pb.TransferYearStatusFailedResponse {
	var TransferRecords []*pb.TransferYearStatusFailedResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseYearStatusFailed(Transfer))
	}

	return TransferRecords
}

func (m *transferProtoMapper) ToResponseTransferMonthAmount(s *response.TransferMonthAmountResponse) *pb.TransferMonthAmountResponse {
	return &pb.TransferMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (m *transferProtoMapper) ToResponseTransferMonthAmounts(s []*response.TransferMonthAmountResponse) []*pb.TransferMonthAmountResponse {
	var protoResponses []*pb.TransferMonthAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.ToResponseTransferMonthAmount(transfer))
	}
	return protoResponses
}

func (m *transferProtoMapper) ToResponseTransferYearAmount(s *response.TransferYearAmountResponse) *pb.TransferYearAmountResponse {
	return &pb.TransferYearAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (m *transferProtoMapper) ToResponseTransferYearAmounts(s []*response.TransferYearAmountResponse) []*pb.TransferYearAmountResponse {
	var protoResponses []*pb.TransferYearAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.ToResponseTransferYearAmount(transfer))
	}
	return protoResponses
}
