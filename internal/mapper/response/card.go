package responsemapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
)

type cardResponseMapper struct {
}

func NewCardResponseMapper() *cardResponseMapper {
	return &cardResponseMapper{}
}

func (s *cardResponseMapper) ToCardResponse(card *record.CardRecord) *response.CardResponse {
	return &response.CardResponse{
		ID:           card.ID,
		UserID:       card.UserID,
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}

func (s *cardResponseMapper) ToCardsResponse(cards []*record.CardRecord) []*response.CardResponse {
	var response []*response.CardResponse

	for _, card := range cards {
		response = append(response, s.ToCardResponse(card))
	}

	return response
}

func (s *cardResponseMapper) ToCardResponseDeleteAt(card *record.CardRecord) *response.CardResponseDeleteAt {
	return &response.CardResponseDeleteAt{
		ID:           card.ID,
		UserID:       card.UserID,
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
		DeletedAt:    *card.DeletedAt,
	}
}

func (s *cardResponseMapper) ToCardsResponseDeleteAt(cards []*record.CardRecord) []*response.CardResponseDeleteAt {
	var response []*response.CardResponseDeleteAt

	for _, card := range cards {
		response = append(response, s.ToCardResponseDeleteAt(card))
	}

	return response
}

func (s *cardResponseMapper) ToGetMonthlyBalance(card *record.CardMonthBalance) *response.CardResponseMonthBalance {
	return &response.CardResponseMonthBalance{
		Month:        card.Month,
		TotalBalance: card.TotalBalance,
	}
}

func (s *cardResponseMapper) ToGetMonthlyBalances(cards []*record.CardMonthBalance) []*response.CardResponseMonthBalance {
	var records []*response.CardResponseMonthBalance

	for _, card := range cards {
		records = append(records, s.ToGetMonthlyBalance(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetYearlyBalance(card *record.CardYearlyBalance) *response.CardResponseYearlyBalance {
	return &response.CardResponseYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

func (s *cardResponseMapper) ToGetYearlyBalances(cards []*record.CardYearlyBalance) []*response.CardResponseYearlyBalance {
	var records []*response.CardResponseYearlyBalance

	for _, card := range cards {
		records = append(records, s.ToGetYearlyBalance(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetMonthlyTopupAmount(card *record.CardMonthTopupAmount) *response.CardResponseMonthTopupAmount {
	return &response.CardResponseMonthTopupAmount{
		Month:       card.Month,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetMonthlyTopupAmounts(cards []*record.CardMonthTopupAmount) []*response.CardResponseMonthTopupAmount {
	var records []*response.CardResponseMonthTopupAmount

	for _, card := range cards {
		records = append(records, s.ToGetMonthlyTopupAmount(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetYearlyTopupAmount(card *record.CardYearlyTopupAmount) *response.CardResponseYearlyTopupAmount {
	return &response.CardResponseYearlyTopupAmount{
		Year:        card.Year,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetYearlyTopupAmounts(cards []*record.CardYearlyTopupAmount) []*response.CardResponseYearlyTopupAmount {
	var records []*response.CardResponseYearlyTopupAmount

	for _, card := range cards {
		records = append(records, s.ToGetYearlyTopupAmount(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetMonthlyWithdrawAmount(card *record.CardMonthWithdrawAmount) *response.CardResponseMonthWithdrawAmount {
	return &response.CardResponseMonthWithdrawAmount{
		Month:       card.Month,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetMonthlyWithdrawAmounts(cards []*record.CardMonthWithdrawAmount) []*response.CardResponseMonthWithdrawAmount {
	var records []*response.CardResponseMonthWithdrawAmount

	for _, card := range cards {
		records = append(records, s.ToGetMonthlyWithdrawAmount(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetYearlyWithdrawAmount(card *record.CardYearlyWithdrawAmount) *response.CardResponseYearlyWithdrawAmount {
	return &response.CardResponseYearlyWithdrawAmount{
		Year:        card.Year,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetYearlyWithdrawAmounts(cards []*record.CardYearlyWithdrawAmount) []*response.CardResponseYearlyWithdrawAmount {
	var records []*response.CardResponseYearlyWithdrawAmount

	for _, card := range cards {
		records = append(records, s.ToGetYearlyWithdrawAmount(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetMonthlyTransactionAmount(card *record.CardMonthTransactionAmount) *response.CardResponseMonthTransactionAmount {
	return &response.CardResponseMonthTransactionAmount{
		Month:       card.Month,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetMonthlyTransactionAmounts(cards []*record.CardMonthTransactionAmount) []*response.CardResponseMonthTransactionAmount {
	var records []*response.CardResponseMonthTransactionAmount

	for _, card := range cards {
		records = append(records, s.ToGetMonthlyTransactionAmount(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetYearlyTransactionAmount(card *record.CardYearlyTransactionAmount) *response.CardResponseYearlyTransactionAmount {
	return &response.CardResponseYearlyTransactionAmount{
		Year:        card.Year,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetYearlyTransactionAmounts(cards []*record.CardYearlyTransactionAmount) []*response.CardResponseYearlyTransactionAmount {
	var records []*response.CardResponseYearlyTransactionAmount

	for _, card := range cards {
		records = append(records, s.ToGetYearlyTransactionAmount(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetMonthlyTransferSenderAmount(card *record.CardMonthTransferAmount) *response.CardResponseMonthTransferAmount {
	return &response.CardResponseMonthTransferAmount{
		Month:       card.Month,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetMonthlyTransferSenderAmounts(cards []*record.CardMonthTransferAmount) []*response.CardResponseMonthTransferAmount {
	var records []*response.CardResponseMonthTransferAmount

	for _, card := range cards {
		records = append(records, s.ToGetMonthlyTransferSenderAmount(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetYearlyTransferSenderAmount(card *record.CardYearlyTransferAmount) *response.CardResponseYearlyTransferAmount {
	return &response.CardResponseYearlyTransferAmount{
		Year:        card.Year,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetYearlyTransferSenderAmounts(cards []*record.CardYearlyTransferAmount) []*response.CardResponseYearlyTransferAmount {
	var records []*response.CardResponseYearlyTransferAmount

	for _, card := range cards {
		records = append(records, s.ToGetYearlyTransferSenderAmount(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetMonthlyTransferReceiverAmount(card *record.CardMonthTransferAmount) *response.CardResponseMonthTransferAmount {
	return &response.CardResponseMonthTransferAmount{
		Month:       card.Month,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetMonthlyTransferReceiverAmounts(cards []*record.CardMonthTransferAmount) []*response.CardResponseMonthTransferAmount {
	var records []*response.CardResponseMonthTransferAmount

	for _, card := range cards {
		records = append(records, s.ToGetMonthlyTransferReceiverAmount(card))
	}

	return records
}

func (s *cardResponseMapper) ToGetYearlyTransferReceiverAmount(card *record.CardYearlyTransferAmount) *response.CardResponseYearlyTransferAmount {
	return &response.CardResponseYearlyTransferAmount{
		Year:        card.Year,
		TotalAmount: card.TotalAmount,
	}
}

func (s *cardResponseMapper) ToGetYearlyTransferReceiverAmounts(cards []*record.CardYearlyTransferAmount) []*response.CardResponseYearlyTransferAmount {
	var records []*response.CardResponseYearlyTransferAmount

	for _, card := range cards {
		records = append(records, s.ToGetYearlyTransferReceiverAmount(card))
	}

	return records
}
