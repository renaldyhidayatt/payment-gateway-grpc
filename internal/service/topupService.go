package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	methodtopup "MamangRust/paymentgatewaygrpc/pkg/method_topup"
	"errors"
	"time"

	"go.uber.org/zap"
)

type topupService struct {
	topup  repository.TopupRepository
	saldo  repository.SaldoRepository
	user   repository.UserRepository
	logger logger.Logger
}

func NewTopupService(topup repository.TopupRepository, saldo repository.SaldoRepository, user repository.UserRepository, logger logger.Logger) *topupService {
	return &topupService{
		topup:  topup,
		saldo:  saldo,
		user:   user,
		logger: logger,
	}
}

func (s *topupService) FindAll() ([]*db.Topup, error) {
	res, err := s.topup.FindAll()
	if err != nil {
		s.logger.Error("Failed to get topup", zap.Error(err))
		return nil, errors.New("failed get topup")
	}
	return res, nil
}

func (s *topupService) FindById(id int) (*db.Topup, error) {
	res, err := s.topup.FindById(id)
	if err != nil {

		s.logger.Error("Failed to get topup", zap.Error(err))

		return nil, errors.New("failed get topup")
	}
	return res, nil
}

func (s *topupService) FindByUsers(user_id int) ([]*db.Topup, error) {
	_, err := s.user.FindById(user_id)
	if err != nil {

		s.logger.Error("User not found", zap.Error(err))

		return nil, errors.New("user not found")
	}

	res, err := s.topup.FindByUsers(user_id)
	if err != nil {

		s.logger.Error("Failed to get topup", zap.Error(err))

		return nil, errors.New("failed get topup")
	}

	return res, nil
}

func (s *topupService) FindByUsersId(user_id int) (*db.Topup, error) {
	_, err := s.user.FindById(user_id)
	if err != nil {

		s.logger.Error("User not found", zap.Error(err))

		return nil, errors.New("user not found")
	}

	res, err := s.topup.FindByUsersId(user_id)
	if err != nil {

		s.logger.Error("Failed to get topup", zap.Error(err))

		return nil, errors.New("failed get topup")
	}

	return res, nil
}

func (s *topupService) Create(input *requests.CreateTopupRequest) (*db.Topup, error) {
	if input.TopupAmount < 50000 {

		s.logger.Error("Failed to create topup", zap.Error(errors.New("topup amount must be greater than or equal to 50000")))

		return nil, errors.New("topup amount must be greater than or equal to 50000")
	}

	_, err := s.user.FindById(input.UserID)
	if err != nil {

		s.logger.Error("Failed to create topup", zap.Error(errors.New("user not found")))

		return nil, errors.New("user not found")
	}

	if !methodtopup.PaymentMethodValidator(input.TopupMethod) {

		s.logger.Error("Failed to create topup", zap.Error(errors.New("payment method not found")))

		return nil, errors.New("payment method not found")
	}

	request := &db.CreateTopupParams{
		TopupNo:     input.TopupNo,
		TopupAmount: int32(input.TopupAmount),
		TopupMethod: input.TopupMethod,
		UserID:      int32(input.UserID),
		TopupTime:   time.Now(),
	}

	res, err := s.topup.Create(request)
	if err != nil {

		s.logger.Error("Failed to create topup", zap.Error(err))

		return nil, errors.New("failed create topup")
	}

	saldo, err := s.saldo.FindByUserId(input.UserID)
	if err != nil {
		errRollback := s.topup.Delete(int(res.TopupID))

		if errRollback != nil {

			s.logger.Error("Failed to create topup and failed to rollback topup", zap.Error(errRollback))
			return nil, errors.New("failed create topup and failed to rollback topup")
		}

		s.logger.Error("Failed to create topup", zap.Error(errors.New("failed get saldo")))

		return nil, errors.New("failed get saldo")
	}

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		UserID:       int32(input.UserID),
		TotalBalance: saldo.TotalBalance + res.TopupAmount,
	})
	if err != nil {
		errRollback := s.topup.Delete(int(res.TopupID))
		if errRollback != nil {

			s.logger.Error("Failed to create topup and failed to rollback topup and failed to update saldo", zap.Error(errRollback))

			return nil, errors.New("failed create topup and failed to rollback topup and failed to update saldo")
		}

		s.logger.Error("Failed to update saldo", zap.Error(err))

		return nil, errors.New("failed update saldo: " + err.Error())
	}

	return res, nil
}

func (s *topupService) UpdateTopup(input *requests.UpdateTopupRequest) (*db.Topup, error) {
	if input.TopupAmount < 50000 {

		s.logger.Error("Topup amount must be greater than or equal to 50000")

		return nil, errors.New("topup amount must be greater than or equal to 50000")
	}

	_, err := s.user.FindById(input.UserID)
	if err != nil {

		s.logger.Error("User not found", zap.Error(err))

		return nil, errors.New("user not found")
	}

	res, err := s.topup.FindById(input.TopupID)

	if err != nil {
		s.logger.Error("Topup not found", zap.Error(err))

		return nil, errors.New("topup not found")
	}

	saldo, err := s.saldo.FindByUserId(input.UserID)

	if err != nil {
		s.logger.Error("Failed to get saldo", zap.Error(err))

		return nil, errors.New("failed get saldo")
	}

	if !methodtopup.PaymentMethodValidator(input.TopupMethod) {
		s.logger.Error("Payment method not found")

		return nil, errors.New("payment method not found")
	}

	topup, err := s.topup.Update(&db.UpdateTopupParams{
		TopupID:     int32(input.TopupID),
		TopupAmount: int32(input.TopupAmount),
		TopupMethod: input.TopupMethod,
		TopupTime:   time.Now(),
	})

	if err != nil {
		s.logger.Error("Failed to update topup", zap.Error(err))

		return nil, errors.New("failed update topup: " + err.Error())
	}

	_, err = s.saldo.UpdateSaldoBalance(&db.UpdateSaldoBalanceParams{
		UserID:       int32(input.UserID),
		TotalBalance: saldo.TotalBalance - res.TopupAmount,
	})
	if err != nil {
		_, errRollback := s.topup.Update(&db.UpdateTopupParams{
			TopupID:     int32(input.TopupID),
			TopupAmount: topup.TopupAmount,
			TopupMethod: topup.TopupMethod,
			TopupTime:   topup.TopupTime,
		})
		if errRollback != nil {
			s.logger.Error("Failed to rollback topup", zap.Error(errRollback))

			return nil, errors.New("failed update topup and failed to rollback saldo")
		}

		s.logger.Error("Failed to update saldo", zap.Error(err))

		return nil, errors.New("failed update saldo")
	}

	return topup, nil
}

func (s *topupService) DeleteTopup(id int) error {
	res, err := s.user.FindById(id)

	if err != nil {
		s.logger.Error("failed get user id ", zap.Error(err))
		return errors.New("user not found")
	}

	err = s.topup.Delete(int(res.UserID))

	if err != nil {
		s.logger.Error("failed delete topup ", zap.Error(err))

		return errors.New("failed delete topup")
	}

	return nil

}
