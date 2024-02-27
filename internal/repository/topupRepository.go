package repository

import (
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"
	"errors"
	"fmt"
	"time"
)

type topupRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewTopupRepository(db *db.Queries, ctx context.Context) *topupRepository {
	return &topupRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *topupRepository) FindAll() ([]*db.Topup, error) {
	topup, err := r.db.GetAllTopups(r.ctx)

	if err != nil {
		return nil, errors.New("failed get topup")
	}

	return topup, nil
}

func (r *topupRepository) FindById(id int) (*db.Topup, error) {
	topup, err := r.db.GetTopupById(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get topup")
	}

	return topup, nil
}

func (r *topupRepository) FindByUsers(id int) ([]*db.Topup, error) {
	topup, err := r.db.GetTopupByUsers(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get topup " + err.Error())
	}

	return topup, nil
}

func (r *topupRepository) FindByUsersId(id int) (*db.Topup, error) {
	topup, err := r.db.GetTopupByUserId(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get topup " + err.Error())
	}

	return topup, nil
}

func (r *topupRepository) Create(input *db.CreateTopupParams) (*db.Topup, error) {
	var topupRequest db.CreateTopupParams

	topupRequest.TopupNo = input.TopupNo
	topupRequest.TopupAmount = int32(input.TopupAmount)
	topupRequest.TopupMethod = input.TopupMethod
	topupRequest.UserID = int32(input.UserID)
	topupRequest.TopupTime = time.Now()

	topup, err := r.db.CreateTopup(r.ctx, topupRequest)

	if err != nil {
		return nil, errors.New("failed create topup")
	}

	return topup, nil
}

func (r *topupRepository) Update(input *db.UpdateTopupParams) (*db.Topup, error) {
	var topupRequest db.UpdateTopupParams

	topupRequest.TopupID = int32(input.TopupID)
	topupRequest.TopupAmount = int32(input.TopupAmount)
	topupRequest.TopupMethod = input.TopupMethod
	topupRequest.TopupTime = time.Now()

	topup, err := r.db.UpdateTopup(r.ctx, topupRequest)

	if err != nil {
		return nil, errors.New("failed update topup")
	}

	return topup, nil
}

func (r *topupRepository) Delete(id int) error {
	err := r.db.DeleteTopup(r.ctx, int32(id))

	if err != nil {
		return fmt.Errorf("failed error")
	}

	return nil
}
