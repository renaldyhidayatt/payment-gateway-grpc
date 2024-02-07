package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"errors"
	"time"
)

type transferService struct {
	transfer repository.TransferRepository
	saldo    repository.SaldoRepository
	user     repository.UserRepository
}

func NewTransferService(transfer repository.TransferRepository, saldo repository.SaldoRepository, user repository.UserRepository) *transferService {
	return &transferService{
		transfer: transfer,
		saldo:    saldo,
		user:     user,
	}
}

func (s *transferService) FindAll() ([]*db.Transfer, error) {
	res, err := s.transfer.FindAll()

	if err != nil {
		return nil, errors.New("failed get transfer")
	}

	return res, nil
}

func (s *transferService) FindById(id int) (*db.Transfer, error) {
	res, err := s.transfer.FindById(id)

	if err != nil {
		return nil, errors.New("failed get transfer")
	}

	return res, nil
}

func (s *transferService) FindByUsers(user_id int) ([]*db.Transfer, error) {
	_, err := s.user.FindById(user_id)

	if err != nil {
		return nil, errors.New("user not found")
	}

	res, err := s.transfer.FindByUsers(user_id)

	if err != nil {
		return nil, errors.New("failed get transfer")
	}

	return res, nil
}

func (s *transferService) FindByUsersId(user_id int) (*db.Transfer, error) {
	_, err := s.user.FindById(user_id)

	if err != nil {
		return nil, errors.New("user not found")
	}

	res, err := s.transfer.FindByUser(user_id)

	if err != nil {
		return nil, errors.New("failed get transfer")
	}

	return res, nil
}

func (s *transferService) Create(req *requests.CreateTransferRequest) (*db.Transfer, error) {
	if req.TransferAmount < 50000 {
		return nil, errors.New("transfer amount must be greater than or equal to 50000")
	}

	_, err := s.user.FindById(req.TransferFrom)

	if err != nil {
		return nil, errors.New("sender not found")
	}

	_, err = s.user.FindById(req.TransferTo)

	if err != nil {
		return nil, errors.New("receiver not found")
	}

	transfer, err := s.transfer.Create(&db.CreateTransferParams{
		TransferFrom:   int32(req.TransferFrom),
		TransferTo:     int32(req.TransferTo),
		TransferAmount: int32(req.TransferAmount),
		TransferTime:   time.Now(),
	})

	if err != nil {
		return nil, errors.New("failed create transfer")
	}

	senderSaldo, err := s.saldo.FindByUserId(req.TransferFrom)

	if err != nil {
		return nil, errors.New("failed get sender saldo")
	}

	receiverSaldo, err := s.saldo.FindByUserId(req.TransferTo)

	if err != nil {
		return nil, errors.New("failed get receiver saldo")
	}

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		TotalBalance: senderSaldo.TotalBalance - int32(req.TransferAmount),
	})

	if err != nil {
		return nil, errors.New("failed update sender saldo")
	}

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		TotalBalance: receiverSaldo.TotalBalance + int32(req.TransferAmount),
	})

	if err != nil {
		return nil, errors.New("failed update receiver saldo")
	}

	return transfer, nil
}

func (s *transferService) Update(req *requests.UpdateTransferRequest) (*db.Transfer, error) {
	if req.TransferAmount < 50000 {
		return nil, errors.New("transfer amount must be greater than or equal to 50000")
	}

	_, err := s.user.FindById(req.TransferFrom)

	if err != nil {
		return nil, errors.New("sender not found")
	}

	_, err = s.user.FindById(req.TransferTo)

	if err != nil {
		return nil, errors.New("receiver not found")
	}

	transfer, err := s.transfer.Update(&db.UpdateTransferParams{
		TransferID:     int32(req.TransferID),
		TransferAmount: int32(req.TransferAmount),
	})

	if err != nil {
		return nil, errors.New("failed update transfer")
	}

	senderSaldo, err := s.saldo.FindByUserId(req.TransferFrom)

	if err != nil {
		return nil, errors.New("failed get sender saldo")
	}

	receiverSaldo, err := s.saldo.FindByUserId(req.TransferTo)

	if err != nil {
		return nil, errors.New("failed get receiver saldo")
	}

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		TotalBalance: senderSaldo.TotalBalance - int32(req.TransferAmount),
	})

	if err != nil {
		return nil, errors.New("failed update sender saldo")
	}

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		TotalBalance: receiverSaldo.TotalBalance + int32(req.TransferAmount),
	})

	if err != nil {
		return nil, errors.New("failed update receiver saldo")
	}

	return transfer, nil
}

func (s *transferService) Delete(id int) error {
	res, err := s.user.FindById(id)

	if err != nil {
		return errors.New("user not found")
	}

	err = s.transfer.Delete(int(res.UserID))

	if err != nil {
		return errors.New("failed delete transfer")
	}

	return nil

}
