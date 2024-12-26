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
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestFindAllTopups_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindAllTopupRequest{
		Page:     1,
		PageSize: 10,
		Search:   "test",
	}

	mockTopupData := []*response.TopupResponse{
		{
			ID:         1,
			CardNumber: "1234",
		},
	}

	mockPbResponse := []*pb.TopupResponse{
		{
			Id:         1,
			CardNumber: "1234",
		},
	}

	mockTopupService.EXPECT().FindAll(1, 10, "test").Return(mockTopupData, 1, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponsesTopup(mockTopupData).Return(mockPbResponse).Times(1)

	res, err := topupHandler.FindAllTopups(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully fetch topups", res.GetMessage())
	assert.Len(t, res.GetData(), 1)
	assert.Equal(t, int32(1), res.GetPagination().GetTotalPages())
	assert.Equal(t, int32(1), res.GetPagination().GetTotalRecords())
}

func TestFindAllTopups_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindAllTopupRequest{
		Page:     1,
		PageSize: 10,
		Search:   "test",
	}

	mockTopupService.EXPECT().FindAll(1, 10, "test").Return(nil, 0, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to fetch topups",
	}).Times(1)

	res, err := topupHandler.FindAllTopups(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to fetch topups")
}

func TestFindAllTopups_EmptyData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindAllTopupRequest{
		Page:     1,
		PageSize: 10,
		Search:   "test",
	}
	mockTopups := []*response.TopupResponse{}
	mockProtoTopups := []*pb.TopupResponse{}

	mockTopupService.EXPECT().FindAll(1, 10, "test").Return([]*response.TopupResponse{}, 0, nil).Times(1)

	mockProtoMapper.EXPECT().ToResponsesTopup(mockTopups).Return(mockProtoTopups).Times(1)

	res, err := topupHandler.FindAllTopups(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully fetch topups", res.GetMessage())
	assert.Len(t, res.GetData(), 0)
	assert.Equal(t, int32(0), res.GetPagination().GetTotalRecords())
}

func TestFindTopup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByIdTopupRequest{
		TopupId: 1,
	}

	mockTopupData := &response.TopupResponse{
		ID:          1,
		CardNumber:  "1234",
		TopupAmount: 1000,
	}

	mockPbResponse := &pb.TopupResponse{
		Id:          1,
		CardNumber:  "1234",
		TopupAmount: 1000,
	}

	mockTopupService.EXPECT().FindById(1).Return(mockTopupData, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponseTopup(mockTopupData).Return(mockPbResponse).Times(1)

	res, err := topupHandler.FindTopup(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully fetch topup", res.GetMessage())
	assert.Equal(t, mockPbResponse, res.GetData())
}

func TestFindTopup_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByIdTopupRequest{
		TopupId: 1,
	}

	mockTopupService.EXPECT().FindById(1).Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to fetch topup",
	}).Times(1)

	res, err := topupHandler.FindTopup(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to fetch topup")
}

func TestFindByCardNumberTopup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByCardNumberRequest{
		CardNumber: "1234",
	}

	mockTopupData := []*response.TopupResponse{
		{
			ID:         1,
			CardNumber: "1234",
		},
	}

	mockPbResponse := []*pb.TopupResponse{
		{
			Id:         1,
			CardNumber: "1234",
		},
	}

	mockTopupService.EXPECT().FindByCardNumber("1234").Return(mockTopupData, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponsesTopup(mockTopupData).Return(mockPbResponse).Times(1)

	res, err := topupHandler.FindByCardNumber(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully fetch topups", res.GetMessage())
	assert.Len(t, res.GetData(), 1)
}

func TestFindByCardNumberTopup_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByCardNumberRequest{
		CardNumber: "1234",
	}

	mockTopupService.EXPECT().FindByCardNumber("1234").Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to fetch topups",
	}).Times(1)

	res, err := topupHandler.FindByCardNumber(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to fetch topups")
}

func TestFindByActiveTopup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &emptypb.Empty{}

	mockTopupData := []*response.TopupResponse{
		{
			ID:         1,
			CardNumber: "1234",
		},
	}

	mockPbResponse := []*pb.TopupResponse{
		{
			Id:         1,
			CardNumber: "1234",
		},
	}

	mockTopupService.EXPECT().FindByActive().Return(mockTopupData, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponsesTopup(mockTopupData).Return(mockPbResponse).Times(1)

	res, err := topupHandler.FindByActive(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully fetch topups", res.GetMessage())
	assert.Len(t, res.GetData(), 1)
}

func TestFindByActiveTopup_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &emptypb.Empty{}

	mockTopupService.EXPECT().FindByActive().Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to fetch topups",
	}).Times(1)

	res, err := topupHandler.FindByActive(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to fetch topups")
}

func TestFindByActiveTopup_Empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &emptypb.Empty{}

	mockTopupService.EXPECT().FindByActive().Return([]*response.TopupResponse{}, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponsesTopup([]*response.TopupResponse{}).Return([]*pb.TopupResponse{}).Times(1)

	res, err := topupHandler.FindByActive(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully fetch topups", res.GetMessage())
	assert.Empty(t, res.GetData())
}

func TestFindByTrashedTopup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &emptypb.Empty{}

	mockTopupData := []*response.TopupResponse{
		{
			ID:         1,
			CardNumber: "1234",
		},
	}

	mockPbResponse := []*pb.TopupResponse{
		{
			Id:         1,
			CardNumber: "1234",
		},
	}

	mockTopupService.EXPECT().FindByTrashed().Return(mockTopupData, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponsesTopup(mockTopupData).Return(mockPbResponse).Times(1)

	res, err := topupHandler.FindByTrashed(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully fetch topups", res.GetMessage())
	assert.Len(t, res.GetData(), 1)
}

func TestFindByTrashed_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &emptypb.Empty{}

	mockTopupService.EXPECT().FindByTrashed().Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to fetch topups",
	}).Times(1)

	res, err := topupHandler.FindByTrashed(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to fetch topups")
}

func TestFindByTrashed_Empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &emptypb.Empty{}

	mockTopupService.EXPECT().FindByTrashed().Return([]*response.TopupResponse{}, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponsesTopup([]*response.TopupResponse{}).Return([]*pb.TopupResponse{}).Times(1)

	res, err := topupHandler.FindByTrashed(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully fetch topups", res.GetMessage())
	assert.Empty(t, res.GetData())
}

func TestCreateTopup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.CreateTopupRequest{
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Credit Card",
	}

	mockRequest := &requests.CreateTopupRequest{
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Credit Card",
	}

	mockResponse := &response.TopupResponse{
		ID:          1,
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Credit Card",
	}

	mockPbResponse := &pb.TopupResponse{
		Id:          1,
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Credit Card",
	}

	mockTopupService.EXPECT().CreateTopup(mockRequest).Return(mockResponse, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponseTopup(mockResponse).Return(mockPbResponse).Times(1)

	res, err := topupHandler.CreateTopup(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully created topup", res.GetMessage())
	assert.Equal(t, int32(1), res.GetData().GetId())
	assert.Equal(t, "1234", res.GetData().GetCardNumber())
	assert.Equal(t, "TOPUP123", res.GetData().GetTopupNo())
	assert.Equal(t, int32(5000), res.GetData().GetTopupAmount())
	assert.Equal(t, "Credit Card", res.GetData().GetTopupMethod())
}

func TestCreateTopup_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.CreateTopupRequest{
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Credit Card",
	}

	mockRequest := &requests.CreateTopupRequest{
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Credit Card",
	}

	mockTopupService.EXPECT().CreateTopup(mockRequest).Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to create topup",
	}).Times(1)

	res, err := topupHandler.CreateTopup(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to create topup")
}

func TestUpdateTopup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.UpdateTopupRequest{
		TopupId:     1,
		CardNumber:  "1234",
		TopupAmount: 6000,
		TopupMethod: "Debit Card",
	}

	mockRequest := &requests.UpdateTopupRequest{
		TopupID:     1,
		CardNumber:  "1234",
		TopupAmount: 6000,
		TopupMethod: "Debit Card",
	}

	mockResponse := &response.TopupResponse{
		ID:          1,
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 6000,
		TopupMethod: "Debit Card",
	}

	mockPbResponse := &pb.TopupResponse{
		Id:          1,
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 6000,
		TopupMethod: "Debit Card",
	}

	mockTopupService.EXPECT().UpdateTopup(mockRequest).Return(mockResponse, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponseTopup(mockResponse).Return(mockPbResponse).Times(1)

	res, err := topupHandler.UpdateTopup(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully updated topup", res.GetMessage())
	assert.Equal(t, int32(1), res.GetData().GetId())
	assert.Equal(t, "1234", res.GetData().GetCardNumber())
	assert.Equal(t, "TOPUP123", res.GetData().GetTopupNo())
	assert.Equal(t, int32(6000), res.GetData().GetTopupAmount())
	assert.Equal(t, "Debit Card", res.GetData().GetTopupMethod())
}

func TestUpdateTopup_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.UpdateTopupRequest{
		TopupId:     1,
		CardNumber:  "1234",
		TopupAmount: 6000,
		TopupMethod: "Debit Card",
	}

	mockRequest := &requests.UpdateTopupRequest{
		TopupID:     1,
		CardNumber:  "1234",
		TopupAmount: 6000,
		TopupMethod: "Debit Card",
	}

	mockTopupService.EXPECT().UpdateTopup(mockRequest).Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to update topup",
	}).Times(1)

	res, err := topupHandler.UpdateTopup(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to update topup")
}

func TestTrashedTopup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByIdTopupRequest{
		TopupId: 1,
	}

	mockResponse := &response.TopupResponse{
		ID:          1,
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Debit Card",
	}

	mockPbResponse := &pb.TopupResponse{
		Id:          1,
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Debit Card",
	}

	mockTopupService.EXPECT().TrashedTopup(1).Return(mockResponse, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponseTopup(mockResponse).Return(mockPbResponse).Times(1)

	res, err := topupHandler.TrashedTopup(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully trashed topup", res.GetMessage())
	assert.Equal(t, int32(1), res.GetData().GetId())
}

func TestTrashedTopup_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByIdTopupRequest{
		TopupId: 1,
	}

	mockTopupService.EXPECT().TrashedTopup(1).Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to trash topup",
	}).Times(1)

	res, err := topupHandler.TrashedTopup(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to trash topup")
}

func TestRestoreTopup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByIdTopupRequest{
		TopupId: 1,
	}

	mockResponse := &response.TopupResponse{
		ID:          1,
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Debit Card",
	}

	mockPbResponse := &pb.TopupResponse{
		Id:          1,
		CardNumber:  "1234",
		TopupNo:     "TOPUP123",
		TopupAmount: 5000,
		TopupMethod: "Debit Card",
	}

	mockTopupService.EXPECT().RestoreTopup(1).Return(mockResponse, nil).Times(1)
	mockProtoMapper.EXPECT().ToResponseTopup(mockResponse).Return(mockPbResponse).Times(1)

	res, err := topupHandler.RestoreTopup(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully restored topup", res.GetMessage())
	assert.Equal(t, int32(1), res.GetData().GetId())
}

func TestRestoreTopup_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByIdTopupRequest{
		TopupId: 1,
	}

	mockTopupService.EXPECT().RestoreTopup(1).Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to restore topup",
	}).Times(1)

	res, err := topupHandler.RestoreTopup(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to restore topup")
}

func TestDeleteTopupPermanent_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByIdTopupRequest{
		TopupId: 1,
	}

	mockResponse := &pb.ApiResponseTopupDelete{
		Status:  "success",
		Message: "Successfully deleted topup permanently",
	}

	mockTopupService.EXPECT().DeleteTopupPermanent(1).Return(mockResponse, nil).Times(1)

	res, err := topupHandler.DeleteTopupPermanent(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "success", res.GetStatus())
	assert.Equal(t, "Successfully deleted topup permanently", res.GetMessage())
}

func TestDeleteTopupPermanent_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTopupService := mock_service.NewMockTopupService(ctrl)
	mockProtoMapper := mock_protomapper.NewMockTopupProtoMapper(ctrl)
	topupHandler := gapi.NewTopupHandleGrpc(mockTopupService, mockProtoMapper)

	req := &pb.FindByIdTopupRequest{
		TopupId: 1,
	}

	mockTopupService.EXPECT().DeleteTopupPermanent(1).Return(nil, &response.ErrorResponse{
		Status:  "error",
		Message: "Failed to delete topup permanently",
	}).Times(1)

	res, err := topupHandler.DeleteTopupPermanent(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to delete topup permanently")
}
