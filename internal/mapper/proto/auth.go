package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type authProtoMapper struct {
}

func NewAuthProtoMapper() *authProtoMapper {
	return &authProtoMapper{}
}

func (s *authProtoMapper) ToResponseLogin(token string) *pb.ApiResponseLogin {
	return &pb.ApiResponseLogin{
		Status:  "success",
		Message: "Login successful",
		Token:   token,
	}
}

func (s *authProtoMapper) ToResponseRegister(response response.UserResponse) *pb.ApiResponseRegister {
	return &pb.ApiResponseRegister{
		Status:  "success",
		Message: "Registration successful",
		User: &pb.UserResponse{
			Id:        int32(response.ID),
			Firstname: response.FirstName,
			Lastname:  response.LastName,
			Email:     response.Email,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
		},
	}
}
