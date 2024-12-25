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

	totalPages := (totalRecords + pageSize - 1) / pageSize

	so := s.mapping.ToResponsesCard(cards)

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationCard{
		Status:     "success",
		Message:    "Successfully fetched card records",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *cardHandleGrpc) FindByIdCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	card, err := s.cardService.FindById(int(req.GetCardId()))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponseCard(card)

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    so,
	}, nil
}

func (s *cardHandleGrpc) FindByUserIdCard(ctx context.Context, req *pb.FindByUserIdCardRequest) (*pb.ApiResponseCard, error) {
	res, err := s.cardService.FindByUserID(int(req.GetUserId()))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponseCard(res)

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    so,
	}, nil
}

func (s *cardHandleGrpc) FindByActiveCard(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCards, error) {
	res, err := s.cardService.FindByActive()

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesCard(res)

	return &pb.ApiResponseCards{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    so,
	}, nil
}

func (s *cardHandleGrpc) FindByTrashedCard(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCards, error) {
	res, err := s.cardService.FindByTrashed()

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesCard(res)

	return &pb.ApiResponseCards{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    so,
	}, nil

}

func (s *cardHandleGrpc) FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponseCard, error) {

	res, err := s.cardService.FindByCardNumber(req.GetCardNumber())

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Card not found: " + err.Message,
		})
	}

	so := s.mapping.ToResponseCard(res)

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    so,
	}, nil

}

func (s *cardHandleGrpc) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.ApiResponseCard, error) {
	request := requests.CreateCardRequest{
		UserID:       int(req.UserId),
		CardType:     req.CardType,
		ExpireDate:   req.ExpireDate.AsTime(),
		CVV:          req.Cvv,
		CardProvider: req.CardProvider,
	}

	res, err := s.cardService.CreateCard(&request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create card: " + err.Message,
		})
	}

	so := s.mapping.ToResponseCard(res)

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully created card",
		Data:    so,
	}, nil
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

	res, err := s.cardService.UpdateCard(&request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update card: " + err.Message,
		})
	}

	so := s.mapping.ToResponseCard(res)

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully updated card",
		Data:    so,
	}, nil
}

func (s *cardHandleGrpc) TrashedCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	res, err := s.cardService.TrashedCard(int(req.CardId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed card: " + err.Message,
		})
	}

	so := s.mapping.ToResponseCard(res)

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully trashed card",
		Data:    so,
	}, nil
}

func (s *cardHandleGrpc) RestoreCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	res, err := s.cardService.RestoreCard(int(req.CardId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore card: " + err.Message,
		})
	}

	so := s.mapping.ToResponseCard(res)

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully restored card",
		Data:    so,
	}, nil
}

func (s *cardHandleGrpc) DeleteCardPermanent(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCardDelete, error) {
	_, err := s.cardService.DeleteCardPermanent(int(req.CardId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete card: " + err.Message,
		})
	}

	return &pb.ApiResponseCardDelete{
		Status:  "success",
		Message: "Successfully deleted card",
	}, nil
}
