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

type saldoHandleGrpc struct {
	pb.UnimplementedSaldoServiceServer
	saldo service.SaldoService
}

func NewSaldoHandleGrpc(saldo service.SaldoService) *saldoHandleGrpc {
	return &saldoHandleGrpc{saldo: saldo}
}

func (s *saldoHandleGrpc) GetSaldos(ctx context.Context, req *emptypb.Empty) (*pb.SaldoResponses, error) {
	res, err := s.saldo.FindAll()

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	var pbSaldos []*pb.Saldo

	for _, saldo := range res {
		createdAtProto := timestamppb.New(saldo.CreatedAt.Time)

		var updatedAtProto *timestamppb.Timestamp
		if saldo.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(saldo.UpdatedAt.Time)
		}

		pbSaldos = append(pbSaldos, &pb.Saldo{
			SaldoId:        int32(saldo.SaldoID),
			UserId:         int32(saldo.UserID),
			TotalBalance:   int32(saldo.TotalBalance),
			WithdrawTime:   timestamppb.New(saldo.WithdrawTime.Time),
			WithdrawAmount: saldo.WithdrawAmount.Int32,
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		})
	}

	return &pb.SaldoResponses{Saldos: pbSaldos}, nil
}

func (s *saldoHandleGrpc) GetSaldo(ctx context.Context, req *pb.SaldoRequest) (*pb.SaldoResponse, error) {
	res, err := s.saldo.FindById(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.SaldoResponse{
		Saldo: &pb.Saldo{
			SaldoId:        int32(res.SaldoID),
			UserId:         int32(res.UserID),
			TotalBalance:   int32(res.TotalBalance),
			WithdrawTime:   timestamppb.New(res.WithdrawTime.Time),
			WithdrawAmount: res.WithdrawAmount.Int32,
		},
	}, nil
}

func (s *saldoHandleGrpc) GetSaldoByUsers(ctx context.Context, req *pb.SaldoRequest) (*pb.SaldoResponses, error) {
	res, err := s.saldo.FindByUsersId(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	var pbSaldos []*pb.Saldo

	for _, saldo := range res {
		createdAtProto := timestamppb.New(saldo.CreatedAt.Time)

		var updatedAtProto *timestamppb.Timestamp

		if saldo.UpdatedAt.Valid {
			updatedAtProto = timestamppb.New(saldo.UpdatedAt.Time)
		}

		pbSaldos = append(pbSaldos, &pb.Saldo{
			SaldoId:        int32(saldo.SaldoID),
			UserId:         int32(saldo.UserID),
			TotalBalance:   int32(saldo.TotalBalance),
			WithdrawTime:   timestamppb.New(saldo.WithdrawTime.Time),
			WithdrawAmount: saldo.WithdrawAmount.Int32,
			CreatedAt:      createdAtProto,
			UpdatedAt:      updatedAtProto,
		})
	}

	return &pb.SaldoResponses{Saldos: pbSaldos}, nil
}

func (s *saldoHandleGrpc) GetSaldoByUserId(ctx context.Context, req *pb.SaldoRequest) (*pb.SaldoResponse, error) {
	res, err := s.saldo.FindByUserId(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.SaldoResponse{
		Saldo: &pb.Saldo{
			SaldoId:        int32(res.SaldoID),
			UserId:         int32(res.UserID),
			TotalBalance:   int32(res.TotalBalance),
			WithdrawTime:   timestamppb.New(res.WithdrawTime.Time),
			WithdrawAmount: res.WithdrawAmount.Int32,
		},
	}, nil

}

func (s *saldoHandleGrpc) CreateSaldo(ctx context.Context, req *pb.CreateSaldoRequest) (*pb.SaldoResponse, error) {
	request := &requests.CreateSaldoRequest{
		UserID:       int(req.UserId),
		TotalBalance: int(req.TotalBalance),
	}

	res, err := s.saldo.Create(request)

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.SaldoResponse{
		Saldo: &pb.Saldo{
			SaldoId:        int32(res.SaldoID),
			UserId:         int32(res.UserID),
			TotalBalance:   int32(res.TotalBalance),
			WithdrawTime:   timestamppb.New(res.WithdrawTime.Time),
			WithdrawAmount: res.WithdrawAmount.Int32,
		},
	}, nil
}

func (s *saldoHandleGrpc) UpdateSaldo(ctx context.Context, req *pb.UpdateSaldoRequest) (*pb.SaldoResponse, error) {
	request := &requests.UpdateSaldoRequest{
		SaldoID:        int(req.SaldoId),
		UserID:         int(req.UserId),
		TotalBalance:   int(req.TotalBalance),
		WithdrawAmount: int(req.WithdrawAmount),
		WithdrawTime:   req.WithdrawTime.AsTime(),
	}

	res, err := s.saldo.Update(request)

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.SaldoResponse{
		Saldo: &pb.Saldo{
			SaldoId:        int32(res.SaldoID),
			UserId:         int32(res.UserID),
			TotalBalance:   int32(res.TotalBalance),
			WithdrawTime:   timestamppb.New(res.WithdrawTime.Time),
			WithdrawAmount: res.WithdrawAmount.Int32,
		},
	}, nil
}

func (s *saldoHandleGrpc) DeleteSaldo(ctx context.Context, req *pb.SaldoRequest) (*pb.DeleteSaldoResponse, error) {
	err := s.saldo.Delete(int(req.Id))

	if err != nil {
		return nil, status.Errorf(status.Code(err), err.Error())
	}

	return &pb.DeleteSaldoResponse{
		Success: true,
	}, nil

}
