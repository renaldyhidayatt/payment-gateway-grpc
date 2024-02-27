package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

type saldoService struct {
	user   repository.UserRepository
	saldo  repository.SaldoRepository
	logger logger.Logger
}

func NewSaldoService(saldo repository.SaldoRepository, user repository.UserRepository, logger logger.Logger) *saldoService {
	return &saldoService{
		saldo:  saldo,
		user:   user,
		logger: logger,
	}
}

func (s *saldoService) FindAll() ([]*db.Saldo, error) {
	res, err := s.saldo.FindAll()

	if err != nil {
		s.logger.Error("Failed to get saldo", zap.Error(err))
		return nil, errors.New("failed to get saldo")
	}

	return res, nil
}

func (s *saldoService) FindById(id int) (*db.Saldo, error) {
	res, err := s.saldo.FindById(id)

	if err != nil {
		s.logger.Error("Failed to get saldo by ID", zap.Error(err))
		return nil, errors.New("failed to get saldo")
	}

	return res, nil
}

func (s *saldoService) FindByUserId(id int) (*db.Saldo, error) {
	_, err := s.user.FindById(id)
	if err != nil {
		s.logger.Error("User not found", zap.Error(err))
		return nil, errors.New("user not found")
	}

	res, err := s.saldo.FindByUserId(id)
	if err != nil {
		s.logger.Error("Failed to get saldo by user ID", zap.Error(err))
		return nil, errors.New("failed to get saldo")
	}
	return res, nil
}

func (s *saldoService) FindByUsersId(id int) ([]*db.Saldo, error) {
	_, err := s.user.FindById(id)
	if err != nil {
		s.logger.Error("User not found", zap.Error(err))
		return nil, errors.New("user not found")
	}

	res, err := s.saldo.FindByUsersId(id)

	if err != nil {
		s.logger.Error("Failed to get saldo by users ID", zap.Error(err))
		return nil, errors.New("failed to get saldo")
	}

	return res, nil
}

func (s *saldoService) Create(input *requests.CreateSaldoRequest) (*db.Saldo, error) {
	_, err := s.user.FindById(input.UserID)

	if err != nil {
		s.logger.Error("User not found", zap.Error(err))
		return nil, errors.New("user not found")
	}

	if input.TotalBalance < 50000 {
		s.logger.Error("Total balance must be greater than or equal to 50000")
		return nil, errors.New("total balance must be greater than or equal to 50000")
	}

	request := &db.CreateSaldoParams{
		UserID:       int32(input.UserID),
		TotalBalance: int32(input.TotalBalance),
	}

	res, err := s.saldo.Create(request)

	if err != nil {
		s.logger.Error("Failed to create saldo", zap.Error(err))
		return nil, errors.New("failed create saldo")
	}

	return res, nil
}

func (s *saldoService) Update(input *requests.UpdateSaldoRequest) (*db.Saldo, error) {
	_, err := s.user.FindById(input.UserID)

	if err != nil {
		s.logger.Error("User not found", zap.Error(err))
		return nil, errors.New("user not found")
	}

	if input.TotalBalance < 50000 {
		s.logger.Error("Total balance must be greater than or equal to 50000")
		return nil, errors.New("total balance must be greater than or equal to 50000")
	}

	request := &db.UpdateSaldoParams{
		UserID: int32(input.UserID),
		WithdrawAmount: sql.NullInt32{
			Int32: int32(input.WithdrawAmount),
			Valid: true,
		},
		WithdrawTime: sql.NullTime{
			Time:  input.WithdrawTime,
			Valid: true,
		},
	}

	res, err := s.saldo.Update(request)

	if err != nil {
		s.logger.Error("Failed to update saldo", zap.Error(err))
		return nil, errors.New("failed update saldo")
	}

	return res, nil
}

func (s *saldoService) Delete(id int) error {
	res, err := s.user.FindById(id)

	if err != nil {
		s.logger.Error("User not found", zap.Error(err))
		return errors.New("user not found")
	}

	err = s.saldo.Delete(int(res.UserID))
	if err != nil {
		s.logger.Error("Failed to delete saldo", zap.Error(err))
		return errors.New("failed delete saldo")
	}

	return nil
}
