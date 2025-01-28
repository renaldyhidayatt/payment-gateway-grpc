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

type cardHandleGrpc struct {
	pb.UnimplementedCardServiceServer
	cardService service.CardService
	mapping     protomapper.CardProtoMapper
}

func NewCardHandleGrpc(card service.CardService, mapping protomapper.CardProtoMapper) *cardHandleGrpc {
	return &cardHandleGrpc{cardService: card, mapping: mapping}
}

func (s *cardHandleGrpc) FindAllCard(ctx context.Context, req *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCard, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	cards, totalRecords, err := s.cardService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card records: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationCard(paginationMeta, "success", "Successfully fetched card records", cards)

	return so, nil
}

func (s *cardHandleGrpc) FindByIdCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	if req.GetCardId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid card id",
		})
	}

	card, err := s.cardService.FindById(int(req.GetCardId()))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: ",
		})
	}

	so := s.mapping.ToProtoResponseCard("success", "Successfully fetched card record", card)

	return so, nil
}

func (s *cardHandleGrpc) FindByUserIdCard(ctx context.Context, req *pb.FindByUserIdCardRequest) (*pb.ApiResponseCard, error) {
	if req.GetUserId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid user id",
		})
	}

	res, err := s.cardService.FindByUserID(int(req.GetUserId()))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCard("success", "Successfully fetched card record", res)

	return so, nil
}

func (s *cardHandleGrpc) DashboardCard(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseDashboardCard, error) {

	dashboardCard, err := s.cardService.DashboardCard()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve dashboard card: %v", err)
	}

	so := s.mapping.ToProtoResponseDashboardCard("success", "Dashboard card retrieved successfully", dashboardCard)

	return so, nil
}

func (s *cardHandleGrpc) DashboardCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponseDashboardCardNumber, error) {
	if req.GetCardNumber() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "card number is required")
	}

	dashboardCard, err := s.cardService.DashboardCardCardNumber(req.GetCardNumber())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve dashboard card for card number %s: %v", req.GetCardNumber(), err)
	}

	so := s.mapping.ToProtoResponseDashboardCardCardNumber("success", "Dashboard card for card number retrieved successfully", dashboardCard)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyBalance(ctx context.Context, req *pb.FindYearBalance) (*pb.ApiResponseMonthlyBalance, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyBalance(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly balance: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyBalances("success", "Monthly balance retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyBalance(ctx context.Context, req *pb.FindYearBalance) (*pb.ApiResponseYearlyBalance, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyBalance(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly balance: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyBalances("success", "Yearly balance retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyTopupAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyTopupAmount(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly topup amount: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly topup amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyTopupAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyTopupAmount(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly topup amount: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "Yearly topup amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyWithdrawAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyWithdrawAmount(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly withdraw amount: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly withdraw amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyWithdrawAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyWithdrawAmount(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly withdraw amount: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "Yearly withdraw amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyTransactionAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyTransactionAmount(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly transaction amount: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly transaction amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyTransactionAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyTransactionAmount(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly transaction amount: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "Yearly transaction amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyTransferSenderAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyTransferAmountSender(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly transfer sender amount: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly transfer sender amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyTransferSenderAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyTransferAmountSender(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly transfer sender amount: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "transfer sender amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyTransferReceiverAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyTransferAmountReceiver(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly transfer receiver amount: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly transfer receiver amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyTransferReceiverAmount(ctx context.Context, req *pb.FindYearAmount) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyTransferAmountReceiver(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly transfer receiver amount: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "Yearly transfer receiver amount retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyBalanceByCardNumber(ctx context.Context, req *pb.FindYearBalanceCardNumber) (*pb.ApiResponseMonthlyBalance, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyBalanceByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly balance: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyBalances("success", "Monthly balance retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyBalanceByCardNumber(ctx context.Context, req *pb.FindYearBalanceCardNumber) (*pb.ApiResponseYearlyBalance, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyBalanceByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly balance: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyBalances("success", "Yearly balance retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyTopupAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyTopupAmountByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly topup amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly topup amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyTopupAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyTopupAmountByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly topup amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "Yearly topup amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyWithdrawAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyWithdrawAmountByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly withdraw amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly withdraw amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyWithdrawAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyWithdrawAmountByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly withdraw amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "Yearly withdraw amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyTransactionAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyTransactionAmountByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly transaction amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly transaction amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyTransactionAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyTransactionAmountByCardNumber(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly transaction amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "Yearly transaction amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyTransferSenderAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyTransferAmountBySender(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly transfer sender amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly transfer sender amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyTransferSenderAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyTransferAmountBySender(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly transfer sender amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "Yearly transfer sender amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindMonthlyTransferReceiverAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseMonthlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindMonthlyTransferAmountByReceiver(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find monthly transfer receiver amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseMonthlyAmounts("success", "Monthly transfer receiver amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindYearlyTransferReceiverAmountByCardNumber(ctx context.Context, req *pb.FindYearAmountCardNumber) (*pb.ApiResponseYearlyAmount, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid year: %d", req.GetYear())
	}

	res, err := s.cardService.FindYearlyTransferAmountByReceiver(req.GetCardNumber(), int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find yearly transfer receiver amount by card number: %v", err)
	}

	so := s.mapping.ToProtoResponseYearlyAmounts("success", "Yearly transfer receiver amount by card number retrieved successfully", res)

	return so, nil
}

func (s *cardHandleGrpc) FindByActiveCard(ctx context.Context, req *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCardDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.cardService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationCardDeletedAt(paginationMeta, "success", "Successfully fetched card record", res)

	return so, nil
}

func (s *cardHandleGrpc) FindByTrashedCard(ctx context.Context, req *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCardDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.cardService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationCardDeletedAt(paginationMeta, "success", "Successfully fetched card record", res)

	return so, nil

}

func (s *cardHandleGrpc) FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponseCard, error) {

	res, err := s.cardService.FindByCardNumber(req.GetCardNumber())

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCard("success", "Successfully fetched card record", res)

	return so, nil

}

func (s *cardHandleGrpc) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.ApiResponseCard, error) {
	request := requests.CreateCardRequest{
		UserID:       int(req.UserId),
		CardType:     req.CardType,
		ExpireDate:   req.ExpireDate.AsTime(),
		CVV:          req.Cvv,
		CardProvider: req.CardProvider,
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create card: ",
		})
	}

	res, err := s.cardService.CreateCard(&request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create card: ",
		})
	}

	so := s.mapping.ToProtoResponseCard("success", "Successfully created card", res)

	return so, nil
}

func (s *cardHandleGrpc) UpdateCard(ctx context.Context, req *pb.UpdateCardRequest) (*pb.ApiResponseCard, error) {
	request := requests.UpdateCardRequest{
		CardID:       int(req.CardId),
		UserID:       int(req.UserId),
		CardType:     req.CardType,
		ExpireDate:   req.ExpireDate.AsTime(),
		CVV:          req.Cvv,
		CardProvider: req.CardProvider,
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update card: ",
		})
	}

	res, err := s.cardService.UpdateCard(&request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update card: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCard("success", "Successfully updated card", res)

	return so, nil
}

func (s *cardHandleGrpc) TrashedCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	res, err := s.cardService.TrashedCard(int(req.CardId))

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid Id",
		})
	}

	so := s.mapping.ToProtoResponseCard("success", "Successfully trashed card", res)

	return so, nil
}

func (s *cardHandleGrpc) RestoreCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	if req.CardId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore card: ",
		})
	}

	res, err := s.cardService.RestoreCard(int(req.CardId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore card: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCard("success", "Successfully restored card", res)

	return so, nil
}

func (s *cardHandleGrpc) DeleteCardPermanent(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCardDelete, error) {
	if req.CardId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete card: ",
		})
	}

	_, err := s.cardService.DeleteCardPermanent(int(req.CardId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete card: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCardDeleteAt("success", "Successfully deleted card")

	return so, nil
}

func (s *cardHandleGrpc) RestoreAllCard(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCardAll, error) {
	_, err := s.cardService.RestoreAllCard()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all card: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCardAll("success", "Successfully restore card")

	return so, nil
}

func (s *cardHandleGrpc) DeleteAllCardPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCardAll, error) {
	_, err := s.cardService.DeleteAllCardPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete card permanent: ",
		})
	}

	so := s.mapping.ToProtoResponseCardAll("success", "Successfully delete card permanent")

	return so, nil
}
