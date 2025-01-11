package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type withdrawRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.WithdrawRecordMapping
}

func NewWithdrawRepository(db *db.Queries, ctx context.Context, mapping recordmapper.WithdrawRecordMapping) *withdrawRepository {
	return &withdrawRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *withdrawRepository) FindAll(search string, page, pageSize int) ([]*record.WithdrawRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetWithdrawsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	withdraw, err := r.db.GetWithdraws(r.ctx, req)

	if err != nil {
		return nil, 0, errors.New("failed get withdraw")
	}

	var totalCount int
	if len(withdraw) > 0 {
		totalCount = int(withdraw[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToWithdrawsRecordALl(withdraw), totalCount, nil

}

func (r *withdrawRepository) FindById(id int) (*record.WithdrawRecord, error) {
	withdraw, err := r.db.GetWithdrawByID(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get withdraw")
	}

	return r.mapping.ToWithdrawRecord(withdraw), nil
}

func (r *withdrawRepository) FindByCardNumber(card_number string) ([]*record.WithdrawRecord, error) {
	cardNumberSQL := sql.NullString{
		String: card_number,
		Valid:  card_number != "",
	}

	res, err := r.db.SearchWithdrawByCardNumber(r.ctx, cardNumberSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to find card number: %w", err)
	}

	return r.mapping.ToWithdrawsRecord(res), nil
}
func (r *withdrawRepository) FindByActive(search string, page, pageSize int) ([]*record.WithdrawRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetActiveWithdrawsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetActiveWithdraws(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active withdraw: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToWithdrawsRecordActive(res), totalCount, nil
}

func (r *withdrawRepository) FindByTrashed(search string, page, pageSize int) ([]*record.WithdrawRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTrashedWithdrawsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTrashedWithdraws(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find trashed merchant: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToWithdrawsRecordTrashed(res), totalCount, nil
}

// func (r *withdrawRepository) GetMonthly() {
// 	res, err := r.db.GetMonthlyWithdraws(r.ctx)
// }

// func (r *withdrawRepository) GetYearly() {
// 	res, err := r.db.GetYearlyWithdraws(r.ctx)
// }

// func (r *withdrawRepository) GetMonthlyByCardNumber() {
// 	res, err := r.db.GetMonthlyWithdrawsByCardNumber(r.ctx)
// }

// func (r *withdrawRepository) GetYearlyCardNumber() {
// 	res, err := r.db.GetYearlyWithdrawsByCardNumber(r.ctx)
// }

func (r *withdrawRepository) CountActiveByDate(date time.Time) (int64, error) {
	res, err := r.db.CountActiveWithdrawsByDate(r.ctx, date)

	if err != nil {
		return 0, fmt.Errorf("failed to count active by date: %w", err)
	}

	return int64(res), nil
}

func (r *withdrawRepository) CountAllWithdraws() (*int64, error) {
	res, err := r.db.CountAllWithdraws(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("faield to count withdraw: %w", err)
	}

	return &res, nil
}

func (r *withdrawRepository) CountWithdraws(search string) (*int64, error) {
	res, err := r.db.CountWithdraws(r.ctx, search)

	if err != nil {
		return nil, fmt.Errorf("faield to count withdraw by search: %w", err)
	}

	return &res, nil
}

func (r *withdrawRepository) CreateWithdraw(request *requests.CreateWithdrawRequest) (*record.WithdrawRecord, error) {
	req := db.CreateWithdrawParams{
		CardNumber:     request.CardNumber,
		WithdrawAmount: int32(request.WithdrawAmount),
		WithdrawTime:   request.WithdrawTime,
	}

	res, err := r.db.CreateWithdraw(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update withdraw :%w", err)
	}

	return r.mapping.ToWithdrawRecord(res), nil
}

func (r *withdrawRepository) UpdateWithdraw(request *requests.UpdateWithdrawRequest) (*record.WithdrawRecord, error) {
	req := db.UpdateWithdrawParams{
		WithdrawID:     int32(request.WithdrawID),
		CardNumber:     request.CardNumber,
		WithdrawAmount: int32(request.WithdrawAmount),
		WithdrawTime:   request.WithdrawTime,
	}

	err := r.db.UpdateWithdraw(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update withdraw: %w", err)
	}

	res, err := r.db.GetWithdrawByID(r.ctx, req.WithdrawID)

	if err != nil {
		return nil, fmt.Errorf("failed to find withdraw: %w", err)
	}

	return r.mapping.ToWithdrawRecord(res), nil
}

func (r *withdrawRepository) TrashedWithdraw(WithdrawID int) (*record.WithdrawRecord, error) {
	err := r.db.TrashWithdraw(r.ctx, int32(WithdrawID))

	if err != nil {
		return nil, fmt.Errorf("failed to trash withdraw: %w", err)
	}

	merchant, err := r.db.GetTrashedWithdrawByID(r.ctx, int32(WithdrawID))

	if err != nil {
		return nil, fmt.Errorf("failed to find trashed by id topup: %w", err)
	}

	return r.mapping.ToWithdrawRecord(merchant), nil
}

func (r *withdrawRepository) RestoreWithdraw(WithdrawID int) (*record.WithdrawRecord, error) {
	err := r.db.RestoreWithdraw(r.ctx, int32(WithdrawID))

	if err != nil {
		return nil, fmt.Errorf("failed to restore withdraw: %w", err)
	}

	withdraw, err := r.db.GetWithdrawByID(r.ctx, int32(WithdrawID))

	if err != nil {
		return nil, fmt.Errorf("failed not found withdraw :%w", err)
	}

	return r.mapping.ToWithdrawRecord(withdraw), nil
}

func (r *withdrawRepository) DeleteWithdrawPermanent(WithdrawID int) (bool, error) {
	err := r.db.DeleteWithdrawPermanently(r.ctx, int32(WithdrawID))

	if err != nil {
		return false, fmt.Errorf("failed to delete withdraw: %w", err)
	}

	return true, nil
}

func (r *withdrawRepository) RestoreAllWithdraw() (bool, error) {
	err := r.db.RestoreAllWithdraws(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all withdraws: %w", err)
	}

	return true, nil
}

func (r *withdrawRepository) DeleteAllWithdrawPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentWithdraws(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all withdraws permanently: %w", err)
	}

	return true, nil
}
