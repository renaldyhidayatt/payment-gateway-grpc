package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/auth"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"MamangRust/paymentgatewaygrpc/pkg/hash"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"MamangRust/paymentgatewaygrpc/pkg/randomvcc"
	"errors"

	"go.uber.org/zap"
)

type authService struct {
	auth   repository.UserRepository
	hash   hash.Hashing
	token  auth.TokenManager
	logger logger.Logger
}

func NewAuthService(auth repository.UserRepository, hash hash.Hashing, token auth.TokenManager, logger logger.Logger) *authService {
	return &authService{auth: auth, hash: hash, token: token, logger: logger}
}

func (s *authService) Register(request *requests.CreateUserRequest) (*db.User, error) {
	_, err := s.auth.FindByEmail(request.Email)

	if err == nil {
		s.logger.Error("Email already exists", zap.String("email", request.Email))
		return nil, errors.New("failed email already exist")
	}

	passwordHash, err := s.hash.HashPassword(request.Password)

	if err != nil {
		s.logger.Error("failed to hash password", zap.Error(err))
		return nil, err
	}

	randomVCC, err := randomvcc.RandomVCC()

	if err != nil {
		s.logger.Error("failed to generate random VCC:", zap.Error(err))
		return nil, errors.New("failed generate random vcc")
	}

	user := db.CreateUserParams{
		Firstname:   request.FirstName,
		Lastname:    request.LastName,
		Email:       request.Email,
		Password:    passwordHash,
		NocTransfer: randomVCC,
	}

	res, err := s.auth.Create(&user)

	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, errors.New("failed create user :" + err.Error())
	}

	s.logger.Info("User Registered successfully " + request.Email)
	return res, nil

}

func (s *authService) Login(request *requests.AuthLoginRequest) (*requests.JWTToken, error) {
	res, err := s.auth.FindByEmail(request.Email)

	if err != nil {
		s.logger.Error("failed to get user", zap.Error(err))
		return nil, errors.New("failed get user " + err.Error())
	}

	err = s.hash.ComparePassword(res.Password, request.Password)

	if err != nil {

		s.logger.Error("failed to compare password", zap.Error(err))

		return nil, errors.New("failed compare password " + err.Error())
	}

	token, err := s.createJwt(res.Firstname+" "+res.Lastname, res.UserID)

	if err != nil {
		s.logger.Error("failed to generate jwt token", zap.Error(err))
		return nil, err
	}

	s.logger.Info("User logged in successfully " + request.Email)

	return &requests.JWTToken{
		Token: token,
	}, nil
}

func (s *authService) createJwt(fullname string, id int32) (string, error) {
	token, err := s.token.GenerateToken(fullname, id)

	if err != nil {
		return "", err
	}

	return token, nil
}
