package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
)

type transactionRecordMapper struct{}

func NewTransactionRecordMapper() *transactionRecordMapper {
	return &transactionRecordMapper{}
}

func (s *transactionRecordMapper) ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord {
	var deletedAt *string

	if transaction.DeletedAt.Valid {
		formatedDeletedAt := transaction.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.TransactionRecord{
		ID:              int(transaction.TransactionID),
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime.Format("2006-01-02 15:04:05"),
		CreatedAt:       transaction.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:       transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:       deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionRecord(transaction))
	}
	return transactionRecords
}
