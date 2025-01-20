package requests

import (
	methodtopup "MamangRust/paymentgatewaygrpc/pkg/method_topup"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateTransactionRequest struct {
	CardNumber      string    `json:"card_number" validate:"required,min=1"`
	Amount          int       `json:"amount" validate:"required,min=50000"`
	PaymentMethod   string    `json:"payment_method" validate:"required"`
	MerchantID      *int      `json:"merchant_id" validate:"required,min=1"`
	TransactionTime time.Time `json:"transaction_time" validate:"required"`
}

type UpdateTransactionRequest struct {
	TransactionID   int       `json:"transaction_id" validate:"required,min=1"`
	CardNumber      string    `json:"card_number" validate:"required,min=1"`
	Amount          int       `json:"amount" validate:"required,min=50000"`
	PaymentMethod   string    `json:"payment_method" validate:"required"`
	MerchantID      *int      `json:"merchant_id" validate:"required,min=1"`
	TransactionTime time.Time `json:"transaction_time" validate:"required"`
}

type UpdateTransactionStatus struct {
	TransactionID int    `json:"transaction_id" validate:"required,min=1"`
	Status        string `json:"status" validate:"required"`
}

func (r *CreateTransactionRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if !methodtopup.PaymentMethodValidator(r.PaymentMethod) {
		return fmt.Errorf("payment method not found")
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateTransactionRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if !methodtopup.PaymentMethodValidator(r.PaymentMethod) {
		return fmt.Errorf("payment method not found")
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateTransactionStatus) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}
