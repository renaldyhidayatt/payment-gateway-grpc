package response

type TopupResponse struct {
	ID          int    `json:"id"`
	CardNumber  string `json:"card_number"`
	TopupNo     string `json:"topup_no"`
	TopupAmount int    `json:"topup_amount"`
	TopupMethod string `json:"topup_method"`
	TopupTime   string `json:"topup_time"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TopupResponseDeleteAt struct {
	ID          int    `json:"id"`
	CardNumber  string `json:"card_number"`
	TopupNo     string `json:"topup_no"`
	TopupAmount int    `json:"topup_amount"`
	TopupMethod string `json:"topup_method"`
	TopupTime   string `json:"topup_time"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_At"`
}

type TopupResponseMonthStatusSuccess struct {
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalAmount  int    `json:"total_amount"`
	TotalSuccess int    `json:"total_success"`
}

type TopupResponseYearStatusSuccess struct {
	Year         string `json:"year"`
	TotalAmount  int    `json:"total_amount"`
	TotalSuccess int    `json:"total_success"`
}

type TopupResponseMonthStatusFailed struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
}

type TopupResponseYearStatusFailed struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
	TotalFailed int    `json:"total_failed"`
}

type TopupMonthMethodResponse struct {
	Month       string `json:"month"`
	TopupMethod string `json:"topup_method"`
	TotalTopups int    `json:"total_topups"`
	TotalAmount int    `json:"total_amount"`
}

type TopupYearlyMethodResponse struct {
	Year        string `json:"year"`
	TopupMethod string `json:"topup_method"`
	TotalTopups int    `json:"total_topups"`
	TotalAmount int    `json:"total_amount"`
}

type TopupMonthAmountResponse struct {
	Month       string `json:"month"`
	TotalAmount int    `json:"total_amount"`
}

type TopupYearlyAmountResponse struct {
	Year        string `json:"year"`
	TotalAmount int    `json:"total_amount"`
}
