package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type withdrawService struct {
	userRepository     repository.UserRepository
	saldoRepository    repository.SaldoRepository
	withdrawRepository repository.WithdrawRepository
	logger             logger.LoggerInterface
	mapping            responsemapper.WithdrawResponseMapper
}

func NewWithdrawService(
	userRepository repository.UserRepository,
	withdrawRepository repository.WithdrawRepository, saldoRepository repository.SaldoRepository, logger logger.LoggerInterface, mapping responsemapper.WithdrawResponseMapper) *withdrawService {
	return &withdrawService{
		userRepository:     userRepository,
		saldoRepository:    saldoRepository,
		withdrawRepository: withdrawRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *withdrawService) FindAll(page int, pageSize int, search string) ([]*response.WithdrawResponse, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	withdraws, totalRecords, err := s.withdrawRepository.FindAll(search, page, pageSize)

	if err != nil {
		s.logger.Error("failed to fetch withdraws", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraws",
		}
	}

	withdrawResponse := s.mapping.ToWithdrawsResponse(withdraws)

	return withdrawResponse, totalRecords, nil
}

func (s *withdrawService) FindById(withdrawID int) (*response.WithdrawResponse, *response.ErrorResponse) {
	withdraw, err := s.withdrawRepository.FindById(withdrawID)
	if err != nil {
		s.logger.Error("failed to find withdraw by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw record by ID.",
		}
	}
	so := s.mapping.ToWithdrawResponse(withdraw)

	return so, nil
}

func (s *withdrawService) FindByCardNumber(card_number string) ([]*response.WithdrawResponse, *response.ErrorResponse) {
	withdrawRecords, err := s.withdrawRepository.FindByCardNumber(card_number)

	if err != nil {
		s.logger.Error("Failed to fetch withdraw records by card number", zap.Error(err), zap.String("card_number", card_number))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw records for the given card number",
		}
	}

	withdrawResponses := s.mapping.ToWithdrawsResponse(withdrawRecords)

	return withdrawResponses, nil
}

func (s *withdrawService) FindByActive(page int, pageSize int, search string) ([]*response.WithdrawResponseDeleteAt, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	withdraws, totalRecords, err := s.withdrawRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch active withdraw records", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active withdraw records",
		}
	}

	withdrawResponses := s.mapping.ToWithdrawsResponseDeleteAt(withdraws)

	return withdrawResponses, totalRecords, nil
}

func (s *withdrawService) FindByTrashed(page int, pageSize int, search string) ([]*response.WithdrawResponseDeleteAt, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	withdraws, totalRecords, err := s.withdrawRepository.FindByTrashed(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch trashed withdraw records", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed withdraw records",
		}
	}

	withdrawResponses := s.mapping.ToWithdrawsResponseDeleteAt(withdraws)

	return withdrawResponses, totalRecords, nil
}

func (s *withdrawService) Create(request *requests.CreateWithdrawRequest) (*response.WithdrawResponse, *response.ErrorResponse) {
	saldo, err := s.saldoRepository.FindByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("Failed to find saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch saldo for the user.",
		}
	}

	if saldo == nil {
		s.logger.Error("Saldo not found for user", zap.String("cardNumber", request.CardNumber))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found for the specified user ID.",
		}
	}

	if saldo.TotalBalance < request.WithdrawAmount {
		s.logger.Error("Insufficient balance for user", zap.String("cardNumber", request.CardNumber), zap.Int("requested", request.WithdrawAmount))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for withdrawal.",
		}
	}

	newTotalBalance := saldo.TotalBalance - request.WithdrawAmount

	updateData := &requests.UpdateSaldoWithdraw{
		CardNumber:     request.CardNumber,
		TotalBalance:   newTotalBalance,
		WithdrawAmount: &request.WithdrawAmount,
		WithdrawTime:   &request.WithdrawTime,
	}

	_, err = s.saldoRepository.UpdateSaldoWithdraw(updateData)
	if err != nil {
		s.logger.Error("Failed to update saldo after withdrawal", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo after withdrawal.",
		}
	}

	withdrawRecord, err := s.withdrawRepository.CreateWithdraw(request)
	if err != nil {
		s.logger.Error("Failed to create withdraw record", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create withdraw record.",
		}
	}

	so := s.mapping.ToWithdrawResponse(withdrawRecord)

	return so, nil
}

func (s *withdrawService) Update(request *requests.UpdateWithdrawRequest) (*response.WithdrawResponse, *response.ErrorResponse) {
	_, err := s.withdrawRepository.FindById(request.WithdrawID)
	if err != nil {
		s.logger.Error("Failed to find withdraw record by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Withdraw record not found.",
		}
	}

	saldo, err := s.saldoRepository.FindByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("Failed to fetch saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch saldo for the user.",
		}
	}

	if saldo.TotalBalance < request.WithdrawAmount {
		s.logger.Error("Insufficient balance for user", zap.String("cardNumber", request.CardNumber))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for withdrawal update.",
		}
	}

	// Update saldo baru
	newTotalBalance := saldo.TotalBalance - request.WithdrawAmount
	updateSaldoData := &requests.UpdateSaldoWithdraw{
		CardNumber:     saldo.CardNumber,
		TotalBalance:   newTotalBalance,
		WithdrawAmount: &request.WithdrawAmount,
		WithdrawTime:   &request.WithdrawTime,
	}

	_, err = s.saldoRepository.UpdateSaldoWithdraw(updateSaldoData)
	if err != nil {
		s.logger.Error("Failed to update saldo balance", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo balance.",
		}
	}

	updatedWithdraw, err := s.withdrawRepository.UpdateWithdraw(request)
	if err != nil {
		rollbackData := &requests.UpdateSaldoBalance{
			CardNumber:   saldo.CardNumber,
			TotalBalance: saldo.TotalBalance,
		}
		_, rollbackErr := s.saldoRepository.UpdateSaldoBalance(rollbackData)
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback saldo after withdraw update failure", zap.Error(rollbackErr))
		}
		s.logger.Error("Failed to update withdraw record", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update withdraw record.",
		}
	}

	so := s.mapping.ToWithdrawResponse(updatedWithdraw)

	return so, nil
}

func (s *withdrawService) TrashedWithdraw(withdraw_id int) (*response.WithdrawResponse, *response.ErrorResponse) {
	s.logger.Debug("Trashing withdraw", zap.Int("withdraw_id", withdraw_id))

	res, err := s.withdrawRepository.TrashedWithdraw(withdraw_id)
	if err != nil {
		s.logger.Error("Failed to trash withdraw", zap.Error(err), zap.Int("withdraw_id", withdraw_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash withdraw",
		}
	}

	withdrawResponse := s.mapping.ToWithdrawResponse(res)

	s.logger.Debug("Successfully trashed withdraw", zap.Int("withdraw_id", withdraw_id))

	return withdrawResponse, nil
}

func (s *withdrawService) RestoreWithdraw(withdraw_id int) (*response.WithdrawResponse, *response.ErrorResponse) {
	s.logger.Debug("Restoring withdraw", zap.Int("withdraw_id", withdraw_id))

	res, err := s.withdrawRepository.RestoreWithdraw(withdraw_id)
	if err != nil {
		s.logger.Error("Failed to restore withdraw", zap.Error(err), zap.Int("withdraw_id", withdraw_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore withdraw",
		}
	}

	withdrawResponse := s.mapping.ToWithdrawResponse(res)

	s.logger.Debug("Successfully restored withdraw", zap.Int("withdraw_id", withdraw_id))

	return withdrawResponse, nil
}

func (s *withdrawService) DeleteWithdrawPermanent(withdraw_id int) (interface{}, *response.ErrorResponse) {
	s.logger.Debug("Deleting withdraw permanently", zap.Int("withdraw_id", withdraw_id))

	err := s.withdrawRepository.DeleteWithdrawPermanent(withdraw_id)
	if err != nil {
		s.logger.Error("Failed to delete withdraw permanently", zap.Error(err), zap.Int("withdraw_id", withdraw_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete withdraw permanently",
		}
	}

	s.logger.Debug("Successfully deleted withdraw permanently", zap.Int("withdraw_id", withdraw_id))

	return nil, nil
}
