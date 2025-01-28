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

func (u *userProtoMapper) ToProtoResponseUser(status string, message string, pbResponse *response.UserResponse) *pb.ApiResponseUser {
	return &pb.ApiResponseUser{
		Status:  status,
		Message: message,
		Data:    u.mapResponseUser(pbResponse),
	}
}

func (u *userProtoMapper) ToProtoResponsesUser(status string, message string, pbResponse []*response.UserResponse) *pb.ApiResponsesUser {
	return &pb.ApiResponsesUser{
		Status:  status,
		Message: message,
		Data:    u.mapResponsesUser(pbResponse),
	}
}

func (u *userProtoMapper) ToProtoResponseUserDelete(status string, message string) *pb.ApiResponseUserDelete {
	return &pb.ApiResponseUserDelete{
		Status:  status,
		Message: message,
	}
}

func (u *userProtoMapper) ToProtoResponseUserAll(status string, message string) *pb.ApiResponseUserAll {
	return &pb.ApiResponseUserAll{
		Status:  status,
		Message: message,
	}
}

func (u *userProtoMapper) ToProtoResponsePaginationUserDeleteAt(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponseDeleteAt) *pb.ApiResponsePaginationUserDeleteAt {
	return &pb.ApiResponsePaginationUserDeleteAt{
		Status:     status,
		Message:    message,
		Data:       u.mapResponsesUserDeleteAt(users),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (u *userProtoMapper) ToProtoResponsePaginationUser(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponse) *pb.ApiResponsePaginationUser {
	return &pb.ApiResponsePaginationUser{
		Status:     status,
		Message:    message,
		Data:       u.mapResponsesUser(users),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (u *userProtoMapper) mapResponseUser(user *response.UserResponse) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u *userProtoMapper) mapResponsesUser(users []*response.UserResponse) []*pb.UserResponse {
	var mappedUsers []*pb.UserResponse

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.mapResponseUser(user))
	}

	return mappedUsers
}

func (u *userProtoMapper) mapResponseUserDelete(user *response.UserResponseDeleteAt) *pb.UserResponseWithDeleteAt {
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

func (u *userProtoMapper) mapResponsesUserDeleteAt(users []*response.UserResponseDeleteAt) []*pb.UserResponseWithDeleteAt {
	var mappedUsers []*pb.UserResponseWithDeleteAt

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.mapResponseUserDelete(user))
	}

	return mappedUsers
}

func (u *userProtoMapper) mapProtoResponseUser(status string, message string, pbResponse *response.UserResponse) *pb.ApiResponseUser {
	return &pb.ApiResponseUser{
		Status:  status,
		Message: message,
		Data:    u.mapResponseUser(pbResponse),
	}
}

func mapPaginationMeta(s *pb.PaginationMeta) *pb.PaginationMeta {
	return &pb.PaginationMeta{
		CurrentPage:  int32(s.CurrentPage),
		PageSize:     int32(s.PageSize),
		TotalPages:   int32(s.TotalPages),
		TotalRecords: int32(s.TotalRecords),
	}
}
