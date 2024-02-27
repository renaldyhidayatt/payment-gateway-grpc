package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userHandleGrpc struct {
	pb.UnimplementedUserServiceServer
	user service.UserService
}

func NewUserHandleGrpc(user service.UserService) *userHandleGrpc {
	return &userHandleGrpc{user: user}
}

func (s *userHandleGrpc) GetUsers(ctx context.Context, empty *emptypb.Empty) (*pb.UsersResponse, error) {
	res, err := s.user.FindAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while retrieving users: %v", err)
	}

	return &pb.UsersResponse{Users: s.convertToPbUsers(res)}, nil
}

func (s *userHandleGrpc) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	res, err := s.user.FindById(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found: %v", err)
	}

	return &pb.UserResponse{
		User: s.convertToPbUser(res),
	}, nil
}

func (s *userHandleGrpc) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	request := &requests.CreateUserRequest{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}

	res, err := s.user.Create(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while creating user: %v", err)
	}

	return &pb.UserResponse{
		User: s.convertToPbUser(res),
	}, nil
}

func (s *userHandleGrpc) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	request := &requests.UpdateUserRequest{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}

	res, err := s.user.Update(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while updating user: %v", err)
	}

	return &pb.UserResponse{
		User: s.convertToPbUser(res),
	}, nil
}

func (s *userHandleGrpc) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.DeleteUserResponse, error) {
	err := s.user.Delete(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while deleting user: %v", err)
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}

func (s *userHandleGrpc) convertToPbUsers(users []*db.User) []*pb.User {
	var pbUsers []*pb.User

	for _, user := range users {
		pbUsers = append(pbUsers, s.convertToPbUser(user))
	}

	return pbUsers
}

func (s *userHandleGrpc) convertToPbUser(user *db.User) *pb.User {
	createdAtProto := timestamppb.New(user.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(user.UpdatedAt.Time)
	}

	return &pb.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: createdAtProto,
		UpdatedAt: updatedAtProto,
	}
}
