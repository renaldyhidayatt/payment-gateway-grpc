package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"context"

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
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	var pbTopups []*pb.Topup

	for _, topup := range res {
		updatedAtProto := timestamppb.New(topup.UpdatedAt.Time)

		if topup.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(topup.UpdatedAt.Time)
		}

		pbTopups = append(pbTopups, &pb.Topup{
			TopupId:     int32(topup.TopupID),
			UserId:      int32(topup.UserID),
			TopupNo:     topup.TopupNo,
			TopupAmount: int32(topup.TopupAmount),
			TopupMethod: topup.TopupMethod,
			TopupTime:   timestamppb.New(topup.TopupTime),
			CreatedAt:   timestamppb.New(topup.CreatedAt.Time),
			UpdatedAt:   updatedAtProto,
		})
	}

	return &pb.TopupsResponse{
		Topups: pbTopups,
	}, nil
}

func (s *topupHandleGrpc) GetTopup(ctx context.Context, req *pb.TopupRequest) (*pb.TopupResponse, error) {
	res, err := s.topup.FindById(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	updatedAtProto := timestamppb.New(res.UpdatedAt.Time)

	if res.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(res.UpdatedAt.Time)
	}

	return &pb.TopupResponse{
		Topup: &pb.Topup{
			TopupId:     int32(res.TopupID),
			UserId:      int32(res.UserID),
			TopupNo:     res.TopupNo,
			TopupAmount: int32(res.TopupAmount),
			TopupMethod: res.TopupMethod,
			TopupTime:   timestamppb.New(res.TopupTime),
			CreatedAt:   timestamppb.New(res.CreatedAt.Time),
			UpdatedAt:   updatedAtProto,
		},
	}, nil
}

func (s *topupHandleGrpc) GetTopupByUsers(ctx context.Context, req *pb.TopupRequest) (*pb.TopupsResponse, error) {
	res, err := s.topup.FindByUsers(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	var pbTopups []*pb.Topup

	for _, topup := range res {
		updatedAtProto := timestamppb.New(topup.UpdatedAt.Time)

		if topup.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(topup.UpdatedAt.Time)
		}

		pbTopups = append(pbTopups, &pb.Topup{
			TopupId:     int32(topup.TopupID),
			UserId:      int32(topup.UserID),
			TopupNo:     topup.TopupNo,
			TopupAmount: int32(topup.TopupAmount),
			TopupMethod: topup.TopupMethod,
			TopupTime:   timestamppb.New(topup.TopupTime),
			CreatedAt:   timestamppb.New(topup.CreatedAt.Time),
			UpdatedAt:   updatedAtProto,
		})
	}

	return &pb.TopupsResponse{
		Topups: pbTopups,
	}, nil
}

func (s *topupHandleGrpc) GetTopupByUserId(ctx context.Context, req *pb.TopupRequest) (*pb.TopupResponse, error) {

	res, err := s.topup.FindByUsersId(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	updatedAtProto := timestamppb.New(res.UpdatedAt.Time)

	if res.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(res.UpdatedAt.Time)
	}

	return &pb.TopupResponse{
		Topup: &pb.Topup{
			TopupId:     int32(res.TopupID),
			UserId:      int32(res.UserID),
			TopupNo:     res.TopupNo,
			TopupAmount: int32(res.TopupAmount),
			TopupMethod: res.TopupMethod,
			TopupTime:   timestamppb.New(res.TopupTime),
			CreatedAt:   timestamppb.New(res.CreatedAt.Time),
			UpdatedAt:   updatedAtProto,
		},
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
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.TopupResponse{
		Topup: &pb.Topup{
			TopupId:     int32(res.TopupID),
			UserId:      int32(res.UserID),
			TopupNo:     res.TopupNo,
			TopupAmount: int32(res.TopupAmount),
			TopupMethod: res.TopupMethod,
			TopupTime:   timestamppb.New(res.TopupTime),
			CreatedAt:   timestamppb.New(res.CreatedAt.Time),
			UpdatedAt:   timestamppb.New(res.UpdatedAt.Time),
		},
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
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.TopupResponse{
		Topup: &pb.Topup{
			TopupId:     int32(res.TopupID),
			UserId:      int32(res.UserID),
			TopupNo:     res.TopupNo,
			TopupAmount: int32(res.TopupAmount),
			TopupMethod: res.TopupMethod,
			TopupTime:   timestamppb.New(res.TopupTime),
			CreatedAt:   timestamppb.New(res.CreatedAt.Time),
			UpdatedAt:   timestamppb.New(res.UpdatedAt.Time),
		},
	}, nil
}

func (s *topupHandleGrpc) DeleteTopup(ctx context.Context, req *pb.TopupRequest) (*pb.DeleteTopupResponse, error) {
	err := s.topup.DeleteTopup(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.DeleteTopupResponse{
		Success: true,
	}, nil
}
