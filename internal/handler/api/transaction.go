package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/middlewares"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type transactionHandler struct {
	transaction pb.TransactionServiceClient
	logger      logger.LoggerInterface
}

func NewHandlerTransaction(transaction pb.TransactionServiceClient, merchant pb.MerchantServiceClient, router *echo.Echo, logger logger.LoggerInterface) *transactionHandler {
	transactionHandler := transactionHandler{
		transaction: transaction,
		logger:      logger,
	}

	routerTransaction := router.Group("/api/transaction")

	routerTransaction.GET("", transactionHandler.FindAll)
	routerTransaction.GET("/:id", transactionHandler.FindById)
	routerTransaction.GET("/card/:card_number", transactionHandler.FindByCardNumber)
	routerTransaction.GET("/merchant/:merchant_id", transactionHandler.FindByTransactionMerchantId)
	routerTransaction.GET("/active", transactionHandler.FindByActiveTransaction)
	routerTransaction.GET("/trashed", transactionHandler.FindByTrashedTransaction)
	routerTransaction.POST("/create", middlewares.ApiKeyMiddleware(merchant)(transactionHandler.Create))
	routerTransaction.POST("/update", middlewares.ApiKeyMiddleware(merchant)(transactionHandler.Update))

	routerTransaction.POST("/restore/:id", transactionHandler.RestoreTransaction)
	routerTransaction.POST("/trashed/:id", transactionHandler.TrashedTransaction)

	routerTransaction.DELETE("/permanent/:id", transactionHandler.DeletePermanent)

	return &transactionHandler
}

// @Summary Find all
// @Tags Transaction
// @Description Retrieve a list of all transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} pb.ApiResponsePaginationTransaction "List of transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction/active [get]
func (h *transactionHandler) FindAll(c echo.Context) error {
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

	req := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.transaction.FindAllTransaction(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindById retrieves a transaction record by its ID.
// @Summary Find a transaction by ID
// @Tags Transaction
// @Description Retrieve a transaction record using its ID
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} pb.ApiResponseTransaction "Transaction data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction/{id} [get]
func (h *transactionHandler) FindById(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Invalid transaction ID", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.transaction.FindByIdTransaction(ctx, &pb.FindByIdTransactionRequest{
		TransactionId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindByCardNumber retrieves a transaction record by its card number.
// @Summary Find a transaction by card number
// @Tags Transaction
// @Description Retrieve a transaction record using its card number
// @Accept json
// @Produce json
// @Param card_number query string true "Card number"
// @Success 200 {object} pb.ApiResponseTransactions "Transaction data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction/{card_number} [get]
func (h *transactionHandler) FindByCardNumber(c echo.Context) error {
	cardNumber := c.QueryParam("card_number")

	ctx := c.Request().Context()

	req := &pb.FindByCardNumberTransactionRequest{
		CardNumber: cardNumber,
	}

	res, err := h.transaction.FindByCardNumberTransaction(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindByTransactionMerchantId retrieves transactions associated with a specific merchant ID.
// @Summary Find transactions by merchant ID
// @Tags Transaction
// @Description Retrieve a list of transactions using the merchant ID
// @Accept json
// @Produce json
// @Param merchant_id query string true "Merchant ID"
// @Success 200 {object} pb.ApiResponseTransactions "Transaction data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction/merchant/{merchant_id} [get]
func (h *transactionHandler) FindByTransactionMerchantId(c echo.Context) error {
	merchantId := c.QueryParam("merchant_id")

	merchantIdInt, err := strconv.Atoi(merchantId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindTransactionByMerchantIdRequest{
		MerchantId: int32(merchantIdInt),
	}

	res, err := h.transaction.FindTransactionByMerchantId(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindByActiveTransaction retrieves a list of active transactions.
// @Summary Find active transactions
// @Tags Transaction
// @Description Retrieve a list of active transactions
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseTransactions "List of active transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction/active [get]
func (h *transactionHandler) FindByActiveTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.transaction.FindByActiveTransaction(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindByTrashedTransaction retrieves a list of trashed transactions.
// @Summary Retrieve trashed transactions
// @Tags Transaction
// @Description Retrieve a list of trashed transactions
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseTransactions "List of trashed transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction/trashed [get]
func (h *transactionHandler) FindByTrashedTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.transaction.FindByTrashedTransaction(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// Create handles the creation of a new transaction.
// @Summary Create a new transaction
// @Tags Transaction
// @Description Create a new transaction record with the provided details.
// @Accept json
// @Produce json
// @Param CreateTransactionRequest body requests.CreateTransactionRequest true "Create Transaction Request"
// @Success 200 {object} pb.ApiResponseTransaction "Successfully created transaction record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create transaction"
// @Router /api/transaction/create [post]
func (h *transactionHandler) Create(c echo.Context) error {
	var body requests.CreateTransactionRequest

	apiKey := c.Get("apiKey").(string)

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid request body",
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

	res, err := h.transaction.CreateTransaction(ctx, &pb.CreateTransactionRequest{
		ApiKey:          apiKey,
		CardNumber:      body.CardNumber,
		Amount:          int32(body.Amount),
		PaymentMethod:   body.PaymentMethod,
		MerchantId:      int32(*body.MerchantID),
		TransactionTime: timestamppb.New(body.TransactionTime),
	})

	if err != nil {
		h.logger.Debug("Failed to create transaction", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transaction: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// Update updates an existing transaction with the provided details.
// @Summary Update a transaction
// @Tags Transaction
// @Description Update an existing transaction record using its ID
// @Accept json
// @Produce json
// @Param transaction body requests.UpdateTransactionRequest true "Transaction data"
// @Success 200 {object} pb.ApiResponseTransaction "Updated transaction data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update transaction"
// @Router /api/transaction/update [post]
func (h *transactionHandler) Update(c echo.Context) error {
	var body requests.UpdateTransactionRequest

	apiKey, ok := c.Get("apiKey").(string)
	if !ok {
		h.logger.Debug("Missing or invalid API key")
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid or missing API key",
		})
	}

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid request body",
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

	res, err := h.transaction.UpdateTransaction(ctx, &pb.UpdateTransactionRequest{
		TransactionId:   int32(body.TransactionID),
		CardNumber:      body.CardNumber,
		ApiKey:          apiKey,
		Amount:          int32(body.Amount),
		PaymentMethod:   body.PaymentMethod,
		MerchantId:      int32(*body.MerchantID),
		TransactionTime: timestamppb.New(body.TransactionTime),
	})

	if err != nil {
		h.logger.Debug("Failed to update transaction", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// TrashedTransaction trashes a transaction record.
// @Summary Trash a transaction
// @Tags Transaction
// @Description Trash a transaction record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} pb.ApiResponseTransaction "Successfully trashed transaction record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed transaction"
// @Router /api/transaction/trashed/{id} [post]
func (h *transactionHandler) TrashedTransaction(c echo.Context) error {
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

	res, err := h.transaction.TrashedTransaction(ctx, &pb.FindByIdTransactionRequest{
		TransactionId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to trashed transaction", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed transaction:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// RestoreTransaction restores a trashed transaction record.
// @Summary Restore a trashed transaction
// @Tags Transaction
// @Description Restore a trashed transaction record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} pb.ApiResponseTransaction "Successfully restored transaction record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore transaction:"
// @Router /api/transaction/restore/{id} [post]
func (h *transactionHandler) RestoreTransaction(c echo.Context) error {
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

	res, err := h.transaction.RestoreTransaction(ctx, &pb.FindByIdTransactionRequest{
		TransactionId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to restore transaction", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore transaction:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// DeletePermanent permanently deletes a transaction record by its ID.
// @Summary Permanently delete a transaction
// @Tags Transaction
// @Description Permanently delete a transaction record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} pb.ApiResponseTransactionDelete "Successfully deleted transaction record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete transaction:"
// @Router /api/transaction/delete/{id} [delete]
func (h *transactionHandler) DeletePermanent(c echo.Context) error {
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

	res, err := h.transaction.DeleteTransactionPermanent(ctx, &pb.FindByIdTransactionRequest{
		TransactionId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to delete transaction", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transaction:",
		})
	}

	return c.JSON(http.StatusOK, res)
}
