package service

import (
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/auth"
	"MamangRust/paymentgatewaygrpc/pkg/hash"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
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
	Logger       logger.Logger
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:     NewAuthService(deps.Repositories.User, deps.Hash, deps.Token, deps.Logger),
		User:     NewUserService(deps.Repositories.User, deps.Hash, deps.Logger),
		Saldo:    NewSaldoService(deps.Repositories.Saldo, deps.Repositories.User, deps.Logger),
		Topup:    NewTopupService(deps.Repositories.Topup, deps.Repositories.Saldo, deps.Repositories.User, deps.Logger),
		Transfer: NewTransferService(deps.Repositories.Transfer, deps.Repositories.Saldo, deps.Repositories.User, deps.Logger),
		Withdraw: NewWithdrawService(deps.Repositories.Withdraw, deps.Repositories.Saldo, deps.Repositories.User, deps.Logger),
	}
}
