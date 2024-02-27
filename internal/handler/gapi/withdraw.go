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
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	pbWithdraws := h.convertToPbWithdraws(withdraws)

	return &pb.WithdrawsResponse{Withdraws: pbWithdraws}, nil
}

func (h *withdrawHandleGrpc) GetWithdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	withdraw, err := h.withdraw.FindById(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Withdraw not found: %v", err)
	}

	pbWithdraw := h.convertToPbWithdraw(withdraw)

	return &pb.WithdrawResponse{Withdraw: pbWithdraw}, nil
}

func (h *withdrawHandleGrpc) GetWithdrawByUsers(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawsResponse, error) {
	withdraws, err := h.withdraw.FindByUsers(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Withdraws not found for user: %v", err)
	}

	pbWithdraws := h.convertToPbWithdraws(withdraws)

	return &pb.WithdrawsResponse{Withdraws: pbWithdraws}, nil
}

func (h *withdrawHandleGrpc) GetWithdrawByUserId(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	withdraw, err := h.withdraw.FindByUsersId(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Withdraw not found for user: %v", err)
	}

	pbWithdraw := h.convertToPbWithdraw(withdraw)

	return &pb.WithdrawResponse{Withdraw: pbWithdraw}, nil
}

func (h *withdrawHandleGrpc) CreateWithdraw(ctx context.Context, req *pb.CreateWithdrawRequest) (*pb.WithdrawResponse, error) {
	request := &requests.CreateWithdrawRequest{
		UserID:         int(req.UserId),
		WithdrawAmount: int(req.WithdrawAmount),
		WithdrawTime:   req.WithdrawTime.AsTime(),
	}

	res, err := h.withdraw.Create(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create withdraw: %v", err)
	}

	pbWithdraw := h.convertToPbWithdraw(res)

	return &pb.WithdrawResponse{Withdraw: pbWithdraw}, nil
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
		return nil, status.Errorf(codes.Internal, "Failed to update withdraw: %v", err)
	}

	pbWithdraw := h.convertToPbWithdraw(res)

	return &pb.WithdrawResponse{Withdraw: pbWithdraw}, nil
}

func (h *withdrawHandleGrpc) DeleteWithdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.DeleteWithdrawResponse, error) {
	err := h.withdraw.Delete(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete withdraw: %v", err)
	}

	return &pb.DeleteWithdrawResponse{Success: true}, nil
}

func (h *withdrawHandleGrpc) convertToPbWithdraws(withdraws []*db.Withdraw) []*pb.Withdraw {
	var pbWithdraws []*pb.Withdraw

	for _, withdraw := range withdraws {
		pbWithdraws = append(pbWithdraws, h.convertToPbWithdraw(withdraw))
	}

	return pbWithdraws
}

func (h *withdrawHandleGrpc) convertToPbWithdraw(withdraw *db.Withdraw) *pb.Withdraw {
	createdAtProto := timestamppb.New(withdraw.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp
	if withdraw.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(withdraw.UpdatedAt.Time)
	}

	return &pb.Withdraw{
		WithdrawId:     int32(withdraw.WithdrawID),
		UserId:         int32(withdraw.UserID),
		WithdrawAmount: int32(withdraw.WithdrawAmount),
		WithdrawTime:   timestamppb.New(withdraw.WithdrawTime),
		CreatedAt:      createdAtProto,
		UpdatedAt:      updatedAtProto,
	}
}
