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
		Id:              int32(transaction.ID),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int32(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantId:      int32(transaction.MerchantID),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}

}

func (m *transactionProtoMapper) ToResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse {
	var result []*pb.TransactionResponse
	for _, transaction := range transactions {
		result = append(result, m.ToResponseTransaction(transaction))
	}
	return result
}

func (m *transactionProtoMapper) ToResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *pb.TransactionResponseDeleteAt {
	return &pb.TransactionResponseDeleteAt{
		Id:              int32(transaction.ID),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int32(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantId:      int32(transaction.MerchantID),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
		DeletedAt:       transaction.DeletedAt,
	}

}

func (m *transactionProtoMapper) ToResponsesTransactionDeleteAt(transactions []*response.TransactionResponseDeleteAt) []*pb.TransactionResponseDeleteAt {
	var result []*pb.TransactionResponseDeleteAt

	for _, transaction := range transactions {
		result = append(result, m.ToResponseTransactionDeleteAt(transaction))
	}
	return result
}

func (t *transactionProtoMapper) ToResponseTransactionMonthStatusSuccess(s *response.TransactionResponseMonthStatusSuccess) *pb.TransactionMonthStatusSuccessResponse {
	return &pb.TransactionMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *transactionProtoMapper) ToResponsesTransactionMonthStatusSuccess(Transactions []*response.TransactionResponseMonthStatusSuccess) []*pb.TransactionMonthStatusSuccessResponse {
	var TransactionRecords []*pb.TransactionMonthStatusSuccessResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToResponseTransactionMonthStatusSuccess(Transaction))
	}

	return TransactionRecords
}

func (t *transactionProtoMapper) ToTransactionResponseYearStatusSuccess(s *response.TransactionResponseYearStatusSuccess) *pb.TransactionYearStatusSuccessResponse {
	return &pb.TransactionYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *transactionProtoMapper) ToTransactionResponsesYearStatusSuccess(Transactions []*response.TransactionResponseYearStatusSuccess) []*pb.TransactionYearStatusSuccessResponse {
	var TransactionRecords []*pb.TransactionYearStatusSuccessResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseYearStatusSuccess(Transaction))
	}

	return TransactionRecords
}

func (t *transactionProtoMapper) ToResponseTransactionMonthStatusFailed(s *response.TransactionResponseMonthStatusFailed) *pb.TransactionMonthStatusFailedResponse {
	return &pb.TransactionMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *transactionProtoMapper) ToResponsesTransactionMonthStatusFailed(Transactions []*response.TransactionResponseMonthStatusFailed) []*pb.TransactionMonthStatusFailedResponse {
	var TransactionRecords []*pb.TransactionMonthStatusFailedResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToResponseTransactionMonthStatusFailed(Transaction))
	}

	return TransactionRecords
}

func (t *transactionProtoMapper) ToTransactionResponseYearStatusFailed(s *response.TransactionResponseYearStatusFailed) *pb.TransactionYearStatusFailedResponse {
	return &pb.TransactionYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *transactionProtoMapper) ToTransactionResponsesYearStatusFailed(Transactions []*response.TransactionResponseYearStatusFailed) []*pb.TransactionYearStatusFailedResponse {
	var TransactionRecords []*pb.TransactionYearStatusFailedResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseYearStatusFailed(Transaction))
	}

	return TransactionRecords
}

func (m *transactionProtoMapper) ToResponseTransactionMonthMethod(s *response.TransactionMonthMethodResponse) *pb.TransactionMonthMethodResponse {
	return &pb.TransactionMonthMethodResponse{
		Month:             s.Month,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int32(s.TotalTransactions),
		TotalAmount:       int32(s.TotalAmount),
	}
}

func (m *transactionProtoMapper) ToResponseTransactionMonthMethods(s []*response.TransactionMonthMethodResponse) []*pb.TransactionMonthMethodResponse {
	var protoResponses []*pb.TransactionMonthMethodResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.ToResponseTransactionMonthMethod(transaction))
	}
	return protoResponses
}

func (m *transactionProtoMapper) ToResponseTransactionYearMethod(s *response.TransactionYearMethodResponse) *pb.TransactionYearMethodResponse {
	return &pb.TransactionYearMethodResponse{
		Year:              s.Year,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int32(s.TotalTransactions),
		TotalAmount:       int32(s.TotalAmount),
	}
}

func (m *transactionProtoMapper) ToResponseTransactionYearMethods(s []*response.TransactionYearMethodResponse) []*pb.TransactionYearMethodResponse {
	var protoResponses []*pb.TransactionYearMethodResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.ToResponseTransactionYearMethod(transaction))
	}
	return protoResponses
}

func (m *transactionProtoMapper) ToResponseTransactionMonthAmount(s *response.TransactionMonthAmountResponse) *pb.TransactionMonthAmountResponse {
	return &pb.TransactionMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (m *transactionProtoMapper) ToResponseTransactionMonthAmounts(s []*response.TransactionMonthAmountResponse) []*pb.TransactionMonthAmountResponse {
	var protoResponses []*pb.TransactionMonthAmountResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.ToResponseTransactionMonthAmount(transaction))
	}
	return protoResponses
}

func (m *transactionProtoMapper) ToResponseTransactionYearlyAmount(s *response.TransactionYearlyAmountResponse) *pb.TransactionYearlyAmountResponse {
	return &pb.TransactionYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (m *transactionProtoMapper) ToResponseTransactionYearlyAmounts(s []*response.TransactionYearlyAmountResponse) []*pb.TransactionYearlyAmountResponse {
	var protoResponses []*pb.TransactionYearlyAmountResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.ToResponseTransactionYearlyAmount(transaction))
	}
	return protoResponses
}
