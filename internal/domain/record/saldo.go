package record

type SaldoRecord struct {
	ID             int     `json:"id"`
	CardNumber     string  `json:"card_number"`
	TotalBalance   int     `json:"total_balance"`
	WithdrawAmount int     `json:"withdraw_amount"`
	WithdrawTime   string  `json:"withdraw_time"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedAt      *string `json:"deleted_at"`
}
