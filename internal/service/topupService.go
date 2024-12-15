package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type topupService struct {
	cardRepository  repository.CardRepository
	topupRepository repository.TopupRepository
	saldoRepository repository.SaldoRepository
	logger          *logger.Logger
	mapping         responsemapper.TopupResponseMapper
}

func NewTopupService(cardRepository repository.CardRepository,
	topupRepository repository.TopupRepository,
	saldoRepository repository.SaldoRepository,
	logger *logger.Logger, mapping responsemapper.TopupResponseMapper) *topupService {
	return &topupService{
		topupRepository: topupRepository,
		saldoRepository: saldoRepository,
		cardRepository:  cardRepository,
		logger:          logger,
		mapping:         mapping,
	}
}

func (s *topupService) FindAll(page int, pageSize int, search string) ([]*response.TopupResponse, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	topups, totalRecords, err := s.topupRepository.FindAllTopups(search, page, pageSize)

	if err != nil {
		s.logger.Error("failed to fetch topups", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topups",
		}
	}

	so := s.mapping.ToTopupResponses(topups)

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return so, totalPages, nil
}

func (s *topupService) FindById(topupID int) (*response.TopupResponse, *response.ErrorResponse) {
	topup, err := s.topupRepository.FindById(topupID)
	if err != nil {
		s.logger.Error("failed to find topup by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Topup record not found",
		}
	}

	so := s.mapping.ToTopupResponse(*topup)

	return so, nil
}

func (s *topupService) FindByCardNumber(card_number string) ([]*response.TopupResponse, *response.ErrorResponse) {
	s.logger.Debug("Finding top-up by card number", zap.String("card_number", card_number))

	res, err := s.topupRepository.FindByCardNumber(card_number)
	if err != nil {
		s.logger.Error("Failed to find top-up by card number", zap.Error(err), zap.String("card_number", card_number))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find top-up by card number",
		}
	}

	so := s.mapping.ToTopupResponses(res)

	s.logger.Debug("Successfully found top-up by card number", zap.String("card_number", card_number))

	return so, nil
}

func (s *topupService) FindByActive() ([]*response.TopupResponse, *response.ErrorResponse) {
	s.logger.Info("Finding active top-up records")

	res, err := s.topupRepository.FindByActive()
	if err != nil {
		s.logger.Error("Failed to find active top-up records", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find active top-up records",
		}
	}

	so := s.mapping.ToTopupResponses(res)

	s.logger.Debug("Successfully found active top-up records", zap.Int("count", len(res)))

	return so, nil
}

func (s *topupService) FindByTrashed() ([]*response.TopupResponse, *response.ErrorResponse) {
	s.logger.Info("Finding trashed top-up records")

	res, err := s.topupRepository.FindByTrashed()
	if err != nil {
		s.logger.Error("Failed to find trashed top-up records", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find trashed top-up records",
		}
	}

	so := s.mapping.ToTopupResponses(res)

	s.logger.Debug("Successfully found trashed top-up records", zap.Int("count", len(res)))

	return so, nil
}

func (s *topupService) CreateTopup(request requests.CreateTopupRequest) (*response.TopupResponse, *response.ErrorResponse) {
	card, err := s.cardRepository.FindCardByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("failed to find card by number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}

	topup, err := s.topupRepository.CreateTopup(request)
	if err != nil {
		s.logger.Error("failed to create topup", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create topup record",
		}
	}

	saldo, err := s.saldoRepository.FindByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("failed to find saldo by user id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch user's saldo",
		}
	}

	newBalance := saldo.TotalBalance + request.TopupAmount
	_, err = s.saldoRepository.UpdateSaldoBalance(requests.UpdateSaldoBalance{
		CardNumber:   request.CardNumber,
		TotalBalance: newBalance,
	})
	if err != nil {
		s.logger.Error("failed to update saldo balance", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo balance",
		}
	}

	expireDate, err := time.Parse("2006-01-02", card.ExpireDate)
	if err != nil {
		s.logger.Error("failed to parse expire date", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Invalid expire date format",
		}
	}

	_, err = s.cardRepository.UpdateCard(requests.UpdateCardRequest{
		CardID:       card.ID,
		UserID:       card.UserID,
		CardType:     card.CardType,
		ExpireDate:   expireDate,
		CVV:          card.CVV,
		CardProvider: card.CardProvider,
	})
	if err != nil {
		s.logger.Error("failed to update card expire date", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update card expire date",
		}
	}

	so := s.mapping.ToTopupResponse(*topup)

	return so, nil
}

func (s *topupService) UpdateTopup(request requests.UpdateTopupRequest) (*response.TopupResponse, *response.ErrorResponse) {

	_, err := s.cardRepository.FindCardByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("failed to find card by number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}

	existingTopup, err := s.topupRepository.FindById(request.TopupID)

	if err != nil || existingTopup == nil {
		s.logger.Error("Failed to find topup by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Topup not found",
		}
	}

	topupDifference := request.TopupAmount - existingTopup.TopupAmount

	_, err = s.topupRepository.UpdateTopup(request)
	if err != nil {
		s.logger.Error("Failed to update topup amount", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update topup amount: %v", err),
		}
	}

	currentSaldo, err := s.saldoRepository.FindByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("Failed to retrieve current saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve current saldo: %v", err),
		}
	}

	if currentSaldo == nil {
		s.logger.Error("No saldo found for card number", zap.String("card_number", request.CardNumber))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "card not found",
		}
	}

	newBalance := currentSaldo.TotalBalance + topupDifference

	_, err = s.saldoRepository.UpdateSaldoBalance(requests.UpdateSaldoBalance{
		CardNumber:   request.CardNumber,
		TotalBalance: newBalance,
	})

	if err != nil {
		s.logger.Error("Failed to update saldo balance", zap.Error(err))

		_, rollbackErr := s.topupRepository.UpdateTopupAmount(requests.UpdateTopupAmount{
			TopupID:     request.TopupID,
			TopupAmount: existingTopup.TopupAmount,
		})
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback topup update", zap.Error(rollbackErr))
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update saldo balance: %v", err),
		}
	}

	updatedTopup, err := s.topupRepository.FindById(request.TopupID)

	if err != nil || updatedTopup == nil {
		s.logger.Error("Failed to find updated topup by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Updated topup not found",
		}
	}

	so := s.mapping.ToTopupResponse(*updatedTopup)

	return so, nil
}

func (s *topupService) TrashedTopup(topup_id int) (*response.TopupResponse, *response.ErrorResponse) {
	res, err := s.topupRepository.TrashedTopup(topup_id)
	if err != nil {
		s.logger.Error("Failed to trash topup", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to trash topup: %v", err),
		}
	}

	so := s.mapping.ToTopupResponse(*res)

	return so, nil
}

func (s *topupService) RestoreTopup(topup_id int) (*response.TopupResponse, *response.ErrorResponse) {
	res, err := s.topupRepository.RestoreTopup(topup_id)
	if err != nil {
		s.logger.Error("Failed to restore topup", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to restore topup: %v", err),
		}
	}

	so := s.mapping.ToTopupResponse(*res)

	return so, nil
}

func (s *topupService) DeleteTopupPermanent(topup_id int) (interface{}, *response.ErrorResponse) {
	err := s.topupRepository.DeleteTopupPermanent(topup_id)
	if err != nil {
		s.logger.Error("Failed to delete topup permanently", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to delete topup permanently: %v", err),
		}
	}

	return nil, nil
}
