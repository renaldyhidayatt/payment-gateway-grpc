package repository

import (
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"
	"errors"
	"fmt"
	"time"
)

type withdrawRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewWithdrawRepository(db *db.Queries, ctx context.Context) *withdrawRepository {
	return &withdrawRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *withdrawRepository) FindAll() ([]*db.Withdraw, error) {
	withdraw, err := r.db.GetAllWithdraws(r.ctx)

	if err != nil {
		return nil, errors.New("failed get withdraw")
	}

	return withdraw, nil

}

func (r *withdrawRepository) FindById(id int) (*db.Withdraw, error) {
	withdraw, err := r.db.GetWithdrawById(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get withdraw")
	}

	return withdraw, nil
}

func (r *withdrawRepository) FindByUsers(user_id int) ([]*db.Withdraw, error) {
	withdraw, err := r.db.GetWithdrawByUsers(r.ctx, int32(user_id))

	if err != nil {
		return nil, errors.New("failed get withdraw")
	}

	return withdraw, nil
}

func (r *withdrawRepository) FindByUsersId(user_id int) (*db.Withdraw, error) {
	withdraw, err := r.db.GetWithdrawByUserId(r.ctx, int32(user_id))

	if err != nil {
		return nil, errors.New("failed get withdraw")
	}

	return withdraw, nil
}

func (r *withdrawRepository) Create(input *db.CreateWithdrawParams) (*db.Withdraw, error) {
	var withdrawRequest db.CreateWithdrawParams

	withdrawRequest.UserID = int32(input.UserID)
	withdrawRequest.WithdrawAmount = int32(input.WithdrawAmount)
	withdrawRequest.WithdrawTime = time.Now()

	withdraw, err := r.db.CreateWithdraw(r.ctx, withdrawRequest)

	if err != nil {
		return nil, errors.New("failed create withdraw")
	}

	return withdraw, nil
}

func (r *withdrawRepository) Update(input *db.UpdateWithdrawParams) (*db.Withdraw, error) {
	var withdrawRequest db.UpdateWithdrawParams

	withdrawRequest.WithdrawID = int32(input.WithdrawID)
	withdrawRequest.WithdrawAmount = int32(input.WithdrawAmount)
	withdrawRequest.WithdrawTime = time.Now()

	withdraw, err := r.db.UpdateWithdraw(r.ctx, withdrawRequest)

	if err != nil {
		return nil, errors.New("failed update withdraw")
	}

	return withdraw, nil
}

func (r *withdrawRepository) Delete(id int) error {
	err := r.db.DeleteWithdraw(r.ctx, int32(id))

	if err != nil {
		return fmt.Errorf("failed error")
	}

	return nil

}
