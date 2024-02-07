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
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	var pbTransfers []*pb.Transfer

	for _, transfer := range transfers {
		createdAtProto := timestamppb.New(transfer.CreatedAt.Time)

		var updatedAtProto *timestamppb.Timestamp

		if transfer.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(transfer.UpdatedAt.Time)
		}

		pbTransfers = append(pbTransfers, &pb.Transfer{
			TransferId:     int32(transfer.TransferID),
			TransferFrom:   int32(transfer.TransferFrom),
			TransferTo:     int32(transfer.TransferTo),
			TransferAmount: int32(transfer.TransferAmount),
			TransferTime:   timestamppb.New(transfer.TransferTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		})
	}

	return &pb.TransfersResponse{Transfers: pbTransfers}, nil
}

func (h *transferHandleGrpc) GetTransfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	transfer, err := h.transfer.FindById(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(transfer.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if transfer.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(transfer.UpdatedAt.Time)
	}

	return &pb.TransferResponse{
		Transfer: &pb.Transfer{
			TransferId:     int32(transfer.TransferID),
			TransferFrom:   int32(transfer.TransferFrom),
			TransferTo:     int32(transfer.TransferTo),
			TransferAmount: int32(transfer.TransferAmount),
			TransferTime:   timestamppb.New(transfer.TransferTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		},
	}, nil
}

func (h *transferHandleGrpc) GetTransferByUsers(ctx context.Context, req *pb.TransferRequest) (*pb.TransfersResponse, error) {
	transfers, err := h.transfer.FindByUsers(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	var pbTransfers []*pb.Transfer

	for _, transfer := range transfers {
		createdAtProto := timestamppb.New(transfer.CreatedAt.Time)

		var updatedAtProto *timestamppb.Timestamp

		if transfer.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(transfer.UpdatedAt.Time)
		}

		pbTransfers = append(pbTransfers, &pb.Transfer{
			TransferId:     int32(transfer.TransferID),
			TransferFrom:   int32(transfer.TransferFrom),
			TransferTo:     int32(transfer.TransferTo),
			TransferAmount: int32(transfer.TransferAmount),
			TransferTime:   timestamppb.New(transfer.TransferTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		})
	}

	return &pb.TransfersResponse{Transfers: pbTransfers}, nil
}

func (h *transferHandleGrpc) GetTransferByUserId(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	transfers, err := h.transfer.FindByUsersId(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(transfers.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if transfers.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(transfers.UpdatedAt.Time)
	}

	return &pb.TransferResponse{
		Transfer: &pb.Transfer{
			TransferId:     int32(transfers.TransferID),
			TransferFrom:   int32(transfers.TransferFrom),
			TransferTo:     int32(transfers.TransferTo),
			TransferAmount: int32(transfers.TransferAmount),
			TransferTime:   timestamppb.New(transfers.TransferTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		},
	}, nil
}

func (h *transferHandleGrpc) CreateTransfer(ctx context.Context, req *pb.CreateTransferRequest) (*pb.TransferResponse, error) {
	request := &requests.CreateTransferRequest{
		TransferFrom:   int(req.TransferFrom),
		TransferTo:     int(req.TransferTo),
		TransferAmount: int(req.TransferAmount),
	}

	res, err := h.transfer.Create(request)

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(res.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if res.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(res.UpdatedAt.Time)
	}

	return &pb.TransferResponse{
		Transfer: &pb.Transfer{
			TransferId:     int32(res.TransferID),
			TransferFrom:   int32(res.TransferFrom),
			TransferTo:     int32(res.TransferTo),
			TransferAmount: int32(res.TransferAmount),
			TransferTime:   timestamppb.New(res.TransferTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		},
	}, nil
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
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	createdAtProto := timestamppb.New(res.CreatedAt.Time)

	var updatedAtProto *timestamppb.Timestamp

	if res.UpdatedAt.Valid {
		updatedAtProto = timestamppb.New(res.UpdatedAt.Time)
	}

	return &pb.TransferResponse{
		Transfer: &pb.Transfer{
			TransferId:     int32(res.TransferID),
			TransferFrom:   int32(res.TransferFrom),
			TransferTo:     int32(res.TransferTo),
			TransferAmount: int32(res.TransferAmount),
			TransferTime:   timestamppb.New(res.TransferTime),
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		},
	}, nil
}

func (h *transferHandleGrpc) DeleteTransfer(ctx context.Context, req *pb.TransferRequest) (*pb.DeleteTransferResponse, error) {
	err := h.transfer.Delete(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.DeleteTransferResponse{
		Success: true,
	}, nil
}
