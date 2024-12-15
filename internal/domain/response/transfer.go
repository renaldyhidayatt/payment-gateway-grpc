package response

type TransferResponse struct {
	ID             int    `json:"id"`
	TransferFrom   string `json:"transfer_from"`
	TransferTo     string `json:"transfer_to"`
	TransferAmount int    `json:"transfer_amount"`
	TransferTime   string `json:"transfer_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
