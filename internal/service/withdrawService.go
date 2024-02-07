package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"database/sql"
	"errors"
	"time"
)

type withdrawService struct {
	withdraw repository.WithdrawRepository
	saldo    repository.SaldoRepository
	user     repository.UserRepository
}

func NewWithdrawService(withdraw repository.WithdrawRepository, saldo repository.SaldoRepository, user repository.UserRepository) *withdrawService {
	return &withdrawService{
		withdraw: withdraw,
		saldo:    saldo,
		user:     user,
	}
}

func (s *withdrawService) FindAll() ([]*db.Withdraw, error) {
	res, err := s.withdraw.FindAll()

	if err != nil {
		return nil, errors.New("failed get withdraw")
	}

	return res, nil
}

func (s *withdrawService) FindById(id int) (*db.Withdraw, error) {
	res, err := s.withdraw.FindById(id)

	if err != nil {
		return nil, errors.New("failed get withdraw")
	}

	return res, nil
}

func (s *withdrawService) FindByUsers(user_id int) ([]*db.Withdraw, error) {
	_, err := s.user.FindById(user_id)

	if err != nil {
		return nil, errors.New("user not found")
	}

	res, err := s.withdraw.FindByUsers(user_id)

	if err != nil {
		return nil, errors.New("failed get withdraw")
	}

	return res, nil
}

func (s *withdrawService) FindByUsersId(user_id int) (*db.Withdraw, error) {
	_, err := s.user.FindById(user_id)

	if err != nil {
		return nil, errors.New("user not found")
	}

	res, err := s.withdraw.FindByUsersId(user_id)

	if err != nil {
		return nil, errors.New("failed get withdraw")
	}

	return res, nil
}

func (s *withdrawService) Create(input *requests.CreateWithdrawRequest) (*db.Withdraw, error) {
	_, err := s.user.FindById(input.UserID)

	if err != nil {
		return nil, errors.New("user not found")
	}

	saldo, err := s.saldo.FindByUserId(input.UserID)

	if err != nil {
		return nil, errors.New("failed get saldo")
	}

	if saldo.TotalBalance < int32(input.WithdrawAmount) {
		return nil, errors.New("balance not enough")
	}

	_, err = s.saldo.Update(&db.UpdateSaldoParams{
		UserID: int32(input.UserID),
		WithdrawAmount: sql.NullInt32{
			Int32: int32(input.WithdrawAmount),
			Valid: true,
		},
		WithdrawTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		TotalBalance: int32(input.WithdrawAmount) + saldo.TotalBalance,
	})

	if err != nil {
		return nil, errors.New("failed update saldo")
	}

	request := &db.CreateWithdrawParams{
		WithdrawAmount: int32(input.WithdrawAmount),
		UserID:         int32(input.UserID),
		WithdrawTime:   time.Now(),
	}

	res, err := s.withdraw.Create(request)

	if err != nil {
		return nil, errors.New("failed create withdraw")
	}

	return res, nil
}

func (s *withdrawService) Update(input *requests.UpdateWithdrawRequest) (*db.Withdraw, error) {
	_, err := s.withdraw.FindById(input.WithdrawID)

	if err != nil {
		return nil, errors.New("withdraw not found")
	}

	_, err = s.user.FindById(input.UserID)

	if err != nil {
		return nil, errors.New("user not found")
	}

	saldo, err := s.saldo.FindByUserId(input.UserID)

	if err != nil {
		return nil, errors.New("failed get saldo")
	}

	if saldo.TotalBalance < int32(input.WithdrawAmount) {
		return nil, errors.New("balance not enough")
	}

	_, err = s.saldo.Update(&db.UpdateSaldoParams{
		UserID: int32(input.UserID),
		WithdrawAmount: sql.NullInt32{
			Int32: int32(input.WithdrawAmount),
			Valid: true,
		},
		WithdrawTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		TotalBalance: int32(input.WithdrawAmount) + saldo.TotalBalance,
	})

	if err != nil {
		return nil, errors.New("failed update saldo")
	}

	request := &db.UpdateWithdrawParams{
		WithdrawID:     int32(input.WithdrawID),
		WithdrawAmount: int32(input.WithdrawAmount),
		WithdrawTime:   time.Now(),
	}

	res, err := s.withdraw.Update(request)

	if err != nil {
		return nil, errors.New("failed create withdraw")
	}

	return res, nil

}

func (s *withdrawService) Delete(id int) error {
	res, err := s.user.FindById(id)

	if err != nil {
		return errors.New("user not found")
	}

	err = s.withdraw.Delete(int(res.UserID))

	if err != nil {
		return errors.New("failed delete withdraw")
	}

	return nil
}
