package responsemapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
)

type merchantResponseMapper struct{}

func NewMerchantResponseMapper() *merchantResponseMapper {
	return &merchantResponseMapper{}
}

func (s *merchantResponseMapper) ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse {
	return &response.MerchantResponse{
		ID:        merchant.ID,
		Name:      merchant.Name,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}

func (s *merchantResponseMapper) ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse {
	var response []*response.MerchantResponse
	for _, merchant := range merchants {
		response = append(response, s.ToMerchantResponse(merchant))
	}
	return response
}
