package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"
	"database/sql"
	"fmt"
)

type saldoRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.SaldoRecordMapping
}

func NewSaldoRepository(db *db.Queries, ctx context.Context, mapping recordmapper.SaldoRecordMapping) *saldoRepository {
	return &saldoRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *saldoRepository) FindAllSaldos(search string, page, pageSize int) ([]*record.SaldoRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetSaldosParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	saldos, err := r.db.GetSaldos(r.ctx, req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find saldos: %w", err)
	}

	totalRecords := len(saldos)

	return r.mapping.ToSaldosRecord(saldos), totalRecords, nil
}

func (r *saldoRepository) FindByCardNumber(card_number string) (*record.SaldoRecord, error) {
	res, err := r.db.GetSaldoByCardNumber(r.ctx, card_number)

	if err != nil {
		return nil, fmt.Errorf("failed to find card number saldo: %w", err)
	}

	return r.mapping.ToSaldoRecord(res), nil
}

func (r *saldoRepository) FindById(saldo_id int) (*record.SaldoRecord, error) {
	res, err := r.db.GetSaldoByID(r.ctx, int32(saldo_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find saldo: %w", err)
	}

	return r.mapping.ToSaldoRecord(res), nil
}

func (r *saldoRepository) FindByActive() ([]*record.SaldoRecord, error) {
	res, err := r.db.GetActiveSaldos(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find active: %w", err)
	}

	return r.mapping.ToSaldosRecord(res), nil

}

func (r *saldoRepository) FindByTrashed() ([]*record.SaldoRecord, error) {
	saldos, err := r.db.GetTrashedSaldos(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get trashed saldos: %w", err)
	}

	return r.mapping.ToSaldosRecord(saldos), nil
}

func (r *saldoRepository) CreateSaldo(request *requests.CreateSaldoRequest) (*record.SaldoRecord, error) {
	req := db.CreateSaldoParams{
		CardNumber:   request.CardNumber,
		TotalBalance: int32(request.TotalBalance),
	}
	res, err := r.db.CreateSaldo(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to create saldo")
	}

	return r.mapping.ToSaldoRecord(res), nil
}

func (r *saldoRepository) UpdateSaldo(request *requests.UpdateSaldoRequest) (*record.SaldoRecord, error) {
	req := db.UpdateSaldoParams{
		SaldoID:      int32(request.SaldoID),
		CardNumber:   request.CardNumber,
		TotalBalance: int32(request.TotalBalance),
	}

	err := r.db.UpdateSaldo(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update saldo")
	}

	saldo, err := r.db.GetSaldoByID(r.ctx, req.SaldoID)

	if err != nil {
		return nil, fmt.Errorf("failed to update saldo")
	}

	return r.mapping.ToSaldoRecord(saldo), nil
}

func (r *saldoRepository) UpdateSaldoBalance(request *requests.UpdateSaldoBalance) (*record.SaldoRecord, error) {
	req := db.UpdateSaldoBalanceParams{
		CardNumber:   request.CardNumber,
		TotalBalance: int32(request.TotalBalance),
	}

	err := r.db.UpdateSaldoBalance(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update balance saldo: %w", err)
	}

	res, err := r.db.GetSaldoByCardNumber(r.ctx, request.CardNumber)

	if err != nil {
		return nil, fmt.Errorf("failed to found saldo by card number: %w", err)
	}

	return r.mapping.ToSaldoRecord(res), nil
}

func (r *saldoRepository) TrashedSaldo(saldoID int) (*record.SaldoRecord, error) {
	err := r.db.TrashSaldo(r.ctx, int32(saldoID))
	if err != nil {
		return nil, fmt.Errorf("failed to trash saldo: %w", err)
	}

	saldo, err := r.db.GetTrashedSaldoByID(r.ctx, int32(saldoID))
	if err != nil {
		return nil, fmt.Errorf("saldo not found after trashing: %w", err)
	}

	return r.mapping.ToSaldoRecord(saldo), nil
}

func (r *saldoRepository) RestoreSaldo(saldoID int) (*record.SaldoRecord, error) {
	err := r.db.RestoreSaldo(r.ctx, int32(saldoID))

	if err != nil {
		return nil, fmt.Errorf("failed to restore saldo: %w", err)
	}

	saldo, err := r.db.GetSaldoByID(r.ctx, int32(saldoID))

	if err != nil {
		return nil, fmt.Errorf("saldo not found restore saldo: %w", err)
	}

	return r.mapping.ToSaldoRecord(saldo), nil
}

func (r *saldoRepository) UpdateSaldoWithdraw(request *requests.UpdateSaldoWithdraw) (*record.SaldoRecord, error) {
	withdrawAmount := sql.NullInt32{
		Int32: int32(*request.WithdrawAmount),
		Valid: request.WithdrawAmount != nil,
	}
	var withdrawTime sql.NullTime
	if request.WithdrawTime != nil {
		withdrawTime = sql.NullTime{
			Time:  *request.WithdrawTime,
			Valid: true,
		}
	}

	req := db.UpdateSaldoWithdrawParams{
		CardNumber:     request.CardNumber,
		WithdrawAmount: withdrawAmount,
		WithdrawTime:   withdrawTime,
	}

	err := r.db.UpdateSaldoWithdraw(r.ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update saldo for card number %s: %w", request.CardNumber, err)
	}

	saldo, err := r.db.GetSaldoByCardNumber(r.ctx, request.CardNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated saldo for card number %s: %w", request.CardNumber, err)
	}

	return r.mapping.ToSaldoRecord(saldo), nil
}

func (r *saldoRepository) DeleteSaldoPermanent(saldo_id int) error {
	err := r.db.DeleteSaldoPermanently(r.ctx, int32(saldo_id))

	if err != nil {
		return nil
	}

	return fmt.Errorf("failed delete saldo permanent")
}
