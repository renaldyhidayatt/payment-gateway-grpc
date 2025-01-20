package record

type CardRecord struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	CardNumber   string  `json:"card_number"`
	CardType     string  `json:"card_type"`
	ExpireDate   string  `json:"expire_date"`
	CVV          string  `json:"cvv"`
	CardProvider string  `json:"card_provider"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    *string `json:"deleted_at"`
}

type CardMonthBalance struct {
	Month        string `json:"month"`
	TotalBalance int64  `json:"total_balance"`
}

type CardYearlyBalance struct {
	Year         string `json:"year"`
	TotalBalance int64  `json:"total_balance"`
}

type CardMonthTopupAmount struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

type CardYearlyTopupAmount struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}

type CardMonthWithdrawAmount struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

type CardYearlyWithdrawAmount struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}

type CardMonthTransactionAmount struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

type CardYearlyTransactionAmount struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}

type CardMonthTransferAmount struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

type CardYearlyTransferAmount struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}
