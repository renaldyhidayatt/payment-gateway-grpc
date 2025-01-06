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
