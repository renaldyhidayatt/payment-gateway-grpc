package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	apimapper "MamangRust/paymentgatewaygrpc/internal/mapper/response/api"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type merchantHandleApi struct {
	merchant pb.MerchantServiceClient
	logger   logger.LoggerInterface
	mapping  apimapper.MerchantResponseMapper
}

func NewHandlerMerchant(merchant pb.MerchantServiceClient, router *echo.Echo, logger logger.LoggerInterface, mapping apimapper.MerchantResponseMapper) *merchantHandleApi {
	merchantHandler := &merchantHandleApi{
		merchant: merchant,
		mapping:  mapping,
		logger:   logger,
	}

	routerMerchant := router.Group("/api/merchants")

	routerMerchant.GET("", merchantHandler.FindAll)
	routerMerchant.GET("/:id", merchantHandler.FindById)

	routerMerchant.GET("/transactions", merchantHandler.FindAllTransactions)
	routerMerchant.GET("/transactions/:merchant_id", merchantHandler.FindAllTransactionByMerchant)

	routerMerchant.GET("/monthly-payment-methods", merchantHandler.FindMonthlyPaymentMethodsMerchant)
	routerMerchant.GET("/yearly-payment-methods", merchantHandler.FindYearlyPaymentMethodMerchant)
	routerMerchant.GET("/monthly-amount", merchantHandler.FindMonthlyAmountMerchant)
	routerMerchant.GET("/yearly-amount", merchantHandler.FindYearlyAmountMerchant)

	routerMerchant.GET("/monthly-payment-methods-by-merchant", merchantHandler.FindMonthlyPaymentMethodByMerchants)
	routerMerchant.GET("/yearly-payment-methods-by-merchant", merchantHandler.FindYearlyPaymentMethodByMerchants)
	routerMerchant.GET("/monthly-amount-by-merchant", merchantHandler.FindMonthlyAmountByMerchants)
	routerMerchant.GET("/yearly-amount-by-merchant", merchantHandler.FindYearlyAmountByMerchants)

	routerMerchant.GET("/api-key", merchantHandler.FindByApiKey)
	routerMerchant.GET("/merchant-user", merchantHandler.FindByMerchantUserId)
	routerMerchant.GET("/active", merchantHandler.FindByActive)
	routerMerchant.GET("/trashed", merchantHandler.FindByTrashed)

	routerMerchant.POST("/create", merchantHandler.Create)
	routerMerchant.POST("/updates/:id", merchantHandler.Update)

	routerMerchant.POST("/trashed/:id", merchantHandler.TrashedMerchant)
	routerMerchant.POST("/restore/:id", merchantHandler.RestoreMerchant)
	routerMerchant.DELETE("/permanent/:id", merchantHandler.Delete)

	routerMerchant.POST("/restore/all", merchantHandler.RestoreAllMerchant)
	routerMerchant.POST("/permanent/all", merchantHandler.DeleteAllMerchantPermanent)

	return merchantHandler
}

// FindAll godoc
// @Summary Find all merchants
// @Tags Merchant
// @Security Bearer
// @Description Retrieve a list of all merchants
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} pb.ApiResponsePaginationMerchant "List of merchants"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant [get]
func (h *merchantHandleApi) FindAll(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllMerchantRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.merchant.FindAllMerchant(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve merchant data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	so := h.mapping.ToApiResponsesMerchant(res)

	return c.JSON(http.StatusOK, so)

}

// FindAllTransactions godoc
// @Summary Find all transactions
// @Tags Transaction
// @Security Bearer
// @Description Retrieve a list of all transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} pb.ApiResponsePaginationTransaction "List of transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction [get]
func (h *merchantHandleApi) FindAllTransactions(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllMerchantRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.merchant.FindAllTransactionMerchant(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	so := h.mapping.ToApiResponseMerchantsTransactionResponse(res)

	return c.JSON(http.StatusOK, so)
}

// FindAllTransactionByMerchant godoc
// @Summary Find all transactions by merchant ID
// @Tags Transaction
// @Security Bearer
// @Description Retrieve a list of transactions for a specific merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} pb.ApiResponsePaginationTransaction "List of transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/merchant/{merchant_id}/transaction [get]
func (h *merchantHandleApi) FindAllTransactionByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllMerchantTransaction{
		MerchantId: int32(merchantID),
		Page:       int32(page),
		PageSize:   int32(pageSize),
		Search:     search,
	}

	res, err := h.merchant.FindAllTransactionByMerchant(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	so := h.mapping.ToApiResponseMerchantsTransactionResponse(res)

	return c.JSON(http.StatusOK, so)
}

// FindById godoc
// @Summary Find a merchant by ID
// @Tags Merchant
// @Security Bearer
// @Description Retrieve a merchant by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} pb.ApiResponseMerchant "Merchant data"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/{id} [get]
func (h *merchantHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdMerchantRequest{
		MerchantId: int32(id),
	}

	res, err := h.merchant.FindByIdMerchant(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve merchant data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseMerchant(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyPaymentMethodsMerchant godoc
// @Summary Find monthly payment methods for a merchant
// @Tags Merchant
// @Security Bearer
// @Description Retrieve monthly payment methods for a merchant by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMerchantMonthlyPaymentMethod "Monthly payment methods"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly payment methods"
// @Router /api/merchant/monthly-payment-methods [get]
func (h *merchantHandleApi) FindMonthlyPaymentMethodsMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year",
		})
	}

	ctx := c.Request().Context()
	req := &pb.FindYearMerchant{
		Year: int32(year),
	}

	res, err := h.merchant.FindMonthlyPaymentMethodsMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to find monthly payment methods for merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find monthly payment methods for merchant: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseMonthlyPaymentMethods(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyPaymentMethodMerchant godoc.
// @Summary Find yearly payment methods for a merchant
// @Tags Merchant
// @Security Bearer
// @Description Retrieve yearly payment methods for a merchant by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMerchantYearlyPaymentMethod "Yearly payment methods"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly payment methods"
// @Router /api/merchant/monthly-amount [get]
func (h *merchantHandleApi) FindYearlyPaymentMethodMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year",
		})
	}

	ctx := c.Request().Context()
	req := &pb.FindYearMerchant{
		Year: int32(year),
	}

	res, err := h.merchant.FindYearlyPaymentMethodMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to find yearly payment methods for merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find yearly payment methods for merchant: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseYearlyPaymentMethods(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyAmountMerchant godoc
// @Summary Find monthly transaction amounts for a merchant
// @Tags Merchant
// @Security Bearer
// @Description Retrieve monthly transaction amounts for a merchant by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMerchantMonthlyAmount "Monthly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
// @Router /api/merchant/monthly-amount [get]
func (h *merchantHandleApi) FindMonthlyAmountMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year",
		})
	}

	ctx := c.Request().Context()
	req := &pb.FindYearMerchant{
		Year: int32(year),
	}

	res, err := h.merchant.FindMonthlyAmountMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to find monthly amount for merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find monthly amount for merchant: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyAmountMerchant godoc.
// @Summary Find yearly transaction amounts for a merchant
// @Tags Merchant
// @Security Bearer
// @Description Retrieve yearly transaction amounts for a merchant by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseYearlyAmount "Yearly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
// @Router /api/merchant/yearly-amount [get]
func (h *merchantHandleApi) FindYearlyAmountMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year",
		})
	}

	ctx := c.Request().Context()
	req := &pb.FindYearMerchant{
		Year: int32(year),
	}

	res, err := h.merchant.FindYearlyAmountMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to find yearly amount for merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find yearly amount for merchant: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyPaymentMethodByMerchants godoc.
// @Summary Find monthly payment methods for a specific merchant
// @Tags Merchant
// @Security Bearer
// @Description Retrieve monthly payment methods for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMerchantMonthlyPaymentMethod "Monthly payment methods"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly payment methods"
// @Router /api/merchant/monthly-payment-methods-by-merchant [get]
func (h *merchantHandleApi) FindMonthlyPaymentMethodByMerchants(c echo.Context) error {
	merchantIDStr := c.QueryParam("merchant_id")
	yearStr := c.QueryParam("year")

	merchantID, err := strconv.Atoi(merchantIDStr)
	if err != nil || merchantID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year",
		})
	}

	ctx := c.Request().Context()
	req := &pb.FindYearMerchantById{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.merchant.FindMonthlyPaymentMethodByMerchants(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to find monthly payment methods by merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find monthly payment methods by merchant: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseMonthlyPaymentMethods(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyPaymentMethodByMerchants godoc.
// @Summary Find yearly payment methods for a specific merchant
// @Tags Merchant
// @Security Bearer
// @Description Retrieve yearly payment methods for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMerchantYearlyPaymentMethod "Yearly payment methods"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly payment methods"
// @Router /api/merchant/payment-methods/yearly/by-merchant [get]
func (h *merchantHandleApi) FindYearlyPaymentMethodByMerchants(c echo.Context) error {
	merchantIDStr := c.QueryParam("merchant_id")
	yearStr := c.QueryParam("year")

	merchantID, err := strconv.Atoi(merchantIDStr)
	if err != nil || merchantID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year",
		})
	}

	ctx := c.Request().Context()
	req := &pb.FindYearMerchantById{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.merchant.FindYearlyPaymentMethodByMerchants(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to find yearly payment methods by merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find yearly payment methods by merchant: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseYearlyPaymentMethods(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyAmountByMerchants godoc.
// @Summary Find monthly transaction amounts for a specific merchant
// @Tags Merchant
// @Security Bearer
// @Description Retrieve monthly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMerchantMonthlyAmount "Monthly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
// @Router /api/merchant/amount/monthly/by-merchant [get]
func (h *merchantHandleApi) FindMonthlyAmountByMerchants(c echo.Context) error {
	merchantIDStr := c.QueryParam("merchant_id")
	yearStr := c.QueryParam("year")

	merchantID, err := strconv.Atoi(merchantIDStr)
	if err != nil || merchantID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year",
		})
	}

	ctx := c.Request().Context()
	req := &pb.FindYearMerchantById{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.merchant.FindMonthlyAmountByMerchants(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to find monthly amount by merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find monthly amount by merchant: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyAmountByMerchants godoc.
// @Summary Find yearly transaction amounts for a specific merchant
// @Tags Merchant
// @Security Bearer
// @Description Retrieve yearly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMerchantYearlyAmount "Yearly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
// @Router /api/merchant/yearly-payment-methods-by-merchant [get]
func (h *merchantHandleApi) FindYearlyAmountByMerchants(c echo.Context) error {
	merchantIDStr := c.QueryParam("merchant_id")
	yearStr := c.QueryParam("year")

	merchantID, err := strconv.Atoi(merchantIDStr)
	if err != nil || merchantID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year",
		})
	}

	ctx := c.Request().Context()
	req := &pb.FindYearMerchantById{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.merchant.FindYearlyAmountByMerchants(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to find yearly amount by merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find yearly amount by merchant: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindByApiKey godoc
// @Summary Find a merchant by API key
// @Tags Merchant
// @Security Bearer
// @Description Retrieve a merchant by its API key
// @Accept json
// @Produce json
// @Param api_key query string true "API key"
// @Success 200 {object} pb.ApiResponseMerchant "Merchant data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/api-key [get]
func (h *merchantHandleApi) FindByApiKey(c echo.Context) error {
	apiKey := c.QueryParam("api_key")

	ctx := c.Request().Context()

	req := &pb.FindByApiKeyRequest{
		ApiKey: apiKey,
	}

	res, err := h.merchant.FindByApiKey(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve merchant data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	so := h.mapping.ToApiResponseMerchant(res)

	return c.JSON(http.StatusOK, so)
}

// FindByMerchantUserId godoc.
// @Summary Find a merchant by user ID
// @Tags Merchant
// @Security Bearer
// @Description Retrieve a merchant by its user ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} pb.ApiResponsesMerchant "Merchant data"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/merchant-user [get]
func (h *merchantHandleApi) FindByMerchantUserId(c echo.Context) error {
	id, ok := c.Get("user_id").(int32)

	if !ok {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByMerchantUserIdRequest{
		UserId: id,
	}

	res, err := h.merchant.FindByMerchantUserId(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve merchant data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	so := h.mapping.ToApiResponseMerchants(res)

	return c.JSON(http.StatusOK, so)
}

// FindByActive godoc
// @Summary Find active merchants
// @Tags Merchant
// @Security Bearer
// @Description Retrieve a list of active merchants
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsesMerchant "List of active merchants"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/active [get]
func (h *merchantHandleApi) FindByActive(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllMerchantRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.merchant.FindByActive(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve merchant data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	so := h.mapping.ToApiResponsesMerchantDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// FindByTrashed godoc
// @Summary Find trashed merchants
// @Tags Merchant
// @Security Bearer
// @Description Retrieve a list of trashed merchants
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsesMerchant "List of trashed merchants"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/trashed [get]
func (h *merchantHandleApi) FindByTrashed(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllMerchantRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.merchant.FindByTrashed(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve merchant data", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	so := h.mapping.ToApiResponsesMerchantDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// Create godoc
// @Summary Create a new merchant
// @Tags Merchant
// @Security Bearer
// @Description Create a new merchant with the given name and user ID
// @Accept json
// @Produce json
// @Param body body requests.CreateMerchantRequest true "Create merchant request"
// @Success 200 {object} pb.ApiResponseMerchant "Created merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create merchant"
// @Router /api/merchant/create [post]
func (h *merchantHandleApi) Create(c echo.Context) error {
	var body requests.CreateMerchantRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	req := &pb.CreateMerchantRequest{
		Name:   body.Name,
		UserId: int32(body.UserID),
	}

	res, err := h.merchant.CreateMerchant(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to create merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create merchant:",
		})
	}

	so := h.mapping.ToApiResponseMerchant(res)

	return c.JSON(http.StatusOK, so)
}

// Update godoc
// @Summary Update a merchant
// @Tags Merchant
// @Security Bearer
// @Description Update a merchant with the given ID
// @Accept json
// @Produce json
// @Param body body requests.UpdateMerchantRequest true "Update merchant request"
// @Success 200 {object} pb.ApiResponseMerchant "Updated merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update merchant"
// @Router /api/merchant/update/{id} [post]
func (h *merchantHandleApi) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	var body requests.UpdateMerchantRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: ",
		})
	}

	body.MerchantID = id

	ctx := c.Request().Context()
	req := &pb.UpdateMerchantRequest{
		MerchantId: int32(body.MerchantID),
		Name:       body.Name,
		UserId:     int32(body.UserID),
		Status:     body.Status,
	}

	res, err := h.merchant.UpdateMerchant(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to update merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update merchant: ",
		})
	}

	so := h.mapping.ToApiResponseMerchant(res)

	return c.JSON(http.StatusOK, so)
}

// TrashedMerchant godoc
// @Summary Trashed a merchant
// @Tags Merchant
// @Security Bearer
// @Description Trashed a merchant by its ID
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} pb.ApiResponseMerchant "Trashed merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed merchant"
// @Router /api/merchant/trashed/{id} [post]
func (h *merchantHandleApi) TrashedMerchant(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.merchant.TrashedMerchant(ctx, &pb.FindByIdMerchantRequest{
		MerchantId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to trashed merchant", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed merchant:",
		})
	}

	so := h.mapping.ToApiResponseMerchant(res)

	return c.JSON(http.StatusOK, so)
}

// RestoreMerchant godoc
// @Summary Restore a merchant
// @Tags Merchant
// @Security Bearer
// @Description Restore a merchant by its ID
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} pb.ApiResponseMerchant "Restored merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore merchant"
// @Router /api/merchant/restore/{id} [post]
func (h *merchantHandleApi) RestoreMerchant(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.merchant.RestoreMerchant(ctx, &pb.FindByIdMerchantRequest{
		MerchantId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to restore merchant", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore merchant:",
		})
	}

	so := h.mapping.ToApiResponseMerchant(res)

	return c.JSON(http.StatusOK, so)
}

// Delete godoc
// @Summary Delete a merchant permanently
// @Tags Merchant
// @Security Bearer
// @Description Delete a merchant by its ID permanently
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} pb.ApiResponseMerchantDelete "Deleted merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete merchant"
// @Router /api/merchant/{id} [delete]
func (h *merchantHandleApi) Delete(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.merchant.DeleteMerchantPermanent(ctx, &pb.FindByIdMerchantRequest{
		MerchantId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to delete merchant", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete merchant:",
		})
	}

	so := h.mapping.ToApiResponseMerchantDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// RestoreAllMerchant godoc.
// @Summary Restore all merchant records
// @Tags Merchant
// @Security Bearer
// @Description Restore all merchant records that were previously deleted.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseMerchantAll "Successfully restored all merchant records"
// @Failure 500 {object} response.ErrorResponse "Failed to restore all merchant records"
// @Router /api/merchant/restore/all [post]
func (h *merchantHandleApi) RestoreAllMerchant(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.merchant.RestoreAllMerchant(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to restore all merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently restore all merchant",
		})
	}

	h.logger.Debug("Successfully restored all merchant")

	so := h.mapping.ToApiResponseMerchantAll(res)

	return c.JSON(http.StatusOK, so)
}

// DeleteAllMerchantPermanent godoc.
// @Summary Permanently delete all merchant records
// @Tags Merchant
// @Security Bearer
// @Description Permanently delete all merchant records from the database.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseMerchantAll "Successfully deleted all merchant records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to permanently delete all merchant records"
// @Router /api/merchant/permanent/all [delete]
func (h *merchantHandleApi) DeleteAllMerchantPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.merchant.DeleteAllMerchantPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to permanently delete all merchant", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all merchant",
		})
	}

	h.logger.Debug("Successfully deleted all merchant permanently")

	so := h.mapping.ToApiResponseMerchantAll(res)

	return c.JSON(http.StatusOK, so)
}
