package responsemapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/interfaces.go
type CardResponseMapper interface {
	ToCardResponse(card *record.CardRecord) *response.CardResponse
	ToCardsResponse(cards []*record.CardRecord) []*response.CardResponse

	ToCardResponseDeleteAt(card *record.CardRecord) *response.CardResponseDeleteAt
	ToCardsResponseDeleteAt(cards []*record.CardRecord) []*response.CardResponseDeleteAt
}

type UserResponseMapper interface {
	ToUserResponse(user *record.UserRecord) *response.UserResponse
	ToUsersResponse(users []*record.UserRecord) []*response.UserResponse

	ToUserResponseDeleteAt(user *record.UserRecord) *response.UserResponseDeleteAt
	ToUsersResponseDeleteAt(users []*record.UserRecord) []*response.UserResponseDeleteAt
}

type RoleResponseMapper interface {
	ToRoleResponse(role *record.RoleRecord) *response.RoleResponse
	ToRolesResponse(roles []*record.RoleRecord) []*response.RoleResponse

	ToRoleResponseDeleteAt(role *record.RoleRecord) *response.RoleResponseDeleteAt
	ToRolesResponseDeleteAt(roles []*record.RoleRecord) []*response.RoleResponseDeleteAt
}

type RefreshTokenResponseMapper interface {
	ToRefreshTokenResponse(refresh *record.RefreshTokenRecord) *response.RefreshTokenResponse
	ToRefreshTokenResponses(refreshs []*record.RefreshTokenRecord) []*response.RefreshTokenResponse
}

type SaldoResponseMapper interface {
	ToSaldoResponse(saldo *record.SaldoRecord) *response.SaldoResponse
	ToSaldoResponses(saldos []*record.SaldoRecord) []*response.SaldoResponse

	ToSaldoResponseDeleteAt(saldo *record.SaldoRecord) *response.SaldoResponseDeleteAt
	ToSaldoResponsesDeleteAt(saldos []*record.SaldoRecord) []*response.SaldoResponseDeleteAt
}

type TopupResponseMapper interface {
	ToTopupResponse(topup *record.TopupRecord) *response.TopupResponse
	ToTopupResponses(topups []*record.TopupRecord) []*response.TopupResponse

	ToTopupResponseDeleteAt(topup *record.TopupRecord) *response.TopupResponseDeleteAt
	ToTopupResponsesDeleteAt(topups []*record.TopupRecord) []*response.TopupResponseDeleteAt
}

type TransactionResponseMapper interface {
	ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse
	ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse

	ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt
	ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt
}

type TransferResponseMapper interface {
	ToTransferResponse(transfer *record.TransferRecord) *response.TransferResponse
	ToTransfersResponse(transfers []*record.TransferRecord) []*response.TransferResponse

	ToTransferResponseDeleteAt(transfer *record.TransferRecord) *response.TransferResponseDeleteAt
	ToTransfersResponseDeleteAt(transfers []*record.TransferRecord) []*response.TransferResponseDeleteAt
}

type WithdrawResponseMapper interface {
	ToWithdrawResponse(withdraw *record.WithdrawRecord) *response.WithdrawResponse
	ToWithdrawsResponse(withdraws []*record.WithdrawRecord) []*response.WithdrawResponse

	ToWithdrawResponseDeleteAt(withdraw *record.WithdrawRecord) *response.WithdrawResponseDeleteAt
	ToWithdrawsResponseDeleteAt(withdraws []*record.WithdrawRecord) []*response.WithdrawResponseDeleteAt
}

type MerchantResponseMapper interface {
	ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse
	ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse

	ToMerchantResponseDeleteAt(merchant *record.MerchantRecord) *response.MerchantResponseDeleteAt
	ToMerchantsResponseDeleteAt(merchants []*record.MerchantRecord) []*response.MerchantResponseDeleteAt
}
