package response

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ApiResponseLogin struct {
	Status  string         `json:"status"`
	Message string         `json:"messsage"`
	Data    *TokenResponse `json:"data"`
}

type ApiResponseRegister struct {
	Status  string        `json:"status"`
	Message string        `json:"messsage"`
	Data    *UserResponse `json:"data"`
}

type ApiResponseRefreshToken struct {
	Status  string         `json:"status"`
	Message string         `json:"messsage"`
	Data    *TokenResponse `json:"data"`
}

type ApiResponseGetMe struct {
	Status  string        `json:"status"`
	Message string        `json:"messsage"`
	Data    *UserResponse `json:"data"`
}
