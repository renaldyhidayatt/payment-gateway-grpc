package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type topupHandleGrpc struct {
	pb.UnimplementedTopupServiceServer
	topup service.TopupService
}

func NewTopupHandleGrpc(topup service.TopupService) *topupHandleGrpc {
	return &topupHandleGrpc{
		topup: topup,
	}
}

func (s *topupHandleGrpc) GetTopups(ctx context.Context, empty *emptypb.Empty) (*pb.TopupsResponse, error) {
	res, err := s.topup.FindAll()

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get topups: %v", err)
	}

	return &pb.TopupsResponse{
		Topups: s.convertToPbTopups(res),
	}, nil
}

func (s *topupHandleGrpc) GetTopup(ctx context.Context, req *pb.TopupRequest) (*pb.TopupResponse, error) {
	res, err := s.topup.FindById(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get topup: %v", err)
	}

	return &pb.TopupResponse{
		Topup: s.convertToPbTopup(res),
	}, nil
}

func (s *topupHandleGrpc) GetTopupByUsers(ctx context.Context, req *pb.TopupRequest) (*pb.TopupsResponse, error) {
	res, err := s.topup.FindByUsers(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get topups by users: %v", err)
	}

	return &pb.TopupsResponse{
		Topups: s.convertToPbTopups(res),
	}, nil
}

func (s *topupHandleGrpc) GetTopupByUserId(ctx context.Context, req *pb.TopupRequest) (*pb.TopupResponse, error) {
	res, err := s.topup.FindByUsersId(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get topup by user ID: %v", err)
	}

	return &pb.TopupResponse{
		Topup: s.convertToPbTopup(res),
	}, nil
}

func (s *topupHandleGrpc) CreateTopup(ctx context.Context, req *pb.CreateTopupRequest) (*pb.TopupResponse, error) {
	request := &requests.CreateTopupRequest{
		UserID:      int(req.UserId),
		TopupNo:     req.TopupNo,
		TopupAmount: int(req.TopupAmount),
		TopupMethod: req.TopupMethod,
	}

	res, err := s.topup.Create(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create topup: %v", err)
	}

	return &pb.TopupResponse{
		Topup: s.convertToPbTopup(res),
	}, nil
}

func (s *topupHandleGrpc) UpdateTopup(ctx context.Context, req *pb.UpdateTopupRequest) (*pb.TopupResponse, error) {
	request := &requests.UpdateTopupRequest{
		UserID:      int(req.UserId),
		TopupID:     int(req.TopupId),
		TopupAmount: int(req.TopupAmount),
		TopupMethod: req.TopupMethod,
	}

	res, err := s.topup.UpdateTopup(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update topup: %v", err)
	}

	return &pb.TopupResponse{
		Topup: s.convertToPbTopup(res),
	}, nil
}

func (s *topupHandleGrpc) DeleteTopup(ctx context.Context, req *pb.TopupRequest) (*pb.DeleteTopupResponse, error) {
	err := s.topup.DeleteTopup(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete topup: %v", err)
	}

	return &pb.DeleteTopupResponse{
		Success: true,
	}, nil
}

func (s *topupHandleGrpc) convertToPbTopups(topups []*db.Topup) []*pb.Topup {
	var pbTopups []*pb.Topup

	for _, topup := range topups {
		pbTopups = append(pbTopups, s.convertToPbTopup(topup))
	}

	return pbTopups
}

func (s *topupHandleGrpc) convertToPbTopup(topup *db.Topup) *pb.Topup {
	createdAtProto := timestamppb.New(topup.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if topup.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(topup.UpdatedAt.Time)
	}

	return &pb.Topup{
		TopupId:     int32(topup.TopupID),
		UserId:      int32(topup.UserID),
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   timestamppb.New(topup.TopupTime),
		CreatedAt:   createdAtProto,
		UpdatedAt:   updatedAtProto,
	}
}
