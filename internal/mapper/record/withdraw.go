package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
)

type withdrawRecordMapper struct{}

func NewWithdrawRecordMapper() *withdrawRecordMapper {
	return &withdrawRecordMapper{}
}

func (s *withdrawRecordMapper) ToWithdrawRecord(withdraw *db.Withdraw) *record.WithdrawRecord {
	var deletedAt *string

	if withdraw.DeletedAt.Valid {
		formatedDeletedAt := withdraw.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.WithdrawRecord{
		ID:             int(withdraw.WithdrawID),
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime.String(),
		CreatedAt:      withdraw.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      withdraw.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

func (s *withdrawRecordMapper) ToWithdrawsRecord(withdraws []*db.Withdraw) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecord(withdraw))
	}

	return withdrawRecords
}
