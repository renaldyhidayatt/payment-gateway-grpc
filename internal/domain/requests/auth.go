package requests

import "github.com/go-playground/validator/v10"

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthLoginRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

type JWTToken struct {
	Token string `json:"token"`
}
