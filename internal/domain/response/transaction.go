package response

type TransactionResponse struct {
	ID              int    `json:"id"`
	TransactionNo   string `json:"transaction_no"`
	CardNumber      string `json:"card_number"`
	Amount          int    `json:"amount"`
	PaymentMethod   string `json:"payment_method"`
	MerchantID      int    `json:"merchant_id"`
	TransactionTime string `json:"transaction_time"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type TransactionResponseDeleteAt struct {
	ID              int    `json:"id"`
	TransactionNo   string `json:"transaction_no"`
	CardNumber      string `json:"card_number"`
	Amount          int    `json:"amount"`
	PaymentMethod   string `json:"payment_method"`
	MerchantID      int    `json:"merchant_id"`
	TransactionTime string `json:"transaction_time"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}

type TransactionResponseMonthStatusSuccess struct {
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalAmount  int    `json:"total_amount"`
	TotalSuccess int    `json:"total_success"`
}

type TransactionResponseYearStatusSuccess struct {
	Year         string `json:"year"`
	TotalAmount  int    `json:"total_amount"`
	TotalSuccess int    `json:"total_success"`
}

type TransactionResponseMonthStatusFailed struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
}

type TransactionResponseYearStatusFailed struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
	TotalFailed int    `json:"total_failed"`
}

type TransactionMonthMethodResponse struct {
	Month             string `json:"month"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type TransactionYearMethodResponse struct {
	Year              string `json:"year"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type TransactionMonthAmountResponse struct {
	Month       string `json:"month"`
	TotalAmount int    `json:"total_amount"`
}

type TransactionYearlyAmountResponse struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
}
