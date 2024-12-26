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

type merchantHandleGrpc struct {
	pb.UnimplementedMerchantServiceServer
	merchantService service.MerchantService
	mapping         protomapper.MerchantProtoMapper
}

func NewMerchantHandleGrpc(merchantService service.MerchantService, mapping protomapper.MerchantProtoMapper) *merchantHandleGrpc {
	return &merchantHandleGrpc{merchantService: merchantService, mapping: mapping}
}

func (s *merchantHandleGrpc) FindAll(ctx context.Context, req *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchant, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.merchantService.FindAll(page, pageSize, search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card records: " + err.Message,
		})
	}

	totalPages := (totalRecords + pageSize - 1) / pageSize

	so := s.mapping.ToResponsesMerchant(merchants)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationMerchant{
		Status:     "success",
		Message:    "Successfully fetched merchant record",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *merchantHandleGrpc) FindById(ctx context.Context, req *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error) {
	merchant, err := s.merchantService.FindById(int(req.GetMerchantId()))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found: ",
		})
	}

	so := s.mapping.ToResponseMerchant(merchant)

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully fetched merchant record",
		Data:    so,
	}, nil
}

func (s *merchantHandleGrpc) FindByApiKey(ctx context.Context, req *pb.FindByApiKeyRequest) (*pb.ApiResponseMerchant, error) {
	merchant, err := s.merchantService.FindByApiKey(req.ApiKey)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found: ",
		})
	}

	so := s.mapping.ToResponseMerchant(merchant)

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully fetched merchant record",
		Data:    so,
	}, nil
}

func (s *merchantHandleGrpc) FindByMerchantUserId(ctx context.Context, req *pb.FindByMerchantUserIdRequest) (*pb.ApiResponsesMerchant, error) {
	res, err := s.merchantService.FindByMerchantUserId(int(req.GetUserId()))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesMerchant(res)

	return &pb.ApiResponsesMerchant{
		Status:  "success",
		Message: "Successfully fetched merchant record",
		Data:    so,
	}, nil
}

func (s *merchantHandleGrpc) FindByActive(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesMerchant, error) {
	res, err := s.merchantService.FindByActive()

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesMerchant(res)

	return &pb.ApiResponsesMerchant{
		Status:  "success",
		Message: "Successfully fetched merchant record",
		Data:    so,
	}, nil
}

func (s *merchantHandleGrpc) FindByTrashed(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesMerchant, error) {
	res, err := s.merchantService.FindByTrashed()

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesMerchant(res)

	return &pb.ApiResponsesMerchant{
		Status:  "success",
		Message: "Successfully fetched merchant record",
		Data:    so,
	}, nil
}

func (s *merchantHandleGrpc) CreateMerchant(ctx context.Context, req *pb.CreateMerchantRequest) (*pb.ApiResponseMerchant, error) {
	request := requests.CreateMerchantRequest{
		Name:   req.GetName(),
		UserID: int(req.GetUserId()),
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create merchant: " + err.Error(),
		})
	}

	merchant, err := s.merchantService.CreateMerchant(&request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create merchant: " + err.Message,
		})
	}

	so := s.mapping.ToResponseMerchant(merchant)

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully created merchant",
		Data:    so,
	}, nil

}

func (s *merchantHandleGrpc) UpdateMerchant(ctx context.Context, req *pb.UpdateMerchantRequest) (*pb.ApiResponseMerchant, error) {
	request := requests.UpdateMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Name:       req.GetName(),
		UserID:     int(req.GetUserId()),
		Status:     req.GetStatus(),
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update merchant: " + err.Error(),
		})
	}

	merchant, err := s.merchantService.UpdateMerchant(&request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update merchant: " + err.Message,
		})
	}

	so := s.mapping.ToResponseMerchant(merchant)

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully updated merchant",
		Data:    so,
	}, nil
}

func (s *merchantHandleGrpc) TrashedMerchant(ctx context.Context, req *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error) {
	merchant, err := s.merchantService.TrashedMerchant(int(req.GetMerchantId()))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponseMerchant(merchant)

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully trashed merchant",
		Data:    so,
	}, nil
}

func (s *merchantHandleGrpc) RestoreMerchant(ctx context.Context, req *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error) {
	merchant, err := s.merchantService.RestoreMerchant(int(req.GetMerchantId()))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponseMerchant(merchant)

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully restored merchant",
		Data:    so,
	}, nil
}

func (s *merchantHandleGrpc) DeleteMerchant(ctx context.Context, req *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchatDelete, error) {
	_, err := s.merchantService.DeleteMerchantPermanent(int(req.GetMerchantId()))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found: " + err.Message,
		})
	}

	return &pb.ApiResponseMerchatDelete{
		Status:  "success",
		Message: "Successfully deleted merchant",
	}, nil
}
