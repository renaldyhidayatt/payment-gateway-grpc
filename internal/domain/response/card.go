package response

type CardResponse struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	CardNumber   string `json:"card_number"`
	CardType     string `json:"card_type"`
	ExpireDate   string `json:"expire_date"`
	CVV          string `json:"cvv"`
	CardProvider string `json:"card_provider"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CardResponseDeleteAt struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	CardNumber   string `json:"card_number"`
	CardType     string `json:"card_type"`
	ExpireDate   string `json:"expire_date"`
	CVV          string `json:"cvv"`
	CardProvider string `json:"card_provider"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}

type DashboardCard struct {
	TotalBalance     *int64 `json:"total_balance"`
	TotalTopup       *int64 `json:"total_topup"`
	TotalWithdraw    *int64 `json:"total_withdraw"`
	TotalTransaction *int64 `json:"total_transaction"`
	TotalTransfer    *int64 `json:"total_transfer"`
}

type DashboardCardCardNumber struct {
	TotalBalance          *int64 `json:"total_balance"`
	TotalTopup            *int64 `json:"total_topup"`
	TotalWithdraw         *int64 `json:"total_withdraw"`
	TotalTransaction      *int64 `json:"total_transaction"`
	TotalTransferSent     *int64 `json:"total_transfer_send"`
	TotalTransferReceiver *int64 `json:"total_transfer_receiver"`
}

type CardResponseMonthBalance struct {
	Month        string `json:"month"`
	TotalBalance int64  `json:"total_balance"`
}

type CardResponseYearlyBalance struct {
	Year         string `json:"year"`
	TotalBalance int64  `json:"total_amount"`
}

type CardResponseMonthTopupAmount struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

type CardResponseYearlyTopupAmount struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}

type CardResponseMonthWithdrawAmount struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

type CardResponseYearlyWithdrawAmount struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}

type CardResponseMonthTransactionAmount struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

type CardResponseYearlyTransactionAmount struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}

type CardResponseMonthTransferAmount struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

type CardResponseYearlyTransferAmount struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}
