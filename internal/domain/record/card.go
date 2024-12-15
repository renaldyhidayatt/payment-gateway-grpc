package record

type CardRecord struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	CardNumber   string  `json:"card_number"`
	CardType     string  `json:"card_type"`
	ExpireDate   string  `json:"expire_date"`
	CVV          string  `json:"cvv"`
	CardProvider string  `json:"card_provider"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    *string `json:"deleted_at"`
}
