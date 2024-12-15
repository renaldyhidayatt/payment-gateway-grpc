package protomapper

type ProtoMapper struct {
	AuthProtoMapper        AuthProtoMapper
	CardProtoMapper        CardProtoMapper
	MerchantProtoMapper    MerchantProtoMapper
	SaldoProtoMapper       SaldoProtoMapper
	TopupProtoMapper       TopupProtoMapper
	TransactionProtoMapper TransactionProtoMapper
	TransferProtoMapper    TransferProtoMapper
	UserProtoMapper        UserProtoMapper
	WithdrawalProtoMapper  WithdrawalProtoMapper
}

func NewProtoMapper() *ProtoMapper {
	return &ProtoMapper{
		AuthProtoMapper:        NewAuthProtoMapper(),
		CardProtoMapper:        NewCardProtoMapper(),
		MerchantProtoMapper:    NewMerchantProtoMapper(),
		SaldoProtoMapper:       NewSaldoProtoMapper(),
		TopupProtoMapper:       NewTopupProtoMapper(),
		TransactionProtoMapper: NewTransactionProtoMapper(),
		TransferProtoMapper:    NewTransferProtoMapper(),
		UserProtoMapper:        NewUserProtoMapper(),
		WithdrawalProtoMapper:  NewWithdrawProtoMapper(),
	}
}
