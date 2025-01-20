package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
)

type UserRecordMapping interface {
	ToUserRecord(user *db.User) *record.UserRecord
	ToUserRecordPagination(user *db.GetUsersWithPaginationRow) *record.UserRecord
	ToUsersRecordPagination(users []*db.GetUsersWithPaginationRow) []*record.UserRecord

	ToUserRecordActivePagination(user *db.GetActiveUsersWithPaginationRow) *record.UserRecord
	ToUsersRecordActivePagination(users []*db.GetActiveUsersWithPaginationRow) []*record.UserRecord
	ToUserRecordTrashedPagination(user *db.GetTrashedUsersWithPaginationRow) *record.UserRecord
	ToUsersRecordTrashedPagination(users []*db.GetTrashedUsersWithPaginationRow) []*record.UserRecord
}

type RoleRecordMapping interface {
	ToRoleRecord(role *db.Role) *record.RoleRecord
	ToRolesRecord(roles []*db.Role) []*record.RoleRecord

	ToRoleRecordAll(role *db.GetRolesRow) *record.RoleRecord
	ToRolesRecordAll(roles []*db.GetRolesRow) []*record.RoleRecord

	ToRoleRecordActive(role *db.GetActiveRolesRow) *record.RoleRecord
	ToRolesRecordActive(roles []*db.GetActiveRolesRow) []*record.RoleRecord
	ToRoleRecordTrashed(role *db.GetTrashedRolesRow) *record.RoleRecord
	ToRolesRecordTrashed(roles []*db.GetTrashedRolesRow) []*record.RoleRecord
}

type UserRoleRecordMapping interface {
	ToUserRoleRecord(userRole *db.UserRole) *record.UserRoleRecord
}

type RefreshTokenRecordMapping interface {
	ToRefreshTokenRecord(refreshToken *db.RefreshToken) *record.RefreshTokenRecord
	ToRefreshTokensRecord(refreshTokens []*db.RefreshToken) []*record.RefreshTokenRecord
}

type SaldoRecordMapping interface {
	ToSaldoRecord(saldo *db.Saldo) *record.SaldoRecord
	ToSaldosRecord(saldos []*db.Saldo) []*record.SaldoRecord

	ToSaldoMonthTotalBalance(ss *db.GetMonthlyTotalSaldoBalanceRow) *record.SaldoMonthTotalBalance
	ToSaldoMonthTotalBalances(ss []*db.GetMonthlyTotalSaldoBalanceRow) []*record.SaldoMonthTotalBalance

	ToSaldoYearTotalBalance(ss *db.GetYearlyTotalSaldoBalancesRow) *record.SaldoYearTotalBalance
	ToSaldoYearTotalBalances(ss []*db.GetYearlyTotalSaldoBalancesRow) []*record.SaldoYearTotalBalance

	ToSaldoMonthBalance(ss *db.GetMonthlySaldoBalancesRow) *record.SaldoMonthSaldoBalance
	ToSaldoMonthBalances(ss []*db.GetMonthlySaldoBalancesRow) []*record.SaldoMonthSaldoBalance
	ToSaldoYearSaldoBalance(ss *db.GetYearlySaldoBalancesRow) *record.SaldoYearSaldoBalance
	ToSaldoYearSaldoBalances(ss []*db.GetYearlySaldoBalancesRow) []*record.SaldoYearSaldoBalance

	ToSaldoRecordAll(saldo *db.GetSaldosRow) *record.SaldoRecord
	ToSaldosRecordAll(saldos []*db.GetSaldosRow) []*record.SaldoRecord

	ToSaldoRecordActive(saldo *db.GetActiveSaldosRow) *record.SaldoRecord
	ToSaldosRecordActive(saldos []*db.GetActiveSaldosRow) []*record.SaldoRecord

	ToSaldoRecordTrashed(saldo *db.GetTrashedSaldosRow) *record.SaldoRecord
	ToSaldosRecordTrashed(saldos []*db.GetTrashedSaldosRow) []*record.SaldoRecord
}

type TopupRecordMapping interface {
	ToTopupRecord(topup *db.Topup) *record.TopupRecord
	ToTopupRecords(topups []*db.Topup) []*record.TopupRecord

	ToTopupRecordMonthStatusSuccess(s *db.GetMonthTopupStatusSuccessRow) *record.TopupRecordMonthStatusSuccess
	ToTopupRecordsMonthStatusSuccess(topups []*db.GetMonthTopupStatusSuccessRow) []*record.TopupRecordMonthStatusSuccess
	ToTopupRecordYearStatusSuccess(s *db.GetYearlyTopupStatusSuccessRow) *record.TopupRecordYearStatusSuccess
	ToTopupRecordsYearStatusSuccess(topups []*db.GetYearlyTopupStatusSuccessRow) []*record.TopupRecordYearStatusSuccess

	ToTopupRecordMonthStatusFailed(s *db.GetMonthTopupStatusFailedRow) *record.TopupRecordMonthStatusFailed
	ToTopupRecordsMonthStatusFailed(topups []*db.GetMonthTopupStatusFailedRow) []*record.TopupRecordMonthStatusFailed
	ToTopupRecordYearStatusFailed(s *db.GetYearlyTopupStatusFailedRow) *record.TopupRecordYearStatusFailed
	ToTopupRecordsYearStatusFailed(topups []*db.GetYearlyTopupStatusFailedRow) []*record.TopupRecordYearStatusFailed

	ToTopupMonthlyMethod(s *db.GetMonthlyTopupMethodsRow) *record.TopupMonthMethod
	ToTopupMonthlyMethods(s []*db.GetMonthlyTopupMethodsRow) []*record.TopupMonthMethod
	ToTopupYearlyMethod(s *db.GetYearlyTopupMethodsRow) *record.TopupYearlyMethod
	ToTopupYearlyMethods(s []*db.GetYearlyTopupMethodsRow) []*record.TopupYearlyMethod
	ToTopupMonthlyAmount(s *db.GetMonthlyTopupAmountsRow) *record.TopupMonthAmount
	ToTopupMonthlyAmounts(s []*db.GetMonthlyTopupAmountsRow) []*record.TopupMonthAmount
	ToTopupYearlyAmount(s *db.GetYearlyTopupAmountsRow) *record.TopupYearlyAmount
	ToTopupYearlyAmounts(s []*db.GetYearlyTopupAmountsRow) []*record.TopupYearlyAmount

	ToTopupMonthlyMethodByCardNumber(s *db.GetMonthlyTopupMethodsByCardNumberRow) *record.TopupMonthMethod
	ToTopupMonthlyMethodsByCardNumber(s []*db.GetMonthlyTopupMethodsByCardNumberRow) []*record.TopupMonthMethod
	ToTopupYearlyMethodByCardNumber(s *db.GetYearlyTopupMethodsByCardNumberRow) *record.TopupYearlyMethod
	ToTopupYearlyMethodsByCardNumber(s []*db.GetYearlyTopupMethodsByCardNumberRow) []*record.TopupYearlyMethod
	ToTopupMonthlyAmountByCardNumber(s *db.GetMonthlyTopupAmountsByCardNumberRow) *record.TopupMonthAmount
	ToTopupMonthlyAmountsByCardNumber(s []*db.GetMonthlyTopupAmountsByCardNumberRow) []*record.TopupMonthAmount
	ToTopupYearlyAmountByCardNumber(s *db.GetYearlyTopupAmountsByCardNumberRow) *record.TopupYearlyAmount
	ToTopupYearlyAmountsByCardNumber(s []*db.GetYearlyTopupAmountsByCardNumberRow) []*record.TopupYearlyAmount

	ToTopupRecordAll(topup *db.GetTopupsRow) *record.TopupRecord
	ToTopupRecordsAll(topups []*db.GetTopupsRow) []*record.TopupRecord

	ToTopupRecordActive(topup *db.GetActiveTopupsRow) *record.TopupRecord
	ToTopupRecordsActive(topups []*db.GetActiveTopupsRow) []*record.TopupRecord
	ToTopupRecordTrashed(topup *db.GetTrashedTopupsRow) *record.TopupRecord
	ToTopupRecordsTrashed(topups []*db.GetTrashedTopupsRow) []*record.TopupRecord
}

type TransferRecordMapping interface {
	ToTransferRecord(transfer *db.Transfer) *record.TransferRecord
	ToTransfersRecord(transfers []*db.Transfer) []*record.TransferRecord

	ToTransferRecordMonthStatusSuccess(s *db.GetMonthTransferStatusSuccessRow) *record.TransferRecordMonthStatusSuccess
	ToTransferRecordsMonthStatusSuccess(Transfers []*db.GetMonthTransferStatusSuccessRow) []*record.TransferRecordMonthStatusSuccess
	ToTransferRecordYearStatusSuccess(s *db.GetYearlyTransferStatusSuccessRow) *record.TransferRecordYearStatusSuccess
	ToTransferRecordsYearStatusSuccess(Transfers []*db.GetYearlyTransferStatusSuccessRow) []*record.TransferRecordYearStatusSuccess

	ToTransferRecordMonthStatusFailed(s *db.GetMonthTransferStatusFailedRow) *record.TransferRecordMonthStatusFailed
	ToTransferRecordsMonthStatusFailed(Transfers []*db.GetMonthTransferStatusFailedRow) []*record.TransferRecordMonthStatusFailed
	ToTransferRecordYearStatusFailed(s *db.GetYearlyTransferStatusFailedRow) *record.TransferRecordYearStatusFailed
	ToTransferRecordsYearStatusFailed(Transfers []*db.GetYearlyTransferStatusFailedRow) []*record.TransferRecordYearStatusFailed

	ToTransferMonthAmount(ss *db.GetMonthlyTransferAmountsRow) *record.TransferMonthAmount
	ToTransferMonthAmounts(ss []*db.GetMonthlyTransferAmountsRow) []*record.TransferMonthAmount
	ToTransferYearAmount(ss *db.GetYearlyTransferAmountsRow) *record.TransferYearAmount
	ToTransferYearAmounts(ss []*db.GetYearlyTransferAmountsRow) []*record.TransferYearAmount

	ToTransferMonthAmountSender(ss *db.GetMonthlyTransferAmountsBySenderCardNumberRow) *record.TransferMonthAmount
	ToTransferMonthAmountsSender(ss []*db.GetMonthlyTransferAmountsBySenderCardNumberRow) []*record.TransferMonthAmount
	ToTransferYearAmountSender(ss *db.GetYearlyTransferAmountsBySenderCardNumberRow) *record.TransferYearAmount
	ToTransferYearAmountsSender(ss []*db.GetYearlyTransferAmountsBySenderCardNumberRow) []*record.TransferYearAmount

	ToTransferMonthAmountReceiver(ss *db.GetMonthlyTransferAmountsByReceiverCardNumberRow) *record.TransferMonthAmount
	ToTransferMonthAmountsReceiver(ss []*db.GetMonthlyTransferAmountsByReceiverCardNumberRow) []*record.TransferMonthAmount
	ToTransferYearAmountReceiver(ss *db.GetYearlyTransferAmountsByReceiverCardNumberRow) *record.TransferYearAmount
	ToTransferYearAmountsReceiver(ss []*db.GetYearlyTransferAmountsByReceiverCardNumberRow) []*record.TransferYearAmount

	ToTransferRecordAll(transfer *db.GetTransfersRow) *record.TransferRecord
	ToTransfersRecordAll(transfers []*db.GetTransfersRow) []*record.TransferRecord

	ToTransferRecordActive(transfer *db.GetActiveTransfersRow) *record.TransferRecord
	ToTransfersRecordActive(transfers []*db.GetActiveTransfersRow) []*record.TransferRecord
	ToTransferRecordTrashed(transfer *db.GetTrashedTransfersRow) *record.TransferRecord
	ToTransfersRecordTrashed(transfers []*db.GetTrashedTransfersRow) []*record.TransferRecord
}

type WithdrawRecordMapping interface {
	ToWithdrawRecord(withdraw *db.Withdraw) *record.WithdrawRecord
	ToWithdrawsRecord(withdraws []*db.Withdraw) []*record.WithdrawRecord

	ToWithdrawRecordMonthStatusSuccess(s *db.GetMonthWithdrawStatusSuccessRow) *record.WithdrawRecordMonthStatusSuccess
	ToWithdrawRecordsMonthStatusSuccess(Withdraws []*db.GetMonthWithdrawStatusSuccessRow) []*record.WithdrawRecordMonthStatusSuccess
	ToWithdrawRecordYearStatusSuccess(s *db.GetYearlyWithdrawStatusSuccessRow) *record.WithdrawRecordYearStatusSuccess
	ToWithdrawRecordsYearStatusSuccess(Withdraws []*db.GetYearlyWithdrawStatusSuccessRow) []*record.WithdrawRecordYearStatusSuccess

	ToWithdrawRecordMonthStatusFailed(s *db.GetMonthWithdrawStatusFailedRow) *record.WithdrawRecordMonthStatusFailed
	ToWithdrawRecordsMonthStatusFailed(Withdraws []*db.GetMonthWithdrawStatusFailedRow) []*record.WithdrawRecordMonthStatusFailed
	ToWithdrawRecordYearStatusFailed(s *db.GetYearlyWithdrawStatusFailedRow) *record.WithdrawRecordYearStatusFailed
	ToWithdrawRecordsYearStatusFailed(Withdraws []*db.GetYearlyWithdrawStatusFailedRow) []*record.WithdrawRecordYearStatusFailed

	ToWithdrawAmountMonthly(ss *db.GetMonthlyWithdrawsRow) *record.WithdrawMonthlyAmount
	ToWithdrawsAmountMonthly(ss []*db.GetMonthlyWithdrawsRow) []*record.WithdrawMonthlyAmount

	ToWithdrawAmountYearly(ss *db.GetYearlyWithdrawsRow) *record.WithdrawYearlyAmount
	ToWithdrawsAmountYearly(ss []*db.GetYearlyWithdrawsRow) []*record.WithdrawYearlyAmount

	ToWithdrawAmountMonthlyByCardNumber(ss *db.GetMonthlyWithdrawsByCardNumberRow) *record.WithdrawMonthlyAmount
	ToWithdrawsAmountMonthlyByCardNumber(ss []*db.GetMonthlyWithdrawsByCardNumberRow) []*record.WithdrawMonthlyAmount

	ToWithdrawAmountYearlyByCardNumber(ss *db.GetYearlyWithdrawsByCardNumberRow) *record.WithdrawYearlyAmount
	ToWithdrawsAmountYearlyByCardNumber(ss []*db.GetYearlyWithdrawsByCardNumberRow) []*record.WithdrawYearlyAmount

	ToWithdrawRecordAll(withdraw *db.GetWithdrawsRow) *record.WithdrawRecord
	ToWithdrawsRecordALl(withdraws []*db.GetWithdrawsRow) []*record.WithdrawRecord

	ToWithdrawRecordActive(withdraw *db.GetActiveWithdrawsRow) *record.WithdrawRecord
	ToWithdrawsRecordActive(withdraws []*db.GetActiveWithdrawsRow) []*record.WithdrawRecord

	ToWithdrawRecordTrashed(withdraw *db.GetTrashedWithdrawsRow) *record.WithdrawRecord
	ToWithdrawsRecordTrashed(withdraws []*db.GetTrashedWithdrawsRow) []*record.WithdrawRecord
}

type CardRecordMapping interface {
	ToCardRecord(card *db.Card) *record.CardRecord
	ToCardsRecord(cards []*db.GetCardsRow) []*record.CardRecord

	ToCardRecordActive(card *db.GetActiveCardsWithCountRow) *record.CardRecord
	ToCardRecordsActive(cards []*db.GetActiveCardsWithCountRow) []*record.CardRecord

	ToCardRecordTrashed(card *db.GetTrashedCardsWithCountRow) *record.CardRecord
	ToCardRecordsTrashed(cards []*db.GetTrashedCardsWithCountRow) []*record.CardRecord

	ToMonthlyBalance(card *db.GetMonthlyBalancesRow) *record.CardMonthBalance
	ToMonthlyBalances(cards []*db.GetMonthlyBalancesRow) []*record.CardMonthBalance

	ToYearlyBalance(card *db.GetYearlyBalancesRow) *record.CardYearlyBalance
	ToYearlyBalances(cards []*db.GetYearlyBalancesRow) []*record.CardYearlyBalance

	ToMonthlyTopupAmount(card *db.GetMonthlyTopupAmountRow) *record.CardMonthTopupAmount
	ToMonthlyTopupAmounts(cards []*db.GetMonthlyTopupAmountRow) []*record.CardMonthTopupAmount

	ToYearlyTopupAmount(card *db.GetYearlyTopupAmountRow) *record.CardYearlyTopupAmount
	ToYearlyTopupAmounts(cards []*db.GetYearlyTopupAmountRow) []*record.CardYearlyTopupAmount

	ToMonthlyWithdrawAmount(card *db.GetMonthlyWithdrawAmountRow) *record.CardMonthWithdrawAmount
	ToMonthlyWithdrawAmounts(cards []*db.GetMonthlyWithdrawAmountRow) []*record.CardMonthWithdrawAmount

	ToYearlyWithdrawAmount(card *db.GetYearlyWithdrawAmountRow) *record.CardYearlyWithdrawAmount
	ToYearlyWithdrawAmounts(cards []*db.GetYearlyWithdrawAmountRow) []*record.CardYearlyWithdrawAmount

	ToMonthlyTransactionAmount(card *db.GetMonthlyTransactionAmountRow) *record.CardMonthTransactionAmount
	ToMonthlyTransactionAmounts(cards []*db.GetMonthlyTransactionAmountRow) []*record.CardMonthTransactionAmount

	ToYearlyTransactionAmount(card *db.GetYearlyTransactionAmountRow) *record.CardYearlyTransactionAmount
	ToYearlyTransactionAmounts(cards []*db.GetYearlyTransactionAmountRow) []*record.CardYearlyTransactionAmount

	ToMonthlyTransferSenderAmount(card *db.GetMonthlyTransferAmountSenderRow) *record.CardMonthTransferAmount
	ToMonthlyTransferSenderAmounts(cards []*db.GetMonthlyTransferAmountSenderRow) []*record.CardMonthTransferAmount

	ToYearlyTransferSenderAmount(card *db.GetYearlyTransferAmountSenderRow) *record.CardYearlyTransferAmount
	ToYearlyTransferSenderAmounts(cards []*db.GetYearlyTransferAmountSenderRow) []*record.CardYearlyTransferAmount

	ToMonthlyTransferReceiverAmount(card *db.GetMonthlyTransferAmountReceiverRow) *record.CardMonthTransferAmount
	ToMonthlyTransferReceiverAmounts(cards []*db.GetMonthlyTransferAmountReceiverRow) []*record.CardMonthTransferAmount

	ToYearlyTransferReceiverAmount(card *db.GetYearlyTransferAmountReceiverRow) *record.CardYearlyTransferAmount
	ToYearlyTransferReceiverAmounts(cards []*db.GetYearlyTransferAmountReceiverRow) []*record.CardYearlyTransferAmount

	ToMonthlyBalanceCardNumber(card *db.GetMonthlyBalancesByCardNumberRow) *record.CardMonthBalance
	ToMonthlyBalancesCardNumber(cards []*db.GetMonthlyBalancesByCardNumberRow) []*record.CardMonthBalance

	ToYearlyBalanceCardNumber(card *db.GetYearlyBalancesByCardNumberRow) *record.CardYearlyBalance
	ToYearlyBalancesCardNumber(cards []*db.GetYearlyBalancesByCardNumberRow) []*record.CardYearlyBalance

	ToMonthlyTopupAmountByCardNumber(card *db.GetMonthlyTopupAmountByCardNumberRow) *record.CardMonthTopupAmount
	ToMonthlyTopupAmountsByCardNumber(cards []*db.GetMonthlyTopupAmountByCardNumberRow) []*record.CardMonthTopupAmount

	ToYearlyTopupAmountByCardNumber(card *db.GetYearlyTopupAmountByCardNumberRow) *record.CardYearlyTopupAmount
	ToYearlyTopupAmountsByCardNumber(cards []*db.GetYearlyTopupAmountByCardNumberRow) []*record.CardYearlyTopupAmount

	ToMonthlyWithdrawAmountByCardNumber(card *db.GetMonthlyWithdrawAmountByCardNumberRow) *record.CardMonthWithdrawAmount
	ToMonthlyWithdrawAmountsByCardNumber(cards []*db.GetMonthlyWithdrawAmountByCardNumberRow) []*record.CardMonthWithdrawAmount

	ToYearlyWithdrawAmountByCardNumber(card *db.GetYearlyWithdrawAmountByCardNumberRow) *record.CardYearlyWithdrawAmount
	ToYearlyWithdrawAmountsByCardNumber(cards []*db.GetYearlyWithdrawAmountByCardNumberRow) []*record.CardYearlyWithdrawAmount

	ToMonthlyTransactionAmountByCardNumber(card *db.GetMonthlyTransactionAmountByCardNumberRow) *record.CardMonthTransactionAmount
	ToMonthlyTransactionAmountsByCardNumber(cards []*db.GetMonthlyTransactionAmountByCardNumberRow) []*record.CardMonthTransactionAmount

	ToYearlyTransactionAmountByCardNumber(card *db.GetYearlyTransactionAmountByCardNumberRow) *record.CardYearlyTransactionAmount
	ToYearlyTransactionAmountsByCardNumber(cards []*db.GetYearlyTransactionAmountByCardNumberRow) []*record.CardYearlyTransactionAmount

	ToMonthlyTransferSenderAmountByCardNumber(card *db.GetMonthlyTransferAmountBySenderRow) *record.CardMonthTransferAmount
	ToMonthlyTransferSenderAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountBySenderRow) []*record.CardMonthTransferAmount

	ToYearlyTransferSenderAmountByCardNumber(card *db.GetYearlyTransferAmountBySenderRow) *record.CardYearlyTransferAmount
	ToYearlyTransferSenderAmountsByCardNumber(cards []*db.GetYearlyTransferAmountBySenderRow) []*record.CardYearlyTransferAmount

	ToMonthlyTransferReceiverAmountByCardNumber(card *db.GetMonthlyTransferAmountByReceiverRow) *record.CardMonthTransferAmount
	ToMonthlyTransferReceiverAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountByReceiverRow) []*record.CardMonthTransferAmount

	ToYearlyTransferReceiverAmountByCardNumber(card *db.GetYearlyTransferAmountByReceiverRow) *record.CardYearlyTransferAmount
	ToYearlyTransferReceiverAmountsByCardNumber(cards []*db.GetYearlyTransferAmountByReceiverRow) []*record.CardYearlyTransferAmount
}

type TransactionRecordMapping interface {
	ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord
	ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord

	ToTransactionRecordMonthStatusSuccess(s *db.GetMonthTransactionStatusSuccessRow) *record.TransactionRecordMonthStatusSuccess
	ToTransactionRecordsMonthStatusSuccess(Transactions []*db.GetMonthTransactionStatusSuccessRow) []*record.TransactionRecordMonthStatusSuccess
	ToTransactionRecordYearStatusSuccess(s *db.GetYearlyTransactionStatusSuccessRow) *record.TransactionRecordYearStatusSuccess
	ToTransactionRecordsYearStatusSuccess(Transactions []*db.GetYearlyTransactionStatusSuccessRow) []*record.TransactionRecordYearStatusSuccess

	ToTransactionRecordMonthStatusFailed(s *db.GetMonthTransactionStatusFailedRow) *record.TransactionRecordMonthStatusFailed
	ToTransactionRecordsMonthStatusFailed(Transactions []*db.GetMonthTransactionStatusFailedRow) []*record.TransactionRecordMonthStatusFailed
	ToTransactionRecordYearStatusFailed(s *db.GetYearlyTransactionStatusFailedRow) *record.TransactionRecordYearStatusFailed
	ToTransactionRecordsYearStatusFailed(Transactions []*db.GetYearlyTransactionStatusFailedRow) []*record.TransactionRecordYearStatusFailed

	ToTransactionMonthlyMethod(ss *db.GetMonthlyPaymentMethodsRow) *record.TransactionMonthMethod
	ToTransactionMonthlyMethods(ss []*db.GetMonthlyPaymentMethodsRow) []*record.TransactionMonthMethod
	ToTransactionYearlyMethod(ss *db.GetYearlyPaymentMethodsRow) *record.TransactionYearMethod
	ToTransactionYearlyMethods(ss []*db.GetYearlyPaymentMethodsRow) []*record.TransactionYearMethod

	ToTransactionMonthlyAmount(ss *db.GetMonthlyAmountsRow) *record.TransactionMonthAmount
	ToTransactionMonthlyAmounts(ss []*db.GetMonthlyAmountsRow) []*record.TransactionMonthAmount
	ToTransactionYearlyAmount(ss *db.GetYearlyAmountsRow) *record.TransactionYearlyAmount
	ToTransactionYearlyAmounts(ss []*db.GetYearlyAmountsRow) []*record.TransactionYearlyAmount

	ToTransactionMonthlyMethodByCardNumber(ss *db.GetMonthlyPaymentMethodsByCardNumberRow) *record.TransactionMonthMethod
	ToTransactionMonthlyMethodsByCardNumber(ss []*db.GetMonthlyPaymentMethodsByCardNumberRow) []*record.TransactionMonthMethod
	ToTransactionYearlyMethodByCardNumber(ss *db.GetYearlyPaymentMethodsByCardNumberRow) *record.TransactionYearMethod
	ToTransactionYearlyMethodsByCardNumber(ss []*db.GetYearlyPaymentMethodsByCardNumberRow) []*record.TransactionYearMethod

	ToTransactionMonthlyAmountByCardNumber(ss *db.GetMonthlyAmountsByCardNumberRow) *record.TransactionMonthAmount
	ToTransactionMonthlyAmountsByCardNumber(ss []*db.GetMonthlyAmountsByCardNumberRow) []*record.TransactionMonthAmount
	ToTransactionYearlyAmountByCardNumber(ss *db.GetYearlyAmountsByCardNumberRow) *record.TransactionYearlyAmount
	ToTransactionYearlyAmountsByCardNumber(ss []*db.GetYearlyAmountsByCardNumberRow) []*record.TransactionYearlyAmount

	ToTransactionRecordAll(transaction *db.GetTransactionsRow) *record.TransactionRecord
	ToTransactionsRecordAll(transactions []*db.GetTransactionsRow) []*record.TransactionRecord

	ToTransactionRecordActive(transaction *db.GetActiveTransactionsRow) *record.TransactionRecord
	ToTransactionsRecordActive(transactions []*db.GetActiveTransactionsRow) []*record.TransactionRecord
	ToTransactionRecordTrashed(transaction *db.GetTrashedTransactionsRow) *record.TransactionRecord
	ToTransactionsRecordTrashed(transactions []*db.GetTrashedTransactionsRow) []*record.TransactionRecord
}

type MerchantRecordMapping interface {
	ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord
	ToMerchantsRecord(merchants []*db.Merchant) []*record.MerchantRecord

	ToMerchantGetAllRecord(merchant *db.GetMerchantsRow) *record.MerchantRecord
	ToMerchantsGetAllRecord(merchants []*db.GetMerchantsRow) []*record.MerchantRecord

	ToMerchantMonthlyPaymentMethod(ms *db.GetMonthlyPaymentMethodsMerchantRow) *record.MerchantMonthlyPaymentMethod
	ToMerchantMonthlyPaymentMethods(ms []*db.GetMonthlyPaymentMethodsMerchantRow) []*record.MerchantMonthlyPaymentMethod
	ToMerchantYearlyPaymentMethod(ms *db.GetYearlyPaymentMethodMerchantRow) *record.MerchantYearlyPaymentMethod
	ToMerchantYearlyPaymentMethods(ms []*db.GetYearlyPaymentMethodMerchantRow) []*record.MerchantYearlyPaymentMethod

	ToMerchantMonthlyAmount(ms *db.GetMonthlyAmountMerchantRow) *record.MerchantMonthlyAmount
	ToMerchantMonthlyAmounts(ms []*db.GetMonthlyAmountMerchantRow) []*record.MerchantMonthlyAmount
	ToMerchantYearlyAmount(ms *db.GetYearlyAmountMerchantRow) *record.MerchantYearlyAmount
	ToMerchantYearlyAmounts(ms []*db.GetYearlyAmountMerchantRow) []*record.MerchantYearlyAmount

	ToMerchantMonthlyPaymentMethodByMerchant(ms *db.GetMonthlyPaymentMethodByMerchantsRow) *record.MerchantMonthlyPaymentMethod
	ToMerchantMonthlyPaymentMethodsByMerchant(ms []*db.GetMonthlyPaymentMethodByMerchantsRow) []*record.MerchantMonthlyPaymentMethod
	ToMerchantYearlyPaymentMethodByMerchant(ms *db.GetYearlyPaymentMethodByMerchantsRow) *record.MerchantYearlyPaymentMethod
	ToMerchantYearlyPaymentMethodsByMerchant(ms []*db.GetYearlyPaymentMethodByMerchantsRow) []*record.MerchantYearlyPaymentMethod

	ToMerchantMonthlyAmountByMerchant(ms *db.GetMonthlyAmountByMerchantsRow) *record.MerchantMonthlyAmount
	ToMerchantMonthlyAmountsByMerchant(ms []*db.GetMonthlyAmountByMerchantsRow) []*record.MerchantMonthlyAmount
	ToMerchantYearlyAmountByMerchant(ms *db.GetYearlyAmountByMerchantsRow) *record.MerchantYearlyAmount
	ToMerchantYearlyAmountsMerchant(ms []*db.GetYearlyAmountByMerchantsRow) []*record.MerchantYearlyAmount

	ToMerchantActiveRecord(merchant *db.GetActiveMerchantsRow) *record.MerchantRecord
	ToMerchantsActiveRecord(merchants []*db.GetActiveMerchantsRow) []*record.MerchantRecord
	ToMerchantTrashedRecord(merchant *db.GetTrashedMerchantsRow) *record.MerchantRecord
	ToMerchantsTrashedRecord(merchants []*db.GetTrashedMerchantsRow) []*record.MerchantRecord
}
