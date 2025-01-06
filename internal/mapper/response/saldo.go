package responsemapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
)

type saldoResponseMapper struct {
}

func NewSaldoResponseMapper() *saldoResponseMapper {
	return &saldoResponseMapper{}
}

func (s *saldoResponseMapper) ToSaldoResponse(saldo *record.SaldoRecord) *response.SaldoResponse {
	return &response.SaldoResponse{
		ID:             saldo.ID,
		CardNumber:     saldo.CardNumber,
		TotalBalance:   saldo.TotalBalance,
		WithdrawAmount: saldo.WithdrawAmount,
		WithdrawTime:   saldo.WithdrawTime,
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}

func (s *saldoResponseMapper) ToSaldoResponses(saldos []*record.SaldoRecord) []*response.SaldoResponse {
	var responses []*response.SaldoResponse

	for _, response := range saldos {
		responses = append(responses, s.ToSaldoResponse(response))
	}

	return responses
}

func (s *saldoResponseMapper) ToSaldoResponseDeleteAt(saldo *record.SaldoRecord) *response.SaldoResponseDeleteAt {
	return &response.SaldoResponseDeleteAt{
		ID:             saldo.ID,
		CardNumber:     saldo.CardNumber,
		TotalBalance:   saldo.TotalBalance,
		WithdrawAmount: saldo.WithdrawAmount,
		WithdrawTime:   saldo.WithdrawTime,
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
		DeletedAt:      *saldo.DeletedAt,
	}
}

func (s *saldoResponseMapper) ToSaldoResponsesDeleteAt(saldos []*record.SaldoRecord) []*response.SaldoResponseDeleteAt {
	var responses []*response.SaldoResponseDeleteAt

	for _, response := range saldos {
		responses = append(responses, s.ToSaldoResponseDeleteAt(response))
	}

	return responses
}
