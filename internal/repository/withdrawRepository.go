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

	req := db.SearchWithdrawsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	withdraw, err := r.db.SearchWithdraws(r.ctx, req)

	if err != nil {
		return nil, 0, errors.New("failed get withdraw")
	}

	totalRecords := len(withdraw)

	return r.mapping.ToWithdrawsRecord(withdraw), totalRecords, nil

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
func (r *withdrawRepository) FindByActive() ([]*record.WithdrawRecord, error) {
	res, err := r.db.GetActiveWithdraws(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find active withdraw: %w", err)
	}

	return r.mapping.ToWithdrawsRecord(res), nil
}

func (r *withdrawRepository) FindByTrashed() ([]*record.WithdrawRecord, error) {
	res, err := r.db.GetTrashedWithdraws(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find trashed merchant: %w", err)
	}

	return r.mapping.ToWithdrawsRecord(res), nil
}

func (r *withdrawRepository) CountActiveByDate(date time.Time) (int64, error) {
	res, err := r.db.CountActiveWithdrawsByDate(r.ctx, date)

	if err != nil {
		return 0, fmt.Errorf("failed to count active by date: %w", err)
	}

	return int64(res), nil
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

func (r *withdrawRepository) DeleteWithdrawPermanent(WithdrawID int) error {
	err := r.db.DeleteWithdrawPermanently(r.ctx, int32(WithdrawID))

	if err != nil {
		return nil
	}

	return fmt.Errorf("failed to delete withdraw: %w", err)
}
