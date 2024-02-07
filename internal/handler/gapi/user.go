package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"context"

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
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	var pbUsers []*pb.User

	for _, user := range res {
		createdAtProto := timestamppb.New(user.CreatedAt.Time)

		var updatedAtProto *timestamppb.Timestamp
		if user.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(user.UpdatedAt.Time)
		}

		pbUsers = append(pbUsers, &pb.User{
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
		})
	}

	return &pb.UsersResponse{Users: pbUsers}, nil
}

func (s *userHandleGrpc) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	res, err := s.user.FindById(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.UserResponse{
		User: &pb.User{
			Firstname: res.Firstname,
			Lastname:  res.Lastname,
			Email:     res.Email,
		},
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
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(res.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if res.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(res.UpdatedAt.Time)
	}

	return &pb.UserResponse{
		User: &pb.User{
			Firstname: res.Firstname,
			Lastname:  res.Lastname,
			Email:     res.Email,
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
		},
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
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(res.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if res.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(res.UpdatedAt.Time)
	}

	return &pb.UserResponse{
		User: &pb.User{
			Firstname: res.Firstname,
			Lastname:  res.Lastname,
			Email:     res.Email,
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
		},
	}, nil
}

func (s *userHandleGrpc) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.DeleteUserResponse, error) {
	err := s.user.Delete(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}
