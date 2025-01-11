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

type transactionHandleGrpc struct {
	pb.UnimplementedTransactionServiceServer
	transactionService service.TransactionService
	mapping            protomapper.TransactionProtoMapper
}

func NewTransactionHandleGrpc(transactionService service.TransactionService, mapping protomapper.TransactionProtoMapper) *transactionHandleGrpc {
	return &transactionHandleGrpc{
		transactionService: transactionService,
		mapping:            mapping,
	}
}

func (t *transactionHandleGrpc) FindAllTransaction(ctx context.Context, request *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransaction, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	transactions, totalRecords, err := t.transactionService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transactions: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := t.mapping.ToResponsesTransaction(transactions)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationTransaction{
		Status:     "success",
		Message:    "Transactions fetched successfully",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (t *transactionHandleGrpc) FindTransactionById(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.TransactionResponse, error) {
	if request.GetTransactionId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	id := request.GetTransactionId()

	transaction, err := t.transactionService.FindById(int(id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transaction: " + err.Message,
		})
	}

	so := t.mapping.ToResponseTransaction(transaction)

	return so, nil
}

func (t *transactionHandleGrpc) FindByCardNumberTransaction(ctx context.Context, request *pb.FindByCardNumberTransactionRequest) (*pb.ApiResponseTransactions, error) {
	cardNumber := request.GetCardNumber()

	transactions, err := t.transactionService.FindByCardNumber(cardNumber)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transactions: " + err.Message,
		})
	}

	so := t.mapping.ToResponsesTransaction(transactions)

	return &pb.ApiResponseTransactions{
		Status:  "success",
		Message: "Successfully fetch transactions",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindTransactionByMerchantIdRequest(ctx context.Context, request *pb.FindTransactionByMerchantIdRequest) (*pb.ApiResponseTransactions, error) {
	if request.GetMerchantId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	merchantId := request.GetMerchantId()

	transactions, err := t.transactionService.FindTransactionByMerchantId(int(merchantId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transactions: " + err.Message,
		})
	}

	so := t.mapping.ToResponsesTransaction(transactions)

	return &pb.ApiResponseTransactions{
		Status:  "success",
		Message: "Successfully fetch transactions",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindByActiveTransaction(ctx context.Context, req *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransactionDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	transactions, totalRecords, err := t.transactionService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transactions: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := t.mapping.ToResponsesTransactionDeleteAt(transactions)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:     "success",
		Message:    "Successfully fetch transactions",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (t *transactionHandleGrpc) FindByTrashedTransaction(ctx context.Context, req *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransactionDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	transactions, totalRecords, err := t.transactionService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transactions: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := t.mapping.ToResponsesTransactionDeleteAt(transactions)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:     "success",
		Message:    "Successfully fetch transactions",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (t *transactionHandleGrpc) CreateTransaction(ctx context.Context, request *pb.CreateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	transactionTime := request.GetTransactionTime().AsTime()
	merchantID := int(request.GetMerchantId())

	req := requests.CreateTransactionRequest{
		CardNumber:      request.GetCardNumber(),
		Amount:          int(request.GetAmount()),
		PaymentMethod:   request.GetPaymentMethod(),
		MerchantID:      &merchantID,
		TransactionTime: transactionTime,
	}

	res, err := t.transactionService.Create(request.ApiKey, &req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transaction: " + err.Message,
		})
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully created transaction",
		Data:    t.mapping.ToResponseTransaction(res),
	}, nil
}

func (t *transactionHandleGrpc) UpdateTransaction(ctx context.Context, request *pb.UpdateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	if request.GetTransactionId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	transactionTime := request.GetTransactionTime().AsTime()
	merchantID := int(request.GetMerchantId())

	req := requests.UpdateTransactionRequest{
		TransactionID:   int(request.GetTransactionId()),
		CardNumber:      request.GetCardNumber(),
		Amount:          int(request.GetAmount()),
		PaymentMethod:   request.GetPaymentMethod(),
		MerchantID:      &merchantID,
		TransactionTime: transactionTime,
	}

	res, err := t.transactionService.Update(request.ApiKey, &req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction: " + err.Message,
		})
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully updated transaction",
		Data:    t.mapping.ToResponseTransaction(res),
	}, nil
}

func (t *transactionHandleGrpc) TrashedTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransaction, error) {
	if request.GetTransactionId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	res, err := t.transactionService.TrashedTransaction(int(request.GetTransactionId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transaction: " + err.Message,
		})
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully trashed transaction",
		Data:    t.mapping.ToResponseTransaction(res),
	}, nil
}

func (t *transactionHandleGrpc) RestoreTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransaction, error) {
	if request.GetTransactionId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	res, err := t.transactionService.RestoreTransaction(int(request.GetTransactionId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transaction: " + err.Message,
		})
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully restored transaction",
		Data:    t.mapping.ToResponseTransaction(res),
	}, nil
}

func (t *transactionHandleGrpc) DeleteTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDelete, error) {
	if request.GetTransactionId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	_, err := t.transactionService.DeleteTransactionPermanent(int(request.GetTransactionId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transaction: " + err.Message,
		})
	}

	return &pb.ApiResponseTransactionDelete{
		Status:  "success",
		Message: "Successfully deleted transaction",
	}, nil

}

func (s *transactionHandleGrpc) RestoreAllTransaction(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.RestoreAllTransaction()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all transaction: ",
		})
	}

	return &pb.ApiResponseTransactionAll{
		Status:  "success",
		Message: "Successfully restore all transaction",
	}, nil
}

func (s *transactionHandleGrpc) DeleteAllMerchantPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.DeleteAllTransactionPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transaction permanent: ",
		})
	}

	return &pb.ApiResponseTransactionAll{
		Status:  "success",
		Message: "Successfully delete transaction permanent",
	}, nil
}
