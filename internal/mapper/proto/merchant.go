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

func (m *merchantProto) ToResponseMonthlyPaymentMethod(ms *response.MerchantResponseMonthlyPaymentMethod) *pb.MerchantResponseMonthlyPaymentMethod {
	return &pb.MerchantResponseMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int64(ms.TotalAmount),
	}
}

func (m *merchantProto) ToResponseMonthlyPaymentMethods(ms []*response.MerchantResponseMonthlyPaymentMethod) []*pb.MerchantResponseMonthlyPaymentMethod {
	var response []*pb.MerchantResponseMonthlyPaymentMethod
	for _, merchant := range ms {
		response = append(response, m.ToResponseMonthlyPaymentMethod(merchant))
	}
	return response
}

func (m *merchantProto) ToResponseYearlyPaymentMethod(ms *response.MerchantResponseYearlyPaymentMethod) *pb.MerchantResponseYearlyPaymentMethod {
	return &pb.MerchantResponseYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int64(ms.TotalAmount),
	}
}

func (m *merchantProto) ToResponseYearlyPaymentMethods(ms []*response.MerchantResponseYearlyPaymentMethod) []*pb.MerchantResponseYearlyPaymentMethod {
	var response []*pb.MerchantResponseYearlyPaymentMethod
	for _, merchant := range ms {
		response = append(response, m.ToResponseYearlyPaymentMethod(merchant))
	}
	return response
}

func (m *merchantProto) ToResponseMonthlyAmount(ms *response.MerchantResponseMonthlyAmount) *pb.MerchantResponseMonthlyAmount {
	return &pb.MerchantResponseMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int64(ms.TotalAmount),
	}
}
func (m *merchantProto) ToResponseMonthlyAmounts(ms []*response.MerchantResponseMonthlyAmount) []*pb.MerchantResponseMonthlyAmount {
	var response []*pb.MerchantResponseMonthlyAmount
	for _, merchant := range ms {
		response = append(response, m.ToResponseMonthlyAmount(merchant))
	}
	return response
}

func (m *merchantProto) ToResponseYearlyAmount(ms *response.MerchantResponseYearlyAmount) *pb.MerchantResponseYearlyAmount {
	return &pb.MerchantResponseYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int64(ms.TotalAmount),
	}
}

func (m *merchantProto) ToResponseYearlyAmounts(ms []*response.MerchantResponseYearlyAmount) []*pb.MerchantResponseYearlyAmount {
	var response []*pb.MerchantResponseYearlyAmount
	for _, merchant := range ms {
		response = append(response, m.ToResponseYearlyAmount(merchant))
	}
	return response
}
