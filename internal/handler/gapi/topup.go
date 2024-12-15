package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	protomapper "MamangRust/paymentgatewaygrpc/internal/mapper/proto"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type topupHandleGrpc struct {
	pb.UnimplementedTopupServiceServer
	topupService service.TopupService
	mapping      protomapper.TopupProtoMapper
}

func NewTopupHandleGrpc(topup service.TopupService, mapping protomapper.TopupProtoMapper) *topupHandleGrpc {
	return &topupHandleGrpc{
		topupService: topup,
		mapping:      mapping,
	}
}

func (s *topupHandleGrpc) FindAllTopups(ctx context.Context, req *pb.FindAllTopupRequest) (*pb.ApiResponsePaginationTopup, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	topups, totalRecords, err := s.topupService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topups: " + err.Message,
		})
	}

	totalPages := (totalRecords + pageSize - 1) / pageSize

	so := s.mapping.ToResponsesTopup(topups)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationTopup{
		Status:     "success",
		Message:    "Successfully fetch topups",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *topupHandleGrpc) FindTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.TopupResponse, error) {
	id := req.GetTopupId()

	topup, err := s.topupService.FindById(int(id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topup: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopup(topup)

	return so, nil
}

func (s *topupHandleGrpc) FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponsesTopup, error) {
	cardNumber := req.GetCardNumber()

	topups, err := s.topupService.FindByCardNumber(cardNumber)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topups: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesTopup(topups)

	return &pb.ApiResponsesTopup{
		Status:  "success",
		Message: "Successfully fetch topups",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindByActive(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesTopup, error) {
	topups, err := s.topupService.FindByActive()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topups: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesTopup(topups)

	return &pb.ApiResponsesTopup{
		Status:  "success",
		Message: "Successfully fetch topups",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindByTrashed(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesTopup, error) {
	topups, err := s.topupService.FindByTrashed()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topups: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesTopup(topups)

	return &pb.ApiResponsesTopup{
		Status:  "success",
		Message: "Successfully fetch topups",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) CreateTopup(ctx context.Context, req *pb.CreateTopupRequest) (*pb.ApiResponseTopup, error) {
	request := requests.CreateTopupRequest{
		CardNumber:  req.GetCardNumber(),
		TopupNo:     req.GetTopupNo(),
		TopupAmount: int(req.GetTopupAmount()),
		TopupMethod: req.GetTopupMethod(),
	}

	res, err := s.topupService.CreateTopup(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create topup: " + err.Message,
		})
	}

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully created topup",
		Data:    s.mapping.ToResponseTopup(res),
	}, nil
}

func (s *topupHandleGrpc) UpdateTopup(ctx context.Context, req *pb.UpdateTopupRequest) (*pb.ApiResponseTopup, error) {
	request := requests.UpdateTopupRequest{
		TopupID:     int(req.GetTopupId()),
		CardNumber:  req.GetCardNumber(),
		TopupAmount: int(req.GetTopupAmount()),
		TopupMethod: req.GetTopupMethod(),
	}

	res, err := s.topupService.UpdateTopup(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update topup: " + err.Message,
		})
	}

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully updated topup",
		Data:    s.mapping.ToResponseTopup(res),
	}, nil
}

func (s *topupHandleGrpc) TrashedTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopup, error) {
	res, err := s.topupService.TrashedTopup(int(req.GetTopupId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash topup: " + err.Message,
		})
	}

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully trashed topup",
		Data:    s.mapping.ToResponseTopup(res),
	}, nil
}

func (s *topupHandleGrpc) RestoreTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopup, error) {
	res, err := s.topupService.RestoreTopup(int(req.GetTopupId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore topup: " + err.Message,
		})
	}

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully restored topup",
		Data:    s.mapping.ToResponseTopup(res),
	}, nil
}

func (s *topupHandleGrpc) DeleteTopupPermanent(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopupDelete, error) {
	_, err := s.topupService.DeleteTopupPermanent(int(req.GetTopupId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete topup permanently: " + err.Message,
		})
	}

	return &pb.ApiResponseTopupDelete{
		Status:  "success",
		Message: "Successfully deleted topup permanently",
	}, nil
}
