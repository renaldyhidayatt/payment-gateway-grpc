package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
)

type topupRecordMapper struct {
}

func NewTopupRecordMapper() *topupRecordMapper {
	return &topupRecordMapper{}
}

func (t *topupRecordMapper) ToTopupRecord(topup *db.Topup) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

func (t *topupRecordMapper) ToTopupRecords(topups []*db.Topup) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecord(topup))
	}

	return topupRecords
}

func (t *topupRecordMapper) ToTopupRecordAll(topup *db.GetTopupsRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

func (t *topupRecordMapper) ToTopupRecordsAll(topups []*db.GetTopupsRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordAll(topup))
	}

	return topupRecords
}

func (t *topupRecordMapper) ToTopupRecordActive(topup *db.GetActiveTopupsRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

func (t *topupRecordMapper) ToTopupRecordsActive(topups []*db.GetActiveTopupsRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordActive(topup))
	}

	return topupRecords
}

func (t *topupRecordMapper) ToTopupRecordTrashed(topup *db.GetTrashedTopupsRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

func (t *topupRecordMapper) ToTopupRecordsTrashed(topups []*db.GetTrashedTopupsRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordTrashed(topup))
	}

	return topupRecords
}
