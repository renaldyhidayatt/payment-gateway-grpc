// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
	"time"
)

type Querier interface {
	AssignRoleToUser(ctx context.Context, arg AssignRoleToUserParams) (*UserRole, error)
	CountActiveRoles(ctx context.Context, dollar_1 sql.NullString) (int64, error)
	// Count All Active Users
	CountActiveUsers(ctx context.Context) (int64, error)
	// Count Active Withdraws by Date
	CountActiveWithdrawsByDate(ctx context.Context, withdrawTime time.Time) (int64, error)
	CountAllActiveRoles(ctx context.Context) (int64, error)
	CountAllRoles(ctx context.Context) (int64, error)
	CountAllSaldos(ctx context.Context) (int64, error)
	// Count All Topups
	CountAllTopups(ctx context.Context) (int64, error)
	// Count All Transactions
	CountAllTransactions(ctx context.Context) (int64, error)
	// Count All Transfers
	CountAllTransfers(ctx context.Context) (int64, error)
	CountAllTrashedRoles(ctx context.Context) (int64, error)
	CountAllWithdraws(ctx context.Context) (int64, error)
	CountRoles(ctx context.Context, dollar_1 string) (int64, error)
	CountSaldos(ctx context.Context, dollar_1 string) (int64, error)
	CountTopups(ctx context.Context, dollar_1 string) (int64, error)
	// Count Topups by Date
	CountTopupsByDate(ctx context.Context, dollar_1 time.Time) (int64, error)
	CountTransactions(ctx context.Context, dollar_1 string) (int64, error)
	// Count Transactions by Date
	CountTransactionsByDate(ctx context.Context, dollar_1 time.Time) (int64, error)
	CountTransfers(ctx context.Context, dollar_1 string) (int64, error)
	// Count Transfers by Date
	CountTransfersByDate(ctx context.Context, dollar_1 time.Time) (int64, error)
	CountTrashedRoles(ctx context.Context, dollar_1 sql.NullString) (int64, error)
	CountWithdraws(ctx context.Context, dollar_1 string) (int64, error)
	// Create Card
	CreateCard(ctx context.Context, arg CreateCardParams) (*Card, error)
	// Create Merchant
	CreateMerchant(ctx context.Context, arg CreateMerchantParams) (*Merchant, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (*RefreshToken, error)
	CreateRole(ctx context.Context, roleName string) (*Role, error)
	// Create Saldo
	CreateSaldo(ctx context.Context, arg CreateSaldoParams) (*Saldo, error)
	// Create Topup
	CreateTopup(ctx context.Context, arg CreateTopupParams) (*Topup, error)
	// Create Transaction
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (*Transaction, error)
	// Create Transfer
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (*Transfer, error)
	// Create User
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	// Create Withdraw
	CreateWithdraw(ctx context.Context, arg CreateWithdrawParams) (*Withdraw, error)
	// Delete Card Permanently
	DeleteCardPermanently(ctx context.Context, cardID int32) error
	// Delete Merchant Permanently
	DeleteMerchantPermanently(ctx context.Context, merchantID int32) error
	DeletePermanentRole(ctx context.Context, roleID int32) error
	DeleteRefreshToken(ctx context.Context, token string) error
	DeleteRefreshTokenByUserId(ctx context.Context, userID int32) error
	// Delete Saldo Permanently
	DeleteSaldoPermanently(ctx context.Context, saldoID int32) error
	// Delete Topup Permanently
	DeleteTopupPermanently(ctx context.Context, topupID int32) error
	// Delete Transaction Permanently
	DeleteTransactionPermanently(ctx context.Context, transactionID int32) error
	// Delete Transfer Permanently
	DeleteTransferPermanently(ctx context.Context, transferID int32) error
	// Delete User Permanently
	DeleteUserPermanently(ctx context.Context, userID int32) error
	// Delete Withdraw Permanently
	DeleteWithdrawPermanently(ctx context.Context, withdrawID int32) error
	FindAllTransactionsByMerchantID(ctx context.Context, merchantID int32) ([]*FindAllTransactionsByMerchantIDRow, error)
	FindAllTransfersByCardNumberAsReceiver(ctx context.Context, transferTo string) ([]*Transfer, error)
	FindAllTransfersByCardNumberAsSender(ctx context.Context, transferFrom string) ([]*Transfer, error)
	FindAllWithdrawsByCardNumber(ctx context.Context, cardNumber string) ([]*Withdraw, error)
	FindRefreshTokenByToken(ctx context.Context, token string) (*RefreshToken, error)
	FindRefreshTokenByUserId(ctx context.Context, userID int32) (*RefreshToken, error)
	GetActiveCardsWithCount(ctx context.Context, arg GetActiveCardsWithCountParams) ([]*GetActiveCardsWithCountRow, error)
	GetActiveMerchants(ctx context.Context, arg GetActiveMerchantsParams) ([]*GetActiveMerchantsRow, error)
	// Get All Active Roles
	GetActiveRoles(ctx context.Context, arg GetActiveRolesParams) ([]*GetActiveRolesRow, error)
	// Get All Active Saldos with Pagination, Search, and Total Count
	GetActiveSaldos(ctx context.Context, arg GetActiveSaldosParams) ([]*GetActiveSaldosRow, error)
	// Get All Active Topups with Pagination and Search
	GetActiveTopups(ctx context.Context, arg GetActiveTopupsParams) ([]*GetActiveTopupsRow, error)
	// Get Active Transactions with Pagination, Search, and Count
	GetActiveTransactions(ctx context.Context, arg GetActiveTransactionsParams) ([]*GetActiveTransactionsRow, error)
	// Get Active Transfers with Search, Pagination, and Total Count
	GetActiveTransfers(ctx context.Context, arg GetActiveTransfersParams) ([]*GetActiveTransfersRow, error)
	// Get Active Users with Pagination and Total Count
	GetActiveUsersWithPagination(ctx context.Context, arg GetActiveUsersWithPaginationParams) ([]*GetActiveUsersWithPaginationRow, error)
	// Get Active Withdraws with Search, Pagination, and Total Count
	GetActiveWithdraws(ctx context.Context, arg GetActiveWithdrawsParams) ([]*GetActiveWithdrawsRow, error)
	GetAllBalances(ctx context.Context) ([]*GetAllBalancesRow, error)
	// Get Card by Card Number
	GetCardByCardNumber(ctx context.Context, cardNumber string) (*Card, error)
	// Get Card by ID
	GetCardByID(ctx context.Context, cardID int32) (*Card, error)
	// Get a single Card by User ID
	GetCardByUserID(ctx context.Context, userID int32) (*Card, error)
	// Search Cards with Pagination and Total Count
	GetCards(ctx context.Context, arg GetCardsParams) ([]*GetCardsRow, error)
	// Get Merchant by API Key
	GetMerchantByApiKey(ctx context.Context, apiKey string) (*Merchant, error)
	// Get Merchant by ID
	GetMerchantByID(ctx context.Context, merchantID int32) (*Merchant, error)
	// Get Merchant by Name
	GetMerchantByName(ctx context.Context, name string) (*Merchant, error)
	GetMerchants(ctx context.Context, arg GetMerchantsParams) ([]*GetMerchantsRow, error)
	// Get Merchants by User ID
	GetMerchantsByUserID(ctx context.Context, userID int32) ([]*Merchant, error)
	GetMonthlyAmountMerchant(ctx context.Context, transactionTime time.Time) ([]*GetMonthlyAmountMerchantRow, error)
	GetMonthlyAmounts(ctx context.Context, transactionTime time.Time) ([]*GetMonthlyAmountsRow, error)
	GetMonthlyAmountsByMerchant(ctx context.Context, arg GetMonthlyAmountsByMerchantParams) ([]*GetMonthlyAmountsByMerchantRow, error)
	GetMonthlyBalances(ctx context.Context, createdAt sql.NullTime) ([]*GetMonthlyBalancesRow, error)
	GetMonthlyPaymentMethods(ctx context.Context, transactionTime time.Time) ([]*GetMonthlyPaymentMethodsRow, error)
	GetMonthlyPaymentMethodsMerchant(ctx context.Context, transactionTime time.Time) ([]*GetMonthlyPaymentMethodsMerchantRow, error)
	GetMonthlyTopupAmounts(ctx context.Context, topupTime time.Time) ([]*GetMonthlyTopupAmountsRow, error)
	GetMonthlyTopupMethods(ctx context.Context, topupTime time.Time) ([]*GetMonthlyTopupMethodsRow, error)
	GetMonthlyTotalBalance(ctx context.Context, createdAt sql.NullTime) ([]*GetMonthlyTotalBalanceRow, error)
	GetMonthlyTransferAmounts(ctx context.Context, transferTime time.Time) ([]*GetMonthlyTransferAmountsRow, error)
	GetMonthlyWithdrawsAll(ctx context.Context, withdrawTime time.Time) ([]*GetMonthlyWithdrawsAllRow, error)
	GetMonthlyWithdrawsByCardNumber(ctx context.Context, arg GetMonthlyWithdrawsByCardNumberParams) ([]*GetMonthlyWithdrawsByCardNumberRow, error)
	GetRole(ctx context.Context, roleID int32) (*Role, error)
	GetRoleByName(ctx context.Context, roleName string) (*Role, error)
	GetRoles(ctx context.Context, arg GetRolesParams) ([]*GetRolesRow, error)
	// Get Saldo by Card Number
	GetSaldoByCardNumber(ctx context.Context, cardNumber string) (*Saldo, error)
	// Get Saldo by ID
	GetSaldoByID(ctx context.Context, saldoID int32) (*Saldo, error)
	// Search Saldos with Pagination and Total Count
	GetSaldos(ctx context.Context, arg GetSaldosParams) ([]*GetSaldosRow, error)
	// Get Topup by ID
	GetTopupByID(ctx context.Context, topupID int32) (*Topup, error)
	// Search Topups with Pagination
	GetTopups(ctx context.Context, arg GetTopupsParams) ([]*GetTopupsRow, error)
	// Get Topups by Card Number
	GetTopupsByCardNumber(ctx context.Context, cardNumber string) ([]*Topup, error)
	// Get Transaction by ID
	GetTransactionByID(ctx context.Context, transactionID int32) (*Transaction, error)
	// Search Transactions with Pagination
	GetTransactions(ctx context.Context, arg GetTransactionsParams) ([]*GetTransactionsRow, error)
	// Get Transactions by Card Number
	GetTransactionsByCardNumber(ctx context.Context, cardNumber string) ([]*Transaction, error)
	// Get Transactions by Merchant ID
	GetTransactionsByMerchantID(ctx context.Context, merchantID int32) ([]*Transaction, error)
	// Get Transfer by ID
	GetTransferByID(ctx context.Context, transferID int32) (*Transfer, error)
	// Search Transfers with Pagination
	GetTransfers(ctx context.Context, arg GetTransfersParams) ([]*GetTransfersRow, error)
	// Get Transfers by Card Number (Source or Destination)
	GetTransfersByCardNumber(ctx context.Context, transferFrom string) ([]*Transfer, error)
	// Get Transfers by Destination Card
	GetTransfersByDestinationCard(ctx context.Context, transferTo string) ([]*Transfer, error)
	// Get Transfers by Source Card
	GetTransfersBySourceCard(ctx context.Context, transferFrom string) ([]*Transfer, error)
	// Get Trashed By Card ID
	GetTrashedCardByID(ctx context.Context, cardID int32) (*Card, error)
	GetTrashedCardsWithCount(ctx context.Context, arg GetTrashedCardsWithCountParams) ([]*GetTrashedCardsWithCountRow, error)
	// Get Trashed By Merchant ID
	GetTrashedMerchantByID(ctx context.Context, merchantID int32) (*Merchant, error)
	GetTrashedMerchants(ctx context.Context, arg GetTrashedMerchantsParams) ([]*GetTrashedMerchantsRow, error)
	// Get All Trashed Roles
	GetTrashedRoles(ctx context.Context, arg GetTrashedRolesParams) ([]*GetTrashedRolesRow, error)
	// Get Trashed By Saldo ID
	GetTrashedSaldoByID(ctx context.Context, saldoID int32) (*Saldo, error)
	// Get Trashed Saldos with Pagination, Search, and Total Count
	GetTrashedSaldos(ctx context.Context, arg GetTrashedSaldosParams) ([]*GetTrashedSaldosRow, error)
	// Get Trashed By Topup ID
	GetTrashedTopupByID(ctx context.Context, topupID int32) (*Topup, error)
	// Get Trashed Topups with Pagination and Search
	GetTrashedTopups(ctx context.Context, arg GetTrashedTopupsParams) ([]*GetTrashedTopupsRow, error)
	// Get Trashed By Transaction ID
	GetTrashedTransactionByID(ctx context.Context, transactionID int32) (*Transaction, error)
	// Get Trashed Transactions with Pagination, Search, and Count
	GetTrashedTransactions(ctx context.Context, arg GetTrashedTransactionsParams) ([]*GetTrashedTransactionsRow, error)
	// Get Trashed By Transfer ID
	GetTrashedTransferByID(ctx context.Context, transferID int32) (*Transfer, error)
	// Get Trashed Transfers with Search, Pagination, and Total Count
	GetTrashedTransfers(ctx context.Context, arg GetTrashedTransfersParams) ([]*GetTrashedTransfersRow, error)
	// Get Trashed By User ID
	GetTrashedUserByID(ctx context.Context, userID int32) (*User, error)
	GetTrashedUserRoles(ctx context.Context, userID int32) ([]*GetTrashedUserRolesRow, error)
	// Get Trashed Users with Pagination and Total Count
	GetTrashedUsersWithPagination(ctx context.Context, arg GetTrashedUsersWithPaginationParams) ([]*GetTrashedUsersWithPaginationRow, error)
	// Get Trashed By Withdraw ID
	GetTrashedWithdrawByID(ctx context.Context, withdrawID int32) (*Withdraw, error)
	// Get Trashed Withdraws with Search, Pagination, and Total Count
	GetTrashedWithdraws(ctx context.Context, arg GetTrashedWithdrawsParams) ([]*GetTrashedWithdrawsRow, error)
	// Get User by Email
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	// Get User by ID
	GetUserByID(ctx context.Context, userID int32) (*User, error)
	GetUserRoles(ctx context.Context, userID int32) ([]*Role, error)
	// Search Users with Pagination and Total Count
	GetUsersWithPagination(ctx context.Context, arg GetUsersWithPaginationParams) ([]*GetUsersWithPaginationRow, error)
	// Get Withdraw by ID
	GetWithdrawByID(ctx context.Context, withdrawID int32) (*Withdraw, error)
	// Search Withdraws with Pagination
	GetWithdraws(ctx context.Context, arg GetWithdrawsParams) ([]*GetWithdrawsRow, error)
	GetYearlyAmountMerchant(ctx context.Context) ([]*GetYearlyAmountMerchantRow, error)
	GetYearlyAmounts(ctx context.Context) ([]*GetYearlyAmountsRow, error)
	GetYearlyAmountsByMerchant(ctx context.Context, merchantID int32) ([]*GetYearlyAmountsByMerchantRow, error)
	GetYearlyBalances(ctx context.Context) ([]*GetYearlyBalancesRow, error)
	GetYearlyPaymentMethodMerchant(ctx context.Context) ([]*GetYearlyPaymentMethodMerchantRow, error)
	GetYearlyPaymentMethods(ctx context.Context) ([]*GetYearlyPaymentMethodsRow, error)
	GetYearlyTopupAmounts(ctx context.Context) ([]*GetYearlyTopupAmountsRow, error)
	GetYearlyTopupMethods(ctx context.Context) ([]*GetYearlyTopupMethodsRow, error)
	GetYearlyTotalBalance(ctx context.Context) ([]*GetYearlyTotalBalanceRow, error)
	GetYearlyTransferAmounts(ctx context.Context) ([]*GetYearlyTransferAmountsRow, error)
	GetYearlyWithdrawsAll(ctx context.Context) ([]*GetYearlyWithdrawsAllRow, error)
	RemoveRoleFromUser(ctx context.Context, arg RemoveRoleFromUserParams) error
	// Restore Trashed Card
	RestoreCard(ctx context.Context, cardID int32) error
	// Restore Trashed Merchant
	RestoreMerchant(ctx context.Context, merchantID int32) error
	RestoreRole(ctx context.Context, roleID int32) error
	// Restore Trashed Saldo
	RestoreSaldo(ctx context.Context, saldoID int32) error
	// Restore Trashed Topup
	RestoreTopup(ctx context.Context, topupID int32) error
	// Restore Trashed Transaction
	RestoreTransaction(ctx context.Context, transactionID int32) error
	// Restore Trashed Transfer
	RestoreTransfer(ctx context.Context, transferID int32) error
	// Restore Trashed User
	RestoreUser(ctx context.Context, userID int32) error
	RestoreUserRole(ctx context.Context, userRoleID int32) error
	// Restore Withdraw (Undelete)
	RestoreWithdraw(ctx context.Context, withdrawID int32) error
	// Search Users by Email
	SearchUsersByEmail(ctx context.Context, dollar_1 sql.NullString) ([]*User, error)
	// Search Withdraw by Card Number
	SearchWithdrawByCardNumber(ctx context.Context, dollar_1 sql.NullString) ([]*Withdraw, error)
	Topup_CountAll(ctx context.Context) (int64, error)
	Transaction_CountAll(ctx context.Context) (int64, error)
	Transfer_CountAll(ctx context.Context) (int64, error)
	// Trash Card
	TrashCard(ctx context.Context, cardID int32) error
	// Trash Merchant
	TrashMerchant(ctx context.Context, merchantID int32) error
	TrashRole(ctx context.Context, roleID int32) error
	// Trash Saldo
	TrashSaldo(ctx context.Context, saldoID int32) error
	// Trash Topup
	TrashTopup(ctx context.Context, topupID int32) error
	// Trash Transaction
	TrashTransaction(ctx context.Context, transactionID int32) error
	// Trash Transfer
	TrashTransfer(ctx context.Context, transferID int32) error
	// Trash User
	TrashUser(ctx context.Context, userID int32) error
	TrashUserRole(ctx context.Context, userRoleID int32) error
	// Trash Withdraw (Soft Delete)
	TrashWithdraw(ctx context.Context, withdrawID int32) error
	// Update Card
	UpdateCard(ctx context.Context, arg UpdateCardParams) error
	// Update Merchant
	UpdateMerchant(ctx context.Context, arg UpdateMerchantParams) error
	UpdateRefreshTokenByUserId(ctx context.Context, arg UpdateRefreshTokenByUserIdParams) error
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (*Role, error)
	// Update Saldo
	UpdateSaldo(ctx context.Context, arg UpdateSaldoParams) error
	// Update Saldo Balance
	UpdateSaldoBalance(ctx context.Context, arg UpdateSaldoBalanceParams) error
	// Update Saldo Withdraw
	UpdateSaldoWithdraw(ctx context.Context, arg UpdateSaldoWithdrawParams) error
	// Update Topup
	UpdateTopup(ctx context.Context, arg UpdateTopupParams) error
	// Update Topup Amount
	UpdateTopupAmount(ctx context.Context, arg UpdateTopupAmountParams) error
	// Update Transaction
	UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) error
	// Update Transfer
	UpdateTransfer(ctx context.Context, arg UpdateTransferParams) error
	// Update Transfer Amount
	UpdateTransferAmount(ctx context.Context, arg UpdateTransferAmountParams) error
	// Update User
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	// Update Withdraw
	UpdateWithdraw(ctx context.Context, arg UpdateWithdrawParams) error
}

var _ Querier = (*Queries)(nil)
