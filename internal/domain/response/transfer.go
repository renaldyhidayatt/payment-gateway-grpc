package response

type TransferResponse struct {
	ID             int    `json:"id"`
	TransferNo     string `json:"transfer_no"`
	TransferFrom   string `json:"transfer_from"`
	TransferTo     string `json:"transfer_to"`
	TransferAmount int    `json:"transfer_amount"`
	TransferTime   string `json:"transfer_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type TransferResponseDeleteAt struct {
	ID             int    `json:"id"`
	TransferNo     string `json:"transfer_no"`
	TransferFrom   string `json:"transfer_from"`
	TransferTo     string `json:"transfer_to"`
	TransferAmount int    `json:"transfer_amount"`
	TransferTime   string `json:"transfer_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}

type TransferResponseMonthStatusSuccess struct {
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalAmount  int    `json:"total_amount"`
	TotalSuccess int    `json:"total_success"`
}

type TransferResponseYearStatusSuccess struct {
	Year         string `json:"year"`
	TotalAmount  int    `json:"total_amount"`
	TotalSuccess int    `json:"total_success"`
}

type TransferResponseMonthStatusFailed struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
}

type TransferResponseYearStatusFailed struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
	TotalFailed int    `json:"total_failed"`
}

type TransferMonthAmountResponse struct {
	Month       string `json:"month"`
	TotalAmount int    `json:"total_amount"`
}

type TransferYearAmountResponse struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
}
