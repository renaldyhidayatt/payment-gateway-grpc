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
