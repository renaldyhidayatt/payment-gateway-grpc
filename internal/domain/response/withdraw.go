package response

type WithdrawResponse struct {
	ID             int    `json:"id"`
	CardNumber     string `json:"card_number"`
	WithdrawAmount int    `json:"withdraw_amount"`
	WithdrawTime   string `json:"withdraw_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
