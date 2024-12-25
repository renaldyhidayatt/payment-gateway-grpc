package test

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/handler/gapi"
	mock_protomapper "MamangRust/paymentgatewaygrpc/internal/mapper/proto/mocks"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	mock_service "MamangRust/paymentgatewaygrpc/internal/service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLoginUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mock_service.NewMockAuthService(ctrl)
	mockMapper := mock_protomapper.NewMockAuthProtoMapper(ctrl)

	loginRequest := &pb.LoginRequest{Email: "test@example.com", Password: "password123"}
	loginRequestService := &requests.AuthRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	loginResponse := "some-jwt-token"

	mockAuthService.EXPECT().Login(loginRequestService).Return(&loginResponse, nil)

	handler := gapi.NewAuthHandleGrpc(mockAuthService, mockMapper)

	resp, err := handler.LoginUser(context.Background(), loginRequest)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "success", resp.Status)
	assert.Equal(t, "Login successful", resp.Message)
	assert.Equal(t, "some-jwt-token", resp.Token)
}

func TestLoginUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mock_service.NewMockAuthService(ctrl)
	mockMapper := mock_protomapper.NewMockAuthProtoMapper(ctrl)

	loginRequest := &pb.LoginRequest{Email: "test@example.com", Password: "wrong-password"}
	loginRequestService := &requests.AuthRequest{
		Email:    "test@example.com",
		Password: "wrong-password",
	}

	mockAuthService.EXPECT().Login(loginRequestService).Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "invalid credentials",
	})

	handler := gapi.NewAuthHandleGrpc(mockAuthService, mockMapper)

	resp, errRes := handler.LoginUser(context.Background(), loginRequest)

	assert.NotNil(t, errRes)
	assert.Nil(t, resp)
	assert.Contains(t, errRes.Error(), "Login failed")
}
