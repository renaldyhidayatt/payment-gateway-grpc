package record

type MerchantRecord struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	ApiKey    string  `json:"api_key"`
	UserID    int     `json:"user_id"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
