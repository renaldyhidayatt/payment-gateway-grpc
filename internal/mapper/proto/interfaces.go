package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type AuthProtoMapper interface {
	ToResponseLogin(token string) *pb.ApiResponseLogin
	ToResponseRegister(response response.UserResponse) *pb.ApiResponseRegister
}

type CardProtoMapper interface {
	ToResponseCard(card *response.CardResponse) *pb.CardResponse
	ToResponsesCard(cards []*response.CardResponse) []*pb.CardResponse
}

type MerchantProtoMapper interface {
	ToResponseMerchant(merchant *response.MerchantResponse) *pb.MerchantResponse
	ToResponsesMerchant(merchants []*response.MerchantResponse) []*pb.MerchantResponse
}

type SaldoProtoMapper interface {
	ToResponseSaldo(saldo *response.SaldoResponse) *pb.SaldoResponse
	ToResponsesSaldo(saldos []*response.SaldoResponse) []*pb.SaldoResponse
}

type TopupProtoMapper interface {
	ToResponseTopup(topup *response.TopupResponse) *pb.TopupResponse
	ToResponsesTopup(topups []*response.TopupResponse) []*pb.TopupResponse
}

type TransactionProtoMapper interface {
	ToResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse
	ToResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse
}

type TransferProtoMapper interface {
	ToResponseTransfer(transfer *response.TransferResponse) *pb.TransferResponse
	ToResponsesTransfer(transfers []*response.TransferResponse) []*pb.TransferResponse
}

type UserProtoMapper interface {
	ToResponseUser(user *response.UserResponse) *pb.UserResponse
	ToResponsesUser(users []*response.UserResponse) []*pb.UserResponse
}

type WithdrawalProtoMapper interface {
	ToResponseWithdrawal(withdrawal *response.WithdrawResponse) *pb.WithdrawResponse
	ToResponsesWithdrawal(withdrawals []*response.WithdrawResponse) []*pb.WithdrawResponse
}
