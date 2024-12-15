package record

type TransactionRecord struct {
	ID              int     `json:"id"`
	CardNumber      string  `json:"card_number"`
	Amount          int     `json:"amount"`
	PaymentMethod   string  `json:"payment_method"`
	MerchantID      int     `json:"merchant_id"`
	TransactionTime string  `json:"transaction_time"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       *string `json:"deleted_at"`
}
