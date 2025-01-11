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

type transferRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.TransferRecordMapping
}

func NewTransferRepository(db *db.Queries, ctx context.Context, mapping recordmapper.TransferRecordMapping) *transferRepository {
	return &transferRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *transferRepository) FindAll(search string, page, pageSize int) ([]*record.TransferRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTransfersParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransfers(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find transfers: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransfersRecordAll(res), totalCount, nil
}

func (r *transferRepository) FindById(id int) (*record.TransferRecord, error) {
	transfer, err := r.db.GetTransferByID(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find by transfer: %w", err)
	}

	return r.mapping.ToTransferRecord(transfer), nil
}

func (r *transferRepository) FindTransferByTransferFrom(transfer_from string) ([]*record.TransferRecord, error) {
	res, err := r.db.GetTransfersBySourceCard(r.ctx, transfer_from)

	if err != nil {
		return nil, fmt.Errorf("failed to find transfer by transfer from: %w", err)
	}

	return r.mapping.ToTransfersRecord(res), nil
}

func (r *transferRepository) FindTransferByTransferTo(transfer_to string) ([]*record.TransferRecord, error) {
	res, err := r.db.GetTransfersByDestinationCard(r.ctx, transfer_to)

	if err != nil {
		return nil, fmt.Errorf("failed to find transfer by transfer to: %w", err)
	}

	return r.mapping.ToTransfersRecord(res), nil
}

func (r *transferRepository) CountTransfersByDate(date string) (int, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, fmt.Errorf("invalid date format: %w", err)
	}

	res, err := r.db.CountTransfersByDate(r.ctx, parsedDate)
	if err != nil {
		return 0, fmt.Errorf("failed to count transfers by date %s: %w", date, err)
	}

	return int(res), nil
}

// func (r *transferRepository) GetMonthlyAmounts() {
// 	res, err := r.db.GetMonthlyTransferAmounts(r.ctx)
// }

// func (r *transferRepository) GetYearlyAmounts() {
// 	res, err := r.db.GetYearlyTransferAmounts(r.ctx)

// }

// func (r *transferRepository) GetMonthlyTransferAmountsBySenderCardNumber() {
// 	res, err := r.db.GetMonthlyTransferAmountsBySenderCardNumber(r.ctx)
// }

// func (r *transferRepository) GetYearlyTransferAmountsBySenderCardNumber() {
// 	res, err := r.db.GetYearlyTransferAmountsBySenderCardNumber(r.ctx)
// }

// func (r *transferRepository) GetMonthlyTransferAmountsByReceiverCardNumber() {
// 	res, err := r.db.GetMonthlyTransferAmountsByReceiverCardNumber(r.ctx)
// }

// func (r *transferRepository) GetYearlyTransferAmountsByReceiverCardNumber() {
// 	res, err := r.db.GetYearlyTransferAmountsByReceiverCardNumber(r.ctx)
// }

func (r *transferRepository) FindByActive(search string, page, pageSize int) ([]*record.TransferRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetActiveTransfersParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetActiveTransfers(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active merchant: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransfersRecordActive(res), totalCount, nil
}

func (r *transferRepository) FindByTrashed(search string, page, pageSize int) ([]*record.TransferRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTrashedTransfersParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTrashedTransfers(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find trashed merchant: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransfersRecordTrashed(res), totalCount, nil
}

func (r *transferRepository) CountAllTransfers() (*int64, error) {
	res, err := r.db.CountAllTransfers(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("faield to count transfer: %w", err)
	}

	return &res, nil
}

func (r *transferRepository) CountTransfers(search string) (*int64, error) {
	res, err := r.db.CountTransfers(r.ctx, search)

	if err != nil {
		return nil, fmt.Errorf("faield to count transfer by search: %w", err)
	}

	return &res, nil
}

func (r *transferRepository) CreateTransfer(request *requests.CreateTransferRequest) (*record.TransferRecord, error) {
	req := db.CreateTransferParams{
		TransferFrom:   request.TransferFrom,
		TransferTo:     request.TransferTo,
		TransferAmount: int32(request.TransferAmount),
	}

	res, err := r.db.CreateTransfer(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to create transfer: %w", err)
	}

	return r.mapping.ToTransferRecord(res), nil
}

func (r *transferRepository) UpdateTransfer(request *requests.UpdateTransferRequest) (*record.TransferRecord, error) {
	req := db.UpdateTransferParams{
		TransferID:     int32(request.TransferID),
		TransferFrom:   request.TransferFrom,
		TransferTo:     request.TransferTo,
		TransferAmount: int32(request.TransferAmount),
	}

	err := r.db.UpdateTransfer(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update transfer: %w", err)
	}

	res, err := r.db.GetTransferByID(r.ctx, int32(request.TransferID))

	if err != nil {
		return nil, fmt.Errorf("failed to find transfer: %w", err)
	}

	return r.mapping.ToTransferRecord(res), nil

}

func (r *transferRepository) UpdateTransferAmount(request *requests.UpdateTransferAmountRequest) (*record.TransferRecord, error) {
	req := db.UpdateTransferAmountParams{
		TransferID:     int32(request.TransferID),
		TransferAmount: int32(request.TransferAmount),
	}

	err := r.db.UpdateTransferAmount(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update transfer amount: %w", err)
	}

	res, err := r.db.GetTransferByID(r.ctx, int32(request.TransferID))

	if err != nil {
		return nil, fmt.Errorf("failed to find transfer: %w", err)
	}

	return r.mapping.ToTransferRecord(res), nil
}

func (r *transferRepository) TrashedTransfer(transfer_id int) (*record.TransferRecord, error) {
	err := r.db.TrashTransfer(r.ctx, int32(transfer_id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash transfer: %w", err)
	}

	merchant, err := r.db.GetTrashedTransferByID(r.ctx, int32(transfer_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find trashed by id transfer: %w", err)
	}

	return r.mapping.ToTransferRecord(merchant), nil
}

func (r *transferRepository) RestoreTransfer(transfer_id int) (*record.TransferRecord, error) {
	err := r.db.RestoreTransfer(r.ctx, int32(transfer_id))

	if err != nil {
		return nil, fmt.Errorf("failed to restore transfer: %w", err)
	}

	transfer, err := r.db.GetTransferByID(r.ctx, int32(transfer_id))

	if err != nil {
		return nil, fmt.Errorf("failed not found transfer :%w", err)
	}

	return r.mapping.ToTransferRecord(transfer), nil
}

func (r *transferRepository) DeleteTransferPermanent(topup_id int) (bool, error) {
	err := r.db.DeleteTransferPermanently(r.ctx, int32(topup_id))
	if err != nil {
		return false, fmt.Errorf("failed to delete transfer: %w", err)
	}
	return true, nil
}

func (r *transferRepository) RestoreAllTransfer() (bool, error) {
	err := r.db.RestoreAllTransfers(r.ctx)
	if err != nil {
		return false, fmt.Errorf("failed to restore all transfers: %w", err)
	}
	return true, nil
}

func (r *transferRepository) DeleteAllTransferPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentTransfers(r.ctx)
	if err != nil {
		return false, fmt.Errorf("failed to delete all transfers permanently: %w", err)
	}
	return true, nil
}
