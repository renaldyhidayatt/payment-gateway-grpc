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

type withdrawHandleGrpc struct {
	pb.UnimplementedWithdrawServiceServer
	withdraw service.WithdrawService
}

func NewWithdrawHandleGrpc(withdraw service.WithdrawService) *withdrawHandleGrpc {
	return &withdrawHandleGrpc{
		withdraw: withdraw,
	}
}

func (h *withdrawHandleGrpc) GetWithdraws(ctx context.Context, req *emptypb.Empty) (*pb.WithdrawsResponse, error) {
	withdraws, err := h.withdraw.FindAll()

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	var pbWithdraws []*pb.Withdraw

	for _, withdraw := range withdraws {
		createdAtProto := timestamppb.New(withdraw.CreatedAt.Time)

		var updatedAtProto *timestamppb.Timestamp

		if withdraw.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(withdraw.UpdatedAt.Time)
		}

		pbWithdraws = append(pbWithdraws, &pb.Withdraw{
			WithdrawId:     int32(withdraw.WithdrawID),
			UserId:         int32(withdraw.UserID),
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   timestamppb.New(withdraw.WithdrawTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		})
	}

	return &pb.WithdrawsResponse{Withdraws: pbWithdraws}, nil
}

func (h *withdrawHandleGrpc) GetWithdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	withdraw, err := h.withdraw.FindById(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(withdraw.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if withdraw.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(withdraw.UpdatedAt.Time)
	}

	return &pb.WithdrawResponse{
		Withdraw: &pb.Withdraw{
			WithdrawId:     int32(withdraw.WithdrawID),
			UserId:         int32(withdraw.UserID),
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   timestamppb.New(withdraw.WithdrawTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		},
	}, nil
}

func (h *withdrawHandleGrpc) GetWithdrawByUsers(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawsResponse, error) {
	withdraws, err := h.withdraw.FindByUsers(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	var pbWithdraws []*pb.Withdraw

	for _, withdraw := range withdraws {
		createdAtProto := timestamppb.New(withdraw.CreatedAt.Time)

		var updatedAtProto *timestamppb.Timestamp

		if withdraw.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(withdraw.UpdatedAt.Time)
		}

		pbWithdraws = append(pbWithdraws, &pb.Withdraw{
			WithdrawId:     int32(withdraw.WithdrawID),
			UserId:         int32(withdraw.UserID),
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   timestamppb.New(withdraw.WithdrawTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		})
	}

	return &pb.WithdrawsResponse{Withdraws: pbWithdraws}, nil
}

func (h *withdrawHandleGrpc) GetWithdrawByUserId(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	withdraw, err := h.withdraw.FindByUsersId(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(withdraw.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if withdraw.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(withdraw.UpdatedAt.Time)
	}

	return &pb.WithdrawResponse{
		Withdraw: &pb.Withdraw{
			WithdrawId:     int32(withdraw.WithdrawID),
			UserId:         int32(withdraw.UserID),
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   timestamppb.New(withdraw.WithdrawTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		},
	}, nil
}

func (h *withdrawHandleGrpc) CreateWithdraw(ctx context.Context, req *pb.CreateWithdrawRequest) (*pb.WithdrawResponse, error) {
	request := &requests.CreateWithdrawRequest{
		UserID:         int(req.UserId),
		WithdrawAmount: int(req.WithdrawAmount),
		WithdrawTime:   req.WithdrawTime.AsTime(),
	}

	res, err := h.withdraw.Create(request)

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(res.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if res.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(res.UpdatedAt.Time)
	}

	return &pb.WithdrawResponse{
		Withdraw: &pb.Withdraw{
			WithdrawId:     int32(res.WithdrawID),
			UserId:         int32(res.UserID),
			WithdrawAmount: int32(res.WithdrawAmount),
			WithdrawTime:   timestamppb.New(res.WithdrawTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		},
	}, nil
}

func (h *withdrawHandleGrpc) UpdateWithdraw(ctx context.Context, req *pb.UpdateWithdrawRequest) (*pb.WithdrawResponse, error) {
	request := &requests.UpdateWithdrawRequest{
		WithdrawID:     int(req.WithdrawId),
		UserID:         int(req.UserId),
		WithdrawAmount: int(req.WithdrawAmount),
		WithdrawTime:   req.WithdrawTime.AsTime(),
	}

	res, err := h.withdraw.Update(request)

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(res.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if res.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(res.UpdatedAt.Time)
	}

	return &pb.WithdrawResponse{
		Withdraw: &pb.Withdraw{
			WithdrawId:     int32(res.WithdrawID),
			UserId:         int32(res.UserID),
			WithdrawAmount: int32(res.WithdrawAmount),
			WithdrawTime:   timestamppb.New(res.WithdrawTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		},
	}, nil
}

func (h *withdrawHandleGrpc) DeleteWithdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.DeleteWithdrawResponse, error) {
	err := h.withdraw.Delete(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.DeleteWithdrawResponse{
		Success: true,
	}, nil
}
