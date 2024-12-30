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

func TestRegisterUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mock_service.NewMockAuthService(ctrl)
	mockMapper := mock_protomapper.NewMockAuthProtoMapper(ctrl)

	request := &requests.CreateUserRequest{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	registerRequest := &pb.RegisterRequest{
		Firstname:       request.FirstName,
		Lastname:        request.LastName,
		Email:           request.Email,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}

	expectedResponse := &response.UserResponse{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@example.com",
	}

	myexpected := &pb.ApiResponseRegister{
		Status:  "success",
		Message: "User registered successfully",
		User: &pb.UserResponse{
			Id:        int32(expectedResponse.ID),
			Firstname: expectedResponse.FirstName,
			Lastname:  expectedResponse.LastName,
			Email:     expectedResponse.Email,
		},
	}

	mockAuthService.EXPECT().Register(&requests.CreateUserRequest{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	}).Return(expectedResponse, nil)

	mockMapper.EXPECT().ToResponseRegister(expectedResponse).Return(myexpected)

	handler := gapi.NewAuthHandleGrpc(mockAuthService, mockMapper)

	resp, err := handler.RegisterUser(context.Background(), registerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "success", resp.Status)
	assert.Equal(t, "User registered successfully", resp.Message)

	assert.Equal(t, int32(expectedResponse.ID), resp.User.Id)
	assert.Equal(t, expectedResponse.FirstName, resp.User.Firstname)
	assert.Equal(t, expectedResponse.LastName, resp.User.Lastname)
	assert.Equal(t, expectedResponse.Email, resp.User.Email)

}

func TestRegisterUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mock_service.NewMockAuthService(ctrl)
	mockMapper := mock_protomapper.NewMockAuthProtoMapper(ctrl)

	request := &requests.CreateUserRequest{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	registerRequest := &pb.RegisterRequest{
		Firstname:       request.FirstName,
		Lastname:        request.LastName,
		Email:           request.Email,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}

	mockAuthService.EXPECT().Register(&requests.CreateUserRequest{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "test@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	}).Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "registration failed",
	})

	handler := gapi.NewAuthHandleGrpc(mockAuthService, mockMapper)

	resp, err := handler.RegisterUser(context.Background(), registerRequest)

	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "Registration failed")

}
