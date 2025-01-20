package response

type MerchantResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type MerchantResponseDeleteAt struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type MerchantResponseMonthlyPaymentMethod struct {
	Month         string `json:"month"`
	PaymentMethod string `json:"payment_method"`
	TotalAmount   int    `json:"total_amount"`
}

type MerchantResponseYearlyPaymentMethod struct {
	Year          string `json:"year"`
	PaymentMethod string `json:"payment_method"`
	TotalAmount   int    `json:"total_amount"`
}

type MerchantResponseMonthlyAmount struct {
	Month       string `json:"month"`
	TotalAmount int    `json:"total_amount"`
}

type MerchantResponseYearlyAmount struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
}
