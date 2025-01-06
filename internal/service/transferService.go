package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type transferService struct {
	userRepository     repository.UserRepository
	cardRepository     repository.CardRepository
	saldoRepository    repository.SaldoRepository
	transferRepository repository.TransferRepository
	logger             logger.LoggerInterface
	mapping            responsemapper.TransferResponseMapper
}

func NewTransferService(
	userRepository repository.UserRepository,
	cardRepository repository.CardRepository,
	transferRepository repository.TransferRepository,
	saldoRepository repository.SaldoRepository, logger logger.LoggerInterface, mapping responsemapper.TransferResponseMapper) *transferService {
	return &transferService{
		userRepository:     userRepository,
		transferRepository: transferRepository,
		saldoRepository:    saldoRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *transferService) FindAll(page int, pageSize int, search string) ([]*response.TransferResponse, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	transfers, totalRecords, err := s.transferRepository.FindAll(search, page, pageSize)

	if err != nil {
		s.logger.Error("failed to fetch transfers", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfers",
		}
	}

	so := s.mapping.ToTransfersResponse(transfers)

	return so, totalRecords, nil
}

func (s *transferService) FindById(transferId int) (*response.TransferResponse, *response.ErrorResponse) {
	transfer, err := s.transferRepository.FindById(transferId)
	if err != nil {
		s.logger.Error("failed to find transfer by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transfer not found",
		}
	}

	so := s.mapping.ToTransferResponse(transfer)

	return so, nil
}

func (s *transferService) FindByActive(page int, pageSize int, search string) ([]*response.TransferResponseDeleteAt, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	transfers, totalRecords, err := s.transferRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch active transaction records", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No active transaction records found",
		}
	}

	so := s.mapping.ToTransfersResponseDeleteAt(transfers)

	s.logger.Debug("Successfully fetched active transaction records", zap.Int("record_count", len(transfers)))

	return so, totalRecords, nil
}

func (s *transferService) FindByTrashed(page int, pageSize int, search string) ([]*response.TransferResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Info("Fetching trashed transaction records")

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	transfers, totalRecords, err := s.transferRepository.FindByTrashed(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch trashed transaction records", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No trashed transaction records found",
		}
	}

	so := s.mapping.ToTransfersResponseDeleteAt(transfers)

	s.logger.Debug("Successfully fetched trashed transaction records", zap.Int("record_count", len(transfers)))

	return so, totalRecords, nil
}

func (s *transferService) FindTransferByTransferFrom(transfer_from string) ([]*response.TransferResponse, *response.ErrorResponse) {
	res, err := s.transferRepository.FindTransferByTransferFrom(transfer_from)
	if err != nil {
		s.logger.Error("Failed to fetch transfers by transfer_from", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfers by transfer_from",
		}
	}

	so := s.mapping.ToTransfersResponse(res)

	return so, nil
}

func (s *transferService) FindTransferByTransferTo(transfer_to string) ([]*response.TransferResponse, *response.ErrorResponse) {
	res, err := s.transferRepository.FindTransferByTransferTo(transfer_to)
	if err != nil {
		s.logger.Error("Failed to fetch transfers by transfer_to", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfers by transfer_to",
		}
	}

	so := s.mapping.ToTransfersResponse(res)

	return so, nil
}

func (s *transferService) CreateTransaction(request *requests.CreateTransferRequest) (*response.TransferResponse, *response.ErrorResponse) {
	_, err := s.cardRepository.FindCardByCardNumber(request.TransferFrom)
	if err != nil {
		s.logger.Error("failed to find sender card by Number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Sender card not found",
		}
	}

	_, err = s.cardRepository.FindCardByCardNumber(request.TransferTo)
	if err != nil {
		s.logger.Error("failed to find receiver card by number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Receiver card not found",
		}
	}

	senderSaldo, err := s.saldoRepository.FindByCardNumber(request.TransferFrom)
	if err != nil {
		s.logger.Error("failed to find sender saldo by card number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find sender saldo",
		}
	}

	receiverSaldo, err := s.saldoRepository.FindByCardNumber(request.TransferTo)
	if err != nil {
		s.logger.Error("failed to find receiver saldo by card number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find receiver saldo",
		}
	}

	if senderSaldo.TotalBalance < request.TransferAmount {
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for sender",
		}
	}

	senderSaldo.TotalBalance -= request.TransferAmount
	receiverSaldo.TotalBalance += request.TransferAmount

	_, err = s.saldoRepository.UpdateSaldoBalance(&requests.UpdateSaldoBalance{
		CardNumber:   senderSaldo.CardNumber,
		TotalBalance: senderSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("failed to update sender saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update sender saldo",
		}
	}

	_, err = s.saldoRepository.UpdateSaldoBalance(&requests.UpdateSaldoBalance{
		CardNumber:   receiverSaldo.CardNumber,
		TotalBalance: receiverSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("failed to update receiver saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update receiver saldo",
		}
	}

	transfer, err := s.transferRepository.CreateTransfer(request)
	if err != nil {
		s.logger.Error("failed to create transfer", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transfer",
		}
	}

	so := s.mapping.ToTransferResponse(transfer)

	return so, nil
}

func (s *transferService) UpdateTransaction(request *requests.UpdateTransferRequest) (*response.TransferResponse, *response.ErrorResponse) {
	transfer, err := s.transferRepository.FindById(request.TransferID)
	if err != nil {
		s.logger.Error("Failed to find transfer by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Transfer with ID %d not found: %v", request.TransferID, err),
		}
	}

	amountDifference := request.TransferAmount - transfer.TransferAmount

	senderSaldo, err := s.saldoRepository.FindByCardNumber(transfer.TransferFrom)
	if err != nil {
		s.logger.Error("Failed to find sender's saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to find sender's saldo: %v", err),
		}
	}

	newSenderBalance := senderSaldo.TotalBalance - amountDifference
	if newSenderBalance < 0 {
		s.logger.Error("Insufficient balance for sender", zap.String("senderID", transfer.TransferFrom))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for sender",
		}
	}

	senderSaldo.TotalBalance = newSenderBalance
	_, err = s.saldoRepository.UpdateSaldoBalance(&requests.UpdateSaldoBalance{
		CardNumber:   senderSaldo.CardNumber,
		TotalBalance: senderSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("Failed to update sender's saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update sender's saldo: %v", err),
		}
	}

	// Update receiver's saldo
	receiverSaldo, err := s.saldoRepository.FindByCardNumber(transfer.TransferTo)
	if err != nil {
		s.logger.Error("Failed to find receiver's saldo by user ID", zap.Error(err))

		// Rollback the sender's saldo if the receiver's saldo update fails
		rollbackSenderBalance := &requests.UpdateSaldoBalance{
			CardNumber:   transfer.TransferFrom,
			TotalBalance: senderSaldo.TotalBalance,
		}
		_, rollbackErr := s.saldoRepository.UpdateSaldoBalance(rollbackSenderBalance)
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback sender's saldo after receiver lookup failure", zap.Error(rollbackErr))
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to find receiver's saldo: %v", err),
		}
	}

	newReceiverBalance := receiverSaldo.TotalBalance + amountDifference
	receiverSaldo.TotalBalance = newReceiverBalance

	_, err = s.saldoRepository.UpdateSaldoBalance(&requests.UpdateSaldoBalance{
		CardNumber:   receiverSaldo.CardNumber,
		TotalBalance: receiverSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("Failed to update receiver's saldo", zap.Error(err))

		rollbackSenderBalance := &requests.UpdateSaldoBalance{
			CardNumber:   transfer.TransferFrom,
			TotalBalance: senderSaldo.TotalBalance + amountDifference,
		}
		rollbackReceiverBalance := &requests.UpdateSaldoBalance{
			CardNumber:   transfer.TransferTo,
			TotalBalance: receiverSaldo.TotalBalance - amountDifference,
		}

		// Handle rollback sender balance
		if _, err := s.saldoRepository.UpdateSaldoBalance(rollbackSenderBalance); err != nil {
			s.logger.Error("Failed to rollback sender's saldo after receiver update failure", zap.Error(err))

		}

		if _, err := s.saldoRepository.UpdateSaldoBalance(rollbackReceiverBalance); err != nil {
			s.logger.Error("Failed to rollback receiver's saldo after sender update failure", zap.Error(err))

		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update receiver's saldo, rollback attempted but may be incomplete",
		}

	}

	updatedTransfer, err := s.transferRepository.UpdateTransfer(request)
	if err != nil {
		s.logger.Error("Failed to update transfer", zap.Error(err))

		rollbackSenderBalance := &requests.UpdateSaldoBalance{
			CardNumber:   transfer.TransferFrom,
			TotalBalance: senderSaldo.TotalBalance + amountDifference,
		}
		rollbackReceiverBalance := &requests.UpdateSaldoBalance{
			CardNumber:   transfer.TransferTo,
			TotalBalance: receiverSaldo.TotalBalance - amountDifference,
		}

		if _, err := s.saldoRepository.UpdateSaldoBalance(rollbackSenderBalance); err != nil {
			s.logger.Error("Failed to rollback sender's saldo after receiver update failure", zap.Error(err))
		}
		if _, err := s.saldoRepository.UpdateSaldoBalance(rollbackReceiverBalance); err != nil {
			s.logger.Error("Failed to rollback receiver's saldo after sender update failure", zap.Error(err))
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update transfer: %v", err),
		}
	}

	so := s.mapping.ToTransferResponse(updatedTransfer)

	return so, nil
}

func (s *transferService) TrashedTransfer(transfer_id int) (*response.TransferResponse, *response.ErrorResponse) {
	res, err := s.transferRepository.TrashedTransfer(transfer_id)
	if err != nil {
		s.logger.Error("Failed to trash transfer", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash transfer",
		}
	}

	so := s.mapping.ToTransferResponse(res)

	return so, nil
}

func (s *transferService) RestoreTransfer(transfer_id int) (*response.TransferResponse, *response.ErrorResponse) {
	res, err := s.transferRepository.RestoreTransfer(transfer_id)
	if err != nil {
		s.logger.Error("Failed to restore transfer", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore transfer",
		}
	}

	so := s.mapping.ToTransferResponse(res)

	return so, nil
}

func (s *transferService) DeleteTransferPermanent(transfer_id int) (interface{}, *response.ErrorResponse) {
	err := s.transferRepository.DeleteTransferPermanent(transfer_id)
	if err != nil {
		s.logger.Error("Failed to permanently delete transfer", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete transfer",
		}
	}

	return nil, nil
}
