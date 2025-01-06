package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
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

func (s *withdrawRecordMapper) ToWithdrawRecordAll(withdraw *db.GetWithdrawsRow) *record.WithdrawRecord {
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

func (s *withdrawRecordMapper) ToWithdrawsRecordALl(withdraws []*db.GetWithdrawsRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecordAll(withdraw))
	}

	return withdrawRecords
}

func (s *withdrawRecordMapper) ToWithdrawRecordActive(withdraw *db.GetActiveWithdrawsRow) *record.WithdrawRecord {
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

func (s *withdrawRecordMapper) ToWithdrawsRecordActive(withdraws []*db.GetActiveWithdrawsRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecordActive(withdraw))
	}

	return withdrawRecords
}

func (s *withdrawRecordMapper) ToWithdrawRecordTrashed(withdraw *db.GetTrashedWithdrawsRow) *record.WithdrawRecord {
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

func (s *withdrawRecordMapper) ToWithdrawsRecordTrashed(withdraws []*db.GetTrashedWithdrawsRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecordTrashed(withdraw))
	}

	return withdrawRecords
}
