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

func (s *transactionHandleGrpc) FindMonthlyTransactionStatusSuccess(ctx context.Context, req *pb.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthStatusSuccess, error) {
	if req.GetYear() <= 0 || req.GetMonth() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year or month",
		})
	}

	year := req.GetYear()
	month := req.GetMonth()

	records, errResponse := s.transactionService.FindMonthTransactionStatusSuccess(int(year), int(month))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly Transaction status success: " + errResponse.Message,
		})
	}

	so := s.mapping.ToResponsesTransactionMonthStatusSuccess(records)

	return &pb.ApiResponseTransactionMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly Transaction status success",
		Data:    so,
	}, nil
}

func (s *transactionHandleGrpc) FindYearlyTransactionStatusSuccess(ctx context.Context, req *pb.FindYearTransaction) (*pb.ApiResponseTransactionYearStatusSuccess, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	year := req.GetYear()

	records, errResponse := s.transactionService.FindYearlyTransactionStatusSuccess(int(year))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly Transaction status success: " + errResponse.Message,
		})
	}

	so := s.mapping.ToTransactionResponsesYearStatusSuccess(records)

	return &pb.ApiResponseTransactionYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly Transaction status success",
		Data:    so,
	}, nil
}

func (s *transactionHandleGrpc) FindMonthlyTransactionStatusFailed(ctx context.Context, req *pb.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthStatusFailed, error) {
	if req.GetYear() <= 0 || req.GetMonth() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year or month",
		})
	}

	year := req.GetYear()
	month := req.GetMonth()

	records, errResponse := s.transactionService.FindMonthTransactionStatusFailed(int(year), int(month))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly Transaction status Failed: " + errResponse.Message,
		})
	}

	so := s.mapping.ToResponsesTransactionMonthStatusFailed(records)

	return &pb.ApiResponseTransactionMonthStatusFailed{
		Status:  "Failed",
		Message: "Failedfully fetched monthly Transaction status Failed",
		Data:    so,
	}, nil
}

func (s *transactionHandleGrpc) FindYearlyTransactionStatusFailed(ctx context.Context, req *pb.FindYearTransaction) (*pb.ApiResponseTransactionYearStatusFailed, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	year := req.GetYear()

	records, errResponse := s.transactionService.FindYearlyTransactionStatusFailed(int(year))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly Transaction status Failed: " + errResponse.Message,
		})
	}

	so := s.mapping.ToTransactionResponsesYearStatusFailed(records)

	return &pb.ApiResponseTransactionYearStatusFailed{
		Status:  "Failed",
		Message: "Failedfully fetched yearly Transaction status Failed",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindMonthlyPaymentMethods(ctx context.Context, req *pb.FindYearTransaction) (*pb.ApiResponseTransactionMonthMethod, error) {
	methods, err := t.transactionService.FindMonthlyPaymentMethods(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly payment methods: " + err.Message,
		})
	}

	so := t.mapping.ToResponseTransactionMonthMethods(methods)

	return &pb.ApiResponseTransactionMonthMethod{
		Status:  "success",
		Message: "Successfully fetched monthly payment methods",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindYearlyPaymentMethods(ctx context.Context, req *pb.FindYearTransaction) (*pb.ApiResponseTransactionYearMethod, error) {
	methods, err := t.transactionService.FindYearlyPaymentMethods(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly payment methods: " + err.Message,
		})
	}

	so := t.mapping.ToResponseTransactionYearMethods(methods)

	return &pb.ApiResponseTransactionYearMethod{
		Status:  "success",
		Message: "Successfully fetched yearly payment methods",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindMonthlyAmounts(ctx context.Context, req *pb.FindYearTransaction) (*pb.ApiResponseTransactionMonthAmount, error) {
	amounts, err := t.transactionService.FindMonthlyAmounts(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly amounts: " + err.Message,
		})
	}

	so := t.mapping.ToResponseTransactionMonthAmounts(amounts)

	return &pb.ApiResponseTransactionMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amounts",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindYearlyAmounts(ctx context.Context, req *pb.FindYearTransaction) (*pb.ApiResponseTransactionYearAmount, error) {
	amounts, err := t.transactionService.FindYearlyAmounts(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly amounts: " + err.Message,
		})
	}

	so := t.mapping.ToResponseTransactionYearlyAmounts(amounts)

	return &pb.ApiResponseTransactionYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amounts",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindMonthlyPaymentMethodsByCardNumber(ctx context.Context, req *pb.FindByYearCardNumberTransactionRequest) (*pb.ApiResponseTransactionMonthMethod, error) {
	methods, err := t.transactionService.FindMonthlyPaymentMethodsByCardNumber(req.GetCardNumber(), int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly payment methods by card number: " + err.Message,
		})
	}

	so := t.mapping.ToResponseTransactionMonthMethods(methods)

	return &pb.ApiResponseTransactionMonthMethod{
		Status:  "success",
		Message: "Successfully fetched monthly payment methods by card number",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindYearlyPaymentMethodsByCardNumber(ctx context.Context, req *pb.FindByYearCardNumberTransactionRequest) (*pb.ApiResponseTransactionYearMethod, error) {
	methods, err := t.transactionService.FindYearlyPaymentMethodsByCardNumber(req.GetCardNumber(), int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly payment methods by card number: " + err.Message,
		})
	}

	so := t.mapping.ToResponseTransactionYearMethods(methods)

	return &pb.ApiResponseTransactionYearMethod{
		Status:  "success",
		Message: "Successfully fetched yearly payment methods by card number",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindMonthlyAmountsByCardNumber(ctx context.Context, req *pb.FindByYearCardNumberTransactionRequest) (*pb.ApiResponseTransactionMonthAmount, error) {
	amounts, err := t.transactionService.FindMonthlyAmountsByCardNumber(req.GetCardNumber(), int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly amounts by card number: " + err.Message,
		})
	}

	so := t.mapping.ToResponseTransactionMonthAmounts(amounts)

	return &pb.ApiResponseTransactionMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amounts by card number",
		Data:    so,
	}, nil
}

func (t *transactionHandleGrpc) FindYearlyAmountsByCardNumber(ctx context.Context, req *pb.FindByYearCardNumberTransactionRequest) (*pb.ApiResponseTransactionYearAmount, error) {
	amounts, err := t.transactionService.FindYearlyAmountsByCardNumber(req.GetCardNumber(), int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly amounts by card number: " + err.Message,
		})
	}

	so := t.mapping.ToResponseTransactionYearlyAmounts(amounts)

	return &pb.ApiResponseTransactionYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amounts by card number",
		Data:    so,
	}, nil
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
