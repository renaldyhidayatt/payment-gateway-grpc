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

type cardService struct {
	cardRepository repository.CardRepository
	userRepository repository.UserRepository
	logger         logger.LoggerInterface
	mapping        responsemapper.CardResponseMapper
}

func NewCardService(
	cardRepository repository.CardRepository,
	userRepository repository.UserRepository,
	logger logger.LoggerInterface,
	mapper responsemapper.CardResponseMapper,

) *cardService {
	return &cardService{
		cardRepository: cardRepository,
		userRepository: userRepository,
		logger:         logger,
		mapping:        mapper,
	}
}

func (s *cardService) FindAll(page int, pageSize int, search string) ([]*response.CardResponse, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching card records",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	cards, totalRecords, err := s.cardRepository.FindAllCards(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch card records",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card records",
		}
	}

	responseData := s.mapping.ToCardsResponse(cards)

	s.logger.Debug("Successfully fetched card records",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return responseData, totalRecords, nil
}

func (s *cardService) FindById(card_id int) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching card by ID", zap.Int("card_id", card_id))

	res, err := s.cardRepository.FindById(card_id)

	if err != nil {
		s.logger.Error("Failed to fetch card by ID", zap.Error(err), zap.Int("card_id", card_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card record not found",
		}
	}

	so := s.mapping.ToCardResponse(res)

	s.logger.Debug("Successfully fetched card", zap.Int("card_id", card_id))

	return so, nil
}

func (s *cardService) FindByUserID(userID int) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching card by user ID", zap.Int("userID", userID))

	res, err := s.cardRepository.FindCardByUserId(userID)

	if err != nil {
		s.logger.Error("Failed to fetch cards by user ID", zap.Error(err), zap.Int("userID", userID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch cards by user ID",
		}
	}

	so := s.mapping.ToCardResponse(res)

	s.logger.Debug("Successfully fetched card records by user ID", zap.Int("userID", userID))

	return so, nil
}

func (s *cardService) DashboardCard() (*response.DashboardCard, *response.ErrorResponse) {
	totalBalance, err := s.cardRepository.GetTotalBalances()
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: "Failed to retrieve total balance",
			Status:  "error",
		}
	}

	totalTopup, err := s.cardRepository.GetTotalTopAmount()
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: "Failed to retrieve total top-up amount",
			Status:  "error",
		}
	}

	totalWithdraw, err := s.cardRepository.GetTotalWithdrawAmount()
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: "Failed to retrieve total withdrawal amount",
			Status:  "error",
		}
	}

	totalTransaction, err := s.cardRepository.GetTotalTransactionAmount()
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: "Failed to retrieve total transaction amount",
			Status:  "error",
		}
	}

	totalTransfer, err := s.cardRepository.GetTotalTransferAmount()
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: "Failed to retrieve total transfer amount",
			Status:  "error",
		}
	}

	return &response.DashboardCard{
		TotalBalance:     totalBalance,
		TotalTopup:       totalTopup,
		TotalWithdraw:    totalWithdraw,
		TotalTransaction: totalTransaction,
		TotalTransfer:    totalTransfer,
	}, nil
}

func (s *cardService) DashboardCardCardNumber(cardNumber string) (*response.DashboardCardCardNumber, *response.ErrorResponse) {
	totalBalance, err := s.cardRepository.GetTotalBalanceByCardNumber(cardNumber)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: fmt.Sprintf("Failed to retrieve total balance for card %s", cardNumber),
			Status:  "error",
		}
	}

	totalTopup, err := s.cardRepository.GetTotalTopupAmountByCardNumber(cardNumber)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: fmt.Sprintf("Failed to retrieve total top-up amount for card %s", cardNumber),
			Status:  "error",
		}
	}

	totalWithdraw, err := s.cardRepository.GetTotalWithdrawAmountByCardNumber(cardNumber)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: fmt.Sprintf("Failed to retrieve total withdrawal amount for card %s", cardNumber),
			Status:  "error",
		}
	}

	totalTransaction, err := s.cardRepository.GetTotalTransactionAmountByCardNumber(cardNumber)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: fmt.Sprintf("Failed to retrieve total transaction amount for card %s", cardNumber),
			Status:  "error",
		}
	}

	totalTransferSent, err := s.cardRepository.GetTotalTransferAmountBySender(cardNumber)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: fmt.Sprintf("Failed to retrieve total transfer amount sent by card %s", cardNumber),
			Status:  "error",
		}
	}

	totalTransferReceived, err := s.cardRepository.GetTotalTransferAmountByReceiver(cardNumber)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message: fmt.Sprintf("Failed to retrieve total transfer amount received by card %s", cardNumber),
			Status:  "error",
		}
	}

	return &response.DashboardCardCardNumber{
		TotalBalance:          totalBalance,
		TotalTopup:            totalTopup,
		TotalWithdraw:         totalWithdraw,
		TotalTransaction:      totalTransaction,
		TotalTransferSent:     totalTransferSent,
		TotalTransferReceiver: totalTransferReceived,
	}, nil
}

func (s *cardService) FindMonthlyBalance(year int) ([]*response.CardResponseMonthBalance, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyBalance called", zap.Int("year", year))

	res, err := s.cardRepository.GetMonthlyBalance(year)

	if err != nil {
		s.logger.Error("Failed to retrieve monthly balance",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly balance: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyBalances(res)

	s.logger.Debug("Monthly balance retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyBalance(year int) ([]*response.CardResponseYearlyBalance, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyBalance called", zap.Int("year", year))

	res, err := s.cardRepository.GetYearlyBalance(year)

	if err != nil {
		s.logger.Error("Failed to retrieve yearly balance",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly balance: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyBalances(res)

	s.logger.Debug("Yearly balance retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyTopupAmount(year int) ([]*response.CardResponseMonthTopupAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyTopupAmount called", zap.Int("year", year))

	res, err := s.cardRepository.GetMonthlyTopupAmount(year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly topup amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly topup amount: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyTopupAmounts(res)

	s.logger.Debug("Monthly topup amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyTopupAmount(year int) ([]*response.CardResponseYearlyTopupAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyTopupAmount called", zap.Int("year", year))

	res, err := s.cardRepository.GetYearlyTopupAmount(year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly topup amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly topup amount: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyTopupAmounts(res)

	s.logger.Debug("Yearly topup amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyWithdrawAmount(year int) ([]*response.CardResponseMonthWithdrawAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyWithdrawAmount called", zap.Int("year", year))

	res, err := s.cardRepository.GetMonthlyWithdrawAmount(year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly withdraw amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly withdraw amount: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyWithdrawAmounts(res)

	s.logger.Debug("Monthly withdraw amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyWithdrawAmount(year int) ([]*response.CardResponseYearlyWithdrawAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyWithdrawAmount called", zap.Int("year", year))

	res, err := s.cardRepository.GetYearlyWithdrawAmount(year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly withdraw amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly withdraw amount: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyWithdrawAmounts(res)

	s.logger.Debug("Yearly withdraw amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyTransactionAmount(year int) ([]*response.CardResponseMonthTransactionAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyTransactionAmount called", zap.Int("year", year))

	res, err := s.cardRepository.GetMonthlyTransactionAmount(year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly transaction amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transaction amount: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyTransactionAmounts(res)

	s.logger.Debug("Monthly transaction amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyTransactionAmount(year int) ([]*response.CardResponseYearlyTransactionAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyTransactionAmount called", zap.Int("year", year))

	res, err := s.cardRepository.GetYearlyTransactionAmount(year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly transaction amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transaction amount: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyTransactionAmounts(res)

	s.logger.Debug("Yearly transaction amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyTransferAmountSender(year int) ([]*response.CardResponseMonthTransferAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyTransferAmountSender called", zap.Int("year", year))

	res, err := s.cardRepository.GetMonthlyTransferAmountSender(year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly transfer sender amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transfer sender amount: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyTransferSenderAmounts(res)

	s.logger.Debug("Monthly transfer sender amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyTransferAmountSender(year int) ([]*response.CardResponseYearlyTransferAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyTransferAmountSender called", zap.Int("year", year))

	res, err := s.cardRepository.GetYearlyTransferAmountSender(year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly transfer sender amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transfer sender amount: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyTransferSenderAmounts(res)

	s.logger.Debug("Yearly transfer sender amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyTransferAmountReceiver(year int) ([]*response.CardResponseMonthTransferAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyTransferAmountReceiver called", zap.Int("year", year))

	res, err := s.cardRepository.GetMonthlyTransferAmountReceiver(year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly transfer receiver amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transfer receiver amount: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyTransferReceiverAmounts(res)

	s.logger.Debug("Monthly transfer receiver amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyTransferAmountReceiver(year int) ([]*response.CardResponseYearlyTransferAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyTransferAmountReceiver called", zap.Int("year", year))

	res, err := s.cardRepository.GetYearlyTransferAmountReceiver(year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly transfer receiver amount",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transfer receiver amount: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyTransferReceiverAmounts(res)

	s.logger.Debug("Yearly transfer receiver amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyBalanceByCardNumber(card_number string, year int) ([]*response.CardResponseMonthBalance, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyBalance called", zap.Int("year", year))

	res, err := s.cardRepository.GetMonthlyBalancesByCardNumber(card_number, year)

	if err != nil {
		s.logger.Error("Failed to retrieve monthly balance",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly balance: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyBalances(res)

	s.logger.Debug("Monthly balance retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyBalanceByCardNumber(card_number string, year int) ([]*response.CardResponseYearlyBalance, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyBalance called", zap.Int("year", year))

	res, err := s.cardRepository.GetYearlyBalanceByCardNumber(card_number, year)

	if err != nil {
		s.logger.Error("Failed to retrieve yearly balance",
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly balance: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyBalances(res)

	s.logger.Debug("Yearly balance retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyTopupAmountByCardNumber(cardNumber string, year int) ([]*response.CardResponseMonthTopupAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyTopupAmountByCardNumber called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetMonthlyTopupAmountByCardNumber(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly topup amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly topup amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyTopupAmounts(res)

	s.logger.Debug("Monthly topup amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyTopupAmountByCardNumber(cardNumber string, year int) ([]*response.CardResponseYearlyTopupAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyTopupAmountByCardNumber called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetYearlyTopupAmountByCardNumber(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly topup amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly topup amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyTopupAmounts(res)

	s.logger.Debug("Yearly topup amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyWithdrawAmountByCardNumber(cardNumber string, year int) ([]*response.CardResponseMonthWithdrawAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyWithdrawAmountByCardNumber called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetMonthlyWithdrawAmountByCardNumber(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly withdraw amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly withdraw amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyWithdrawAmounts(res)

	s.logger.Debug("Monthly withdraw amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyWithdrawAmountByCardNumber(cardNumber string, year int) ([]*response.CardResponseYearlyWithdrawAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyWithdrawAmountByCardNumber called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetYearlyWithdrawAmountByCardNumber(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly withdraw amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly withdraw amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyWithdrawAmounts(res)

	s.logger.Debug("Yearly withdraw amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyTransactionAmountByCardNumber(cardNumber string, year int) ([]*response.CardResponseMonthTransactionAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyTransactionAmountByCardNumber called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetMonthlyTransactionAmountByCardNumber(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly transaction amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transaction amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyTransactionAmounts(res)

	s.logger.Debug("Monthly transaction amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyTransactionAmountByCardNumber(cardNumber string, year int) ([]*response.CardResponseYearlyTransactionAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyTransactionAmountByCardNumber called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetYearlyTransactionAmountByCardNumber(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly transaction amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transaction amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyTransactionAmounts(res)

	s.logger.Debug("Yearly transaction amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyTransferAmountBySender(cardNumber string, year int) ([]*response.CardResponseMonthTransferAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyTransferAmountBySender called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetMonthlyTransferAmountBySender(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly transfer sender amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transfer sender amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyTransferSenderAmounts(res)

	s.logger.Debug("Monthly transfer sender amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyTransferAmountBySender(cardNumber string, year int) ([]*response.CardResponseYearlyTransferAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyTransferAmountBySender called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetYearlyTransferAmountBySender(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly transfer sender amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transfer sender amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyTransferSenderAmounts(res)

	s.logger.Debug("Yearly transfer sender amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindMonthlyTransferAmountByReceiver(cardNumber string, year int) ([]*response.CardResponseMonthTransferAmount, *response.ErrorResponse) {
	s.logger.Debug("FindMonthlyTransferAmountByReceiver called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetMonthlyTransferAmountByReceiver(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve monthly transfer receiver amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transfer receiver amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetMonthlyTransferReceiverAmounts(res)

	s.logger.Debug("Monthly transfer receiver amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindYearlyTransferAmountByReceiver(cardNumber string, year int) ([]*response.CardResponseYearlyTransferAmount, *response.ErrorResponse) {
	s.logger.Debug("FindYearlyTransferAmountByReceiver called",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
	)

	res, err := s.cardRepository.GetYearlyTransferAmountByReceiver(cardNumber, year)
	if err != nil {
		s.logger.Error("Failed to retrieve yearly transfer receiver amount by card number",
			zap.String("card_number", cardNumber),
			zap.Int("year", year),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transfer receiver amount by card number: %v", err),
		}
	}

	so := s.mapping.ToGetYearlyTransferReceiverAmounts(res)

	s.logger.Debug("Yearly transfer receiver amount by card number retrieved successfully",
		zap.String("card_number", cardNumber),
		zap.Int("year", year),
		zap.Int("result_count", len(so)),
	)

	return so, nil
}

func (s *cardService) FindByActive(page int, pageSize int, search string) ([]*response.CardResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching active card records",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.cardRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch active card records",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active card records",
		}
	}

	responseData := s.mapping.ToCardsResponseDeleteAt(res)

	s.logger.Debug("Successfully fetched active card records",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return responseData, totalRecords, nil
}

func (s *cardService) FindByTrashed(page int, pageSize int, search string) ([]*response.CardResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching trashed card records",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.cardRepository.FindByTrashed(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to fetch trashed card records",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed card records",
		}
	}

	responseData := s.mapping.ToCardsResponseDeleteAt(res)

	s.logger.Debug("Successfully fetched trashed card records",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return responseData, totalRecords, nil
}

func (s *cardService) FindByCardNumber(card_number string) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching card record by card number", zap.String("card_number", card_number))

	res, err := s.cardRepository.FindCardByCardNumber(card_number)

	if err != nil {
		s.logger.Error("Failed to fetch card by card number", zap.Error(err), zap.String("card_number", card_number))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card record not found for the given card number",
		}
	}

	so := s.mapping.ToCardResponse(res)

	s.logger.Debug("Successfully fetched card record by card number", zap.String("card_number", card_number))

	return so, nil
}

func (s *cardService) CreateCard(request *requests.CreateCardRequest) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new card", zap.Any("request", request))

	_, err := s.userRepository.FindById(request.UserID)

	if err != nil {
		s.logger.Error("Failed to find user by ID", zap.Error(err), zap.Int("userID", request.UserID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	res, err := s.cardRepository.CreateCard(request)

	if err != nil {
		s.logger.Error("Failed to create card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create card",
		}
	}

	so := s.mapping.ToCardResponse(res)

	s.logger.Debug("Successfully created new card", zap.Int("card_id", so.ID))

	return so, nil
}

func (s *cardService) UpdateCard(request *requests.UpdateCardRequest) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating card", zap.Int("card_id", request.CardID), zap.Any("request", request))

	_, err := s.userRepository.FindById(request.UserID)

	if err != nil {
		s.logger.Error("Failed to find user by ID", zap.Error(err), zap.Int("userID", request.UserID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	res, err := s.cardRepository.UpdateCard(request)
	if err != nil {
		s.logger.Error("Failed to update card", zap.Error(err), zap.Int("cardID", request.CardID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update card",
		}
	}

	so := s.mapping.ToCardResponse(res)

	s.logger.Debug("Successfully updated card", zap.Int("cardID", so.ID))

	return so, nil
}

func (s *cardService) TrashedCard(cardId int) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Trashing card", zap.Int("cardID", cardId))

	res, err := s.cardRepository.TrashedCard(cardId)
	if err != nil {
		s.logger.Error("Failed to trash card", zap.Error(err), zap.Int("cardID", cardId))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash card",
		}
	}

	so := s.mapping.ToCardResponse(res)

	s.logger.Debug("Successfully trashed card", zap.Int("cardID", so.ID))

	return so, nil
}

func (s *cardService) RestoreCard(cardId int) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Restoring card", zap.Int("cardID", cardId))

	res, err := s.cardRepository.RestoreCard(cardId)

	if err != nil {
		s.logger.Error("Failed to restore card", zap.Error(err), zap.Int("cardID", cardId))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore card",
		}
	}

	so := s.mapping.ToCardResponse(res)

	s.logger.Debug("Successfully restored card", zap.Int("cardID", so.ID))

	return so, nil
}

func (s *cardService) DeleteCardPermanent(cardId int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting card", zap.Int("cardID", cardId))

	_, err := s.cardRepository.DeleteCardPermanent(cardId)
	if err != nil {
		s.logger.Error("Failed to permanently delete card", zap.Error(err), zap.Int("cardID", cardId))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete card: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully deleted card permanently", zap.Int("cardID", cardId))

	return true, nil
}

func (s *cardService) RestoreAllCard() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all cards")

	_, err := s.cardRepository.RestoreAllCard()
	if err != nil {

		s.logger.Error("Failed to restore all cards", zap.Error(err))

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all cards: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully restored all cards")
	return true, nil
}

func (s *cardService) DeleteAllCardPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all cards")

	_, err := s.cardRepository.DeleteAllCardPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all cards", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all cards: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully deleted all cards permanently")

	return true, nil
}
