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
