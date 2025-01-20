package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	protomapper "MamangRust/paymentgatewaygrpc/internal/mapper/proto"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"context"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type withdrawHandleGrpc struct {
	pb.UnimplementedWithdrawServiceServer
	withdrawService service.WithdrawService
	mapping         protomapper.WithdrawalProtoMapper
}

func NewWithdrawHandleGrpc(withdraw service.WithdrawService, mapping protomapper.WithdrawalProtoMapper) *withdrawHandleGrpc {
	return &withdrawHandleGrpc{
		withdrawService: withdraw,
		mapping:         mapping,
	}
}

func (w *withdrawHandleGrpc) FindAllWithdraw(ctx context.Context, req *pb.FindAllWithdrawRequest) (*pb.ApiResponsePaginationWithdraw, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	withdraws, totalRecords, err := w.withdrawService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraws: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := w.mapping.ToResponsesWithdrawal(withdraws)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationWithdraw{
		Status:     "success",
		Message:    "Withdraws fetched successfully",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (w *withdrawHandleGrpc) FindByIdWithdraw(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdraw, error) {
	withdraw, err := w.withdrawService.FindById(int(req.GetWithdrawId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw: " + err.Message,
		})
	}

	so := w.mapping.ToResponseWithdrawal(withdraw)

	return &pb.ApiResponseWithdraw{
		Status:  "success",
		Message: "Successfully fetched withdraw",
		Data:    so,
	}, nil
}

func (s *withdrawHandleGrpc) FindMonthlyWithdrawStatusSuccess(ctx context.Context, req *pb.FindMonthlyWithdrawStatus) (*pb.ApiResponseWithdrawMonthStatusSuccess, error) {
	if req.GetYear() <= 0 || req.GetMonth() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year or month",
		})
	}

	year := req.GetYear()
	month := req.GetMonth()

	records, errResponse := s.withdrawService.FindMonthWithdrawStatusSuccess(int(year), int(month))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly Withdraw status success: " + errResponse.Message,
		})
	}

	so := s.mapping.ToResponsesWithdrawMonthStatusSuccess(records)

	return &pb.ApiResponseWithdrawMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly Withdraw status success",
		Data:    so,
	}, nil
}

func (s *withdrawHandleGrpc) FindYearlyWithdrawStatusSuccess(ctx context.Context, req *pb.FindYearWithdraw) (*pb.ApiResponseWithdrawYearStatusSuccess, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	year := req.GetYear()

	records, errResponse := s.withdrawService.FindYearlyWithdrawStatusSuccess(int(year))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly Withdraw status success: " + errResponse.Message,
		})
	}

	so := s.mapping.ToWithdrawResponsesYearStatusSuccess(records)

	return &pb.ApiResponseWithdrawYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly Withdraw status success",
		Data:    so,
	}, nil
}

func (s *withdrawHandleGrpc) FindMonthlyWithdrawStatusFailed(ctx context.Context, req *pb.FindMonthlyWithdrawStatus) (*pb.ApiResponseWithdrawMonthStatusFailed, error) {
	if req.GetYear() <= 0 || req.GetMonth() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year or month",
		})
	}

	year := req.GetYear()
	month := req.GetMonth()

	records, errResponse := s.withdrawService.FindMonthWithdrawStatusFailed(int(year), int(month))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly Withdraw status Failed: " + errResponse.Message,
		})
	}

	so := s.mapping.ToResponsesWithdrawMonthStatusFailed(records)

	return &pb.ApiResponseWithdrawMonthStatusFailed{
		Status:  "Failed",
		Message: "Failedfully fetched monthly Withdraw status Failed",
		Data:    so,
	}, nil
}

func (s *withdrawHandleGrpc) FindYearlyWithdrawStatusFailed(ctx context.Context, req *pb.FindYearWithdraw) (*pb.ApiResponseWithdrawYearStatusFailed, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	year := req.GetYear()

	records, errResponse := s.withdrawService.FindYearlyWithdrawStatusFailed(int(year))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly Withdraw status Failed: " + errResponse.Message,
		})
	}

	so := s.mapping.ToWithdrawResponsesYearStatusFailed(records)

	return &pb.ApiResponseWithdrawYearStatusFailed{
		Status:  "Failed",
		Message: "Failedfully fetched yearly Withdraw status Failed",
		Data:    so,
	}, nil
}

func (w *withdrawHandleGrpc) FindMonthlyWithdraws(ctx context.Context, req *pb.FindYearWithdraw) (*pb.ApiResponseWithdrawMonthAmount, error) {
	withdraws, err := w.withdrawService.FindMonthlyWithdraws(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly withdraws: " + err.Message,
		})
	}

	so := w.mapping.ToResponseWithdrawMonthlyAmounts(withdraws)

	return &pb.ApiResponseWithdrawMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly withdraws",
		Data:    so,
	}, nil
}

func (w *withdrawHandleGrpc) FindYearlyWithdraws(ctx context.Context, req *pb.FindYearWithdraw) (*pb.ApiResponseWithdrawYearAmount, error) {
	withdraws, err := w.withdrawService.FindYearlyWithdraws(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly withdraws: " + err.Message,
		})
	}

	so := w.mapping.ToResponseWithdrawYearlyAmounts(withdraws)

	return &pb.ApiResponseWithdrawYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly withdraws",
		Data:    so,
	}, nil
}

func (w *withdrawHandleGrpc) FindMonthlyWithdrawsByCardNumber(ctx context.Context, req *pb.FindYearWithdrawCardNumber) (*pb.ApiResponseWithdrawMonthAmount, error) {
	withdraws, err := w.withdrawService.FindMonthlyWithdrawsByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly withdraws by card number: " + err.Message,
		})
	}

	so := w.mapping.ToResponseWithdrawMonthlyAmounts(withdraws)

	return &pb.ApiResponseWithdrawMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly withdraws by card number",
		Data:    so,
	}, nil
}

func (w *withdrawHandleGrpc) FindYearlyWithdrawsByCardNumber(ctx context.Context, req *pb.FindYearWithdrawCardNumber) (*pb.ApiResponseWithdrawYearAmount, error) {
	withdraws, err := w.withdrawService.FindYearlyWithdrawsByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly withdraws by card number: " + err.Message,
		})
	}

	so := w.mapping.ToResponseWithdrawYearlyAmounts(withdraws)

	return &pb.ApiResponseWithdrawYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly withdraws by card number",
		Data:    so,
	}, nil
}

func (w *withdrawHandleGrpc) FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponsesWithdraw, error) {

	withdraws, err := w.withdrawService.FindByCardNumber(req.GetCardNumber())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraws: " + err.Message,
		})
	}

	so := w.mapping.ToResponsesWithdrawal(withdraws)

	return &pb.ApiResponsesWithdraw{
		Status:  "success",
		Message: "Successfully fetched withdraws",
		Data:    so,
	}, nil
}

func (w *withdrawHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllWithdrawRequest) (*pb.ApiResponsePaginationWithdrawDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := w.withdrawService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraws: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))
	so := w.mapping.ToResponsesWithdrawalDeleteAt(res)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationWithdrawDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched withdraws",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (w *withdrawHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllWithdrawRequest) (*pb.ApiResponsePaginationWithdrawDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := w.withdrawService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraws: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := w.mapping.ToResponsesWithdrawalDeleteAt(res)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationWithdrawDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched withdraws",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (w *withdrawHandleGrpc) CreateWithdraw(ctx context.Context, req *pb.CreateWithdrawRequest) (*pb.ApiResponseWithdraw, error) {
	request := &requests.CreateWithdrawRequest{
		CardNumber:     req.CardNumber,
		WithdrawAmount: int(req.WithdrawAmount),
		WithdrawTime:   req.WithdrawTime.AsTime(),
	}

	withdraw, err := w.withdrawService.Create(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create withdraw: " + err.Message,
		})
	}

	return &pb.ApiResponseWithdraw{
		Status:  "success",
		Message: "Successfully created withdraw",
		Data:    w.mapping.ToResponseWithdrawal(withdraw),
	}, nil

}

func (w *withdrawHandleGrpc) UpdateWithdraw(ctx context.Context, req *pb.UpdateWithdrawRequest) (*pb.ApiResponseWithdraw, error) {
	request := &requests.UpdateWithdrawRequest{
		WithdrawID:     int(req.WithdrawId),
		CardNumber:     req.CardNumber,
		WithdrawAmount: int(req.WithdrawAmount),
		WithdrawTime:   req.WithdrawTime.AsTime(),
	}

	withdraw, err := w.withdrawService.Update(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update withdraw: " + err.Message,
		})
	}

	return &pb.ApiResponseWithdraw{
		Status:  "success",
		Message: "Successfully updated withdraw",
		Data:    w.mapping.ToResponseWithdrawal(withdraw),
	}, nil
}

func (w *withdrawHandleGrpc) TrashedWithdraw(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdraw, error) {
	if req.WithdrawId <= 0 {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid withdraw id",
		})
	}

	withdraw, err := w.withdrawService.TrashedWithdraw(int(req.WithdrawId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw: " + err.Message,
		})
	}

	return &pb.ApiResponseWithdraw{
		Status:  "success",
		Message: "Successfully trashed withdraw",
		Data:    w.mapping.ToResponseWithdrawal(withdraw),
	}, nil
}

func (w *withdrawHandleGrpc) RestoreWithdraw(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdraw, error) {
	if req.WithdrawId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid withdraw id",
		})
	}

	withdraw, err := w.withdrawService.RestoreWithdraw(int(req.WithdrawId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw: " + err.Message,
		})
	}

	return &pb.ApiResponseWithdraw{
		Status:  "success",
		Message: "Successfully restored withdraw",
		Data:    w.mapping.ToResponseWithdrawal(withdraw),
	}, nil
}

func (w *withdrawHandleGrpc) DeleteWithdrawPermanent(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdrawDelete, error) {
	if req.WithdrawId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "invalid withdraw id",
		})
	}

	_, err := w.withdrawService.DeleteWithdrawPermanent(int(req.WithdrawId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw: " + err.Message,
		})
	}

	return &pb.ApiResponseWithdrawDelete{
		Status:  "success",
		Message: "Successfully deleted withdraw permanently",
	}, nil
}

func (s *withdrawHandleGrpc) RestoreAllWithdraw(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseWithdrawAll, error) {
	_, err := s.withdrawService.RestoreAllWithdraw()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all withdraw: ",
		})
	}

	return &pb.ApiResponseWithdrawAll{
		Status:  "success",
		Message: "Successfully restore all withdraw",
	}, nil
}

func (s *withdrawHandleGrpc) DeleteAllWithdrawPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseWithdrawAll, error) {
	_, err := s.withdrawService.DeleteAllWithdrawPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete withdraw permanent: ",
		})
	}

	return &pb.ApiResponseWithdrawAll{
		Status:  "success",
		Message: "Successfully delete withdraw permanent",
	}, nil
}
