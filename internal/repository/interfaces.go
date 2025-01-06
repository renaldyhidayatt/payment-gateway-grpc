package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"time"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type UserRepository interface {
	FindAllUsers(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	FindById(user_id int) (*record.UserRecord, error)
	FindByEmail(email string) (*record.UserRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	CreateUser(request *requests.CreateUserRequest) (*record.UserRecord, error)
	UpdateUser(request *requests.UpdateUserRequest) (*record.UserRecord, error)
	TrashedUser(user_id int) (*record.UserRecord, error)
	RestoreUser(user_id int) (*record.UserRecord, error)
	DeleteUserPermanent(user_id int) error
}

type RoleRepository interface {
	FindAllRoles(page int, pageSize int, search string) ([]*record.RoleRecord, int, error)
	FindById(role_id int) (*record.RoleRecord, error)
	FindByName(name string) (*record.RoleRecord, error)
	FindByUserId(user_id int) ([]*record.RoleRecord, error)
	FindByActiveRole(page int, pageSize int, search string) ([]*record.RoleRecord, int, error)
	FindByTrashedRole(page int, pageSize int, search string) ([]*record.RoleRecord, int, error)
	CreateRole(request *requests.CreateRoleRequest) (*record.RoleRecord, error)
	UpdateRole(request *requests.UpdateRoleRequest) (*record.RoleRecord, error)
	TrashedRole(role_id int) (*record.RoleRecord, error)
	RestoreRole(role_id int) (*record.RoleRecord, error)
	DeleteRolePermanent(role_id int) error
}

type RefreshTokenRepository interface {
	FindByToken(token string) (*record.RefreshTokenRecord, error)
	FindByUserId(user_id int) (*record.RefreshTokenRecord, error)
	CreateRefreshToken(req *requests.CreateRefreshToken) (*record.RefreshTokenRecord, error)
	UpdateRefreshToken(req *requests.UpdateRefreshToken) (*record.RefreshTokenRecord, error)
	DeleteRefreshToken(token string) error
	DeleteRefreshTokenByUserId(user_id int) error
}

type UserRoleRepository interface {
	AssignRoleToUser(req *requests.CreateUserRoleRequest) (*record.UserRoleRecord, error)
	RemoveRoleFromUser(req *requests.RemoveUserRoleRequest) error
}

type CardRepository interface {
	FindAllCards(search string, page, pageSize int) ([]*record.CardRecord, int, error)
	FindById(card_id int) (*record.CardRecord, error)
	FindCardByUserId(user_id int) (*record.CardRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.CardRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.CardRecord, int, error)
	FindCardByCardNumber(card_number string) (*record.CardRecord, error)
	CreateCard(request *requests.CreateCardRequest) (*record.CardRecord, error)
	UpdateCard(request *requests.UpdateCardRequest) (*record.CardRecord, error)
	TrashedCard(cardId int) (*record.CardRecord, error)
	RestoreCard(cardId int) (*record.CardRecord, error)
	DeleteCardPermanent(card_id int) error
}

type MerchantRepository interface {
	FindAllMerchants(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
	FindById(merchant_id int) (*record.MerchantRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
	FindByApiKey(api_key string) (*record.MerchantRecord, error)
	FindByName(name string) (*record.MerchantRecord, error)
	FindByMerchantUserId(user_id int) ([]*record.MerchantRecord, error)
	CreateMerchant(request *requests.CreateMerchantRequest) (*record.MerchantRecord, error)
	UpdateMerchant(request *requests.UpdateMerchantRequest) (*record.MerchantRecord, error)
	TrashedMerchant(merchantId int) (*record.MerchantRecord, error)
	RestoreMerchant(merchant_id int) (*record.MerchantRecord, error)
	DeleteMerchantPermanent(merchant_id int) error
}

type SaldoRepository interface {
	FindAllSaldos(search string, page, pageSize int) ([]*record.SaldoRecord, int, error)
	FindById(saldo_id int) (*record.SaldoRecord, error)
	FindByCardNumber(card_number string) (*record.SaldoRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.SaldoRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.SaldoRecord, int, error)
	CreateSaldo(request *requests.CreateSaldoRequest) (*record.SaldoRecord, error)
	UpdateSaldo(request *requests.UpdateSaldoRequest) (*record.SaldoRecord, error)
	UpdateSaldoBalance(request *requests.UpdateSaldoBalance) (*record.SaldoRecord, error)
	UpdateSaldoWithdraw(request *requests.UpdateSaldoWithdraw) (*record.SaldoRecord, error)
	TrashedSaldo(saldoID int) (*record.SaldoRecord, error)
	RestoreSaldo(saldoID int) (*record.SaldoRecord, error)
	DeleteSaldoPermanent(saldo_id int) error
}

type TopupRepository interface {
	FindAllTopups(search string, page, pageSize int) ([]*record.TopupRecord, int, error)
	FindById(topup_id int) (*record.TopupRecord, error)
	FindByCardNumber(card_number string) ([]*record.TopupRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.TopupRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.TopupRecord, int, error)
	CountTopupsByDate(date string) (int, error)
	CountAllTopups() (*int64, error)
	CreateTopup(request *requests.CreateTopupRequest) (*record.TopupRecord, error)
	UpdateTopup(request *requests.UpdateTopupRequest) (*record.TopupRecord, error)
	UpdateTopupAmount(request *requests.UpdateTopupAmount) (*record.TopupRecord, error)
	TrashedTopup(topup_id int) (*record.TopupRecord, error)
	RestoreTopup(topup_id int) (*record.TopupRecord, error)
	DeleteTopupPermanent(topup_id int) error
}

type TransactionRepository interface {
	FindAllTransactions(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindById(transaction_id int) (*record.TransactionRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindByCardNumber(card_number string) ([]*record.TransactionRecord, error)
	FindTransactionByMerchantId(merchant_id int) ([]*record.TransactionRecord, error)
	CountTransactionsByDate(date string) (int, error)
	CountAllTransactions() (*int64, error)
	CreateTransaction(request *requests.CreateTransactionRequest) (*record.TransactionRecord, error)
	UpdateTransaction(request *requests.UpdateTransactionRequest) (*record.TransactionRecord, error)
	TrashedTransaction(transaction_id int) (*record.TransactionRecord, error)
	RestoreTransaction(topup_id int) (*record.TransactionRecord, error)
	DeleteTransactionPermanent(topup_id int) error
}

type TransferRepository interface {
	FindAll(search string, page, pageSize int) ([]*record.TransferRecord, int, error)
	FindById(id int) (*record.TransferRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.TransferRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.TransferRecord, int, error)
	FindTransferByTransferFrom(transfer_from string) ([]*record.TransferRecord, error)
	FindTransferByTransferTo(transfer_to string) ([]*record.TransferRecord, error)
	CountTransfersByDate(date string) (int, error)
	CountAllTransfers() (*int64, error)
	CreateTransfer(request *requests.CreateTransferRequest) (*record.TransferRecord, error)
	UpdateTransfer(request *requests.UpdateTransferRequest) (*record.TransferRecord, error)
	UpdateTransferAmount(request *requests.UpdateTransferAmountRequest) (*record.TransferRecord, error)
	TrashedTransfer(transfer_id int) (*record.TransferRecord, error)
	RestoreTransfer(transfer_id int) (*record.TransferRecord, error)
	DeleteTransferPermanent(topup_id int) error
}

type WithdrawRepository interface {
	FindAll(search string, page, pageSize int) ([]*record.WithdrawRecord, int, error)
	FindById(id int) (*record.WithdrawRecord, error)
	FindByCardNumber(card_number string) ([]*record.WithdrawRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.WithdrawRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.WithdrawRecord, int, error)
	CountActiveByDate(date time.Time) (int64, error)
	CreateWithdraw(request *requests.CreateWithdrawRequest) (*record.WithdrawRecord, error)
	UpdateWithdraw(request *requests.UpdateWithdrawRequest) (*record.WithdrawRecord, error)
	TrashedWithdraw(WithdrawID int) (*record.WithdrawRecord, error)
	RestoreWithdraw(WithdrawID int) (*record.WithdrawRecord, error)
	DeleteWithdrawPermanent(WithdrawID int) error
}
