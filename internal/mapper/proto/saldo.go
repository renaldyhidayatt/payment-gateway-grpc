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

func (s *saldoProtoMapper) ToSaldoMonthTotalBalanceResponse(ss *response.SaldoMonthTotalBalanceResponse) *pb.SaldoMonthTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = ss.TotalBalance
	}

	return &pb.SaldoMonthTotalBalanceResponse{
		Month:        ss.Month,
		Year:         ss.Year,
		TotalBalance: int32(totalBalance),
	}
}

func (s *saldoProtoMapper) ToSaldoMonthTotalBalanceResponses(ss []*response.SaldoMonthTotalBalanceResponse) []*pb.SaldoMonthTotalBalanceResponse {
	var saldoProtos []*pb.SaldoMonthTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.ToSaldoMonthTotalBalanceResponse(saldo))
	}
	return saldoProtos
}

func (s *saldoProtoMapper) ToSaldoYearTotalBalanceResponse(ss *response.SaldoYearTotalBalanceResponse) *pb.SaldoYearTotalBalanceResponse {
	return &pb.SaldoYearTotalBalanceResponse{
		Year:         ss.Year,
		TotalBalance: int32(ss.TotalBalance),
	}
}

func (s *saldoProtoMapper) ToSaldoYearTotalBalanceResponses(ss []*response.SaldoYearTotalBalanceResponse) []*pb.SaldoYearTotalBalanceResponse {
	var saldoProtos []*pb.SaldoYearTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.ToSaldoYearTotalBalanceResponse(saldo))
	}
	return saldoProtos
}

func (s *saldoProtoMapper) ToSaldoMonthBalanceResponse(ss *response.SaldoMonthBalanceResponse) *pb.SaldoMonthBalanceResponse {
	return &pb.SaldoMonthBalanceResponse{
		Month:        ss.Month,
		TotalBalance: int32(ss.TotalBalance),
	}
}

func (s *saldoProtoMapper) ToSaldoMonthBalanceResponses(ss []*response.SaldoMonthBalanceResponse) []*pb.SaldoMonthBalanceResponse {
	var saldoProtos []*pb.SaldoMonthBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.ToSaldoMonthBalanceResponse(saldo))
	}
	return saldoProtos
}

func (s *saldoProtoMapper) ToSaldoYearBalanceResponse(ss *response.SaldoYearBalanceResponse) *pb.SaldoYearBalanceResponse {
	return &pb.SaldoYearBalanceResponse{
		Year:         ss.Year,
		TotalBalance: int32(ss.TotalBalance),
	}
}

func (s *saldoProtoMapper) ToSaldoYearBalanceResponses(ss []*response.SaldoYearBalanceResponse) []*pb.SaldoYearBalanceResponse {
	var saldoProtos []*pb.SaldoYearBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.ToSaldoYearBalanceResponse(saldo))
	}
	return saldoProtos
}
