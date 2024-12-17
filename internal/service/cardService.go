package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type cardService struct {
	cardRepository repository.CardRepository
	userRepository repository.UserRepository
	logger         *logger.Logger
	mapping        responsemapper.CardResponseMapper
}

func NewCardService(
	cardRepository repository.CardRepository,
	userRepository repository.UserRepository,
	logger *logger.Logger,
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
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	s.logger.Debug("Fetching all card records", zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))
	cards, totalRecords, err := s.cardRepository.FindAllCards(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to fetch all card records", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card records",
		}
	}

	if len(cards) == 0 {
		s.logger.Debug("No card records found", zap.String("search", search))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No card records found",
		}
	}

	so := s.mapping.ToCardsResponse(cards)
	totalPages := (totalRecords + pageSize - 1) / pageSize

	s.logger.Debug("Successfully fetched card records", zap.Int("totalRecords", totalRecords), zap.Int("totalPages", totalPages))

	return so, totalPages, nil
}

func (s *cardService) FindById(card_id int) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching card record by ID", zap.Int("card_id", card_id))
	res, err := s.cardRepository.FindById(card_id)
	if err != nil {
		s.logger.Error("Failed to fetch card by ID", zap.Error(err), zap.Int("card_id", card_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card record not found",
		}
	}

	so := s.mapping.ToCardResponse(*res)
	s.logger.Debug("Successfully fetched card record", zap.Int("card_id", card_id))
	return so, nil
}

func (s *cardService) FindByUserID(userID int) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching card records by user ID", zap.Int("userID", userID))
	res, err := s.cardRepository.FindCardByUserId(userID)
	if err != nil {
		s.logger.Error("Failed to fetch cards by user ID", zap.Error(err), zap.Int("userID", userID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch cards by user ID",
		}
	}

	so := s.mapping.ToCardResponse(*res)

	s.logger.Debug("Successfully fetched card records by user ID", zap.Int("userID", userID))

	return so, nil
}

func (s *cardService) FindByActive() ([]*response.CardResponse, *response.ErrorResponse) {
	res, err := s.cardRepository.FindByActive()
	if err != nil {
		s.logger.Error("Failed to fetch active cards", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active card records",
		}
	}

	if len(res) == 0 {
		s.logger.Debug("No active cards found")
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No active cards found",
		}
	}

	so := s.mapping.ToCardsResponse(res)

	s.logger.Debug("Successfully fetched active card records")

	return so, nil
}

func (s *cardService) FindByTrashed() ([]*response.CardResponse, *response.ErrorResponse) {
	s.logger.Info("Fetching trashed card records")
	res, err := s.cardRepository.FindByTrashed()
	if err != nil {
		s.logger.Error("Failed to fetch trashed cards", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed card records",
		}
	}

	if len(res) == 0 {
		s.logger.Debug("No trashed cards found")
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No trashed card records found",
		}
	}

	so := s.mapping.ToCardsResponse(res)

	s.logger.Info("Successfully fetched trashed card records")

	return so, nil
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

	so := s.mapping.ToCardResponse(*res)

	s.logger.Debug("Successfully fetched card record by card number", zap.String("card_number", card_number))

	return so, nil
}

func (s *cardService) CreateCard(request requests.CreateCardRequest) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new card", zap.Int("userID", request.UserID))

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

	so := s.mapping.ToCardResponse(*res)

	s.logger.Debug("Successfully created new card", zap.Int("cardID", so.ID))

	return so, nil
}

func (s *cardService) UpdateCard(request requests.UpdateCardRequest) (*response.CardResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating card", zap.Int("userID", request.UserID), zap.Int("cardID", request.CardID))

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

	so := s.mapping.ToCardResponse(*res)

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

	so := s.mapping.ToCardResponse(*res)

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

	so := s.mapping.ToCardResponse(*res)
	s.logger.Debug("Successfully restored card", zap.Int("cardID", so.ID))
	return so, nil
}

func (s *cardService) DeleteCardPermanent(cardId int) (interface{}, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting card", zap.Int("cardID", cardId))

	err := s.cardRepository.DeleteCardPermanent(cardId)
	if err != nil {
		s.logger.Error("Failed to permanently delete card", zap.Error(err), zap.Int("cardID", cardId))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete card: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully deleted card permanently", zap.Int("cardID", cardId))

	return nil, nil
}
