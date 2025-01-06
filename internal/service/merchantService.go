package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type merchantService struct {
	merchantRepository repository.MerchantRepository
	logger             logger.LoggerInterface
	mapping            responsemapper.MerchantResponseMapper
}

func NewMerchantService(
	merchantRepository repository.MerchantRepository,
	logger logger.LoggerInterface,
	mapping responsemapper.MerchantResponseMapper,
) *merchantService {
	return &merchantService{
		merchantRepository: merchantRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *merchantService) FindAll(page int, pageSize int, search string) ([]*response.MerchantResponse, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.merchantRepository.FindAllMerchants(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to fetch merchant records", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch merchant records",
		}
	}

	merchantResponses := s.mapping.ToMerchantsResponse(merchants)

	return merchantResponses, totalRecords, nil
}

func (s *merchantService) FindById(merchant_id int) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Finding merchant by ID", zap.Int("merchant_id", merchant_id))

	res, err := s.merchantRepository.FindById(merchant_id)
	if err != nil {
		s.logger.Error("Failed to find merchant by ID", zap.Error(err), zap.Int("merchant_id", merchant_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found",
		}
	}

	so := s.mapping.ToMerchantResponse(res)

	s.logger.Debug("Successfully found merchant by ID", zap.Int("merchant_id", merchant_id))

	return so, nil
}

func (s *merchantService) FindByActive(page int, pageSize int, search string) ([]*response.MerchantResponseDeleteAt, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	s.logger.Info("Fetching active merchants")

	merchants, totalRecords, err := s.merchantRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch active merchants", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active merchants",
		}
	}

	so := s.mapping.ToMerchantsResponseDeleteAt(merchants)

	s.logger.Info("Successfully fetched active merchants")

	return so, totalRecords, nil
}

func (s *merchantService) FindByTrashed(page int, pageSize int, search string) ([]*response.MerchantResponseDeleteAt, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	s.logger.Info("Fetching trashed merchants")

	merchants, totalRecords, err := s.merchantRepository.FindByTrashed(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to fetch trashed merchants", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed merchants",
		}
	}

	so := s.mapping.ToMerchantsResponseDeleteAt(merchants)

	s.logger.Info("Successfully fetched trashed merchants")

	return so, totalRecords, nil
}

func (s *merchantService) FindByApiKey(api_key string) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Finding merchant by API key", zap.String("api_key", api_key))

	res, err := s.merchantRepository.FindByApiKey(api_key)
	if err != nil {
		s.logger.Error("Failed to find merchant by API key", zap.Error(err), zap.String("api_key", api_key))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found by API key",
		}
	}

	so := s.mapping.ToMerchantResponse(res)

	s.logger.Debug("Successfully found merchant by API key", zap.String("api_key", api_key))

	return so, nil
}

func (s *merchantService) FindByMerchantUserId(user_id int) ([]*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Finding merchant by user ID", zap.Int("user_id", user_id))

	res, err := s.merchantRepository.FindByMerchantUserId(user_id)
	if err != nil {
		s.logger.Error("Failed to find merchant by user ID", zap.Error(err), zap.Int("user_id", user_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found by user ID",
		}
	}

	so := s.mapping.ToMerchantsResponse(res)

	s.logger.Debug("Successfully found merchant by user ID", zap.Int("user_id", user_id))

	return so, nil
}

func (s *merchantService) CreateMerchant(request *requests.CreateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new merchant", zap.String("merchant_name", request.Name))

	res, err := s.merchantRepository.CreateMerchant(request)
	if err != nil {
		s.logger.Error("Failed to create merchant", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create merchant",
		}
	}

	so := s.mapping.ToMerchantResponse(res)

	s.logger.Debug("Successfully created merchant", zap.Int("merchant_id", res.ID))

	return so, nil
}

func (s *merchantService) UpdateMerchant(request *requests.UpdateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating merchant", zap.Int("merchant_id", request.MerchantID))

	_, err := s.merchantRepository.FindById(request.MerchantID)

	if err != nil {
		s.logger.Error("Merchant not found for update", zap.Error(err), zap.Int("merchant_id", request.MerchantID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found",
		}
	}

	res, err := s.merchantRepository.UpdateMerchant(request)
	if err != nil {
		s.logger.Error("Failed to update merchant", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update merchant",
		}
	}

	so := s.mapping.ToMerchantResponse(res)

	s.logger.Debug("Successfully updated merchant", zap.Int("merchant_id", res.ID))

	return so, nil
}

func (s *merchantService) TrashedMerchant(merchant_id int) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Trashing merchant", zap.Int("merchant_id", merchant_id))

	res, err := s.merchantRepository.TrashedMerchant(merchant_id)

	if err != nil {
		s.logger.Error("Failed to trash merchant", zap.Error(err), zap.Int("merchant_id", merchant_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash merchant",
		}
	}

	so := s.mapping.ToMerchantResponse(res)

	s.logger.Debug("Successfully trashed merchant", zap.Int("merchant_id", merchant_id))

	return so, nil
}

func (s *merchantService) RestoreMerchant(merchant_id int) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Restoring merchant", zap.Int("merchant_id", merchant_id))

	res, err := s.merchantRepository.RestoreMerchant(merchant_id)
	if err != nil {
		s.logger.Error("Failed to restore merchant", zap.Error(err), zap.Int("merchant_id", merchant_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore merchant",
		}
	}

	so := s.mapping.ToMerchantResponse(res)

	s.logger.Debug("Successfully restored merchant", zap.Int("merchant_id", merchant_id))

	return so, nil
}

func (s *merchantService) DeleteMerchantPermanent(merchant_id int) (interface{}, *response.ErrorResponse) {
	s.logger.Debug("Deleting merchant permanently", zap.Int("merchant_id", merchant_id))

	err := s.merchantRepository.DeleteMerchantPermanent(merchant_id)
	if err != nil {
		s.logger.Error("Failed to delete merchant permanently", zap.Error(err), zap.Int("merchant_id", merchant_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete merchant permanently",
		}
	}

	s.logger.Debug("Successfully deleted merchant permanently", zap.Int("merchant_id", merchant_id))

	return nil, nil
}
