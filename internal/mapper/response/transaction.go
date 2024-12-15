package responsemapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
)

type transactionResponseMapper struct {
}

func NewTransactionResponseMapper() *transactionResponseMapper {
	return &transactionResponseMapper{}
}

func (s *transactionResponseMapper) ToTransactionResponse(transaction record.TransactionRecord) *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:              transaction.ID,
		CardNumber:      transaction.CardNumber,
		Amount:          transaction.Amount,
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}
}

func (s *transactionResponseMapper) ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse {
	responses := make([]*response.TransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		responses = append(responses, s.ToTransactionResponse(*transaction))
	}
	return responses
}
