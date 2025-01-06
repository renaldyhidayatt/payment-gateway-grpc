package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type merchantProto struct{}

func NewMerchantProtoMapper() *merchantProto {
	return &merchantProto{}
}

func (m *merchantProto) ToResponseMerchant(merchant *response.MerchantResponse) *pb.MerchantResponse {
	return &pb.MerchantResponse{
		Id:        int32(merchant.ID),
		Name:      merchant.Name,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}

func (m *merchantProto) ToResponsesMerchant(merchants []*response.MerchantResponse) []*pb.MerchantResponse {
	var responseMerchants []*pb.MerchantResponse
	for _, merchant := range merchants {
		responseMerchants = append(responseMerchants, m.ToResponseMerchant(merchant))
	}
	return responseMerchants
}

func (m *merchantProto) ToResponseMerchantDeleteAt(merchant *response.MerchantResponseDeleteAt) *pb.MerchantResponseDeleteAt {
	return &pb.MerchantResponseDeleteAt{
		Id:        int32(merchant.ID),
		Name:      merchant.Name,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		DeletedAt: merchant.DeletedAt,
	}
}

func (m *merchantProto) ToResponsesMerchantDeleteAt(merchants []*response.MerchantResponseDeleteAt) []*pb.MerchantResponseDeleteAt {
	var responseMerchants []*pb.MerchantResponseDeleteAt
	for _, merchant := range merchants {
		responseMerchants = append(responseMerchants, m.ToResponseMerchantDeleteAt(merchant))
	}
	return responseMerchants
}
