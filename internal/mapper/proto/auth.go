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

func (s *authProtoMapper) ToResponseLogin(response *response.TokenResponse) *pb.ApiResponseLogin {
	return &pb.ApiResponseLogin{
		Status:  "success",
		Message: "Login successful",
		Data: &pb.TokenResponse{
			AccessToken:  response.AccessToken,
			RefreshToken: response.RefreshToken,
		},
	}
}

func (s *authProtoMapper) ToResponseRegister(response *response.UserResponse) *pb.ApiResponseRegister {
	return &pb.ApiResponseRegister{
		Status:  "success",
		Message: "Registration successful",
		Data: &pb.UserResponse{
			Id:        int32(response.ID),
			Firstname: response.FirstName,
			Lastname:  response.LastName,
			Email:     response.Email,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
		},
	}
}

func (s *authProtoMapper) ToResponseRefreshToken(response *response.TokenResponse) *pb.ApiResponseRefreshToken {
	return &pb.ApiResponseRefreshToken{
		Status:  "success",
		Message: "Refresh token successful",
		Data: &pb.TokenResponse{
			AccessToken:  response.AccessToken,
			RefreshToken: response.RefreshToken,
		},
	}
}

func (s *authProtoMapper) ToResponseGetMe(response *response.UserResponse) *pb.ApiResponseGetMe {
	return &pb.ApiResponseGetMe{
		Status:  "success",
		Message: "Get me successful",
		Data: &pb.UserResponse{
			Id:        int32(response.ID),
			Firstname: response.FirstName,
			Lastname:  response.LastName,
			Email:     response.Email,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
		},
	}
}
