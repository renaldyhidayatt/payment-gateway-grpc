package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
)

type UserRecordMapping interface {
	ToUserRecord(user *db.User) *record.UserRecord
	ToUsersRecord(users []*db.User) []*record.UserRecord
}

type RoleRecordMapping interface {
	ToRoleRecord(role *db.Role) *record.RoleRecord
	ToRolesRecord(roles []*db.Role) []*record.RoleRecord
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
}

type TopupRecordMapping interface {
	ToTopupRecord(topup *db.Topup) *record.TopupRecord
	ToTopupRecords(topups []*db.Topup) []*record.TopupRecord
}

type TransferRecordMapping interface {
	ToTransferRecord(transfer *db.Transfer) *record.TransferRecord
	ToTransfersRecord(transfers []*db.Transfer) []*record.TransferRecord
}

type WithdrawRecordMapping interface {
	ToWithdrawRecord(withdraw *db.Withdraw) *record.WithdrawRecord
	ToWithdrawsRecord(withdraws []*db.Withdraw) []*record.WithdrawRecord
}

type CardRecordMapping interface {
	ToCardRecord(card *db.Card) *record.CardRecord
	ToCardsRecord(cards []*db.Card) []*record.CardRecord
}

type TransactionRecordMapping interface {
	ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord
	ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord
}

type MerchantRecordMapping interface {
	ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord
	ToMerchantsRecord(merchants []*db.Merchant) []*record.MerchantRecord
}
