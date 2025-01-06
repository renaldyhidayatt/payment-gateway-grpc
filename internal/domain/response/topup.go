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
