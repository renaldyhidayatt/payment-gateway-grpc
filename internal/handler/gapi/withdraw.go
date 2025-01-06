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
		Status:  "success",
		Message: "Successfully fetched withdraws",
		Data:    so,
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
