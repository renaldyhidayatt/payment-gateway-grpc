package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"MamangRust/paymentgatewaygrpc/pkg/randomvcc"
	"context"
	"fmt"
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

func (r *cardRepository) DeleteCardPermanent(card_id int) error {
	err := r.db.DeleteCardPermanently(r.ctx, int32(card_id))

	if err != nil {
		return nil
	}

	return fmt.Errorf("failed delete card permanent")
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
