package repository

import (
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type saldoRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewSaldoRepository(db *db.Queries, ctx context.Context) *saldoRepository {
	return &saldoRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *saldoRepository) FindAll() ([]*db.Saldo, error) {
	saldo, err := r.db.GetAllSaldo(r.ctx)

	if err != nil {
		return nil, errors.New("failed get saldo")
	}

	return saldo, nil
}

func (r *saldoRepository) FindById(id int) (*db.Saldo, error) {
	saldo, err := r.db.GetSaldoById(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get saldo")
	}

	return saldo, nil
}

func (r *saldoRepository) FindByUserId(user_id int) (*db.Saldo, error) {
	saldo, err := r.db.GetSaldoByUserId(r.ctx, int32(user_id))

	if err != nil {
		return nil, errors.New("failed get saldo")
	}

	return saldo, nil
}

func (r *saldoRepository) FindByUsersId(user_id int) ([]*db.Saldo, error) {
	saldo, err := r.db.GetSaldoByUsers(r.ctx, int32(user_id))

	if err != nil {
		return nil, errors.New("failed get saldo")
	}

	return saldo, nil
}

func (r *saldoRepository) Create(input *db.CreateSaldoParams) (*db.Saldo, error) {
	var saldoRequest db.CreateSaldoParams

	saldoRequest.UserID = int32(input.UserID)
	saldoRequest.TotalBalance = int32(input.TotalBalance)

	saldo, err := r.db.CreateSaldo(r.ctx, saldoRequest)

	if err != nil {
		return nil, errors.New("failed create saldo")
	}

	return saldo, nil
}

func (r *saldoRepository) Update(input *db.UpdateSaldoParams) (*db.Saldo, error) {
	var saldoRequest db.UpdateSaldoParams

	saldoRequest.UserID = int32(input.UserID)

	if input.WithdrawAmount.Int32 != 0 {
		saldoRequest.WithdrawAmount = sql.NullInt32{
			Int32: int32(input.WithdrawAmount.Int32),
			Valid: true,
		}
	} else {
		saldoRequest.WithdrawAmount = sql.NullInt32{Int32: 0, Valid: false}
	}

	if input.WithdrawTime.Time != (time.Time{}) {
		saldoRequest.WithdrawTime = sql.NullTime{
			Time:  input.WithdrawTime.Time,
			Valid: true,
		}
	} else {
		saldoRequest.WithdrawTime = sql.NullTime{Valid: false}
	}

	saldoRequest.TotalBalance = int32(input.TotalBalance)

	saldo, err := r.db.UpdateSaldo(r.ctx, saldoRequest)

	if err != nil {
		return nil, errors.New("failed update saldo: " + err.Error())
	}

	return saldo, nil
}

func (r *saldoRepository) UpdateSaldoBalance(input *db.UpdateSaldoBalanceParams) (*db.Saldo, error) {
	request := db.UpdateSaldoBalanceParams{
		UserID:       input.UserID,
		TotalBalance: input.TotalBalance,
	}

	saldo, err := r.db.UpdateSaldoBalance(r.ctx, request)

	if err != nil {
		return nil, fmt.Errorf("failed error: " + err.Error())
	}

	return saldo, nil
}

func (r *saldoRepository) Delete(id int) error {
	err := r.db.DeleteSaldo(r.ctx, int32(id))

	if err != nil {
		return fmt.Errorf("failed error")
	}

	return nil
}
