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

func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethod(ms *db.GetMonthlyPaymentMethodsMerchantRow) *record.MerchantMonthlyPaymentMethod {
	return &record.MerchantMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethods(ms []*db.GetMonthlyPaymentMethodsMerchantRow) []*record.MerchantMonthlyPaymentMethod {
	var records []*record.MerchantMonthlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyPaymentMethod(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethod(ms *db.GetYearlyPaymentMethodMerchantRow) *record.MerchantYearlyPaymentMethod {
	return &record.MerchantYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethods(ms []*db.GetYearlyPaymentMethodMerchantRow) []*record.MerchantYearlyPaymentMethod {
	var records []*record.MerchantYearlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyPaymentMethod(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantMonthlyAmount(ms *db.GetMonthlyAmountMerchantRow) *record.MerchantMonthlyAmount {
	return &record.MerchantMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

func (m *merchantRecordMapper) ToMerchantMonthlyAmounts(ms []*db.GetMonthlyAmountMerchantRow) []*record.MerchantMonthlyAmount {
	var records []*record.MerchantMonthlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyAmount(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantYearlyAmount(ms *db.GetYearlyAmountMerchantRow) *record.MerchantYearlyAmount {
	return &record.MerchantYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

func (m *merchantRecordMapper) ToMerchantYearlyAmounts(ms []*db.GetYearlyAmountMerchantRow) []*record.MerchantYearlyAmount {
	var records []*record.MerchantYearlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyAmount(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodByMerchant(ms *db.GetMonthlyPaymentMethodByMerchantsRow) *record.MerchantMonthlyPaymentMethod {
	return &record.MerchantMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodsByMerchant(ms []*db.GetMonthlyPaymentMethodByMerchantsRow) []*record.MerchantMonthlyPaymentMethod {
	var records []*record.MerchantMonthlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyPaymentMethodByMerchant(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodByMerchant(ms *db.GetYearlyPaymentMethodByMerchantsRow) *record.MerchantYearlyPaymentMethod {
	return &record.MerchantYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodsByMerchant(ms []*db.GetYearlyPaymentMethodByMerchantsRow) []*record.MerchantYearlyPaymentMethod {
	var records []*record.MerchantYearlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyPaymentMethodByMerchant(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantMonthlyAmountByMerchant(ms *db.GetMonthlyAmountByMerchantsRow) *record.MerchantMonthlyAmount {
	return &record.MerchantMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

func (m *merchantRecordMapper) ToMerchantMonthlyAmountsByMerchant(ms []*db.GetMonthlyAmountByMerchantsRow) []*record.MerchantMonthlyAmount {
	var records []*record.MerchantMonthlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyAmountByMerchant(merchant))
	}
	return records
}

func (m *merchantRecordMapper) ToMerchantYearlyAmountByMerchant(ms *db.GetYearlyAmountByMerchantsRow) *record.MerchantYearlyAmount {
	return &record.MerchantYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

func (m *merchantRecordMapper) ToMerchantYearlyAmountsMerchant(ms []*db.GetYearlyAmountByMerchantsRow) []*record.MerchantYearlyAmount {
	var records []*record.MerchantYearlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyAmountByMerchant(merchant))
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
