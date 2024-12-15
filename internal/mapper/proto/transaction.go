package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type transactionProtoMapper struct{}

func NewTransactionProtoMapper() *transactionProtoMapper {
	return &transactionProtoMapper{}
}

func (m *transactionProtoMapper) ToResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse {
	return &pb.TransactionResponse{
		Id:            int32(transaction.ID),
		CardNumber:    transaction.CardNumber,
		Amount:        int32(transaction.Amount),
		PaymentMethod: transaction.PaymentMethod,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}

}

func (m *transactionProtoMapper) ToResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse {
	var result []*pb.TransactionResponse
	for _, transaction := range transactions {
		result = append(result, m.ToResponseTransaction(transaction))
	}
	return result
}
