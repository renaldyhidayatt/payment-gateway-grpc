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
}

type TransactionRecordMapping interface {
	ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord
	ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord

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

	ToMerchantActiveRecord(merchant *db.GetActiveMerchantsRow) *record.MerchantRecord
	ToMerchantsActiveRecord(merchants []*db.GetActiveMerchantsRow) []*record.MerchantRecord
	ToMerchantTrashedRecord(merchant *db.GetTrashedMerchantsRow) *record.MerchantRecord
	ToMerchantsTrashedRecord(merchants []*db.GetTrashedMerchantsRow) []*record.MerchantRecord
}
