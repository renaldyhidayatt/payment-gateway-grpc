package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responseservice "MamangRust/paymentgatewaygrpc/internal/mapper/response/service"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type saldoService struct {
	cardRepository  repository.CardRepository
	saldoRepository repository.SaldoRepository
	logger          logger.LoggerInterface
	mapping         responseservice.SaldoResponseMapper
}

func NewSaldoService(saldo repository.SaldoRepository, card repository.CardRepository, logger logger.LoggerInterface, mapping responseservice.SaldoResponseMapper) *saldoService {
	return &saldoService{
		saldoRepository: saldo,
		cardRepository:  card,
		logger:          logger,
		mapping:         mapping,
	}
}

func (s *saldoService) FindAll(page int, pageSize int, search string) ([]*response.SaldoResponse, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching saldo",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	s.logger.Debug("Fetching all saldo records", zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))

	res, totalRecords, err := s.saldoRepository.FindAllSaldos(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch saldo",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Unable to fetch saldo records",
		}
	}

	so := s.mapping.ToSaldoResponses(res)

	s.logger.Error("Failed to fetch saldo",
		zap.Error(err),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

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

func (s *saldoService) FindMonthlyTotalSaldoBalance(year int, month int) ([]*response.SaldoMonthTotalBalanceResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly total saldo balance", zap.Int("year", year), zap.Int("month", month))

	res, err := s.saldoRepository.GetMonthlyTotalSaldoBalance(year, month)
	if err != nil {
		s.logger.Error("Failed to fetch monthly total saldo balance", zap.Error(err), zap.Int("year", year), zap.Int("month", month))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly total saldo balance",
		}
	}

	responses := s.mapping.ToSaldoMonthTotalBalanceResponses(res)

	for i, row := range responses {
		fmt.Printf("Row %d: %+v\n", i, *row)
	}

	s.logger.Debug("Successfully fetched monthly total saldo balance", zap.Int("year", year), zap.Int("month", month))

	return responses, nil
}

func (s *saldoService) FindYearTotalSaldoBalance(year int) ([]*response.SaldoYearTotalBalanceResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly total saldo balance", zap.Int("year", year))

	res, err := s.saldoRepository.GetYearTotalSaldoBalance(year)

	if err != nil {
		s.logger.Error("Failed to fetch yearly total saldo balance", zap.Error(err), zap.Int("year", year))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly total saldo balance",
		}
	}

	s.logger.Debug("Successfully fetched yearly total saldo balance", zap.Int("year", year))

	so := s.mapping.ToSaldoYearTotalBalanceResponses(res)

	return so, nil
}

func (s *saldoService) FindMonthlySaldoBalances(year int) ([]*response.SaldoMonthBalanceResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly saldo balances", zap.Int("year", year))

	res, err := s.saldoRepository.GetMonthlySaldoBalances(year)
	if err != nil {
		s.logger.Error("Failed to fetch monthly saldo balances", zap.Error(err), zap.Int("year", year))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch monthly saldo balances",
		}
	}

	responses := s.mapping.ToSaldoMonthBalanceResponses(res)

	s.logger.Debug("Successfully fetched monthly saldo balances", zap.Int("year", year))

	return responses, nil
}

func (s *saldoService) FindYearlySaldoBalances(year int) ([]*response.SaldoYearBalanceResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly saldo balances", zap.Int("year", year))

	res, err := s.saldoRepository.GetYearlySaldoBalances(year)

	if err != nil {
		s.logger.Error("Failed to fetch yearly saldo balances", zap.Error(err), zap.Int("year", year))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch yearly saldo balances",
		}
	}

	responses := s.mapping.ToSaldoYearBalanceResponses(res)

	s.logger.Debug("Successfully fetched yearly saldo balances", zap.Int("year", year))

	return responses, nil
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
	s.logger.Debug("Fetching saldo record",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.saldoRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch saldo",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No active saldo records found for the given ID",
		}
	}

	so := s.mapping.ToSaldoResponsesDeleteAt(res)

	s.logger.Debug("Successfully fetched saldo",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

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

func (s *saldoService) DeleteSaldoPermanent(saldo_id int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Deleting saldo permanently", zap.Int("saldo_id", saldo_id))

	_, err := s.saldoRepository.DeleteSaldoPermanent(saldo_id)

	if err != nil {
		s.logger.Error("Failed to delete saldo permanently", zap.Error(err), zap.Int("saldo_id", saldo_id))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete saldo permanently",
		}
	}

	s.logger.Debug("Successfully deleted saldo permanently", zap.Int("saldo_id", saldo_id))

	return true, nil
}

func (s *saldoService) RestoreAllSaldo() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all saldo")

	_, err := s.saldoRepository.RestoreAllSaldo()

	if err != nil {
		s.logger.Error("Failed to restore all saldo", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all saldo: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully restored all saldo")
	return true, nil
}

func (s *saldoService) DeleteAllSaldoPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all saldo")

	_, err := s.saldoRepository.DeleteAllSaldoPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all saldo", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all saldo: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully deleted all saldo permanently")
	return true, nil
}
