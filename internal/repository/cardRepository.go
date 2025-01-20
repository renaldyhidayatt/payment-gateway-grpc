package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"MamangRust/paymentgatewaygrpc/pkg/randomvcc"
	"context"
	"fmt"
	"time"
)

type cardRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.CardRecordMapping
}

func NewCardRepository(db *db.Queries, ctx context.Context, mapping recordmapper.CardRecordMapping) *cardRepository {
	return &cardRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *cardRepository) FindAllCards(search string, page, pageSize int) ([]*record.CardRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCardsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	cards, err := r.db.GetCards(r.ctx, req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find cards: %w", err)
	}

	var totalCount int
	if len(cards) > 0 {
		totalCount = int(cards[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCardsRecord(cards), totalCount, nil
}

func (r *cardRepository) FindById(card_id int) (*record.CardRecord, error) {
	res, err := r.db.GetCardByID(r.ctx, int32(card_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find card: %w", err)
	}

	return r.mapping.ToCardRecord(res), nil
}

func (r *cardRepository) FindCardByUserId(user_id int) (*record.CardRecord, error) {
	res, err := r.db.GetCardByUserID(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed no found card in user_id :%w, ", err)
	}

	return r.mapping.ToCardRecord(res), nil
}

func (r *cardRepository) FindCardByCardNumber(card_number string) (*record.CardRecord, error) {
	res, err := r.db.GetCardByCardNumber(r.ctx, card_number)

	if err != nil {
		return nil, fmt.Errorf("failed to not found card in card number :%w", err)
	}

	return r.mapping.ToCardRecord(res), nil
}

func (r *cardRepository) GetTotalBalances() (*int64, error) {
	res, err := r.db.GetTotalBalance(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get total balance: %w", err)
	}

	return &res, nil
}

func (r *cardRepository) GetTotalTopAmount() (*int64, error) {
	res, err := r.db.GetTotalTopupAmount(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get total top-up amount: %w", err)
	}
	return &res, nil
}

func (r *cardRepository) GetTotalWithdrawAmount() (*int64, error) {
	res, err := r.db.GetTotalWithdrawAmount(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get total withdrawal amount: %w", err)
	}
	return &res, nil
}

func (r *cardRepository) GetTotalTransactionAmount() (*int64, error) {
	res, err := r.db.GetTotalTransactionAmount(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get total transaction amount: %w", err)
	}
	return &res, nil
}

func (r *cardRepository) GetTotalTransferAmount() (*int64, error) {
	res, err := r.db.GetTotalTransferAmount(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get total transfer amount: %w", err)
	}
	return &res, nil
}

func (r *cardRepository) GetTotalBalanceByCardNumber(cardNumber string) (*int64, error) {
	res, err := r.db.GetTotalBalanceByCardNumber(r.ctx, cardNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get total balance for card %s: %w", cardNumber, err)
	}
	return &res, nil
}

func (r *cardRepository) GetTotalTopupAmountByCardNumber(cardNumber string) (*int64, error) {
	res, err := r.db.GetTotalTopupAmountByCardNumber(r.ctx, cardNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get total top-up amount for card %s: %w", cardNumber, err)
	}
	return &res, nil
}

func (r *cardRepository) GetTotalWithdrawAmountByCardNumber(cardNumber string) (*int64, error) {
	res, err := r.db.GetTotalWithdrawAmountByCardNumber(r.ctx, cardNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get total withdrawal amount for card %s: %w", cardNumber, err)
	}
	return &res, nil
}

func (r *cardRepository) GetTotalTransactionAmountByCardNumber(cardNumber string) (*int64, error) {
	res, err := r.db.GetTotalTransactionAmountByCardNumber(r.ctx, cardNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get total transaction amount for card %s: %w", cardNumber, err)
	}
	return &res, nil
}

func (r *cardRepository) GetTotalTransferAmountBySender(senderCardNumber string) (*int64, error) {
	res, err := r.db.GetTotalTransferAmountBySender(r.ctx, senderCardNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get total transfer amount sent by card %s: %w", senderCardNumber, err)
	}
	return &res, nil
}

func (r *cardRepository) GetTotalTransferAmountByReceiver(receiverCardNumber string) (*int64, error) {
	res, err := r.db.GetTotalTransferAmountByReceiver(r.ctx, receiverCardNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get total transfer amount received by card %s: %w", receiverCardNumber, err)
	}
	return &res, nil
}

func (r *cardRepository) GetMonthlyBalance(year int) ([]*record.CardMonthBalance, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyBalances(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly balances: %w", err)
	}
	return r.mapping.ToMonthlyBalances(res), nil
}

func (r *cardRepository) GetYearlyBalance(year int) ([]*record.CardYearlyBalance, error) {
	res, err := r.db.GetYearlyBalances(r.ctx, int32(year))

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly balances: %w", err)
	}

	return r.mapping.ToYearlyBalances(res), nil
}

func (r *cardRepository) GetMonthlyTopupAmount(year int) ([]*record.CardMonthTopupAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTopupAmount(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly topup amount: %w", err)
	}

	return r.mapping.ToMonthlyTopupAmounts(res), nil
}

func (r *cardRepository) GetYearlyTopupAmount(year int) ([]*record.CardYearlyTopupAmount, error) {
	res, err := r.db.GetYearlyTopupAmount(r.ctx, int32(year))
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly topup amount: %w", err)
	}

	return r.mapping.ToYearlyTopupAmounts(res), nil
}

func (r *cardRepository) GetMonthlyWithdrawAmount(year int) ([]*record.CardMonthWithdrawAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyWithdrawAmount(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly withdraw amount: %w", err)
	}

	return r.mapping.ToMonthlyWithdrawAmounts(res), nil
}

func (r *cardRepository) GetYearlyWithdrawAmount(year int) ([]*record.CardYearlyWithdrawAmount, error) {
	res, err := r.db.GetYearlyWithdrawAmount(r.ctx, int32(year))
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly withdraw amount: %w", err)
	}

	return r.mapping.ToYearlyWithdrawAmounts(res), nil
}

func (r *cardRepository) GetMonthlyTransactionAmount(year int) ([]*record.CardMonthTransactionAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransactionAmount(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly transaction amount: %w", err)
	}

	return r.mapping.ToMonthlyTransactionAmounts(res), nil
}

func (r *cardRepository) GetYearlyTransactionAmount(year int) ([]*record.CardYearlyTransactionAmount, error) {
	res, err := r.db.GetYearlyTransactionAmount(r.ctx, int32(year))
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly transaction amount: %w", err)
	}

	return r.mapping.ToYearlyTransactionAmounts(res), nil
}

func (r *cardRepository) GetMonthlyTransferAmountSender(year int) ([]*record.CardMonthTransferAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransferAmountSender(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly transfer sender amount: %w", err)
	}

	return r.mapping.ToMonthlyTransferSenderAmounts(res), nil
}

func (r *cardRepository) GetYearlyTransferAmountSender(year int) ([]*record.CardYearlyTransferAmount, error) {
	res, err := r.db.GetYearlyTransferAmountSender(r.ctx, int32(year))
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly transfer sender amount: %w", err)
	}

	return r.mapping.ToYearlyTransferSenderAmounts(res), nil
}

func (r *cardRepository) GetMonthlyTransferAmountReceiver(year int) ([]*record.CardMonthTransferAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransferAmountReceiver(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly transfer receiver amount: %w", err)
	}

	return r.mapping.ToMonthlyTransferReceiverAmounts(res), nil
}

func (r *cardRepository) GetYearlyTransferAmountReceiver(year int) ([]*record.CardYearlyTransferAmount, error) {
	res, err := r.db.GetYearlyTransferAmountReceiver(r.ctx, int32(year))
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly transfer receiver amount: %w", err)
	}

	return r.mapping.ToYearlyTransferReceiverAmounts(res), nil
}

func (r *cardRepository) GetMonthlyBalancesByCardNumber(card_number string, year int) ([]*record.CardMonthBalance, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyBalancesByCardNumber(r.ctx, db.GetMonthlyBalancesByCardNumberParams{
		Column1:    yearStart,
		CardNumber: card_number,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly balances: %w", err)
	}
	return r.mapping.ToMonthlyBalancesCardNumber(res), nil
}

func (r *cardRepository) GetYearlyBalanceByCardNumber(card_number string, year int) ([]*record.CardYearlyBalance, error) {
	res, err := r.db.GetYearlyBalancesByCardNumber(r.ctx, db.GetYearlyBalancesByCardNumberParams{
		Column1:    year,
		CardNumber: card_number,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly balances: %w", err)
	}

	return r.mapping.ToYearlyBalancesCardNumber(res), nil
}

func (r *cardRepository) GetMonthlyTopupAmountByCardNumber(cardNumber string, year int) ([]*record.CardMonthTopupAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTopupAmountByCardNumber(r.ctx, db.GetMonthlyTopupAmountByCardNumberParams{
		Column2:    yearStart,
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly topup amount by card number: %w", err)
	}

	return r.mapping.ToMonthlyTopupAmountsByCardNumber(res), nil
}

func (r *cardRepository) GetYearlyTopupAmountByCardNumber(cardNumber string, year int) ([]*record.CardYearlyTopupAmount, error) {
	res, err := r.db.GetYearlyTopupAmountByCardNumber(r.ctx, db.GetYearlyTopupAmountByCardNumberParams{
		Column2:    int32(year),
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly topup amount by card number: %w", err)
	}

	return r.mapping.ToYearlyTopupAmountsByCardNumber(res), nil
}

func (r *cardRepository) GetMonthlyWithdrawAmountByCardNumber(cardNumber string, year int) ([]*record.CardMonthWithdrawAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyWithdrawAmountByCardNumber(r.ctx, db.GetMonthlyWithdrawAmountByCardNumberParams{
		Column2:    yearStart,
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly withdraw amount by card number: %w", err)
	}

	return r.mapping.ToMonthlyWithdrawAmountsByCardNumber(res), nil
}

func (r *cardRepository) GetYearlyWithdrawAmountByCardNumber(cardNumber string, year int) ([]*record.CardYearlyWithdrawAmount, error) {
	res, err := r.db.GetYearlyWithdrawAmountByCardNumber(r.ctx, db.GetYearlyWithdrawAmountByCardNumberParams{
		Column2:    int32(year),
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly withdraw amount by card number: %w", err)
	}

	return r.mapping.ToYearlyWithdrawAmountsByCardNumber(res), nil
}

func (r *cardRepository) GetMonthlyTransactionAmountByCardNumber(cardNumber string, year int) ([]*record.CardMonthTransactionAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransactionAmountByCardNumber(r.ctx, db.GetMonthlyTransactionAmountByCardNumberParams{
		Column2:    yearStart,
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly transaction amount by card number: %w", err)
	}

	return r.mapping.ToMonthlyTransactionAmountsByCardNumber(res), nil
}

func (r *cardRepository) GetYearlyTransactionAmountByCardNumber(cardNumber string, year int) ([]*record.CardYearlyTransactionAmount, error) {
	res, err := r.db.GetYearlyTransactionAmountByCardNumber(r.ctx, db.GetYearlyTransactionAmountByCardNumberParams{
		Column2:    int32(year),
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly transaction amount by card number: %w", err)
	}

	return r.mapping.ToYearlyTransactionAmountsByCardNumber(res), nil
}

func (r *cardRepository) GetMonthlyTransferAmountBySender(cardNumber string, year int) ([]*record.CardMonthTransferAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransferAmountBySender(r.ctx, db.GetMonthlyTransferAmountBySenderParams{
		Column2:      yearStart,
		TransferFrom: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly transfer sender amount by card number: %w", err)
	}

	return r.mapping.ToMonthlyTransferSenderAmountsByCardNumber(res), nil
}

func (r *cardRepository) GetYearlyTransferAmountBySender(cardNumber string, year int) ([]*record.CardYearlyTransferAmount, error) {
	res, err := r.db.GetYearlyTransferAmountBySender(r.ctx, db.GetYearlyTransferAmountBySenderParams{
		Column2:      int32(year),
		TransferFrom: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly transfer sender amount by card number: %w", err)
	}

	return r.mapping.ToYearlyTransferSenderAmountsByCardNumber(res), nil
}

func (r *cardRepository) GetMonthlyTransferAmountByReceiver(cardNumber string, year int) ([]*record.CardMonthTransferAmount, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransferAmountByReceiver(r.ctx, db.GetMonthlyTransferAmountByReceiverParams{
		Column2:    yearStart,
		TransferTo: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly transfer receiver amount by card number: %w", err)
	}

	return r.mapping.ToMonthlyTransferReceiverAmountsByCardNumber(res), nil
}

func (r *cardRepository) GetYearlyTransferAmountByReceiver(cardNumber string, year int) ([]*record.CardYearlyTransferAmount, error) {
	res, err := r.db.GetYearlyTransferAmountByReceiver(r.ctx, db.GetYearlyTransferAmountByReceiverParams{
		Column2:    int32(year),
		TransferTo: cardNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly transfer receiver amount by card number: %w", err)
	}

	return r.mapping.ToYearlyTransferReceiverAmountsByCardNumber(res), nil
}

func (r *cardRepository) FindByActive(search string, page, pageSize int) ([]*record.CardRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetActiveCardsWithCountParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetActiveCardsWithCount(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCardRecordsActive(res), totalCount, nil

}

func (r *cardRepository) FindByTrashed(search string, page, pageSize int) ([]*record.CardRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTrashedCardsWithCountParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTrashedCardsWithCount(r.ctx, req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get trashed saldos: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCardRecordsTrashed(res), totalCount, nil
}

func (r *cardRepository) CreateCard(request *requests.CreateCardRequest) (*record.CardRecord, error) {
	number, err := randomvcc.RandomCardNumber()

	if err != nil {
		return nil, fmt.Errorf("failed to generate card number: %w", err)
	}

	req := db.CreateCardParams{
		UserID:       int32(request.UserID),
		CardNumber:   number,
		CardType:     request.CardType,
		ExpireDate:   request.ExpireDate,
		Cvv:          request.CVV,
		CardProvider: request.CardProvider,
	}

	res, err := r.db.CreateCard(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to create card")
	}

	return r.mapping.ToCardRecord(res), nil
}
func (r *cardRepository) UpdateCard(request *requests.UpdateCardRequest) (*record.CardRecord, error) {
	req := db.UpdateCardParams{
		CardID:       int32(request.CardID),
		CardType:     request.CardType,
		ExpireDate:   request.ExpireDate,
		Cvv:          request.CVV,
		CardProvider: request.CardProvider,
	}

	err := r.db.UpdateCard(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update card: %w", err)
	}

	card, err := r.db.GetCardByID(r.ctx, int32(request.CardID))

	if err != nil {
		return nil, fmt.Errorf("failed to find card: %w", err)
	}

	return r.mapping.ToCardRecord(card), nil
}

func (r *cardRepository) TrashedCard(saldoID int) (*record.CardRecord, error) {
	err := r.db.TrashSaldo(r.ctx, int32(saldoID))
	if err != nil {
		return nil, fmt.Errorf("failed to trash card: %w", err)
	}

	card, err := r.db.GetTrashedCardByID(r.ctx, int32(saldoID))

	if err != nil {
		return nil, fmt.Errorf("card not found after trashing: %w", err)
	}

	return r.mapping.ToCardRecord(card), nil
}

func (r *cardRepository) RestoreCard(cardId int) (*record.CardRecord, error) {
	err := r.db.RestoreCard(r.ctx, int32(cardId))

	if err != nil {
		return nil, fmt.Errorf("failed to restore card: %w", err)
	}

	card, err := r.db.GetCardByID(r.ctx, int32(cardId))

	if err != nil {
		return nil, fmt.Errorf("card not found restore card: %w", err)
	}

	return r.mapping.ToCardRecord(card), nil
}

func (r *cardRepository) DeleteCardPermanent(card_id int) (bool, error) {
	err := r.db.DeleteCardPermanently(r.ctx, int32(card_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete card permanently: %w", err)
	}

	return true, nil
}

func (r *cardRepository) RestoreAllCard() (bool, error) {
	err := r.db.RestoreAllCards(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all cards: %w", err)
	}

	return true, nil
}

func (r *cardRepository) DeleteAllCardPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentCards(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all cards permanently: %w", err)
	}

	return true, nil
}
