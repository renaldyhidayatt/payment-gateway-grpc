package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
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
	totalRecords := len(res)

	return r.mapping.ToTransfersRecord(res), totalRecords, nil
}

func (r *transferRepository) FindById(id int) (*record.TransferRecord, error) {
	transfer, err := r.db.GetTransferByID(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find by transfer: %w", err)
	}

	return r.mapping.ToTransferRecord(transfer), nil
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

func (r *transferRepository) CountAllTransfers() (int, error) {
	res, err := r.db.CountAllTransfers(r.ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to count all transfers: %w", err)
	}

	return int(res), nil
}

func (r *transferRepository) FindByActive() ([]*record.TransferRecord, error) {
	res, err := r.db.GetActiveTransfers(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find active merchant: %w", err)
	}

	return r.mapping.ToTransfersRecord(res), nil
}

func (r *transferRepository) FindByTrashed() ([]*record.TransferRecord, error) {
	res, err := r.db.GetTrashedTransfers(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find trashed merchant: %w", err)
	}

	return r.mapping.ToTransfersRecord(res), nil
}

func (r *transferRepository) CreateTransfer(request requests.CreateTransferRequest) (*record.TransferRecord, error) {
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

func (r *transferRepository) UpdateTransfer(request requests.UpdateTransferRequest) (*record.TransferRecord, error) {
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

func (r *transferRepository) UpdateTransferAmount(request requests.UpdateTransferAmountRequest) (*record.TransferRecord, error) {
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

func (r *transferRepository) DeleteTransferPermanent(topup_id int) error {
	err := r.db.DeleteTransferPermanently(r.ctx, int32(topup_id))

	if err != nil {
		return nil
	}

	return fmt.Errorf("failed to delete transfer: %w", err)
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
