package service

import (
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/auth"
	"MamangRust/paymentgatewaygrpc/pkg/hash"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
)

type Service struct {
	Auth            AuthService
	User            UserService
	Saldo           SaldoService
	Topup           TopupService
	Transfer        TransferService
	Withdraw        WithdrawService
	Card            CardService
	MerchantService MerchantService
	Transaction     TransactionService
}

type Deps struct {
	Repositories *repository.Repositories
	Token        auth.TokenManager
	Hash         *hash.Hashing
	Logger       *logger.Logger
	Mapper       *responsemapper.ResponseMapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:            NewAuthService(deps.Repositories.User, deps.Hash, deps.Token, deps.Logger, deps.Mapper.UserResponseMapper),
		User:            NewUserService(deps.Repositories.User, deps.Logger, deps.Mapper.UserResponseMapper, deps.Hash),
		Saldo:           NewSaldoService(deps.Repositories.Saldo, deps.Repositories.Card, deps.Logger, deps.Mapper.SaldoResponseMapper),
		Topup:           NewTopupService(deps.Repositories.Card, deps.Repositories.Topup, deps.Repositories.Saldo, deps.Logger, deps.Mapper.TopupResponseMapper),
		Transfer:        NewTransferService(deps.Repositories.User, deps.Repositories.Card, deps.Repositories.Transfer, deps.Repositories.Saldo, deps.Logger, deps.Mapper.TransferResponseMapper),
		Withdraw:        NewWithdrawService(deps.Repositories.User, deps.Repositories.Withdraw, deps.Repositories.Saldo, deps.Logger, deps.Mapper.WithdrawResponseMapper),
		Card:            NewCardService(deps.Repositories.Card, deps.Repositories.User, deps.Logger, deps.Mapper.CardResponseMapper),
		MerchantService: NewMerchantService(deps.Repositories.Merchant, deps.Logger, deps.Mapper.MerchantResponseMapper),
		Transaction:     NewTransactionService(deps.Repositories.Merchant, deps.Repositories.Card, deps.Repositories.Saldo, deps.Repositories.Transaction, deps.Logger, deps.Mapper.TransactionResponseMapper),
	}
}
