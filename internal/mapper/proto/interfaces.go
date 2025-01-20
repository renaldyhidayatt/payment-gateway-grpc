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

type RoleProtoMapper interface {
	ToResponseRole(role *response.RoleResponse) *pb.RoleResponse
	ToResponsesRole(roles []*response.RoleResponse) []*pb.RoleResponse
	ToResponseRoleDeleteAt(role *response.RoleResponseDeleteAt) *pb.RoleResponseDeleteAt
	ToResponsesRoleDeleteAt(roles []*response.RoleResponseDeleteAt) []*pb.RoleResponseDeleteAt
}

type CardProtoMapper interface {
	ToResponseCard(card *response.CardResponse) *pb.CardResponse
	ToResponsesCard(cards []*response.CardResponse) []*pb.CardResponse
	ToResponseCardDeleteAt(card *response.CardResponseDeleteAt) *pb.CardResponseDeleteAt
	ToResponsesCardDeletedAt(cards []*response.CardResponseDeleteAt) []*pb.CardResponseDeleteAt

	ToResponseDashboardCard(dash *response.DashboardCard) *pb.CardResponseDashboard
	ToResponseDashboardCardCardNumber(dash *response.DashboardCardCardNumber) *pb.CardResponseDashboardCardNumber

	ToResponseMonthlyBalance(cards *response.CardResponseMonthBalance) *pb.CardResponseMonthlyBalance
	ToResponseMonthlyBalances(cards []*response.CardResponseMonthBalance) []*pb.CardResponseMonthlyBalance

	ToResponseYearlyBalance(cards *response.CardResponseYearlyBalance) *pb.CardResponseYearlyBalance
	ToResponseYearlyBalances(cards []*response.CardResponseYearlyBalance) []*pb.CardResponseYearlyBalance

	ToResponseMonthlyTopupAmount(cards *response.CardResponseMonthTopupAmount) *pb.CardResponseMonthlyAmount
	ToResponseMonthlyTopupAmounts(cards []*response.CardResponseMonthTopupAmount) []*pb.CardResponseMonthlyAmount

	ToResponseYearlyTopupAmount(cards *response.CardResponseYearlyTopupAmount) *pb.CardResponseYearlyAmount
	ToResponseYearlyTopupAmounts(cards []*response.CardResponseYearlyTopupAmount) []*pb.CardResponseYearlyAmount

	ToResponseMonthlyWithdrawAmount(cards *response.CardResponseMonthWithdrawAmount) *pb.CardResponseMonthlyAmount
	ToResponseMonthlyWithdrawAmounts(cards []*response.CardResponseMonthWithdrawAmount) []*pb.CardResponseMonthlyAmount

	ToResponseYearlyWithdrawAmount(cards *response.CardResponseYearlyWithdrawAmount) *pb.CardResponseYearlyAmount
	ToResponseYearlyWithdrawAmounts(cards []*response.CardResponseYearlyWithdrawAmount) []*pb.CardResponseYearlyAmount

	ToResponseMonthlyTransactionAmount(cards *response.CardResponseMonthTransactionAmount) *pb.CardResponseMonthlyAmount
	ToResponseMonthlyTransactionAmounts(cards []*response.CardResponseMonthTransactionAmount) []*pb.CardResponseMonthlyAmount

	ToResponseYearlyTransactionAmount(cards *response.CardResponseYearlyTransactionAmount) *pb.CardResponseYearlyAmount
	ToResponseYearlyTransactionAmounts(cards []*response.CardResponseYearlyTransactionAmount) []*pb.CardResponseYearlyAmount

	ToResponseMonthlyTransferSenderAmount(cards *response.CardResponseMonthTransferAmount) *pb.CardResponseMonthlyAmount
	ToResponseMonthlyTransferSenderAmounts(cards []*response.CardResponseMonthTransferAmount) []*pb.CardResponseMonthlyAmount

	ToResponseYearlyTransferSenderAmount(cards *response.CardResponseYearlyTransferAmount) *pb.CardResponseYearlyAmount
	ToResponseYearlyTransferSenderAmounts(cards []*response.CardResponseYearlyTransferAmount) []*pb.CardResponseYearlyAmount

	ToResponseMonthlyTransferReceiverAmount(cards *response.CardResponseMonthTransferAmount) *pb.CardResponseMonthlyAmount
	ToResponseMonthlyTransferReceiverAmounts(cards []*response.CardResponseMonthTransferAmount) []*pb.CardResponseMonthlyAmount

	ToResponseYearlyTransferReceiverAmount(cards *response.CardResponseYearlyTransferAmount) *pb.CardResponseYearlyAmount
	ToResponseYearlyTransferReceiverAmounts(cards []*response.CardResponseYearlyTransferAmount) []*pb.CardResponseYearlyAmount
}

type MerchantProtoMapper interface {
	ToResponseMerchant(merchant *response.MerchantResponse) *pb.MerchantResponse
	ToResponsesMerchant(merchants []*response.MerchantResponse) []*pb.MerchantResponse

	ToResponseMonthlyPaymentMethod(ms *response.MerchantResponseMonthlyPaymentMethod) *pb.MerchantResponseMonthlyPaymentMethod
	ToResponseMonthlyPaymentMethods(ms []*response.MerchantResponseMonthlyPaymentMethod) []*pb.MerchantResponseMonthlyPaymentMethod
	ToResponseYearlyPaymentMethod(ms *response.MerchantResponseYearlyPaymentMethod) *pb.MerchantResponseYearlyPaymentMethod
	ToResponseYearlyPaymentMethods(ms []*response.MerchantResponseYearlyPaymentMethod) []*pb.MerchantResponseYearlyPaymentMethod

	ToResponseMonthlyAmount(ms *response.MerchantResponseMonthlyAmount) *pb.MerchantResponseMonthlyAmount
	ToResponseMonthlyAmounts(ms []*response.MerchantResponseMonthlyAmount) []*pb.MerchantResponseMonthlyAmount
	ToResponseYearlyAmount(ms *response.MerchantResponseYearlyAmount) *pb.MerchantResponseYearlyAmount
	ToResponseYearlyAmounts(ms []*response.MerchantResponseYearlyAmount) []*pb.MerchantResponseYearlyAmount

	ToResponseMerchantDeleteAt(merchant *response.MerchantResponseDeleteAt) *pb.MerchantResponseDeleteAt
	ToResponsesMerchantDeleteAt(merchants []*response.MerchantResponseDeleteAt) []*pb.MerchantResponseDeleteAt
}

type SaldoProtoMapper interface {
	ToResponseSaldo(saldo *response.SaldoResponse) *pb.SaldoResponse
	ToResponsesSaldo(saldos []*response.SaldoResponse) []*pb.SaldoResponse

	ToSaldoMonthTotalBalanceResponse(ss *response.SaldoMonthTotalBalanceResponse) *pb.SaldoMonthTotalBalanceResponse
	ToSaldoMonthTotalBalanceResponses(ss []*response.SaldoMonthTotalBalanceResponse) []*pb.SaldoMonthTotalBalanceResponse

	ToSaldoYearTotalBalanceResponse(ss *response.SaldoYearTotalBalanceResponse) *pb.SaldoYearTotalBalanceResponse
	ToSaldoYearTotalBalanceResponses(ss []*response.SaldoYearTotalBalanceResponse) []*pb.SaldoYearTotalBalanceResponse

	ToSaldoMonthBalanceResponse(ss *response.SaldoMonthBalanceResponse) *pb.SaldoMonthBalanceResponse
	ToSaldoMonthBalanceResponses(ss []*response.SaldoMonthBalanceResponse) []*pb.SaldoMonthBalanceResponse

	ToSaldoYearBalanceResponse(ss *response.SaldoYearBalanceResponse) *pb.SaldoYearBalanceResponse
	ToSaldoYearBalanceResponses(ss []*response.SaldoYearBalanceResponse) []*pb.SaldoYearBalanceResponse

	ToResponseSaldoDeleteAt(saldo *response.SaldoResponseDeleteAt) *pb.SaldoResponseDeleteAt
	ToResponsesSaldoDeleteAt(saldos []*response.SaldoResponseDeleteAt) []*pb.SaldoResponseDeleteAt
}

type TopupProtoMapper interface {
	ToResponseTopup(topup *response.TopupResponse) *pb.TopupResponse
	ToResponsesTopup(topups []*response.TopupResponse) []*pb.TopupResponse

	ToResponseTopupMonthStatusSuccess(s *response.TopupResponseMonthStatusSuccess) *pb.TopupMonthStatusSuccessResponse
	ToResponsesTopupMonthStatusSuccess(topups []*response.TopupResponseMonthStatusSuccess) []*pb.TopupMonthStatusSuccessResponse
	ToTopupResponseYearStatusSuccess(s *response.TopupResponseYearStatusSuccess) *pb.TopupYearStatusSuccessResponse
	ToTopupResponsesYearStatusSuccess(topups []*response.TopupResponseYearStatusSuccess) []*pb.TopupYearStatusSuccessResponse

	ToResponseTopupMonthStatusFailed(s *response.TopupResponseMonthStatusFailed) *pb.TopupMonthStatusFailedResponse
	ToResponsesTopupMonthStatusFailed(topups []*response.TopupResponseMonthStatusFailed) []*pb.TopupMonthStatusFailedResponse
	ToTopupResponseYearStatusFailed(s *response.TopupResponseYearStatusFailed) *pb.TopupYearStatusFailedResponse
	ToTopupResponsesYearStatusFailed(topups []*response.TopupResponseYearStatusFailed) []*pb.TopupYearStatusFailedResponse

	ToResponseTopupMonthlyMethod(s *response.TopupMonthMethodResponse) *pb.TopupMonthMethodResponse
	ToResponseTopupMonthlyMethods(s []*response.TopupMonthMethodResponse) []*pb.TopupMonthMethodResponse
	ToResponseTopupYearlyMethod(s *response.TopupYearlyMethodResponse) *pb.TopupYearlyMethodResponse
	ToResponseTopupYearlyMethods(s []*response.TopupYearlyMethodResponse) []*pb.TopupYearlyMethodResponse

	ToResponseTopupMonthlyAmount(s *response.TopupMonthAmountResponse) *pb.TopupMonthAmountResponse
	ToResponseTopupMonthlyAmounts(s []*response.TopupMonthAmountResponse) []*pb.TopupMonthAmountResponse
	ToResponseTopupYearlyAmount(s *response.TopupYearlyAmountResponse) *pb.TopupYearlyAmountResponse
	ToResponseTopupYearlyAmounts(s []*response.TopupYearlyAmountResponse) []*pb.TopupYearlyAmountResponse

	ToResponseTopupDeleteAt(topup *response.TopupResponseDeleteAt) *pb.TopupResponseDeleteAt
	ToResponsesTopupDeleteAt(topups []*response.TopupResponseDeleteAt) []*pb.TopupResponseDeleteAt
}

type TransactionProtoMapper interface {
	ToResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse
	ToResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse

	ToResponseTransactionMonthStatusSuccess(s *response.TransactionResponseMonthStatusSuccess) *pb.TransactionMonthStatusSuccessResponse
	ToResponsesTransactionMonthStatusSuccess(Transactions []*response.TransactionResponseMonthStatusSuccess) []*pb.TransactionMonthStatusSuccessResponse
	ToTransactionResponseYearStatusSuccess(s *response.TransactionResponseYearStatusSuccess) *pb.TransactionYearStatusSuccessResponse
	ToTransactionResponsesYearStatusSuccess(Transactions []*response.TransactionResponseYearStatusSuccess) []*pb.TransactionYearStatusSuccessResponse

	ToResponseTransactionMonthStatusFailed(s *response.TransactionResponseMonthStatusFailed) *pb.TransactionMonthStatusFailedResponse
	ToResponsesTransactionMonthStatusFailed(Transactions []*response.TransactionResponseMonthStatusFailed) []*pb.TransactionMonthStatusFailedResponse
	ToTransactionResponseYearStatusFailed(s *response.TransactionResponseYearStatusFailed) *pb.TransactionYearStatusFailedResponse
	ToTransactionResponsesYearStatusFailed(Transactions []*response.TransactionResponseYearStatusFailed) []*pb.TransactionYearStatusFailedResponse

	ToResponseTransactionMonthMethod(s *response.TransactionMonthMethodResponse) *pb.TransactionMonthMethodResponse
	ToResponseTransactionMonthMethods(s []*response.TransactionMonthMethodResponse) []*pb.TransactionMonthMethodResponse
	ToResponseTransactionYearMethod(s *response.TransactionYearMethodResponse) *pb.TransactionYearMethodResponse
	ToResponseTransactionYearMethods(s []*response.TransactionYearMethodResponse) []*pb.TransactionYearMethodResponse

	ToResponseTransactionMonthAmount(s *response.TransactionMonthAmountResponse) *pb.TransactionMonthAmountResponse
	ToResponseTransactionMonthAmounts(s []*response.TransactionMonthAmountResponse) []*pb.TransactionMonthAmountResponse
	ToResponseTransactionYearlyAmount(s *response.TransactionYearlyAmountResponse) *pb.TransactionYearlyAmountResponse
	ToResponseTransactionYearlyAmounts(s []*response.TransactionYearlyAmountResponse) []*pb.TransactionYearlyAmountResponse

	ToResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *pb.TransactionResponseDeleteAt
	ToResponsesTransactionDeleteAt(transactions []*response.TransactionResponseDeleteAt) []*pb.TransactionResponseDeleteAt
}

type TransferProtoMapper interface {
	ToResponseTransfer(transfer *response.TransferResponse) *pb.TransferResponse
	ToResponsesTransfer(transfers []*response.TransferResponse) []*pb.TransferResponse

	ToResponseTransferMonthStatusSuccess(s *response.TransferResponseMonthStatusSuccess) *pb.TransferMonthStatusSuccessResponse
	ToResponsesTransferMonthStatusSuccess(Transfers []*response.TransferResponseMonthStatusSuccess) []*pb.TransferMonthStatusSuccessResponse
	ToTransferResponseYearStatusSuccess(s *response.TransferResponseYearStatusSuccess) *pb.TransferYearStatusSuccessResponse
	ToTransferResponsesYearStatusSuccess(Transfers []*response.TransferResponseYearStatusSuccess) []*pb.TransferYearStatusSuccessResponse

	ToResponseTransferMonthStatusFailed(s *response.TransferResponseMonthStatusFailed) *pb.TransferMonthStatusFailedResponse
	ToResponsesTransferMonthStatusFailed(Transfers []*response.TransferResponseMonthStatusFailed) []*pb.TransferMonthStatusFailedResponse
	ToTransferResponseYearStatusFailed(s *response.TransferResponseYearStatusFailed) *pb.TransferYearStatusFailedResponse
	ToTransferResponsesYearStatusFailed(Transfers []*response.TransferResponseYearStatusFailed) []*pb.TransferYearStatusFailedResponse

	ToResponseTransferMonthAmount(s *response.TransferMonthAmountResponse) *pb.TransferMonthAmountResponse
	ToResponseTransferMonthAmounts(s []*response.TransferMonthAmountResponse) []*pb.TransferMonthAmountResponse

	ToResponseTransferYearAmount(s *response.TransferYearAmountResponse) *pb.TransferYearAmountResponse
	ToResponseTransferYearAmounts(s []*response.TransferYearAmountResponse) []*pb.TransferYearAmountResponse

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

	ToResponseWithdrawMonthStatusSuccess(s *response.WithdrawResponseMonthStatusSuccess) *pb.WithdrawMonthStatusSuccessResponse
	ToResponsesWithdrawMonthStatusSuccess(Withdraws []*response.WithdrawResponseMonthStatusSuccess) []*pb.WithdrawMonthStatusSuccessResponse
	ToWithdrawResponseYearStatusSuccess(s *response.WithdrawResponseYearStatusSuccess) *pb.WithdrawYearStatusSuccessResponse
	ToWithdrawResponsesYearStatusSuccess(Withdraws []*response.WithdrawResponseYearStatusSuccess) []*pb.WithdrawYearStatusSuccessResponse

	ToResponseWithdrawMonthStatusFailed(s *response.WithdrawResponseMonthStatusFailed) *pb.WithdrawMonthStatusFailedResponse
	ToResponsesWithdrawMonthStatusFailed(Withdraws []*response.WithdrawResponseMonthStatusFailed) []*pb.WithdrawMonthStatusFailedResponse
	ToWithdrawResponseYearStatusFailed(s *response.WithdrawResponseYearStatusFailed) *pb.WithdrawYearStatusFailedResponse
	ToWithdrawResponsesYearStatusFailed(Withdraws []*response.WithdrawResponseYearStatusFailed) []*pb.WithdrawYearStatusFailedResponse

	ToResponseWithdrawMonthlyAmount(s *response.WithdrawMonthlyAmountResponse) *pb.WithdrawMonthlyAmountResponse
	ToResponseWithdrawMonthlyAmounts(s []*response.WithdrawMonthlyAmountResponse) []*pb.WithdrawMonthlyAmountResponse

	ToResponseWithdrawYearlyAmount(s *response.WithdrawYearlyAmountResponse) *pb.WithdrawYearlyAmountResponse
	ToResponseWithdrawYearlyAmounts(s []*response.WithdrawYearlyAmountResponse) []*pb.WithdrawYearlyAmountResponse

	ToResponseWithdrawalDeleteAt(withdraw *response.WithdrawResponseDeleteAt) *pb.WithdrawResponseDeleteAt
	ToResponsesWithdrawalDeleteAt(withdraws []*response.WithdrawResponseDeleteAt) []*pb.WithdrawResponseDeleteAt
}
