package requests

import "github.com/go-playground/validator/v10"

type CreateMerchantRequest struct {
	Name   string `json:"name" validate:"required"`
	UserID int    `json:"user_id" validate:"required,min=1"`
}

type UpdateMerchantRequest struct {
	MerchantID int    `json:"merchant_id" validate:"required,min=1"`
	Name       string `json:"name" validate:"required"`
	UserID     int    `json:"user_id" validate:"required,min=1"`
	Status     string `json:"status" validate:"required"`
}

func (r CreateMerchantRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r UpdateMerchantRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
