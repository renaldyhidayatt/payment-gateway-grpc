package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
)

type AuthService interface {
	Register(request *requests.CreateUserRequest) (*db.User, error)
	Login(request *requests.AuthLoginRequest) (*requests.JWTToken, error)
}

type UserService interface {
	FindAll() ([]*db.User, error)
	FindById(id int) (*db.User, error)
	Create(input *requests.CreateUserRequest) (*db.User, error)
	Update(input *requests.UpdateUserRequest) (*db.User, error)
	Delete(id int) error
}

type SaldoService interface {
	FindAll() ([]*db.Saldo, error)
	FindById(id int) (*db.Saldo, error)
	FindByUserId(id int) (*db.Saldo, error)
	FindByUsersId(id int) ([]*db.Saldo, error)
	Create(input *requests.CreateSaldoRequest) (*db.Saldo, error)
	Update(input *requests.UpdateSaldoRequest) (*db.Saldo, error)
	Delete(id int) error
}

type TopupService interface {
	FindAll() ([]*db.Topup, error)
	FindById(id int) (*db.Topup, error)
	FindByUsers(user_id int) ([]*db.Topup, error)
	FindByUsersId(user_id int) (*db.Topup, error)
	Create(input *requests.CreateTopupRequest) (*db.Topup, error)
	UpdateTopup(input *requests.UpdateTopupRequest) (*db.Topup, error)
	DeleteTopup(id int) error
}

type WithdrawService interface {
	FindAll() ([]*db.Withdraw, error)
	FindById(id int) (*db.Withdraw, error)
	FindByUsers(user_id int) ([]*db.Withdraw, error)
	FindByUsersId(user_id int) (*db.Withdraw, error)
	Create(input *requests.CreateWithdrawRequest) (*db.Withdraw, error)
	Update(input *requests.UpdateWithdrawRequest) (*db.Withdraw, error)
	Delete(id int) error
}

type TransferService interface {
	FindAll() ([]*db.Transfer, error)
	FindById(id int) (*db.Transfer, error)
	FindByUsers(user_id int) ([]*db.Transfer, error)
	FindByUsersId(user_id int) (*db.Transfer, error)
	Create(req *requests.CreateTransferRequest) (*db.Transfer, error)
	Update(req *requests.UpdateTransferRequest) (*db.Transfer, error)
	Delete(id int) error
}
