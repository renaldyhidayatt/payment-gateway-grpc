package repository

import (
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"
	"errors"
	"fmt"
	"time"
)

type transferRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewTransferRepository(db *db.Queries, ctx context.Context) *transferRepository {
	return &transferRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *transferRepository) FindAll() ([]*db.Transfer, error) {
	transfer, err := r.db.GetAllTransfers(r.ctx)

	if err != nil {
		return nil, errors.New("failed get transfer")
	}

	return transfer, nil
}

func (r *transferRepository) FindById(id int) (*db.Transfer, error) {
	transfer, err := r.db.GetTransferById(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get transfer")
	}

	return transfer, nil
}

func (r *transferRepository) FindByUsers(id int) ([]*db.Transfer, error) {
	transfer, err := r.db.GetTransferByUsers(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get transfer")
	}

	return transfer, nil
}

func (r *transferRepository) FindByUser(id int) (*db.Transfer, error) {
	transfer, err := r.db.GetTransferByUserId(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get transfer")
	}

	return transfer, nil
}

func (r *transferRepository) Create(input *db.CreateTransferParams) (*db.Transfer, error) {
	var transferRequest db.CreateTransferParams

	transferRequest.TransferFrom = int32(input.TransferFrom)
	transferRequest.TransferTo = int32(input.TransferTo)
	transferRequest.TransferAmount = int32(input.TransferAmount)
	transferRequest.TransferTime = time.Now()

	transfer, err := r.db.CreateTransfer(r.ctx, transferRequest)

	if err != nil {
		return nil, errors.New("failed create transfer")
	}

	return transfer, nil
}

func (r *transferRepository) Update(input *db.UpdateTransferParams) (*db.Transfer, error) {
	var transferRequest db.UpdateTransferParams

	transferRequest.TransferAmount = int32(input.TransferAmount)
	transferRequest.TransferID = int32(input.TransferID)
	transferRequest.TransferTime = time.Now()

	transfer, err := r.db.UpdateTransfer(r.ctx, transferRequest)

	if err != nil {
		return nil, errors.New("failed update transfer")
	}

	return transfer, nil
}

func (r *transferRepository) Delete(id int) error {
	err := r.db.DeleteTransfer(r.ctx, int32(id))

	if err != nil {
		return fmt.Errorf("failed error")
	}

	return nil
}
