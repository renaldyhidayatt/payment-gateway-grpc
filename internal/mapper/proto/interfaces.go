package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mocks.go
type AuthProtoMapper interface {
	ToResponseLogin(response *response.TokenResponse) *pb.ApiResponseLogin
	ToResponseRegister(response *response.UserResponse) *pb.ApiResponseRegister
	ToResponseRefreshToken(response *response.TokenResponse) *pb.ApiResponseRefreshToken
	ToResponseGetMe(response *response.UserResponse) *pb.ApiResponseGetMe
}

type CardProtoMapper interface {
	ToResponseCard(card *response.CardResponse) *pb.CardResponse
	ToResponsesCard(cards []*response.CardResponse) []*pb.CardResponse
	ToResponseCardDeleteAt(card *response.CardResponseDeleteAt) *pb.CardResponseDeleteAt
	ToResponsesCardDeletedAt(cards []*response.CardResponseDeleteAt) []*pb.CardResponseDeleteAt
}

type MerchantProtoMapper interface {
	ToResponseMerchant(merchant *response.MerchantResponse) *pb.MerchantResponse
	ToResponsesMerchant(merchants []*response.MerchantResponse) []*pb.MerchantResponse

	ToResponseMerchantDeleteAt(merchant *response.MerchantResponseDeleteAt) *pb.MerchantResponseDeleteAt
	ToResponsesMerchantDeleteAt(merchants []*response.MerchantResponseDeleteAt) []*pb.MerchantResponseDeleteAt
}

type SaldoProtoMapper interface {
	ToResponseSaldo(saldo *response.SaldoResponse) *pb.SaldoResponse
	ToResponsesSaldo(saldos []*response.SaldoResponse) []*pb.SaldoResponse

	ToResponseSaldoDeleteAt(saldo *response.SaldoResponseDeleteAt) *pb.SaldoResponseDeleteAt
	ToResponsesSaldoDeleteAt(saldos []*response.SaldoResponseDeleteAt) []*pb.SaldoResponseDeleteAt
}

type TopupProtoMapper interface {
	ToResponseTopup(topup *response.TopupResponse) *pb.TopupResponse
	ToResponsesTopup(topups []*response.TopupResponse) []*pb.TopupResponse

	ToResponseTopupDeleteAt(topup *response.TopupResponseDeleteAt) *pb.TopupResponseDeleteAt
	ToResponsesTopupDeleteAt(topups []*response.TopupResponseDeleteAt) []*pb.TopupResponseDeleteAt
}

type TransactionProtoMapper interface {
	ToResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse
	ToResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse

	ToResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *pb.TransactionResponseDeleteAt
	ToResponsesTransactionDeleteAt(transactions []*response.TransactionResponseDeleteAt) []*pb.TransactionResponseDeleteAt
}

type TransferProtoMapper interface {
	ToResponseTransfer(transfer *response.TransferResponse) *pb.TransferResponse
	ToResponsesTransfer(transfers []*response.TransferResponse) []*pb.TransferResponse

	ToResponseTransferDeleteAt(transfer *response.TransferResponseDeleteAt) *pb.TransferResponseDeleteAt
	ToResponsesTransferDeleteAt(transfers []*response.TransferResponseDeleteAt) []*pb.TransferResponseDeleteAt
}

type UserProtoMapper interface {
	ToResponseUser(user *response.UserResponse) *pb.UserResponse
	ToResponsesUser(users []*response.UserResponse) []*pb.UserResponse

	ToResponseUserDelete(user *response.UserResponseDeleteAt) *pb.UserResponseWithDeleteAt
	ToResponsesUserDeleteAt(users []*response.UserResponseDeleteAt) []*pb.UserResponseWithDeleteAt
}

type WithdrawalProtoMapper interface {
	ToResponseWithdrawal(withdrawal *response.WithdrawResponse) *pb.WithdrawResponse
	ToResponsesWithdrawal(withdrawals []*response.WithdrawResponse) []*pb.WithdrawResponse

	ToResponseWithdrawalDeleteAt(withdraw *response.WithdrawResponseDeleteAt) *pb.WithdrawResponseDeleteAt
	ToResponsesWithdrawalDeleteAt(withdraws []*response.WithdrawResponseDeleteAt) []*pb.WithdrawResponseDeleteAt
}
