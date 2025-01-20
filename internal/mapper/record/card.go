package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
)

type cardRecordMapper struct {
}

func NewCardRecordMapper() *cardRecordMapper {
	return &cardRecordMapper{}
}

func (s *cardRecordMapper) ToCardRecord(card *db.Card) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formatedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.CardRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

func (s *cardRecordMapper) ToCardGetAll(card *db.GetCardsRow) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formatedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.CardRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

func (s *cardRecordMapper) ToCardsRecord(cards []*db.GetCardsRow) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardGetAll(card))
	}
	return records
}

func (s *cardRecordMapper) ToCardRecordActive(card *db.GetActiveCardsWithCountRow) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formattedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formattedDeletedAt
	}

	return &record.CardRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

func (s *cardRecordMapper) ToCardRecordsActive(cards []*db.GetActiveCardsWithCountRow) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardRecordActive(card))
	}
	return records
}

func (s *cardRecordMapper) ToCardRecordTrashed(card *db.GetTrashedCardsWithCountRow) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formattedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formattedDeletedAt
	}

	return &record.CardRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

func (s *cardRecordMapper) ToCardRecordsTrashed(cards []*db.GetTrashedCardsWithCountRow) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardRecordTrashed(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyBalance(card *db.GetMonthlyBalancesRow) *record.CardMonthBalance {
	return &record.CardMonthBalance{
		Month:        card.Month,
		TotalBalance: int64(card.TotalBalance),
	}
}

func (s *cardRecordMapper) ToMonthlyBalances(cards []*db.GetMonthlyBalancesRow) []*record.CardMonthBalance {
	var records []*record.CardMonthBalance
	for _, card := range cards {
		records = append(records, s.ToMonthlyBalance(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyBalance(card *db.GetYearlyBalancesRow) *record.CardYearlyBalance {
	return &record.CardYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

func (s *cardRecordMapper) ToYearlyBalances(cards []*db.GetYearlyBalancesRow) []*record.CardYearlyBalance {
	var records []*record.CardYearlyBalance
	for _, card := range cards {
		records = append(records, s.ToYearlyBalance(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyTopupAmount(card *db.GetMonthlyTopupAmountRow) *record.CardMonthTopupAmount {
	return &record.CardMonthTopupAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTopupAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyTopupAmounts(cards []*db.GetMonthlyTopupAmountRow) []*record.CardMonthTopupAmount {
	var records []*record.CardMonthTopupAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTopupAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyTopupAmount(card *db.GetYearlyTopupAmountRow) *record.CardYearlyTopupAmount {
	return &record.CardYearlyTopupAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTopupAmount,
	}
}

func (s *cardRecordMapper) ToYearlyTopupAmounts(cards []*db.GetYearlyTopupAmountRow) []*record.CardYearlyTopupAmount {
	var records []*record.CardYearlyTopupAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTopupAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyWithdrawAmount(card *db.GetMonthlyWithdrawAmountRow) *record.CardMonthWithdrawAmount {
	return &record.CardMonthWithdrawAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalWithdrawAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyWithdrawAmounts(cards []*db.GetMonthlyWithdrawAmountRow) []*record.CardMonthWithdrawAmount {
	var records []*record.CardMonthWithdrawAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyWithdrawAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyWithdrawAmount(card *db.GetYearlyWithdrawAmountRow) *record.CardYearlyWithdrawAmount {
	return &record.CardYearlyWithdrawAmount{
		Year:        card.Year,
		TotalAmount: card.TotalWithdrawAmount,
	}
}

func (s *cardRecordMapper) ToYearlyWithdrawAmounts(cards []*db.GetYearlyWithdrawAmountRow) []*record.CardYearlyWithdrawAmount {
	var records []*record.CardYearlyWithdrawAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyWithdrawAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyTransactionAmount(card *db.GetMonthlyTransactionAmountRow) *record.CardMonthTransactionAmount {
	return &record.CardMonthTransactionAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTransactionAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyTransactionAmounts(cards []*db.GetMonthlyTransactionAmountRow) []*record.CardMonthTransactionAmount {
	var records []*record.CardMonthTransactionAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransactionAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyTransactionAmount(card *db.GetYearlyTransactionAmountRow) *record.CardYearlyTransactionAmount {
	return &record.CardYearlyTransactionAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTransactionAmount,
	}
}

func (s *cardRecordMapper) ToYearlyTransactionAmounts(cards []*db.GetYearlyTransactionAmountRow) []*record.CardYearlyTransactionAmount {
	var records []*record.CardYearlyTransactionAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransactionAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyTransferSenderAmount(card *db.GetMonthlyTransferAmountSenderRow) *record.CardMonthTransferAmount {
	return &record.CardMonthTransferAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalSentAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyTransferSenderAmounts(cards []*db.GetMonthlyTransferAmountSenderRow) []*record.CardMonthTransferAmount {
	var records []*record.CardMonthTransferAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferSenderAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyTransferSenderAmount(card *db.GetYearlyTransferAmountSenderRow) *record.CardYearlyTransferAmount {
	return &record.CardYearlyTransferAmount{
		Year:        card.Year,
		TotalAmount: card.TotalSentAmount,
	}
}

func (s *cardRecordMapper) ToYearlyTransferSenderAmounts(cards []*db.GetYearlyTransferAmountSenderRow) []*record.CardYearlyTransferAmount {
	var records []*record.CardYearlyTransferAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferSenderAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyTransferReceiverAmount(card *db.GetMonthlyTransferAmountReceiverRow) *record.CardMonthTransferAmount {
	return &record.CardMonthTransferAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalReceivedAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyTransferReceiverAmounts(cards []*db.GetMonthlyTransferAmountReceiverRow) []*record.CardMonthTransferAmount {
	var records []*record.CardMonthTransferAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferReceiverAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyTransferReceiverAmount(card *db.GetYearlyTransferAmountReceiverRow) *record.CardYearlyTransferAmount {
	return &record.CardYearlyTransferAmount{
		Year:        card.Year,
		TotalAmount: card.TotalReceivedAmount,
	}
}

func (s *cardRecordMapper) ToYearlyTransferReceiverAmounts(cards []*db.GetYearlyTransferAmountReceiverRow) []*record.CardYearlyTransferAmount {
	var records []*record.CardYearlyTransferAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferReceiverAmount(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyBalanceCardNumber(card *db.GetMonthlyBalancesByCardNumberRow) *record.CardMonthBalance {
	return &record.CardMonthBalance{
		Month:        card.Month,
		TotalBalance: int64(card.TotalBalance),
	}
}

func (s *cardRecordMapper) ToMonthlyBalancesCardNumber(cards []*db.GetMonthlyBalancesByCardNumberRow) []*record.CardMonthBalance {
	var records []*record.CardMonthBalance
	for _, card := range cards {
		records = append(records, s.ToMonthlyBalanceCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyBalanceCardNumber(card *db.GetYearlyBalancesByCardNumberRow) *record.CardYearlyBalance {
	return &record.CardYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

func (s *cardRecordMapper) ToYearlyBalancesCardNumber(cards []*db.GetYearlyBalancesByCardNumberRow) []*record.CardYearlyBalance {
	var records []*record.CardYearlyBalance
	for _, card := range cards {
		records = append(records, s.ToYearlyBalanceCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyTopupAmountByCardNumber(card *db.GetMonthlyTopupAmountByCardNumberRow) *record.CardMonthTopupAmount {
	return &record.CardMonthTopupAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTopupAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyTopupAmountsByCardNumber(cards []*db.GetMonthlyTopupAmountByCardNumberRow) []*record.CardMonthTopupAmount {
	var records []*record.CardMonthTopupAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTopupAmountByCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyTopupAmountByCardNumber(card *db.GetYearlyTopupAmountByCardNumberRow) *record.CardYearlyTopupAmount {
	return &record.CardYearlyTopupAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTopupAmount,
	}
}

func (s *cardRecordMapper) ToYearlyTopupAmountsByCardNumber(cards []*db.GetYearlyTopupAmountByCardNumberRow) []*record.CardYearlyTopupAmount {
	var records []*record.CardYearlyTopupAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTopupAmountByCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyWithdrawAmountByCardNumber(card *db.GetMonthlyWithdrawAmountByCardNumberRow) *record.CardMonthWithdrawAmount {
	return &record.CardMonthWithdrawAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalWithdrawAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyWithdrawAmountsByCardNumber(cards []*db.GetMonthlyWithdrawAmountByCardNumberRow) []*record.CardMonthWithdrawAmount {
	var records []*record.CardMonthWithdrawAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyWithdrawAmountByCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyWithdrawAmountByCardNumber(card *db.GetYearlyWithdrawAmountByCardNumberRow) *record.CardYearlyWithdrawAmount {
	return &record.CardYearlyWithdrawAmount{
		Year:        card.Year,
		TotalAmount: card.TotalWithdrawAmount,
	}
}

func (s *cardRecordMapper) ToYearlyWithdrawAmountsByCardNumber(cards []*db.GetYearlyWithdrawAmountByCardNumberRow) []*record.CardYearlyWithdrawAmount {
	var records []*record.CardYearlyWithdrawAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyWithdrawAmountByCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyTransactionAmountByCardNumber(card *db.GetMonthlyTransactionAmountByCardNumberRow) *record.CardMonthTransactionAmount {
	return &record.CardMonthTransactionAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTransactionAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyTransactionAmountsByCardNumber(cards []*db.GetMonthlyTransactionAmountByCardNumberRow) []*record.CardMonthTransactionAmount {
	var records []*record.CardMonthTransactionAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransactionAmountByCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyTransactionAmountByCardNumber(card *db.GetYearlyTransactionAmountByCardNumberRow) *record.CardYearlyTransactionAmount {
	return &record.CardYearlyTransactionAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTransactionAmount,
	}
}

func (s *cardRecordMapper) ToYearlyTransactionAmountsByCardNumber(cards []*db.GetYearlyTransactionAmountByCardNumberRow) []*record.CardYearlyTransactionAmount {
	var records []*record.CardYearlyTransactionAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransactionAmountByCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyTransferSenderAmountByCardNumber(card *db.GetMonthlyTransferAmountBySenderRow) *record.CardMonthTransferAmount {
	return &record.CardMonthTransferAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalSentAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyTransferSenderAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountBySenderRow) []*record.CardMonthTransferAmount {
	var records []*record.CardMonthTransferAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferSenderAmountByCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyTransferSenderAmountByCardNumber(card *db.GetYearlyTransferAmountBySenderRow) *record.CardYearlyTransferAmount {
	return &record.CardYearlyTransferAmount{
		Year:        card.Year,
		TotalAmount: card.TotalSentAmount,
	}
}

func (s *cardRecordMapper) ToYearlyTransferSenderAmountsByCardNumber(cards []*db.GetYearlyTransferAmountBySenderRow) []*record.CardYearlyTransferAmount {
	var records []*record.CardYearlyTransferAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferSenderAmountByCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToMonthlyTransferReceiverAmountByCardNumber(card *db.GetMonthlyTransferAmountByReceiverRow) *record.CardMonthTransferAmount {
	return &record.CardMonthTransferAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalReceivedAmount),
	}
}

func (s *cardRecordMapper) ToMonthlyTransferReceiverAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountByReceiverRow) []*record.CardMonthTransferAmount {
	var records []*record.CardMonthTransferAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferReceiverAmountByCardNumber(card))
	}
	return records
}

func (s *cardRecordMapper) ToYearlyTransferReceiverAmountByCardNumber(card *db.GetYearlyTransferAmountByReceiverRow) *record.CardYearlyTransferAmount {
	return &record.CardYearlyTransferAmount{
		Year:        card.Year,
		TotalAmount: card.TotalReceivedAmount,
	}
}

func (s *cardRecordMapper) ToYearlyTransferReceiverAmountsByCardNumber(cards []*db.GetYearlyTransferAmountByReceiverRow) []*record.CardYearlyTransferAmount {
	var records []*record.CardYearlyTransferAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferReceiverAmountByCardNumber(card))
	}
	return records
}
