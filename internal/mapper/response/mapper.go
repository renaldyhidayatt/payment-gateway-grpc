package responsemapper

type ResponseMapper struct {
	CardResponseMapper         CardResponseMapper
	RoleResponseMapper         RoleResponseMapper
	RefreshTokenResponseMapper RefreshTokenResponseMapper
	SaldoResponseMapper        SaldoResponseMapper
	TransactionResponseMapper  TransactionResponseMapper
	TransferResponseMapper     TransferResponseMapper
	TopupResponseMapper        TopupResponseMapper
	WithdrawResponseMapper     WithdrawResponseMapper
	UserResponseMapper         UserResponseMapper
	MerchantResponseMapper     MerchantResponseMapper
}

func NewResponseMapper() *ResponseMapper {
	return &ResponseMapper{
		CardResponseMapper:         NewCardResponseMapper(),
		SaldoResponseMapper:        NewSaldoResponseMapper(),
		TransactionResponseMapper:  NewTransactionResponseMapper(),
		TransferResponseMapper:     NewTransferResponseMapper(),
		TopupResponseMapper:        NewTopupResponseMapper(),
		WithdrawResponseMapper:     NewWithdrawResponseMapper(),
		UserResponseMapper:         NewUserResponseMapper(),
		RefreshTokenResponseMapper: NewRefreshTokenResponseMapper(),
		RoleResponseMapper:         NewRoleResponseMapper(),
		MerchantResponseMapper:     NewMerchantResponseMapper(),
	}
}
