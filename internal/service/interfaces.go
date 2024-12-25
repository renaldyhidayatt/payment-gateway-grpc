package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type AuthService interface {
	Register(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	Login(request *requests.AuthRequest) (*string, *response.ErrorResponse)
}

type CardService interface {
	FindAll(page int, pageSize int, search string) ([]*response.CardResponse, int, *response.ErrorResponse)
	FindById(card_id int) (*response.CardResponse, *response.ErrorResponse)
	FindByUserID(userID int) (*response.CardResponse, *response.ErrorResponse)
	FindByActive() ([]*response.CardResponse, *response.ErrorResponse)
	FindByTrashed() ([]*response.CardResponse, *response.ErrorResponse)
	FindByCardNumber(card_number string) (*response.CardResponse, *response.ErrorResponse)
	CreateCard(request *requests.CreateCardRequest) (*response.CardResponse, *response.ErrorResponse)
	UpdateCard(request *requests.UpdateCardRequest) (*response.CardResponse, *response.ErrorResponse)
	TrashedCard(cardId int) (*response.CardResponse, *response.ErrorResponse)
	RestoreCard(cardId int) (*response.CardResponse, *response.ErrorResponse)
	DeleteCardPermanent(cardId int) (interface{}, *response.ErrorResponse)
}

type MerchantService interface {
	FindAll(page int, pageSize int, search string) ([]*response.MerchantResponse, int, *response.ErrorResponse)
	FindById(merchant_id int) (*response.MerchantResponse, *response.ErrorResponse)
	FindByActive() ([]*response.MerchantResponse, *response.ErrorResponse)
	FindByTrashed() ([]*response.MerchantResponse, *response.ErrorResponse)
	FindByApiKey(api_key string) (*response.MerchantResponse, *response.ErrorResponse)
	FindByMerchantUserId(user_id int) ([]*response.MerchantResponse, *response.ErrorResponse)
	CreateMerchant(request *requests.CreateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse)
	UpdateMerchant(request *requests.UpdateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse)
	TrashedMerchant(merchant_id int) (*response.MerchantResponse, *response.ErrorResponse)
	RestoreMerchant(merchant_id int) (*response.MerchantResponse, *response.ErrorResponse)
	DeleteMerchantPermanent(merchant_id int) (interface{}, *response.ErrorResponse)
}

type SaldoService interface {
	FindAll(page int, pageSize int, search string) ([]*response.SaldoResponse, int, *response.ErrorResponse)
	FindById(saldo_id int) (*response.SaldoResponse, *response.ErrorResponse)
	FindByCardNumber(card_number string) (*response.SaldoResponse, *response.ErrorResponse)
	FindByActive() ([]*response.SaldoResponse, *response.ErrorResponse)
	FindByTrashed() ([]*response.SaldoResponse, *response.ErrorResponse)
	CreateSaldo(request *requests.CreateSaldoRequest) (*response.SaldoResponse, *response.ErrorResponse)
	UpdateSaldo(request *requests.UpdateSaldoRequest) (*response.SaldoResponse, *response.ErrorResponse)
	TrashSaldo(saldo_id int) (*response.SaldoResponse, *response.ErrorResponse)
	RestoreSaldo(saldo_id int) (*response.SaldoResponse, *response.ErrorResponse)
	DeleteSaldoPermanent(saldo_id int) (interface{}, *response.ErrorResponse)
}

type TopupService interface {
	FindAll(page int, pageSize int, search string) ([]*response.TopupResponse, int, *response.ErrorResponse)
	FindById(topupID int) (*response.TopupResponse, *response.ErrorResponse)
	FindByCardNumber(card_number string) ([]*response.TopupResponse, *response.ErrorResponse)
	FindByActive() ([]*response.TopupResponse, *response.ErrorResponse)
	FindByTrashed() ([]*response.TopupResponse, *response.ErrorResponse)
	CreateTopup(request *requests.CreateTopupRequest) (*response.TopupResponse, *response.ErrorResponse)
	UpdateTopup(request *requests.UpdateTopupRequest) (*response.TopupResponse, *response.ErrorResponse)
	TrashedTopup(topup_id int) (*response.TopupResponse, *response.ErrorResponse)
	RestoreTopup(topup_id int) (*response.TopupResponse, *response.ErrorResponse)
	DeleteTopupPermanent(topup_id int) (interface{}, *response.ErrorResponse)
}

type TransactionService interface {
	FindAll(page int, pageSize int, search string) ([]*response.TransactionResponse, int, *response.ErrorResponse)
	FindById(transactionID int) (*response.TransactionResponse, *response.ErrorResponse)
	FindByActive() ([]*response.TransactionResponse, *response.ErrorResponse)
	FindByTrashed() ([]*response.TransactionResponse, *response.ErrorResponse)
	FindByCardNumber(card_number string) ([]*response.TransactionResponse, *response.ErrorResponse)
	FindTransactionByMerchantId(merchant_id int) ([]*response.TransactionResponse, *response.ErrorResponse)
	Create(apiKey string, request *requests.CreateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse)
	Update(apiKey string, request *requests.UpdateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse)
	TrashedTransaction(transaction_id int) (*response.TransactionResponse, *response.ErrorResponse)
	RestoreTransaction(transaction_id int) (*response.TransactionResponse, *response.ErrorResponse)
	DeleteTransactionPermanent(transaction_id int) (interface{}, *response.ErrorResponse)
}

type TransferService interface {
	FindAll(page int, pageSize int, search string) ([]*response.TransferResponse, int, *response.ErrorResponse)
	FindById(transferId int) (*response.TransferResponse, *response.ErrorResponse)
	FindByActive() ([]*response.TransferResponse, *response.ErrorResponse)
	FindByTrashed() ([]*response.TransferResponse, *response.ErrorResponse)
	FindTransferByTransferFrom(transfer_from string) ([]*response.TransferResponse, *response.ErrorResponse)
	FindTransferByTransferTo(transfer_to string) ([]*response.TransferResponse, *response.ErrorResponse)
	CreateTransaction(request *requests.CreateTransferRequest) (*response.TransferResponse, *response.ErrorResponse)
	UpdateTransaction(request *requests.UpdateTransferRequest) (*response.TransferResponse, *response.ErrorResponse)
	TrashedTransfer(transfer_id int) (*response.TransferResponse, *response.ErrorResponse)
	RestoreTransfer(transfer_id int) (*response.TransferResponse, *response.ErrorResponse)
	DeleteTransferPermanent(transfer_id int) (*response.ApiResponse[interface{}], *response.ErrorResponse)
}

type UserService interface {
	FindAll(page int, pageSize int, search string) ([]*response.UserResponse, int, *response.ErrorResponse)
	FindByID(id int) (*response.UserResponse, *response.ErrorResponse)
	FindByActive() ([]*response.UserResponse, *response.ErrorResponse)
	FindByTrashed() ([]*response.UserResponse, *response.ErrorResponse)
	CreateUser(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	UpdateUser(request *requests.UpdateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	TrashedUser(user_id int) (*response.UserResponse, *response.ErrorResponse)
	RestoreUser(user_id int) (*response.UserResponse, *response.ErrorResponse)
	DeleteUserPermanent(user_id int) (interface{}, *response.ErrorResponse)
}

type WithdrawService interface {
	FindAll(page int, pageSize int, search string) ([]*response.WithdrawResponse, int, *response.ErrorResponse)
	FindById(withdrawID int) (*response.WithdrawResponse, *response.ErrorResponse)
	FindByCardNumber(card_number string) ([]*response.WithdrawResponse, *response.ErrorResponse)
	FindByActive() ([]*response.WithdrawResponse, *response.ErrorResponse)
	FindByTrashed() ([]*response.WithdrawResponse, *response.ErrorResponse)
	Create(request *requests.CreateWithdrawRequest) (*response.WithdrawResponse, *response.ErrorResponse)
	Update(request *requests.UpdateWithdrawRequest) (*response.WithdrawResponse, *response.ErrorResponse)
	TrashedWithdraw(withdraw_id int) (*response.WithdrawResponse, *response.ErrorResponse)
	RestoreWithdraw(withdraw_id int) (*response.WithdrawResponse, *response.ErrorResponse)
	DeleteWithdrawPermanent(withdraw_id int) (interface{}, *response.ErrorResponse)
}
