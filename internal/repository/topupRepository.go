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

type topupRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.TopupRecordMapping
}

func NewTopupRepository(db *db.Queries, ctx context.Context, mapping recordmapper.TopupRecordMapping) *topupRepository {
	return &topupRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *topupRepository) FindAllTopups(search string, page, pageSize int) ([]*record.TopupRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTopupsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTopups(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find topups: %w", err)
	}

	totalRecords := len(res)

	return r.mapping.ToTopupRecords(res), totalRecords, nil
}

func (r *topupRepository) FindById(topup_id int) (*record.TopupRecord, error) {
	res, err := r.db.GetTopupByID(r.ctx, int32(topup_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find topup: %w", err)
	}

	return r.mapping.ToTopupRecord(res), nil
}

func (r *topupRepository) CountTopupsByDate(date string) (int, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, fmt.Errorf("invalid date format: %w", err)
	}

	res, err := r.db.CountTopupsByDate(r.ctx, parsedDate)
	if err != nil {
		return 0, fmt.Errorf("failed to count topups by date %s: %w", date, err)
	}

	return int(res), nil
}

func (r *topupRepository) CountAllTopups() (int, error) {
	res, err := r.db.CountAllTopups(r.ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to count all topups: %w", err)
	}

	return int(res), nil
}

func (r *topupRepository) FindByActive() ([]*record.TopupRecord, error) {
	res, err := r.db.GetActiveTopups(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find active merchant: %w", err)
	}

	return r.mapping.ToTopupRecords(res), nil
}

func (r *topupRepository) FindByTrashed() ([]*record.TopupRecord, error) {
	res, err := r.db.GetTrashedTopups(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find trashed merchant: %w", err)
	}

	return r.mapping.ToTopupRecords(res), nil
}

func (r *topupRepository) CreateTopup(request *requests.CreateTopupRequest) (*record.TopupRecord, error) {
	req := db.CreateTopupParams{
		CardNumber:  request.CardNumber,
		TopupNo:     request.TopupNo,
		TopupAmount: int32(request.TopupAmount),
		TopupMethod: request.TopupMethod,
	}

	res, err := r.db.CreateTopup(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to create topup: %w", err)
	}

	return r.mapping.ToTopupRecord(res), nil
}

func (r *topupRepository) UpdateTopup(request *requests.UpdateTopupRequest) (*record.TopupRecord, error) {
	req := db.UpdateTopupParams{
		TopupID:     int32(request.TopupID),
		CardNumber:  request.CardNumber,
		TopupAmount: int32(request.TopupAmount),
		TopupMethod: request.TopupMethod,
	}
	err := r.db.UpdateTopup(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update topup: %w", err)
	}

	res, err := r.db.GetTopupByID(r.ctx, req.TopupID)

	if err != nil {
		return nil, fmt.Errorf("failed to find topup: %w", err)
	}

	return r.mapping.ToTopupRecord(res), nil
}

func (r *topupRepository) UpdateTopupAmount(request *requests.UpdateTopupAmount) (*record.TopupRecord, error) {
	req := db.UpdateTopupAmountParams{
		TopupID:     int32(request.TopupID),
		TopupAmount: int32(request.TopupAmount),
	}

	err := r.db.UpdateTopupAmount(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update topup amount :%w", err)
	}

	res, err := r.db.GetTopupByID(r.ctx, req.TopupID)

	if err != nil {
		return nil, fmt.Errorf("failed to find topup: %w", err)
	}

	return r.mapping.ToTopupRecord(res), nil
}

func (r *topupRepository) TrashedTopup(topup_id int) (*record.TopupRecord, error) {
	err := r.db.TrashTopup(r.ctx, int32(topup_id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash topup: %w", err)
	}

	merchant, err := r.db.GetTrashedTopupByID(r.ctx, int32(topup_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find trashed by id topup: %w", err)
	}

	return r.mapping.ToTopupRecord(merchant), nil
}

func (r *topupRepository) RestoreTopup(topup_id int) (*record.TopupRecord, error) {
	err := r.db.RestoreTopup(r.ctx, int32(topup_id))

	if err != nil {
		return nil, fmt.Errorf("failed to restore topup: %w", err)
	}

	topup, err := r.db.GetTopupByID(r.ctx, int32(topup_id))

	if err != nil {
		return nil, fmt.Errorf("failed not found topup :%w", err)
	}

	return r.mapping.ToTopupRecord(topup), nil
}

func (r *topupRepository) DeleteTopupPermanent(topup_id int) error {
	err := r.db.DeleteTopupPermanently(r.ctx, int32(topup_id))

	if err != nil {
		return nil
	}

	return fmt.Errorf("failed to delete topup: %w", err)
}

func (r *topupRepository) FindByCardNumber(card_number string) ([]*record.TopupRecord, error) {
	res, err := r.db.GetTopupsByCardNumber(r.ctx, card_number)

	if err != nil {
		return nil, fmt.Errorf("failed to find topup by card number: %w", err)
	}

	return r.mapping.ToTopupRecords(res), nil
}
