package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"context"
	"database/sql"
	"fmt"
	"time"
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

	var totalCount int
	if len(saldos) > 0 {
		totalCount = int(saldos[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToSaldosRecordAll(saldos), totalCount, nil
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

func (r *saldoRepository) GetMonthlyTotalSaldoBalance(year int, month int) ([]*record.SaldoMonthTotalBalance, error) {
	currentDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalSaldoBalance(r.ctx, db.GetMonthlyTotalSaldoBalanceParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	for i, row := range res {
		fmt.Printf("Row %d: %+v\n", i, *row)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly total saldo balance: %w", err)
	}
	so := r.mapping.ToSaldoMonthTotalBalances(res)
	return so, nil
}

func (r *saldoRepository) GetYearTotalSaldoBalance(year int) ([]*record.SaldoYearTotalBalance, error) {
	res, err := r.db.GetYearlyTotalSaldoBalances(r.ctx, int32(year))

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total saldo balance: %w", err)
	}

	so := r.mapping.ToSaldoYearTotalBalances(res)

	return so, nil
}

func (r *saldoRepository) GetMonthlySaldoBalances(year int) ([]*record.SaldoMonthSaldoBalance, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlySaldoBalances(r.ctx, yearStart)

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly saldo balances: %w", err)
	}

	so := r.mapping.ToSaldoMonthBalances(res)

	return so, nil
}

func (r *saldoRepository) GetYearlySaldoBalances(year int) ([]*record.SaldoYearSaldoBalance, error) {
	res, err := r.db.GetYearlySaldoBalances(r.ctx, year)

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly saldo balances: %w", err)
	}
	so := r.mapping.ToSaldoYearSaldoBalances(res)

	return so, nil
}

func (r *saldoRepository) FindByActive(search string, page, pageSize int) ([]*record.SaldoRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetActiveSaldosParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetActiveSaldos(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToSaldosRecordActive(res), totalCount, nil

}

func (r *saldoRepository) FindByTrashed(search string, page, pageSize int) ([]*record.SaldoRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTrashedSaldosParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	saldos, err := r.db.GetTrashedSaldos(r.ctx, req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get trashed saldos: %w", err)
	}

	var totalCount int
	if len(saldos) > 0 {
		totalCount = int(saldos[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToSaldosRecordTrashed(saldos), totalCount, nil
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

func (r *saldoRepository) DeleteSaldoPermanent(saldo_id int) (bool, error) {
	err := r.db.DeleteSaldoPermanently(r.ctx, int32(saldo_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete saldo permanently: %w", err)
	}

	return true, nil
}

func (r *saldoRepository) RestoreAllSaldo() (bool, error) {
	err := r.db.RestoreAllSaldos(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all saldos: %w", err)
	}

	return true, nil
}

func (r *saldoRepository) DeleteAllSaldoPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentSaldos(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all saldos permanently: %w", err)
	}

	return true, nil
}
