package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
)

type saldoRecordMapper struct{}

func NewSaldoRecordMapper() *saldoRecordMapper {
	return &saldoRecordMapper{}
}

func (s *saldoRecordMapper) ToSaldoRecord(saldo *db.Saldo) *record.SaldoRecord {
	var deletedAt *string

	if saldo.DeletedAt.Valid {
		formatedDeletedAt := saldo.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.SaldoRecord{
		ID:             int(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawAmount: int(saldo.WithdrawAmount.Int32),
		WithdrawTime:   saldo.WithdrawTime.Time.Format("2006-01-02"),
		CreatedAt:      saldo.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:      saldo.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:      deletedAt,
	}
}

func (s *saldoRecordMapper) ToSaldosRecord(saldos []*db.Saldo) []*record.SaldoRecord {
	var saldoRecords []*record.SaldoRecord
	for _, saldo := range saldos {
		saldoRecords = append(saldoRecords, s.ToSaldoRecord(saldo))
	}
	return saldoRecords
}
