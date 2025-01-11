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
	logger          logger.LoggerInterface
	mapping         responsemapper.TopupResponseMapper
}

func NewTopupService(cardRepository repository.CardRepository,
	topupRepository repository.TopupRepository,
	saldoRepository repository.SaldoRepository,
	logger logger.LoggerInterface, mapping responsemapper.TopupResponseMapper) *topupService {
	return &topupService{
		topupRepository: topupRepository,
		saldoRepository: saldoRepository,
		cardRepository:  cardRepository,
		logger:          logger,
		mapping:         mapping,
	}
}

func (s *topupService) FindAll(page int, pageSize int, search string) ([]*response.TopupResponse, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching topup",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	topups, totalRecords, err := s.topupRepository.FindAllTopups(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch topup",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topups",
		}
	}

	so := s.mapping.ToTopupResponses(topups)

	s.logger.Debug("Successfully fetched topup",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return so, totalRecords, nil
}

func (s *topupService) FindById(topupID int) (*response.TopupResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching topup by ID", zap.Int("topup_id", topupID))

	topup, err := s.topupRepository.FindById(topupID)

	if err != nil {
		s.logger.Error("failed to find topup by id", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Topup record not found",
		}
	}

	so := s.mapping.ToTopupResponse(topup)

	s.logger.Debug("Successfully fetched topup", zap.Int("topup_id", topupID))

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

func (s *topupService) FindByActive(page int, pageSize int, search string) ([]*response.TopupResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching active topup",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	topups, totalRecords, err := s.topupRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch active topup",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find active top-up records",
		}
	}

	so := s.mapping.ToTopupResponsesDeleteAt(topups)

	s.logger.Debug("Successfully fetched active topup",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return so, totalRecords, nil
}

func (s *topupService) FindByTrashed(page int, pageSize int, search string) ([]*response.TopupResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching trashed topup",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	topups, totalRecords, err := s.topupRepository.FindByTrashed(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch trashed topup",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find trashed top-up records",
		}
	}

	so := s.mapping.ToTopupResponsesDeleteAt(topups)

	s.logger.Debug("Successfully fetched trashed topup",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return so, totalRecords, nil
}

func (s *topupService) CreateTopup(request *requests.CreateTopupRequest) (*response.TopupResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting CreateTopup process",
		zap.String("cardNumber", request.CardNumber),
		zap.Float64("topupAmount", float64(request.TopupAmount)),
	)

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
	_, err = s.saldoRepository.UpdateSaldoBalance(&requests.UpdateSaldoBalance{
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

	_, err = s.cardRepository.UpdateCard(&requests.UpdateCardRequest{
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

	so := s.mapping.ToTopupResponse(topup)

	s.logger.Debug("CreateTopup process completed",
		zap.String("cardNumber", request.CardNumber),
		zap.Float64("topupAmount", float64(request.TopupAmount)),
		zap.Float64("newBalance", float64(newBalance)),
	)

	return so, nil
}

func (s *topupService) UpdateTopup(request *requests.UpdateTopupRequest) (*response.TopupResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting UpdateTopup process",
		zap.String("cardNumber", request.CardNumber),
		zap.Int("topupID", request.TopupID),
		zap.Float64("newTopupAmount", float64(request.TopupAmount)),
	)

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

	_, err = s.saldoRepository.UpdateSaldoBalance(&requests.UpdateSaldoBalance{
		CardNumber:   request.CardNumber,
		TotalBalance: newBalance,
	})

	if err != nil {
		s.logger.Error("Failed to update saldo balance", zap.Error(err))

		_, rollbackErr := s.topupRepository.UpdateTopupAmount(&requests.UpdateTopupAmount{
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

	so := s.mapping.ToTopupResponse(updatedTopup)

	s.logger.Debug("UpdateTopup process completed",
		zap.String("cardNumber", request.CardNumber),
		zap.Int("topupID", request.TopupID),
		zap.Float64("newTopupAmount", float64(request.TopupAmount)),
		zap.Float64("newBalance", float64(newBalance)),
	)

	return so, nil
}

func (s *topupService) TrashedTopup(topup_id int) (*response.TopupResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting TrashedTopup process",
		zap.Int("topupID", topup_id),
	)

	res, err := s.topupRepository.TrashedTopup(topup_id)

	if err != nil {
		s.logger.Error("Failed to trash topup", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to trash topup: %v", err),
		}
	}

	so := s.mapping.ToTopupResponse(res)

	s.logger.Debug("TrashedTopup process completed",
		zap.Int("topupID", topup_id),
	)

	return so, nil
}

func (s *topupService) RestoreTopup(topup_id int) (*response.TopupResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting RestoreTopup process",
		zap.Int("topupID", topup_id),
	)

	res, err := s.topupRepository.RestoreTopup(topup_id)

	if err != nil {
		s.logger.Error("Failed to restore topup", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to restore topup: %v", err),
		}
	}

	so := s.mapping.ToTopupResponse(res)

	s.logger.Debug("RestoreTopup process completed",
		zap.Int("topupID", topup_id),
	)

	return so, nil
}

func (s *topupService) DeleteTopupPermanent(topup_id int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Starting DeleteTopupPermanent process",
		zap.Int("topupID", topup_id),
	)

	_, err := s.topupRepository.DeleteTopupPermanent(topup_id)

	if err != nil {
		s.logger.Error("Failed to delete topup permanently", zap.Error(err))

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to delete topup permanently: %v", err),
		}
	}

	s.logger.Debug("DeleteTopupPermanent process completed",
		zap.Int("topupID", topup_id),
	)

	return true, nil
}

func (s *topupService) RestoreAllTopup() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all topups")

	_, err := s.topupRepository.RestoreAllTopup()

	if err != nil {
		s.logger.Error("Failed to restore all topups", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all topups: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully restored all topups")
	return true, nil
}

func (s *topupService) DeleteAllTopupPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all topups")

	_, err := s.topupRepository.DeleteAllTopupPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all topups", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all topups: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully deleted all topups permanently")
	return true, nil
}
