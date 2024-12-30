package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	protomapper "MamangRust/paymentgatewaygrpc/internal/mapper/proto"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authHandleGrpc struct {
	pb.UnimplementedAuthServiceServer
	authService service.AuthService
	mapping     protomapper.AuthProtoMapper
}

func NewAuthHandleGrpc(auth service.AuthService, mapping protomapper.AuthProtoMapper) *authHandleGrpc {
	return &authHandleGrpc{authService: auth, mapping: mapping}
}

func (s *authHandleGrpc) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.ApiResponseLogin, error) {
	request := &requests.AuthRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := s.authService.Login(request)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Login failed: ",
		})
	}

	return &pb.ApiResponseLogin{
		Status:  "success",
		Message: "Login successful",
		Token:   *res,
	}, nil
}

func (s *authHandleGrpc) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.ApiResponseRegister, error) {
	request := &requests.CreateUserRequest{
		FirstName:       req.Firstname,
		LastName:        req.Lastname,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	res, errResp := s.authService.Register(request)
	if errResp != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Registration failed: ",
		})
	}

	so := s.mapping.ToResponseRegister(res)

	return so, nil
}
