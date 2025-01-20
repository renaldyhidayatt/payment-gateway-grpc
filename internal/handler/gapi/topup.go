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

func (s *topupHandleGrpc) FindAllTopup(ctx context.Context, req *pb.FindAllTopupRequest) (*pb.ApiResponsePaginationTopup, error) {
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

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

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

func (s *topupHandleGrpc) FindByIdTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopup, error) {
	if req.GetTopupId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	id := req.GetTopupId()

	topup, err := s.topupService.FindById(int(id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topup: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopup(topup)

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully fetch topup",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindMonthlyTopupStatusSuccess(ctx context.Context, req *pb.FindMonthlyTopupStatus) (*pb.ApiResponseTopupMonthStatusSuccess, error) {
	if req.GetYear() <= 0 || req.GetMonth() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year or month",
		})
	}

	year := req.GetYear()
	month := req.GetMonth()

	records, errResponse := s.topupService.FindMonthTopupStatusSuccess(int(year), int(month))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly topup status success: " + errResponse.Message,
		})
	}

	so := s.mapping.ToResponsesTopupMonthStatusSuccess(records)

	return &pb.ApiResponseTopupMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly topup status success",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindYearlyTopupStatusSuccess(ctx context.Context, req *pb.FindYearTopup) (*pb.ApiResponseTopupYearStatusSuccess, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	year := req.GetYear()

	records, errResponse := s.topupService.FindYearlyTopupStatusSuccess(int(year))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly topup status success: " + errResponse.Message,
		})
	}

	so := s.mapping.ToTopupResponsesYearStatusSuccess(records)

	return &pb.ApiResponseTopupYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly topup status success",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindMonthlyTopupStatusFailed(ctx context.Context, req *pb.FindMonthlyTopupStatus) (*pb.ApiResponseTopupMonthStatusFailed, error) {
	if req.GetYear() <= 0 || req.GetMonth() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year or month",
		})
	}

	year := req.GetYear()
	month := req.GetMonth()

	records, errResponse := s.topupService.FindMonthTopupStatusFailed(int(year), int(month))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly topup status Failed: " + errResponse.Message,
		})
	}

	so := s.mapping.ToResponsesTopupMonthStatusFailed(records)

	return &pb.ApiResponseTopupMonthStatusFailed{
		Status:  "Failed",
		Message: "Failedfully fetched monthly topup status Failed",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindYearlyTopupStatusFailed(ctx context.Context, req *pb.FindYearTopup) (*pb.ApiResponseTopupYearStatusFailed, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	year := req.GetYear()

	records, errResponse := s.topupService.FindYearlyTopupStatusFailed(int(year))
	if errResponse != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly topup status Failed: " + errResponse.Message,
		})
	}

	so := s.mapping.ToTopupResponsesYearStatusFailed(records)

	return &pb.ApiResponseTopupYearStatusFailed{
		Status:  "Failed",
		Message: "Failedfully fetched yearly topup status Failed",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindMonthlyTopupMethods(ctx context.Context, req *pb.FindYearTopup) (*pb.ApiResponseTopupMonthMethod, error) {
	methods, err := s.topupService.FindMonthlyTopupMethods(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly topup methods: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopupMonthlyMethods(methods)

	return &pb.ApiResponseTopupMonthMethod{
		Status:  "success",
		Message: "Successfully fetched monthly topup methods",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindYearlyTopupMethods(ctx context.Context, req *pb.FindYearTopup) (*pb.ApiResponseTopupYearMethod, error) {
	methods, err := s.topupService.FindYearlyTopupMethods(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly topup methods: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopupYearlyMethods(methods)

	return &pb.ApiResponseTopupYearMethod{
		Status:  "success",
		Message: "Successfully fetched yearly topup methods",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindMonthlyTopupAmounts(ctx context.Context, req *pb.FindYearTopup) (*pb.ApiResponseTopupMonthAmount, error) {
	amounts, err := s.topupService.FindMonthlyTopupAmounts(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly topup amounts: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopupMonthlyAmounts(amounts)

	return &pb.ApiResponseTopupMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly topup amounts",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindYearlyTopupAmounts(ctx context.Context, req *pb.FindYearTopup) (*pb.ApiResponseTopupYearAmount, error) {
	amounts, err := s.topupService.FindYearlyTopupAmounts(int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly topup amounts: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopupYearlyAmounts(amounts)

	return &pb.ApiResponseTopupYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly topup amounts",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindMonthlyTopupMethodsByCardNumber(ctx context.Context, req *pb.FindYearTopupCardNumber) (*pb.ApiResponseTopupMonthMethod, error) {
	methods, err := s.topupService.FindMonthlyTopupMethodsByCardNumber(req.GetCardNumber(), int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly topup methods by card number: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopupMonthlyMethods(methods)

	return &pb.ApiResponseTopupMonthMethod{
		Status:  "success",
		Message: "Successfully fetched monthly topup methods by card number",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindYearlyTopupMethodsByCardNumber(ctx context.Context, req *pb.FindYearTopupCardNumber) (*pb.ApiResponseTopupYearMethod, error) {
	methods, err := s.topupService.FindYearlyTopupMethodsByCardNumber(req.GetCardNumber(), int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly topup methods by card number: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopupYearlyMethods(methods)

	return &pb.ApiResponseTopupYearMethod{
		Status:  "success",
		Message: "Successfully fetched yearly topup methods by card number",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindMonthlyTopupAmountsByCardNumber(ctx context.Context, req *pb.FindYearTopupCardNumber) (*pb.ApiResponseTopupMonthAmount, error) {
	amounts, err := s.topupService.FindMonthlyTopupAmountsByCardNumber(req.GetCardNumber(), int(req.GetYear()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly topup amounts by card number: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopupMonthlyAmounts(amounts)

	return &pb.ApiResponseTopupMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly topup amounts by card number",
		Data:    so,
	}, nil
}

func (s *topupHandleGrpc) FindYearlyTopupAmountsByCardNumber(ctx context.Context, req *pb.FindYearTopupCardNumber) (*pb.ApiResponseTopupYearAmount, error) {
	amounts, err := s.topupService.FindYearlyTopupAmountsByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly topup amounts by card number: " + err.Message,
		})
	}

	so := s.mapping.ToResponseTopupYearlyAmounts(amounts)

	return &pb.ApiResponseTopupYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly topup amounts by card number",
		Data:    so,
	}, nil
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

func (s *topupHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllTopupRequest) (*pb.ApiResponsePaginationTopupDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.topupService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topups: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := s.mapping.ToResponsesTopupDeleteAt(res)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationTopupDeleteAt{
		Status:     "success",
		Message:    "Successfully fetch topups",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *topupHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllTopupRequest) (*pb.ApiResponsePaginationTopupDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.topupService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topups: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	so := s.mapping.ToResponsesTopupDeleteAt(res)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationTopupDeleteAt{
		Status:     "success",
		Message:    "Successfully fetch topups",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *topupHandleGrpc) CreateTopup(ctx context.Context, req *pb.CreateTopupRequest) (*pb.ApiResponseTopup, error) {
	request := requests.CreateTopupRequest{
		CardNumber:  req.GetCardNumber(),
		TopupAmount: int(req.GetTopupAmount()),
		TopupMethod: req.GetTopupMethod(),
	}

	res, err := s.topupService.CreateTopup(&request)

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
	if req.GetTopupId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	request := requests.UpdateTopupRequest{
		TopupID:     int(req.GetTopupId()),
		CardNumber:  req.GetCardNumber(),
		TopupAmount: int(req.GetTopupAmount()),
		TopupMethod: req.GetTopupMethod(),
	}

	res, err := s.topupService.UpdateTopup(&request)

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
	if req.GetTopupId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

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
	if req.GetTopupId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

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
	if req.GetTopupId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

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

func (s *topupHandleGrpc) RestoreAllTopup(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTopupAll, error) {
	_, err := s.topupService.RestoreAllTopup()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all topup: ",
		})
	}

	return &pb.ApiResponseTopupAll{
		Status:  "success",
		Message: "Successfully restore all topup",
	}, nil
}

func (s *topupHandleGrpc) DeleteAllTopupPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTopupAll, error) {
	_, err := s.topupService.DeleteAllTopupPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete topup permanent: ",
		})
	}

	return &pb.ApiResponseTopupAll{
		Status:  "success",
		Message: "Successfully delete topup permanent",
	}, nil
}
