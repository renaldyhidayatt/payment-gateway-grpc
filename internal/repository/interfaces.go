package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"time"
)

type UserRepository interface {
	FindAllUsers(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	FindById(user_id int) (*record.UserRecord, error)
	SearchUsersByEmail(email string) ([]*record.UserRecord, error)
	FindByEmail(email string) (*record.UserRecord, error)
	FindByActive() ([]*record.UserRecord, error)
	FindByTrashed() ([]*record.UserRecord, error)
	CreateUser(request requests.CreateUserRequest) (*record.UserRecord, error)
	UpdateUser(request requests.UpdateUserRequest) (*record.UserRecord, error)
	TrashedUser(user_id int) (*record.UserRecord, error)
	RestoreUser(user_id int) (*record.UserRecord, error)
	DeleteUserPermanent(user_id int) error
}

type CardRepository interface {
	FindAllCards(search string, page, pageSize int) ([]*record.CardRecord, int, error)
	FindById(card_id int) (*record.CardRecord, error)
	FindCardByUserId(user_id int) (*record.CardRecord, error)
	FindByActive() ([]*record.CardRecord, error)
	FindByTrashed() ([]*record.CardRecord, error)
	FindCardByCardNumber(card_number string) (*record.CardRecord, error)
	CreateCard(request requests.CreateCardRequest) (*record.CardRecord, error)
	UpdateCard(request requests.UpdateCardRequest) (*record.CardRecord, error)
	TrashedCard(cardId int) (*record.CardRecord, error)
	RestoreCard(cardId int) (*record.CardRecord, error)
	DeleteCardPermanent(card_id int) error
}

type MerchantRepository interface {
	FindAllMerchants(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
	FindById(merchant_id int) (*record.MerchantRecord, error)
	FindByActive() ([]*record.MerchantRecord, error)
	FindByTrashed() ([]*record.MerchantRecord, error)
	FindByApiKey(api_key string) (*record.MerchantRecord, error)
	FindByName(name string) (*record.MerchantRecord, error)
	FindByMerchantUserId(user_id int) ([]*record.MerchantRecord, error)
	CreateMerchant(request requests.CreateMerchantRequest) (*record.MerchantRecord, error)
	UpdateMerchant(request requests.UpdateMerchantRequest) (*record.MerchantRecord, error)
	TrashedMerchant(merchantId int) (*record.MerchantRecord, error)
	RestoreMerchant(merchant_id int) (*record.MerchantRecord, error)
	DeleteMerchantPermanent(merchant_id int) error
}

type SaldoRepository interface {
	FindAllSaldos(search string, page, pageSize int) ([]*record.SaldoRecord, int, error)
	FindById(saldo_id int) (*record.SaldoRecord, error)
	FindByCardNumber(card_number string) (*record.SaldoRecord, error)
	FindByActive() ([]*record.SaldoRecord, error)
	FindByTrashed() ([]*record.SaldoRecord, error)
	CreateSaldo(request requests.CreateSaldoRequest) (*record.SaldoRecord, error)
	UpdateSaldo(request requests.UpdateSaldoRequest) (*record.SaldoRecord, error)
	UpdateSaldoBalance(request requests.UpdateSaldoBalance) (*record.SaldoRecord, error)
	UpdateSaldoWithdraw(request requests.UpdateSaldoWithdraw) (*record.SaldoRecord, error)
	TrashedSaldo(saldoID int) (*record.SaldoRecord, error)
	RestoreSaldo(saldoID int) (*record.SaldoRecord, error)
	DeleteSaldoPermanent(saldo_id int) error
}

type TopupRepository interface {
	FindAllTopups(search string, page, pageSize int) ([]*record.TopupRecord, int, error)
	FindById(topup_id int) (*record.TopupRecord, error)
	FindByCardNumber(card_number string) ([]*record.TopupRecord, error)
	FindByActive() ([]*record.TopupRecord, error)
	FindByTrashed() ([]*record.TopupRecord, error)
	CountTopupsByDate(date string) (int, error)
	CountAllTopups() (int, error)
	CreateTopup(request requests.CreateTopupRequest) (*record.TopupRecord, error)
	UpdateTopup(request requests.UpdateTopupRequest) (*record.TopupRecord, error)
	UpdateTopupAmount(request requests.UpdateTopupAmount) (*record.TopupRecord, error)
	TrashedTopup(topup_id int) (*record.TopupRecord, error)
	RestoreTopup(topup_id int) (*record.TopupRecord, error)
	DeleteTopupPermanent(topup_id int) error
}

type TransactionRepository interface {
	FindAllTransactions(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindById(transaction_id int) (*record.TransactionRecord, error)
	FindByActive() ([]*record.TransactionRecord, error)
	FindByTrashed() ([]*record.TransactionRecord, error)
	FindByCardNumber(card_number string) ([]*record.TransactionRecord, error)
	FindTransactionByMerchantId(merchant_id int) ([]*record.TransactionRecord, error)
	CountTransactionsByDate(date string) (int, error)
	CountAllTransactions() (int, error)
	CreateTransaction(request requests.CreateTransactionRequest) (*record.TransactionRecord, error)
	UpdateTransaction(request requests.UpdateTransactionRequest) (*record.TransactionRecord, error)
	TrashedTransaction(transaction_id int) (*record.TransactionRecord, error)
	RestoreTransaction(topup_id int) (*record.TransactionRecord, error)
	DeleteTransactionPermanent(topup_id int) error
}

type TransferRepository interface {
	FindAll(search string, page, pageSize int) ([]*record.TransferRecord, int, error)
	FindById(id int) (*record.TransferRecord, error)
	FindByActive() ([]*record.TransferRecord, error)
	FindByTrashed() ([]*record.TransferRecord, error)
	FindTransferByTransferFrom(transfer_from string) ([]*record.TransferRecord, error)
	FindTransferByTransferTo(transfer_to string) ([]*record.TransferRecord, error)
	CountTransfersByDate(date string) (int, error)
	CountAllTransfers() (int, error)
	CreateTransfer(request requests.CreateTransferRequest) (*record.TransferRecord, error)
	UpdateTransfer(request requests.UpdateTransferRequest) (*record.TransferRecord, error)
	UpdateTransferAmount(request requests.UpdateTransferAmountRequest) (*record.TransferRecord, error)
	TrashedTransfer(transfer_id int) (*record.TransferRecord, error)
	RestoreTransfer(transfer_id int) (*record.TransferRecord, error)
	DeleteTransferPermanent(topup_id int) error
}

type WithdrawRepository interface {
	FindAll(search string, page, pageSize int) ([]*record.WithdrawRecord, int, error)
	FindById(id int) (*record.WithdrawRecord, error)
	FindByCardNumber(card_number string) ([]*record.WithdrawRecord, error)
	FindByActive() ([]*record.WithdrawRecord, error)
	FindByTrashed() ([]*record.WithdrawRecord, error)
	CountActiveByDate(date time.Time) (int64, error)
	CreateWithdraw(request requests.CreateWithdrawRequest) (*record.WithdrawRecord, error)
	UpdateWithdraw(request requests.UpdateWithdrawRequest) (*record.WithdrawRecord, error)
	TrashedWithdraw(WithdrawID int) (*record.WithdrawRecord, error)
	RestoreWithdraw(WithdrawID int) (*record.WithdrawRecord, error)
	DeleteWithdrawPermanent(WithdrawID int) error
}
