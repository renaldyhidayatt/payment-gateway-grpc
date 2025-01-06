package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type userProtoMapper struct {
}

func NewUserProtoMapper() *userProtoMapper {
	return &userProtoMapper{}
}

func (u *userProtoMapper) ToResponseUser(user *response.UserResponse) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u *userProtoMapper) ToResponsesUser(users []*response.UserResponse) []*pb.UserResponse {
	var mappedUsers []*pb.UserResponse

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.ToResponseUser(user))
	}

	return mappedUsers
}

func (u *userProtoMapper) ToResponseUserDelete(user *response.UserResponseDeleteAt) *pb.UserResponseWithDeleteAt {
	return &pb.UserResponseWithDeleteAt{
		Id:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func (u *userProtoMapper) ToResponsesUserDeleteAt(users []*response.UserResponseDeleteAt) []*pb.UserResponseWithDeleteAt {
	var mappedUsers []*pb.UserResponseWithDeleteAt

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.ToResponseUserDelete(user))
	}

	return mappedUsers
}
