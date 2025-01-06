package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type saldoProtoMapper struct {
}

func NewSaldoProtoMapper() *saldoProtoMapper {
	return &saldoProtoMapper{}
}

func (s *saldoProtoMapper) ToResponseSaldo(saldo *response.SaldoResponse) *pb.SaldoResponse {
	return &pb.SaldoResponse{
		SaldoId:        int32(saldo.ID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int32(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}

func (s *saldoProtoMapper) ToResponsesSaldo(saldos []*response.SaldoResponse) []*pb.SaldoResponse {
	var responseSaldos []*pb.SaldoResponse

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.ToResponseSaldo(saldo))
	}

	return responseSaldos
}

func (s *saldoProtoMapper) ToResponseSaldoDeleteAt(saldo *response.SaldoResponseDeleteAt) *pb.SaldoResponseDeleteAt {
	return &pb.SaldoResponseDeleteAt{
		SaldoId:        int32(saldo.ID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int32(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
		DeletedAt:      saldo.DeletedAt,
	}
}

func (s *saldoProtoMapper) ToResponsesSaldoDeleteAt(saldos []*response.SaldoResponseDeleteAt) []*pb.SaldoResponseDeleteAt {
	var responseSaldos []*pb.SaldoResponseDeleteAt

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.ToResponseSaldoDeleteAt(saldo))
	}

	return responseSaldos
}
