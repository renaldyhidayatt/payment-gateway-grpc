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

type transferHandleGrpc struct {
	pb.UnimplementedTransferServiceServer
	transfer service.TransferService
}

func NewTransferHandleGrpc(transfer service.TransferService) *transferHandleGrpc {
	return &transferHandleGrpc{
		transfer: transfer,
	}
}

func (h *transferHandleGrpc) GetTransfers(ctx context.Context, req *emptypb.Empty) (*pb.TransfersResponse, error) {
	transfers, err := h.transfer.FindAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get transfers: %v", err)
	}

	pbTransfers := h.convertToPbTransfers(transfers)

	return &pb.TransfersResponse{Transfers: pbTransfers}, nil
}

func (h *transferHandleGrpc) GetTransfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	transfer, err := h.transfer.FindById(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Transfer not found: %v", err)
	}

	pbTransfer := h.convertToPbTransfer(transfer)

	return &pb.TransferResponse{Transfer: pbTransfer}, nil
}

func (h *transferHandleGrpc) GetTransferByUsers(ctx context.Context, req *pb.TransferRequest) (*pb.TransfersResponse, error) {
	transfers, err := h.transfer.FindByUsers(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Transfers not found for user: %v", err)
	}

	pbTransfers := h.convertToPbTransfers(transfers)

	return &pb.TransfersResponse{Transfers: pbTransfers}, nil
}

func (h *transferHandleGrpc) GetTransferByUserId(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	transfer, err := h.transfer.FindByUsersId(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Transfer not found for user: %v", err)
	}

	pbTransfer := h.convertToPbTransfer(transfer)

	return &pb.TransferResponse{Transfer: pbTransfer}, nil
}

func (h *transferHandleGrpc) CreateTransfer(ctx context.Context, req *pb.CreateTransferRequest) (*pb.TransferResponse, error) {
	request := &requests.CreateTransferRequest{
		TransferFrom:   int(req.TransferFrom),
		TransferTo:     int(req.TransferTo),
		TransferAmount: int(req.TransferAmount),
	}

	res, err := h.transfer.Create(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create transfer: %v", err)
	}

	pbTransfer := h.convertToPbTransfer(res)

	return &pb.TransferResponse{Transfer: pbTransfer}, nil
}

func (h *transferHandleGrpc) UpdateTransfer(ctx context.Context, req *pb.UpdateTransferRequest) (*pb.TransferResponse, error) {
	request := &requests.UpdateTransferRequest{
		TransferID:     int(req.Id),
		TransferFrom:   int(req.TransferFrom),
		TransferTo:     int(req.TransferTo),
		TransferAmount: int(req.TransferAmount),
	}

	res, err := h.transfer.Update(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update transfer: %v", err)
	}

	pbTransfer := h.convertToPbTransfer(res)

	return &pb.TransferResponse{Transfer: pbTransfer}, nil
}

func (h *transferHandleGrpc) DeleteTransfer(ctx context.Context, req *pb.TransferRequest) (*pb.DeleteTransferResponse, error) {
	err := h.transfer.Delete(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete transfer: %v", err)
	}

	return &pb.DeleteTransferResponse{Success: true}, nil
}

// Convert Database to Proto
func (h *transferHandleGrpc) convertToPbTransfers(transfers []*db.Transfer) []*pb.Transfer {
	var pbTransfers []*pb.Transfer

	for _, transfer := range transfers {
		pbTransfers = append(pbTransfers, h.convertToPbTransfer(transfer))
	}

	return pbTransfers
}

func (h *transferHandleGrpc) convertToPbTransfer(transfer *db.Transfer) *pb.Transfer {
	createdAtProto := timestamppb.New(transfer.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp
	if transfer.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(transfer.UpdatedAt.Time)
	}

	return &pb.Transfer{
		TransferId:     int32(transfer.TransferID),
		TransferFrom:   int32(transfer.TransferFrom),
		TransferTo:     int32(transfer.TransferTo),
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   timestamppb.New(transfer.TransferTime),
		CreatedAt:      createdAtProto,
		UpdatedAt:      updatedAtProto,
	}
}
