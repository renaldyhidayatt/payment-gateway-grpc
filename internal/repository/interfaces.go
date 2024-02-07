package repository

import db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"

type UserRepository interface {
	FindAll() ([]*db.User, error)
	FindById(id int) (*db.User, error)
	Create(input *db.CreateUserParams) (*db.User, error)
	Update(input *db.UpdateUserParams) (*db.User, error)
	Delete(id int) error
	FindByEmail(email string) (*db.User, error)
}

type SaldoRepository interface {
	FindAll() ([]*db.Saldo, error)
	FindById(id int) (*db.Saldo, error)
	FindByUserId(user_id int) (*db.Saldo, error)
	FindByUsersId(user_id int) ([]*db.Saldo, error)
	Create(input *db.CreateSaldoParams) (*db.Saldo, error)
	Update(input *db.UpdateSaldoParams) (*db.Saldo, error)
	UpdateSaldoBalance(input *db.UpdateSaldoBalanceParams) (*db.Saldo, error)
	Delete(id int) error
}

type TopupRepository interface {
	FindAll() ([]*db.Topup, error)
	FindById(id int) (*db.Topup, error)
	FindByUsers(id int) ([]*db.Topup, error)
	FindByUsersId(id int) (*db.Topup, error)
	Create(input *db.CreateTopupParams) (*db.Topup, error)
	Update(input *db.UpdateTopupParams) (*db.Topup, error)
	Delete(id int) error
}

type WithdrawRepository interface {
	FindAll() ([]*db.Withdraw, error)
	FindById(id int) (*db.Withdraw, error)
	FindByUsers(user_id int) ([]*db.Withdraw, error)
	FindByUsersId(user_id int) (*db.Withdraw, error)
	Create(input *db.CreateWithdrawParams) (*db.Withdraw, error)
	Update(input *db.UpdateWithdrawParams) (*db.Withdraw, error)
	Delete(id int) error
}

type TransferRepository interface {
	FindAll() ([]*db.Transfer, error)
	FindById(id int) (*db.Transfer, error)
	FindByUsers(id int) ([]*db.Transfer, error)
	FindByUser(id int) (*db.Transfer, error)
	Create(input *db.CreateTransferParams) (*db.Transfer, error)
	Update(input *db.UpdateTransferParams) (*db.Transfer, error)
	Delete(id int) error
}
