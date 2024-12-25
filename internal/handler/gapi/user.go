package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	protomapper "MamangRust/paymentgatewaygrpc/internal/mapper/proto"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userHandleGrpc struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
	mapping     protomapper.UserProtoMapper
}

func NewUserHandleGrpc(user service.UserService, mapper protomapper.UserProtoMapper) *userHandleGrpc {
	return &userHandleGrpc{userService: user, mapping: mapper}
}

func (s *userHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUser, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := s.userService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch users: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesUser(users)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalRecords / pageSize),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationUser{
		Status:     "success",
		Message:    "Successfully fetched users",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *userHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error) {
	user, err := s.userService.FindByID(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch user: " + err.Message,
		})
	}

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully fetched user",
		User:    s.mapping.ToResponseUser(user),
	}, nil

}

func (s *userHandleGrpc) FindByActive(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesUser, error) {
	users, err := s.userService.FindByActive()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active users: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesUser(users)

	return &pb.ApiResponsesUser{
		Status:  "success",
		Message: "Successfully fetched active users",
		Data:    so,
	}, nil
}

func (s *userHandleGrpc) FindByTrashed(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesUser, error) {
	users, err := s.userService.FindByTrashed()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed users: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesUser(users)

	return &pb.ApiResponsesUser{
		Status:  "success",
		Message: "Successfully fetched trashed users",
		Data:    so,
	}, nil
}

func (s *userHandleGrpc) Create(ctx context.Context, request *pb.CreateUserRequest) (*pb.ApiResponseUser, error) {
	req := &requests.CreateUserRequest{
		FirstName:       request.GetFirstname(),
		LastName:        request.GetLastname(),
		Email:           request.GetEmail(),
		Password:        request.GetPassword(),
		ConfirmPassword: request.GetConfirmPassword(),
	}

	user, err := s.userService.CreateUser(req)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create user: " + err.Message,
		})
	}

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully created user",
		User:    s.mapping.ToResponseUser(user),
	}, nil
}

func (s *userHandleGrpc) Update(ctx context.Context, request *pb.UpdateUserRequest) (*pb.ApiResponseUser, error) {
	req := &requests.UpdateUserRequest{
		UserID:          int(request.GetId()),
		FirstName:       request.GetFirstname(),
		LastName:        request.GetLastname(),
		Email:           request.GetEmail(),
		Password:        request.GetPassword(),
		ConfirmPassword: request.GetConfirmPassword(),
	}

	user, err := s.userService.UpdateUser(req)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update user: " + err.Message,
		})
	}

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully updated user",
		User:    s.mapping.ToResponseUser(user),
	}, nil
}

func (s *userHandleGrpc) TrashedUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error) {
	user, err := s.userService.TrashedUser(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed user: " + err.Message,
		})
	}

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully trashed user",
		User:    s.mapping.ToResponseUser(user),
	}, nil
}

func (s *userHandleGrpc) RestoreUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error) {
	user, err := s.userService.RestoreUser(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore user: " + err.Message,
		})
	}

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully restored user",
		User:    s.mapping.ToResponseUser(user),
	}, nil
}

func (s *userHandleGrpc) DeleteUserPermanent(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUserDelete, error) {
	_, err := s.userService.DeleteUserPermanent(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete user permanently: " + err.Message,
		})
	}

	return &pb.ApiResponseUserDelete{
		Status:  "success",
		Message: "Successfully deleted user permanently",
	}, nil
}
