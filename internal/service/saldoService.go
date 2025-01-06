package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type saldoService struct {
	cardRepository  repository.CardRepository
	saldoRepository repository.SaldoRepository
	logger          logger.LoggerInterface
	mapping         responsemapper.SaldoResponseMapper
}

func NewSaldoService(saldo repository.SaldoRepository, card repository.CardRepository, logger logger.LoggerInterface, mapping responsemapper.SaldoResponseMapper) *saldoService {
	return &saldoService{
		saldoRepository: saldo,
		cardRepository:  card,
		logger:          logger,
		mapping:         mapping,
	}
}

func (s *saldoService) FindAll(page int, pageSize int, search string) ([]*response.SaldoResponse, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	s.logger.Debug("Fetching all saldo records", zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))

	res, totalRecords, err := s.saldoRepository.FindAllSaldos(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to fetch saldo records", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Unable to fetch saldo records",
		}
	}

	so := s.mapping.ToSaldoResponses(res)

	s.logger.Debug("Successfully fetched saldo records", zap.Int("totalRecords", totalRecords), zap.Int("totalPages", totalRecords))

	return so, totalRecords, nil
}

func (s *saldoService) FindById(saldo_id int) (*response.SaldoResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching saldo record by ID", zap.Int("saldo_id", saldo_id))

	res, err := s.saldoRepository.FindById(saldo_id)
	if err != nil {
		s.logger.Error("Failed to fetch saldo by ID", zap.Error(err), zap.Int("saldo_id", saldo_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found for the given ID",
		}
	}

	so := s.mapping.ToSaldoResponse(res)
	s.logger.Debug("Successfully fetched saldo by ID", zap.Int("saldo_id", saldo_id))

	return so, nil
}

func (s *saldoService) FindByCardNumber(card_number string) (*response.SaldoResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching saldo record by card number", zap.String("card_number", card_number))

	res, err := s.saldoRepository.FindByCardNumber(card_number)
	if err != nil {
		s.logger.Error("Failed to fetch saldo by card number", zap.Error(err), zap.String("card_number", card_number))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found for the given card number",
		}
	}

	so := s.mapping.ToSaldoResponse(res)

	s.logger.Debug("Successfully fetched saldo by card number", zap.String("card_number", card_number))

	return so, nil
}

func (s *saldoService) FindByActive(page int, pageSize int, search string) ([]*response.SaldoResponseDeleteAt, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.saldoRepository.FindByActive(search, page, pageSize)

	if err != nil {
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No active saldo records found for the given ID",
		}
	}

	so := s.mapping.ToSaldoResponsesDeleteAt(res)

	return so, totalRecords, nil
}

func (s *saldoService) FindByTrashed(page int, pageSize int, search string) ([]*response.SaldoResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Info("Fetching trashed saldo records")

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.saldoRepository.FindByTrashed(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to fetch trashed saldo records", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No trashed saldo records found",
		}
	}

	so := s.mapping.ToSaldoResponsesDeleteAt(res)

	s.logger.Debug("Successfully fetched trashed saldo records", zap.Int("record_count", len(res)))

	return so, totalRecords, nil
}

func (s *saldoService) CreateSaldo(request *requests.CreateSaldoRequest) (*response.SaldoResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating saldo record", zap.String("card_number", request.CardNumber))

	_, err := s.cardRepository.FindCardByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("Card not found for creating saldo", zap.Error(err), zap.String("card_number", request.CardNumber))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card number not found",
		}
	}

	res, err := s.saldoRepository.CreateSaldo(request)
	if err != nil {
		s.logger.Error("Failed to create saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create saldo record",
		}
	}

	so := s.mapping.ToSaldoResponse(res)

	s.logger.Debug("Successfully created saldo record", zap.String("card_number", request.CardNumber))

	return so, nil
}

func (s *saldoService) UpdateSaldo(request *requests.UpdateSaldoRequest) (*response.SaldoResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating saldo record", zap.String("card_number", request.CardNumber), zap.Float64("amount", float64(request.TotalBalance)))

	_, err := s.cardRepository.FindCardByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("Failed to find card by card number", zap.Error(err), zap.String("card_number", request.CardNumber))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card number not found",
		}
	}

	res, err := s.saldoRepository.UpdateSaldo(request)

	if err != nil {
		s.logger.Error("Failed to update saldo", zap.Error(err), zap.String("card_number", request.CardNumber))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo",
		}
	}

	so := s.mapping.ToSaldoResponse(res)

	s.logger.Debug("Successfully updated saldo", zap.String("card_number", request.CardNumber), zap.Int("saldo_id", res.ID))

	return so, nil
}

func (s *saldoService) TrashSaldo(saldo_id int) (*response.SaldoResponse, *response.ErrorResponse) {
	s.logger.Debug("Trashing saldo record", zap.Int("saldo_id", saldo_id))

	res, err := s.saldoRepository.TrashedSaldo(saldo_id)
	if err != nil {
		s.logger.Error("Failed to trash saldo", zap.Error(err), zap.Int("saldo_id", saldo_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to move saldo to trash",
		}
	}

	so := s.mapping.ToSaldoResponse(res)
	s.logger.Debug("Successfully trashed saldo", zap.Int("saldo_id", saldo_id))

	return so, nil
}

func (s *saldoService) RestoreSaldo(saldo_id int) (*response.SaldoResponse, *response.ErrorResponse) {
	s.logger.Debug("Restoring saldo record from trash", zap.Int("saldo_id", saldo_id))

	res, err := s.saldoRepository.RestoreSaldo(saldo_id)
	if err != nil {
		s.logger.Error("Failed to restore saldo", zap.Error(err), zap.Int("saldo_id", saldo_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore saldo from trash",
		}
	}

	so := s.mapping.ToSaldoResponse(res)
	s.logger.Debug("Successfully restored saldo", zap.Int("saldo_id", saldo_id))

	return so, nil
}

func (s *saldoService) DeleteSaldoPermanent(saldo_id int) (interface{}, *response.ErrorResponse) {
	s.logger.Debug("Deleting saldo permanently", zap.Int("saldo_id", saldo_id))

	err := s.saldoRepository.DeleteSaldoPermanent(saldo_id)
	if err != nil {
		s.logger.Error("Failed to delete saldo permanently", zap.Error(err), zap.Int("saldo_id", saldo_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete saldo permanently",
		}
	}

	s.logger.Debug("Successfully deleted saldo permanently", zap.Int("saldo_id", saldo_id))

	return nil, nil
}
