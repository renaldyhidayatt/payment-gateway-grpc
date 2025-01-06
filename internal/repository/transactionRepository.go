package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"context"
	"fmt"
	"time"
)

type transactionRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.TransactionRecordMapping
}

func NewTransactionRepository(db *db.Queries, ctx context.Context, mapping recordmapper.TransactionRecordMapping) *transactionRepository {
	return &transactionRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *transactionRepository) FindAllTransactions(search string, page, pageSize int) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTransactionsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	transactions, err := r.db.GetTransactions(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find transactions: %w", err)
	}

	var totalCount int
	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionsRecordAll(transactions), totalCount, nil
}

func (r *transactionRepository) FindById(transaction_id int) (*record.TransactionRecord, error) {
	res, err := r.db.GetTransactionByID(r.ctx, int32(transaction_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find transaction: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) FindByActive(search string, page, pageSize int) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetActiveTransactionsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetActiveTransactions(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active transactions: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionsRecordActive(res), totalCount, nil
}

func (r *transactionRepository) FindByTrashed(search string, page, pageSize int) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTrashedTransactionsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTrashedTransactions(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find trashed transactions: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionsRecordTrashed(res), totalCount, nil
}

func (r *transactionRepository) CountTransactionsByDate(date string) (int, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, fmt.Errorf("invalid date format: %w", err)
	}

	res, err := r.db.CountTransactionsByDate(r.ctx, parsedDate)
	if err != nil {
		return 0, fmt.Errorf("failed to count transactions by date %s: %w", date, err)
	}

	return int(res), nil
}

func (r *transactionRepository) CountAllTransactions() (*int64, error) {
	res, err := r.db.CountAllTransactions(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("faield to count transaction: %w", err)
	}

	return &res, nil
}

func (r *transactionRepository) CountTransactions(search string) (*int64, error) {
	res, err := r.db.CountTransactions(r.ctx, search)

	if err != nil {
		return nil, fmt.Errorf("faield to count transaction by search: %w", err)
	}

	return &res, nil
}

func (r *transactionRepository) CreateTransaction(request *requests.CreateTransactionRequest) (*record.TransactionRecord, error) {
	req := db.CreateTransactionParams{
		CardNumber:      request.CardNumber,
		Amount:          int32(request.Amount),
		PaymentMethod:   request.PaymentMethod,
		MerchantID:      int32(*request.MerchantID),
		TransactionTime: request.TransactionTime,
	}

	res, err := r.db.CreateTransaction(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) UpdateTransaction(request *requests.UpdateTransactionRequest) (*record.TransactionRecord, error) {
	req := db.UpdateTransactionParams{
		TransactionID:   int32(request.TransactionID),
		CardNumber:      request.CardNumber,
		Amount:          int32(request.Amount),
		PaymentMethod:   request.PaymentMethod,
		MerchantID:      int32(*request.MerchantID),
		TransactionTime: request.TransactionTime,
	}

	err := r.db.UpdateTransaction(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}

	res, err := r.db.GetTransactionByID(r.ctx, int32(request.TransactionID))

	if err != nil {
		return nil, fmt.Errorf("failed to find transaction: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) TrashedTransaction(transaction_id int) (*record.TransactionRecord, error) {
	err := r.db.TrashTransaction(r.ctx, int32(transaction_id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash transaction: %w", err)
	}

	transaction, err := r.db.GetTrashedTransactionByID(r.ctx, int32(transaction_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find trashed by id transaction: %w", err)
	}

	return r.mapping.ToTransactionRecord(transaction), nil
}

func (r *transactionRepository) RestoreTransaction(topup_id int) (*record.TransactionRecord, error) {
	err := r.db.RestoreTransaction(r.ctx, int32(topup_id))

	if err != nil {
		return nil, fmt.Errorf("failed to restore transaction: %w", err)
	}

	topup, err := r.db.GetTransactionByID(r.ctx, int32(topup_id))

	if err != nil {
		return nil, fmt.Errorf("failed not found transaction :%w", err)
	}

	return r.mapping.ToTransactionRecord(topup), nil
}

func (r *transactionRepository) DeleteTransactionPermanent(topup_id int) error {
	err := r.db.DeleteTransactionPermanently(r.ctx, int32(topup_id))

	if err != nil {
		return nil
	}

	return fmt.Errorf("failed to delete transaction: %w", err)
}

func (r *transactionRepository) FindByCardNumber(card_number string) ([]*record.TransactionRecord, error) {
	res, err := r.db.GetTransactionsByCardNumber(r.ctx, card_number)

	if err != nil {
		return nil, fmt.Errorf("failed to find transaction by card number: %w", err)
	}

	return r.mapping.ToTransactionsRecord(res), nil
}

func (r *transactionRepository) FindTransactionByMerchantId(merchant_id int) ([]*record.TransactionRecord, error) {
	res, err := r.db.GetTransactionsByMerchantID(r.ctx, int32(merchant_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find transaction by merchant id: %w", err)
	}

	return r.mapping.ToTransactionsRecord(res), nil
}
