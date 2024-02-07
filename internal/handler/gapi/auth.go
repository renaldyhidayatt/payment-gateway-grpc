package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"context"

	"google.golang.org/grpc/status"
)

type authHandleGrpc struct {
	pb.UnimplementedAuthServiceServer
	auth service.AuthService
}

func NewAuthHandleGrpc(auth service.AuthService) *authHandleGrpc {
	return &authHandleGrpc{auth: auth}
}

func (s *authHandleGrpc) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	request := &requests.AuthLoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := s.auth.Login(request)

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.LoginResponse{Token: res.Token}, nil
}

func (s *authHandleGrpc) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	request := &requests.CreateUserRequest{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}

	res, err := s.auth.Register(request)

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.RegisterResponse{
		User: &pb.User{
			Firstname: res.Firstname,
			Lastname:  res.Lastname,
			Email:     res.Email,
		},
	}, nil

}
