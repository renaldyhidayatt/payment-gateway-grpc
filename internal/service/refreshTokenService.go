package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type refreshTokenService struct {
	refreshTokenRepository repository.RefreshTokenRepository
	logger                 logger.LoggerInterface
	mapping                responsemapper.RefreshTokenResponseMapper
}

func NewRefreshTokenService(refreshTokenRepository repository.RefreshTokenRepository, logger logger.LoggerInterface, mapping responsemapper.RefreshTokenResponseMapper) *refreshTokenService {
	return &refreshTokenService{
		refreshTokenRepository: refreshTokenRepository,
		logger:                 logger,
		mapping:                mapping,
	}
}

func (r *refreshTokenService) FindByToken(token string) (*response.RefreshTokenResponse, *response.ErrorResponse) {
	refreshToken, err := r.refreshTokenRepository.FindByToken(token)

	if err != nil {
		r.logger.Error("Failed to find refresh token", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find refresh token: ",
		}
	}

	return r.mapping.ToRefreshTokenResponse(refreshToken), nil
}

func (r *refreshTokenService) FindByUserId(user_id int) (*response.RefreshTokenResponse, *response.ErrorResponse) {
	refreshToken, err := r.refreshTokenRepository.FindByUserId(user_id)

	if err != nil {
		r.logger.Error("Failed to find refresh token", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find refresh token: ",
		}
	}

	return r.mapping.ToRefreshTokenResponse(refreshToken), nil
}

func (r *refreshTokenService) UpdateRefreshToken(req *requests.UpdateRefreshToken) (*response.RefreshTokenResponse, *response.ErrorResponse) {
	refreshToken, err := r.refreshTokenRepository.UpdateRefreshToken(req)

	if err != nil {
		r.logger.Error("Failed to update refresh token", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update refresh token: ",
		}
	}

	return r.mapping.ToRefreshTokenResponse(refreshToken), nil
}

func (r *refreshTokenService) DeleteRefreshToken(token string) *response.ErrorResponse {
	err := r.refreshTokenRepository.DeleteRefreshToken(token)

	if err != nil {
		r.logger.Error("Failed to delete refresh token", zap.Error(err))
		return &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete refresh token: ",
		}
	}

	return nil
}
