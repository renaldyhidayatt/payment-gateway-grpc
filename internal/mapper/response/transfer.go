package responsemapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
)

type transferResponseMapper struct {
}

func NewTransferResponseMapper() *transferResponseMapper {
	return &transferResponseMapper{}
}

func (s *transferResponseMapper) ToTransferResponse(transfer *record.TransferRecord) *response.TransferResponse {
	return &response.TransferResponse{
		ID:             transfer.ID,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: transfer.TransferAmount,
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}

func (s *transferResponseMapper) ToTransfersResponse(transfers []*record.TransferRecord) []*response.TransferResponse {
	var responses []*response.TransferResponse

	for _, response := range transfers {
		responses = append(responses, s.ToTransferResponse(response))
	}

	return responses
}
