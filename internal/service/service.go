package service

import (
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/auth"
	"MamangRust/paymentgatewaygrpc/pkg/hash"
)

type Service struct {
	Auth     AuthService
	User     UserService
	Saldo    SaldoService
	Topup    TopupService
	Transfer TransferService
	Withdraw WithdrawService
}

type Deps struct {
	Repositories *repository.Repositories
	Token        auth.TokenManager
	Hash         hash.Hashing
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:     NewAuthService(deps.Repositories.User, deps.Hash, deps.Token),
		User:     NewUserService(deps.Repositories.User, deps.Hash),
		Saldo:    NewSaldoService(deps.Repositories.Saldo, deps.Repositories.User),
		Topup:    NewTopupService(deps.Repositories.Topup, deps.Repositories.Saldo, deps.Repositories.User),
		Transfer: NewTransferService(deps.Repositories.Transfer, deps.Repositories.Saldo, deps.Repositories.User),
		Withdraw: NewWithdrawService(deps.Repositories.Withdraw, deps.Repositories.Saldo, deps.Repositories.User),
	}
}
