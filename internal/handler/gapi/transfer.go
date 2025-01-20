package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	protomapper "MamangRust/paymentgatewaygrpc/internal/mapper/proto"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"context"
	"fmt"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transferHandleGrpc struct {
	pb.UnimplementedTransferServiceServer
	transferService service.TransferService
	mapping         protomapper.TransferProtoMapper
}

func NewTransferHandleGrpc(transferService service.TransferService,
	mapping protomapper.TransferProtoMapper) *transferHandleGrpc {
	return &transferHandleGrpc{
		transferService: transferService,
		mapping:         mapping,
	}
}

func (s *transferHandleGrpc) FindAllTransfer(ctx context.Context, request *pb.FindAllTransferRequest) (*pb.ApiResponsePaginationTransfer, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.transferService.FindAll(page, pageSize, search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfer records: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := s.mapping.ToResponsesTransfer(merchants)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationTransfer{
		Status:     "success",
		Message:    "Successfully fetch transfer records",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *transferHandleGrpc) FindTransferById(ctx context.Context, request *pb.FindByIdTransferRequest) (*pb.ApiResponseTransfer, error) {
	if request.GetTransferId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid transfer id",
		})
	}

	transfer, err := s.transferService.FindById(int(request.GetTransferId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfer record: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTransfer(transfer)

	return &pb.ApiResponseTransfer{
		Status:  "success",
		Message: "Successfully fetch transfer record",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindMonthlyTransferStatusSuccess(ctx context.Context, req *pb.FindMonthlyTransferStatus) (*pb.ApiResponseTransferMonthStatusSuccess, error) {
	if req.GetYear() <= 0 || req.GetMonth() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year or month",
		})
	}

	year := req.GetYear()
	month := req.GetMonth()

	records, errResponse := s.transferService.FindMonthTransferStatusSuccess(int(year), int(month))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly Transfer status success: " + errResponse.Message,
		})
	}

	so := s.mapping.ToResponsesTransferMonthStatusSuccess(records)

	return &pb.ApiResponseTransferMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly Transfer status success",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindYearlyTransferStatusSuccess(ctx context.Context, req *pb.FindYearTransfer) (*pb.ApiResponseTransferYearStatusSuccess, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	year := req.GetYear()

	records, errResponse := s.transferService.FindYearlyTransferStatusSuccess(int(year))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly Transfer status success: " + errResponse.Message,
		})
	}

	so := s.mapping.ToTransferResponsesYearStatusSuccess(records)

	return &pb.ApiResponseTransferYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly Transfer status success",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindMonthlyTransferStatusFailed(ctx context.Context, req *pb.FindMonthlyTransferStatus) (*pb.ApiResponseTransferMonthStatusFailed, error) {
	if req.GetYear() <= 0 || req.GetMonth() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year or month",
		})
	}

	year := req.GetYear()
	month := req.GetMonth()

	records, errResponse := s.transferService.FindMonthTransferStatusFailed(int(year), int(month))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly Transfer status Failed: " + errResponse.Message,
		})
	}

	so := s.mapping.ToResponsesTransferMonthStatusFailed(records)

	return &pb.ApiResponseTransferMonthStatusFailed{
		Status:  "Failed",
		Message: "Failedfully fetched monthly Transfer status Failed",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindYearlyTransferStatusFailed(ctx context.Context, req *pb.FindYearTransfer) (*pb.ApiResponseTransferYearStatusFailed, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	year := req.GetYear()

	records, errResponse := s.transferService.FindYearlyTransferStatusFailed(int(year))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly Transfer status Failed: " + errResponse.Message,
		})
	}

	so := s.mapping.ToTransferResponsesYearStatusFailed(records)

	return &pb.ApiResponseTransferYearStatusFailed{
		Status:  "Failed",
		Message: "Failedfully fetched yearly Transfer status Failed",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindMonthlyTransferAmounts(ctx context.Context, req *pb.FindYearTransfer) (*pb.ApiResponseTransferMonthAmount, error) {
	amounts, err := s.transferService.FindMonthlyTransferAmounts(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly transfer amounts: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTransferMonthAmounts(amounts)

	return &pb.ApiResponseTransferMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly transfer amounts",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindYearlyTransferAmounts(ctx context.Context, req *pb.FindYearTransfer) (*pb.ApiResponseTransferYearAmount, error) {
	amounts, err := s.transferService.FindYearlyTransferAmounts(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly transfer amounts: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTransferYearAmounts(amounts)

	return &pb.ApiResponseTransferYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly transfer amounts",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindMonthlyTransferAmountsBySenderCardNumber(ctx context.Context, req *pb.FindByCardNumberTransferRequest) (*pb.ApiResponseTransferMonthAmount, error) {
	amounts, err := s.transferService.FindMonthlyTransferAmountsBySenderCardNumber(req.GetCardNumber(), int(req.GetYear()))

	fmt.Println("my_cardNumber: ", req.GetCardNumber())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly transfer amounts by sender card number: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTransferMonthAmounts(amounts)

	return &pb.ApiResponseTransferMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly transfer amounts by sender card number",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindMonthlyTransferAmountsByReceiverCardNumber(ctx context.Context, req *pb.FindByCardNumberTransferRequest) (*pb.ApiResponseTransferMonthAmount, error) {
	amounts, err := s.transferService.FindMonthlyTransferAmountsByReceiverCardNumber(req.GetCardNumber(), int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly transfer amounts by receiver card number: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTransferMonthAmounts(amounts)

	return &pb.ApiResponseTransferMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly transfer amounts by receiver card number",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindYearlyTransferAmountsBySenderCardNumber(ctx context.Context, req *pb.FindByCardNumberTransferRequest) (*pb.ApiResponseTransferYearAmount, error) {
	amounts, err := s.transferService.FindYearlyTransferAmountsBySenderCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly transfer amounts by sender card number: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTransferYearAmounts(amounts)

	return &pb.ApiResponseTransferYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly transfer amounts by sender card number",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindYearlyTransferAmountsByReceiverCardNumber(ctx context.Context, req *pb.FindByCardNumberTransferRequest) (*pb.ApiResponseTransferYearAmount, error) {
	amounts, err := s.transferService.FindYearlyTransferAmountsByReceiverCardNumber(req.GetCardNumber(), int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly transfer amounts by receiver card number: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTransferYearAmounts(amounts)

	return &pb.ApiResponseTransferYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly transfer amounts by receiver card number",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindByTransferByTransferFrom(ctx context.Context, request *pb.FindTransferByTransferFromRequest) (*pb.ApiResponseTransfers, error) {
	merchants, err := s.transferService.FindTransferByTransferFrom(request.GetTransferFrom())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfer records: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesTransfer(merchants)

	return &pb.ApiResponseTransfers{
		Status:  "success",
		Message: "Successfully fetch transfer records",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindByTransferByTransferTo(ctx context.Context, request *pb.FindTransferByTransferToRequest) (*pb.ApiResponseTransfers, error) {
	merchants, err := s.transferService.FindTransferByTransferTo(request.GetTransferTo())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfer records: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesTransfer(merchants)

	return &pb.ApiResponseTransfers{
		Status:  "success",
		Message: "Successfully fetch transfer records",
		Data:    so,
	}, nil
}

func (s *transferHandleGrpc) FindByActiveTransfer(ctx context.Context, req *pb.FindAllTransferRequest) (*pb.ApiResponsePaginationTransferDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.transferService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfer records: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := s.mapping.ToResponsesTransferDeleteAt(res)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationTransferDeleteAt{
		Status:     "success",
		Message:    "Successfully fetch transfer records",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *transferHandleGrpc) FindByTrashedTransfer(ctx context.Context, req *pb.FindAllTransferRequest) (*pb.ApiResponsePaginationTransferDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.transferService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfer records: " + err.Message,
		})
	}
	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := s.mapping.ToResponsesTransferDeleteAt(res)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationTransferDeleteAt{
		Status:     "success",
		Message:    "Successfully fetch transfer records",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *transferHandleGrpc) CreateTransfer(ctx context.Context, request *pb.CreateTransferRequest) (*pb.ApiResponseTransfer, error) {
	req := requests.CreateTransferRequest{
		TransferFrom:   request.GetTransferFrom(),
		TransferTo:     request.GetTransferTo(),
		TransferAmount: int(request.GetTransferAmount()),
	}

	res, err := s.transferService.CreateTransaction(&req)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transfer: " + err.Message,
		})
	}

	return &pb.ApiResponseTransfer{
		Status:  "success",
		Message: "Successfully created transfer",
		Data:    s.mapping.ToResponseTransfer(res),
	}, nil
}

func (s *transferHandleGrpc) UpdateTransfer(ctx context.Context, request *pb.UpdateTransferRequest) (*pb.ApiResponseTransfer, error) {
	if request.GetTransferId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Transfer ID is required",
		})
	}

	req := requests.UpdateTransferRequest{
		TransferID:     int(request.GetTransferId()),
		TransferFrom:   request.GetTransferFrom(),
		TransferTo:     request.GetTransferTo(),
		TransferAmount: int(request.GetTransferAmount()),
	}

	res, err := s.transferService.UpdateTransaction(&req)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transfer: " + err.Message,
		})
	}

	return &pb.ApiResponseTransfer{
		Status:  "success",
		Message: "Successfully updated transfer",
		Data:    s.mapping.ToResponseTransfer(res),
	}, nil
}

func (s *transferHandleGrpc) TrashedTransfer(ctx context.Context, request *pb.FindByIdTransferRequest) (*pb.ApiResponseTransfer, error) {
	if request.GetTransferId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Transfer ID is required",
		})
	}

	res, err := s.transferService.TrashedTransfer(int(request.GetTransferId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash transfer: " + err.Message,
		})
	}

	return &pb.ApiResponseTransfer{
		Status:  "success",
		Message: "Successfully trashed transfer",
		Data:    s.mapping.ToResponseTransfer(res),
	}, nil
}

func (s *transferHandleGrpc) RestoreTransfer(ctx context.Context, request *pb.FindByIdTransferRequest) (*pb.ApiResponseTransfer, error) {
	if request.GetTransferId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Transfer ID is required",
		})
	}

	res, err := s.transferService.RestoreTransfer(int(request.GetTransferId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore transfer: " + err.Message,
		})
	}

	return &pb.ApiResponseTransfer{
		Status:  "success",
		Message: "Successfully restored transfer",
		Data:    s.mapping.ToResponseTransfer(res),
	}, nil
}

func (s *transferHandleGrpc) DeleteTransferPermanent(ctx context.Context, request *pb.FindByIdTransferRequest) (*pb.ApiResponseTransferDelete, error) {
	if request.GetTransferId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Transfer ID is required",
		})
	}

	_, err := s.transferService.DeleteTransferPermanent(int(request.GetTransferId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transfer: " + err.Message,
		})
	}

	return &pb.ApiResponseTransferDelete{
		Status:  "success",
		Message: "Successfully deleted transfer",
	}, nil
}

func (s *transferHandleGrpc) RestoreAllTransfer(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransferAll, error) {
	_, err := s.transferService.RestoreAllTransfer()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all transfer: ",
		})
	}

	return &pb.ApiResponseTransferAll{
		Status:  "success",
		Message: "Successfully restore all transfer",
	}, nil
}

func (s *transferHandleGrpc) DeleteAllTransferPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransferAll, error) {
	_, err := s.transferService.DeleteAllTransferPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transfer permanent: ",
		})
	}

	return &pb.ApiResponseTransferAll{
		Status:  "success",
		Message: "Successfully delete transfer permanent",
	}, nil
}
