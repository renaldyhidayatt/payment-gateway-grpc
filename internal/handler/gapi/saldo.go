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

type saldoHandleGrpc struct {
	pb.UnimplementedSaldoServiceServer
	saldoService service.SaldoService
	mapping      protomapper.SaldoProtoMapper
}

func NewSaldoHandleGrpc(saldo service.SaldoService, mapping protomapper.SaldoProtoMapper) *saldoHandleGrpc {
	return &saldoHandleGrpc{
		saldoService: saldo,
		mapping:      mapping,
	}
}

func (s *saldoHandleGrpc) FindAllSaldo(ctx context.Context, req *pb.FindAllSaldoRequest) (*pb.ApiResponsePaginationSaldo, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.saldoService.FindAll(page, pageSize, search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch saldo records: " + err.Message,
		})
	}

	totalPages := (totalRecords + pageSize - 1) / pageSize

	so := s.mapping.ToResponsesSaldo(merchants)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationSaldo{
		Status:     "success",
		Message:    "Successfully fetched saldo record",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *saldoHandleGrpc) FindByIdSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error) {
	if req.GetSaldoId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID",
		})
	}

	id := req.GetSaldoId()

	saldo, err := s.saldoService.FindById(int(id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch saldo record: " + err.Message,
		})
	}

	so := s.mapping.ToResponseSaldo(saldo)

	return &pb.ApiResponseSaldo{Status: "success", Message: "Successfully fetched saldo record", Data: so}, nil
}

func (s *saldoHandleGrpc) FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponseSaldo, error) {
	cardNumber := req.GetCardNumber()
	saldo, err := s.saldoService.FindByCardNumber(cardNumber)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch saldo record: " + err.Message,
		})
	}

	so := s.mapping.ToResponseSaldo(saldo)

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully fetched saldo record",
		Data:    so,
	}, nil
}

func (s *saldoHandleGrpc) FindByActive(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesSaldo, error) {
	res, err := s.saldoService.FindByActive()

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found: " + err.Message,
		})
	}

	return &pb.ApiResponsesSaldo{
		Status:  "success",
		Message: "Successfully fetched saldo record",
		Data:    s.mapping.ToResponsesSaldo(res),
	}, nil
}

func (s *saldoHandleGrpc) FindByTrashed(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesSaldo, error) {
	res, err := s.saldoService.FindByTrashed()

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found: " + err.Message,
		})
	}

	return &pb.ApiResponsesSaldo{
		Status:  "success",
		Message: "Successfully fetched saldo record",
		Data:    s.mapping.ToResponsesSaldo(res),
	}, nil
}

func (s *saldoHandleGrpc) CreateSaldo(ctx context.Context, req *pb.CreateSaldoRequest) (*pb.ApiResponseSaldo, error) {
	request := requests.CreateSaldoRequest{
		CardNumber:   req.GetCardNumber(),
		TotalBalance: int(req.GetTotalBalance()),
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create saldo record: ",
		})
	}

	saldo, err := s.saldoService.CreateSaldo(&request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create saldo record: " + err.Message,
		})
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully created saldo record",
		Data:    s.mapping.ToResponseSaldo(saldo),
	}, nil

}

func (s *saldoHandleGrpc) UpdateSaldo(ctx context.Context, req *pb.UpdateSaldoRequest) (*pb.ApiResponseSaldo, error) {
	if req.GetSaldoId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID",
		})
	}

	request := requests.UpdateSaldoRequest{
		SaldoID:      int(req.GetSaldoId()),
		CardNumber:   req.GetCardNumber(),
		TotalBalance: int(req.GetTotalBalance()),
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo record: ",
		})
	}

	saldo, err := s.saldoService.UpdateSaldo(&request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo record: " + err.Message,
		})
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully updated saldo record",
		Data:    s.mapping.ToResponseSaldo(saldo),
	}, nil
}

func (s *saldoHandleGrpc) TrashSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error) {
	if req.GetSaldoId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID",
		})
	}

	saldo, err := s.saldoService.TrashSaldo(int(req.GetSaldoId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash saldo record: " + err.Message,
		})
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully trashed saldo record",
		Data:    s.mapping.ToResponseSaldo(saldo),
	}, nil
}

func (s *saldoHandleGrpc) RestoreSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error) {
	if req.GetSaldoId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID",
		})
	}

	saldo, err := s.saldoService.RestoreSaldo(int(req.GetSaldoId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore saldo record: " + err.Message,
		})
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully restored saldo record",
		Data:    s.mapping.ToResponseSaldo(saldo),
	}, nil
}

func (s *saldoHandleGrpc) DeleteSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldoDelete, error) {
	if req.GetSaldoId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID",
		})
	}

	_, err := s.saldoService.DeleteSaldoPermanent(int(req.GetSaldoId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete saldo record: " + err.Message,
		})
	}

	return &pb.ApiResponseSaldoDelete{
		Status:  "success",
		Message: "Successfully deleted saldo record",
	}, nil
}
