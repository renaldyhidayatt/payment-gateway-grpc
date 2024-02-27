package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"errors"
	"time"

	"go.uber.org/zap"
)

type transferService struct {
	transfer repository.TransferRepository
	saldo    repository.SaldoRepository
	user     repository.UserRepository
	logger   logger.Logger
}

func NewTransferService(transfer repository.TransferRepository, saldo repository.SaldoRepository, user repository.UserRepository, logger logger.Logger) *transferService {
	return &transferService{
		transfer: transfer,
		saldo:    saldo,
		user:     user,
		logger:   logger,
	}
}

func (s *transferService) FindAll() ([]*db.Transfer, error) {
	res, err := s.transfer.FindAll()

	if err != nil {
		s.logger.Error("Failed to get transfer", zap.Error(err))

		return nil, errors.New("failed get transfer")
	}

	return res, nil
}

func (s *transferService) FindById(id int) (*db.Transfer, error) {
	res, err := s.transfer.FindById(id)

	if err != nil {
		s.logger.Error("Failed to get transfer", zap.Error(err))

		return nil, errors.New("failed get transfer")
	}

	return res, nil
}

func (s *transferService) FindByUsers(user_id int) ([]*db.Transfer, error) {
	_, err := s.user.FindById(user_id)

	if err != nil {
		s.logger.Error("User not found", zap.Error(err))

		return nil, errors.New("user not found")
	}

	res, err := s.transfer.FindByUsers(user_id)

	if err != nil {
		s.logger.Error("Failed to get transfer", zap.Error(err))

		return nil, errors.New("failed get transfer")
	}

	return res, nil
}

func (s *transferService) FindByUsersId(user_id int) (*db.Transfer, error) {
	_, err := s.user.FindById(user_id)

	if err != nil {
		s.logger.Error("User not found", zap.Error(err))

		return nil, errors.New("user not found")
	}

	res, err := s.transfer.FindByUser(user_id)

	if err != nil {
		s.logger.Error("Failed to get transfer", zap.Error(err))

		return nil, errors.New("failed get transfer")
	}

	return res, nil
}

func (s *transferService) Create(req *requests.CreateTransferRequest) (*db.Transfer, error) {
	if req.TransferAmount < 50000 {
		s.logger.Error("Transfer amount must be greater than or equal to 50000")

		return nil, errors.New("transfer amount must be greater than or equal to 50000")
	}

	_, err := s.user.FindById(req.TransferFrom)

	if err != nil {
		s.logger.Error("Sender not found", zap.Error(err))

		return nil, errors.New("sender not found")
	}

	_, err = s.user.FindById(req.TransferTo)

	if err != nil {
		s.logger.Error("Receiver not found", zap.Error(err))

		return nil, errors.New("receiver not found")
	}

	transfer, err := s.transfer.Create(&db.CreateTransferParams{
		TransferFrom:   int32(req.TransferFrom),
		TransferTo:     int32(req.TransferTo),
		TransferAmount: int32(req.TransferAmount),
		TransferTime:   time.Now(),
	})

	if err != nil {
		s.logger.Error("Failed to create transfer", zap.Error(err))

		return nil, errors.New("failed create transfer")
	}

	senderSaldo, err := s.saldo.FindByUserId(req.TransferFrom)

	if err != nil {
		s.logger.Error("Failed to get sender saldo", zap.Error(err))

		return nil, errors.New("failed get sender saldo")
	}

	senderNewBalance := senderSaldo.TotalBalance - int32(req.TransferAmount)

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		UserID:       int32(req.TransferFrom),
		TotalBalance: senderNewBalance,
	})

	if err != nil {
		s.logger.Error("Failed to update sender saldo", zap.Error(err))
		errRollback := s.rollbackTransfer(transfer)

		if errRollback != nil {
			s.logger.Error("Failed to rollback sender", zap.Error(errRollback))
		}
		return nil, errors.New("failed update sender saldo")
	}

	receiverSaldo, err := s.saldo.FindByUserId(req.TransferTo)
	if err != nil {
		s.logger.Error("Failed to get receiver saldo", zap.Error(err))

		errRollback := s.rollbackTransfer(transfer)

		if errRollback != nil {
			s.logger.Error("Failed to rollback sender and receiver", zap.Error(errRollback))
		}
		return nil, errors.New("failed get receiver saldo")
	}

	receiverNewBalance := receiverSaldo.TotalBalance + int32(req.TransferAmount)

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		UserID:       int32(req.TransferTo),
		TotalBalance: receiverNewBalance,
	})

	if err != nil {
		s.logger.Error("Failed to update receiver saldo", zap.Error(err))

		errRollback := s.rollbackTransfer(transfer)

		if errRollback != nil {
			s.logger.Error("Failed to rollback receiver", zap.Error(errRollback))
		}
		return nil, errors.New("failed update receiver saldo")
	}

	return transfer, nil
}

func (s *transferService) Update(req *requests.UpdateTransferRequest) (*db.Transfer, error) {
	if req.TransferAmount < 50000 {
		s.logger.Error("Transfer amount must be greater than or equal to 50000")

		return nil, errors.New("transfer amount must be greater than or equal to 50000")
	}

	_, err := s.user.FindById(req.TransferFrom)

	if err != nil {
		s.logger.Error("Sender not found", zap.Error(err))

		return nil, errors.New("sender not found")
	}

	_, err = s.user.FindById(req.TransferTo)

	if err != nil {

		s.logger.Error("Receiver not found", zap.Error(err))
		return nil, errors.New("receiver not found")
	}

	transfer, err := s.transfer.Update(&db.UpdateTransferParams{
		TransferID:     int32(req.TransferID),
		TransferAmount: int32(req.TransferAmount),
	})

	if err != nil {
		s.logger.Error("Failed to update transfer", zap.Error(err))

		return nil, errors.New("failed to update transfer")
	}

	senderSaldo, err := s.saldo.FindByUserId(req.TransferFrom)

	if err != nil {
		s.logger.Error("Failed to get sender saldo", zap.Error(err))

		return nil, errors.New("failed to get sender saldo")
	}

	senderNewBalance := senderSaldo.TotalBalance - int32(req.TransferAmount)

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		UserID:       int32(req.TransferFrom),
		TotalBalance: senderNewBalance,
	})

	if err != nil {
		s.logger.Error("Failed to update sender saldo", zap.Error(err))

		errRollback := s.rollbackTransferUpdate(int(req.TransferID), int(senderSaldo.TotalBalance))

		if errRollback != nil {
			s.logger.Error("Failed to rollback transfer", zap.Error(errRollback))
		}

		return nil, errors.New("failed update sender saldo")
	}

	receiverSaldo, err := s.saldo.FindByUserId(req.TransferTo)
	if err != nil {
		s.logger.Error("Failed to get receiver saldo", zap.Error(err))

		errRollback := s.rollbackTransferUpdate(int(req.TransferID), int(senderSaldo.TotalBalance))

		if errRollback != nil {
			s.logger.Error("Failed to rollback transfer", zap.Error(errRollback))
		}
		return nil, errors.New("failed get receiver saldo")
	}

	receiverNewBalance := receiverSaldo.TotalBalance + int32(req.TransferAmount)

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		UserID:       int32(req.TransferTo),
		TotalBalance: receiverNewBalance,
	})

	if err != nil {

		s.logger.Error("Failed to update receiver saldo", zap.Error(err))

		errRollback := s.rollbackTransferUpdate(int(req.TransferID), int(senderSaldo.TotalBalance))

		if errRollback != nil {
			s.logger.Error("Failed to rollback transfer", zap.Error(errRollback))
		}

		return nil, errors.New("failed update receiver saldo")
	}

	return transfer, nil
}

func (s *transferService) Delete(id int) error {
	res, err := s.user.FindById(id)

	if err != nil {
		s.logger.Error("User not found", zap.Error(err))

		return errors.New("user not found")
	}

	err = s.transfer.Delete(int(res.UserID))

	if err != nil {
		s.logger.Error("Failed to delete transfer", zap.Error(err))
		return errors.New("failed delete transfer")
	}

	return nil
}

func (s *transferService) rollbackTransfer(transfer *db.Transfer) error {
	err := s.transfer.Delete(int(transfer.TransferID))

	if err != nil {
		return errors.New("failed to rollback transfer")
	}

	return nil
}

func (s *transferService) rollbackTransferUpdate(transerId, amount int) error {
	_, err := s.transfer.Update(&db.UpdateTransferParams{
		TransferID:     int32(transerId),
		TransferAmount: int32(amount),
	})

	if err != nil {
		return errors.New("failed to rollback transfer update")
	}

	return nil
}
