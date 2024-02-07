package repository

import (
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"
)

type Repositories struct {
	User     UserRepository
	Saldo    SaldoRepository
	Topup    TopupRepository
	Withdraw WithdrawRepository
	Transfer TransferRepository
}

func NewRepositories(db *db.Queries, ctx context.Context) *Repositories {
	return &Repositories{
		User:     NewUserRepository(db, ctx),
		Saldo:    NewSaldoRepository(db, ctx),
		Topup:    NewTopupRepository(db, ctx),
		Withdraw: NewWithdrawRepository(db, ctx),
		Transfer: NewTransferRepository(db, ctx),
	}
}
