package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/auth"
	"MamangRust/paymentgatewaygrpc/pkg/hash"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type authService struct {
	auth    repository.UserRepository
	hash    *hash.Hashing
	token   auth.TokenManager
	logger  *logger.Logger
	mapping responsemapper.UserResponseMapper
}

func NewAuthService(auth repository.UserRepository, hash *hash.Hashing, token auth.TokenManager, logger *logger.Logger, mapping responsemapper.UserResponseMapper) *authService {
	return &authService{auth: auth, hash: hash, token: token, logger: logger, mapping: mapping}
}

func (s *authService) Register(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse) {
	_, err := s.auth.FindByEmail(request.Email)
	if err == nil {
		s.logger.Error("Email already exists", zap.String("email", request.Email))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Email already exists",
		}
	}

	passwordHash, err := s.hash.HashPassword(request.Password)
	if err != nil {
		s.logger.Error("Failed to hash password", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to hash password",
		}
	}
	request.Password = passwordHash

	res, err := s.auth.CreateUser(*request)
	if err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create user: " + err.Error(),
		}
	}

	s.logger.Debug("User registered successfully", zap.String("email", request.Email))

	so := s.mapping.ToUserResponse(*res)

	return so, nil
}

func (s *authService) Login(request *requests.AuthRequest) (*string, *response.ErrorResponse) {
	res, err := s.auth.FindByEmail(request.Email)
	if err != nil {
		s.logger.Error("Failed to get user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to get user: " + err.Error(),
		}
	}

	err = s.hash.ComparePassword(res.Password, request.Password)
	if err != nil {
		s.logger.Error("Failed to compare password", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Invalid password",
		}
	}

	token, err := s.createJwt(res.FirstName+" "+res.LastName, int32(res.ID))
	if err != nil {
		s.logger.Error("Failed to generate JWT token", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to generate token: " + err.Error(),
		}
	}

	s.logger.Debug("User logged in successfully", zap.String("email", request.Email))

	return &token, nil
}

func (s *authService) createJwt(fullname string, id int32) (string, error) {
	token, err := s.token.GenerateToken(fullname, id)

	if err != nil {
		return "", err
	}

	return token, nil
}
