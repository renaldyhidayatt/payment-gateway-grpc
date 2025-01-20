package response

type SaldoResponse struct {
	ID             int    `json:"id"`
	CardNumber     string `json:"card_number"`
	TotalBalance   int    `json:"total_balance"`
	WithdrawAmount int    `json:"withdraw_amount"`
	WithdrawTime   string `json:"withdraw_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type SaldoResponseDeleteAt struct {
	ID             int    `json:"id"`
	CardNumber     string `json:"card_number"`
	TotalBalance   int    `json:"total_balance"`
	WithdrawAmount int    `json:"withdraw_amount"`
	WithdrawTime   string `json:"withdraw_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}

type SaldoMonthTotalBalanceResponse struct {
	Month        string `json:"month"`
	Year         string `json:"year"`
	TotalBalance int    `json:"total_balance"`
}

type SaldoYearTotalBalanceResponse struct {
	Year         string `json:"year"`
	TotalBalance int    `json:"total_balance"`
}

type SaldoMonthBalanceResponse struct {
	Month        string `json:"month"`
	TotalBalance int    `json:"total_balance"`
}

type SaldoYearBalanceResponse struct {
	Year         string `json:"year"`
	TotalBalance int    `json:"total_balance"`
}
