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

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTopupRecordsAll(res), totalCount, nil
}

func (r *topupRepository) FindById(topup_id int) (*record.TopupRecord, error) {
	res, err := r.db.GetTopupByID(r.ctx, int32(topup_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find topup: %w", err)
	}

	return r.mapping.ToTopupRecord(res), nil
}

func (r *topupRepository) FindByCardNumber(card_number string) ([]*record.TopupRecord, error) {
	res, err := r.db.GetTopupsByCardNumber(r.ctx, card_number)

	if err != nil {
		return nil, fmt.Errorf("failed to find topup by card number: %w", err)
	}

	return r.mapping.ToTopupRecords(res), nil
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

// func (r *topupRepository) GetMonthlyTopupMethods() {
// 	res, err := r.db.GetMonthlyTopupMethods(r.ctx)
// }

// func (r *topupRepository) GetYearlyTopupMethods() {
// 	res, err := r.db.GetYearlyTopupMethods(r.ctx)
// }

// func (r *topupRepository) GetMonthlyTopupAmounts() {
// 	res, err := r.db.GetMonthlyTopupAmounts(r.ctx)
// }

// func (r *topupRepository) GetYearlyTopupAmounts() {
// 	res, err := r.db.GetYearlyTopupAmounts(r.ctx)
// }

// func (r *topupRepository) GetMonthlyTopupMethodsByCardNumber() {
// 	res, err := r.db.GetMonthlyTopupMethodsByCardNumber(r.ctx)
// }

// func (r *topupRepository) GetMonthlyTopupAmountsByCardNumber() {
// 	res, err := r.db.GetMonthlyTopupAmountsByCardNumber(r.ctx)
// }

// func (r *topupRepository) GetYearlyTopupAmountsByCardNumber() {
// 	res, err := r.db.GetYearlyTopupAmountsByCardNumber(r.ctx)
// }

func (r *topupRepository) CountAllTopups() (*int64, error) {
	res, err := r.db.CountAllTopups(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("faield to count topup: %w", err)
	}

	return &res, nil
}

func (r *topupRepository) CountTopups(search string) (*int64, error) {
	res, err := r.db.CountTopups(r.ctx, search)

	if err != nil {
		return nil, fmt.Errorf("faield to count topup by search: %w", err)
	}

	return &res, nil
}

func (r *topupRepository) FindByActive(search string, page, pageSize int) ([]*record.TopupRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetActiveTopupsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetActiveTopups(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active merchant: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTopupRecordsActive(res), totalCount, nil
}

func (r *topupRepository) FindByTrashed(search string, page, pageSize int) ([]*record.TopupRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTrashedTopupsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTrashedTopups(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find trashed merchant: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTopupRecordsTrashed(res), totalCount, nil
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

func (r *topupRepository) DeleteTopupPermanent(topup_id int) (bool, error) {
	err := r.db.DeleteTopupPermanently(r.ctx, int32(topup_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete topup permanently: %w", err)
	}

	return true, nil
}

func (r *topupRepository) RestoreAllTopup() (bool, error) {
	err := r.db.RestoreAllTopups(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all topups: %w", err)
	}

	return true, nil
}

func (r *topupRepository) DeleteAllTopupPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentTopups(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all topups permanently: %w", err)
	}

	return true, nil
}
