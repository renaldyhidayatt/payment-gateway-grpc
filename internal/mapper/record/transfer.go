package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
)

type transferRecordMapper struct {
}

func NewTransferRecordMapper() *transferRecordMapper {
	return &transferRecordMapper{}
}

func (t *transferRecordMapper) ToTransferRecord(transfer *db.Transfer) *record.TransferRecord {
	var deletedAt *string

	return &record.TransferRecord{
		ID:             int(transfer.TransferID),
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime.String(),
		CreatedAt:      transfer.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      transfer.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

func (s *transferRecordMapper) ToTransfersRecord(transfers []*db.Transfer) []*record.TransferRecord {
	var transferRecords []*record.TransferRecord

	for _, transfer := range transfers {
		transferRecords = append(transferRecords, s.ToTransferRecord(transfer))
	}

	return transferRecords
}
