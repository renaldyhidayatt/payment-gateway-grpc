package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
)

type merchantRecordMapper struct {
}

func NewMerchantRecordMapper() *merchantRecordMapper {
	return &merchantRecordMapper{}
}

func (m *merchantRecordMapper) ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

func (m *merchantRecordMapper) ToMerchantsRecord(merchants []*db.Merchant) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantRecord(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantGetAllRecord(merchant *db.GetMerchantsRow) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

func (m *merchantRecordMapper) ToMerchantsGetAllRecord(merchants []*db.GetMerchantsRow) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantGetAllRecord(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantActiveRecord(merchant *db.GetActiveMerchantsRow) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

func (m *merchantRecordMapper) ToMerchantsActiveRecord(merchants []*db.GetActiveMerchantsRow) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantActiveRecord(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantTrashedRecord(merchant *db.GetTrashedMerchantsRow) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

func (m *merchantRecordMapper) ToMerchantsTrashedRecord(merchants []*db.GetTrashedMerchantsRow) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTrashedRecord(merchant))
	}
	return records
}
