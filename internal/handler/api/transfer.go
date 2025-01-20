package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transferHandleApi struct {
	client pb.TransferServiceClient
	logger logger.LoggerInterface
}

func NewHandlerTransfer(client pb.TransferServiceClient, router *echo.Echo, logger logger.LoggerInterface) *transferHandleApi {
	transferHandler := &transferHandleApi{
		client: client,
		logger: logger,
	}
	routerTransfer := router.Group("/api/transfers")

	routerTransfer.GET("", transferHandler.FindAll)
	routerTransfer.GET("/:id", transferHandler.FindById)

	routerTransfer.GET("/monthly-success", transferHandler.FindMonthlyTransferStatusSuccess)
	routerTransfer.GET("/yearly-success", transferHandler.FindYearlyTransferStatusSuccess)

	routerTransfer.GET("/monthly-failed", transferHandler.FindMonthlyTransferStatusFailed)
	routerTransfer.GET("/yearly-failed", transferHandler.FindYearlyTransferStatusFailed)

	routerTransfer.GET("/monthly", transferHandler.FindMonthlyTransferAmounts)
	routerTransfer.GET("/yearly", transferHandler.FindYearlyTransferAmounts)
	routerTransfer.GET("/monthly-by-sender", transferHandler.FindMonthlyTransferAmountsBySenderCardNumber)
	routerTransfer.GET("/monthly-by-receiver", transferHandler.FindMonthlyTransferAmountsByReceiverCardNumber)
	routerTransfer.GET("/yearly-by-sender", transferHandler.FindYearlyTransferAmountsBySenderCardNumber)
	routerTransfer.GET("/yearly-by-receiver", transferHandler.FindYearlyTransferAmountsByReceiverCardNumber)

	routerTransfer.GET("/transfer_from/:transfer_from", transferHandler.FindByTransferByTransferFrom)
	routerTransfer.GET("/transfer_to/:transfer_to", transferHandler.FindByTransferByTransferTo)

	routerTransfer.GET("/active", transferHandler.FindByActiveTransfer)
	routerTransfer.GET("/trashed", transferHandler.FindByTrashedTransfer)

	routerTransfer.POST("/create", transferHandler.CreateTransfer)
	routerTransfer.POST("/update/:id", transferHandler.UpdateTransfer)
	routerTransfer.POST("/trashed/:id", transferHandler.TrashTransfer)
	routerTransfer.POST("/restore/:id", transferHandler.RestoreTransfer)
	routerTransfer.DELETE("/permanent/:id", transferHandler.DeleteTransferPermanent)

	routerTransfer.POST("/restore/all", transferHandler.RestoreAllTransfer)
	routerTransfer.POST("/permanent/all", transferHandler.DeleteAllTransferPermanent)

	return transferHandler

}

// @Summary Find all transfer records
// @Tags Transfer
// @Security Bearer
// @Description Retrieve a list of all transfer records with pagination
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} pb.ApiResponsePaginationTransfer "List of transfer records"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
// @Router /api/transfer [get]
func (h *transferHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllTransferRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAllTransfer(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve transfer data: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Find a transfer by ID
// @Tags Transfer
// @Security Bearer
// @Description Retrieve a transfer record using its ID
// @Accept json
// @Produce json
// @Param id path string true "Transfer ID"
// @Success 200 {object} pb.ApiResponseTransfer "Transfer data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
// @Router /api/transfer/{id} [get]
func (h *transferHandleApi) FindById(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})

	}

	ctx := c.Request().Context()

	res, err := h.client.FindByIdTransfer(ctx, &pb.FindByIdTransferRequest{
		TransferId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve transfer data: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindMonthlyTransferStatusSuccess retrieves the monthly transfer status for successful transactions.
// @Summary Get monthly transfer status for successful transactions
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the monthly transfer status for successful transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} pb.ApiResponseTransferMonthStatusSuccess "Monthly transfer status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer status for successful transactions"
// @Router /api/transfers/monthly-success [get]
func (h *transferHandleApi) FindMonthlyTransferStatusSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid month",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindMonthlyTransferStatusSuccess(ctx, &pb.FindMonthlyTransferStatus{
		Year:  int32(year),
		Month: int32(month),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve monthly Transfer status success", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve monthly Transfer status success: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindYearlyTransferStatusSuccess retrieves the yearly transfer status for successful transactions.
// @Summary Get yearly transfer status for successful transactions
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the yearly transfer status for successful transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseTransferYearStatusSuccess "Yearly transfer status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer status for successful transactions"
// @Router /api/transfers/yearly-success [get]
func (h *transferHandleApi) FindYearlyTransferStatusSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindYearlyTransferStatusSuccess(ctx, &pb.FindYearTransfer{
		Year: int32(year),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve yearly Transfer status success", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve yearly Transfer status success: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindMonthlyTransferStatusFailed retrieves the monthly transfer status for failed transactions.
// @Summary Get monthly transfer status for failed transactions
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the monthly transfer status for failed transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} pb.ApiResponseTransferMonthStatusFailed "Monthly transfer status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer status for failed transactions"
// @Router /api/transfers/monthly-failed [get]
func (h *transferHandleApi) FindMonthlyTransferStatusFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid month",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindMonthlyTransferStatusFailed(ctx, &pb.FindMonthlyTransferStatus{
		Year:  int32(year),
		Month: int32(month),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve monthly Transfer status Failed", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve monthly Transfer status Failed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindYearlyTransferStatusFailed retrieves the yearly transfer status for failed transactions.
// @Summary Get yearly transfer status for failed transactions
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the yearly transfer status for failed transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseTransferYearStatusFailed "Yearly transfer status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer status for failed transactions"
// @Router /api/transfers/yearly-failed [get]
func (h *transferHandleApi) FindYearlyTransferStatusFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid year",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindYearlyTransferStatusFailed(ctx, &pb.FindYearTransfer{
		Year: int32(year),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve yearly Transfer status Failed", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve yearly Transfer status Failed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindMonthlyTransferAmounts retrieves the monthly transfer amounts for a specific year.
// @Summary Get monthly transfer amounts
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the monthly transfer amounts for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseTransferMonthAmount "Monthly transfer amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer amounts"
// @Router /api/transfers/monthly-amounts [get]
func (h *transferHandleApi) FindMonthlyTransferAmounts(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindMonthlyTransferAmounts(ctx, &pb.FindYearTransfer{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transfer amounts", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve monthly transfer amounts",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindYearlyTransferAmounts retrieves the yearly transfer amounts for a specific year.
// @Summary Get yearly transfer amounts
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the yearly transfer amounts for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseTransferYearAmount "Yearly transfer amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer amounts"
// @Router /api/transfers/yearly-amounts [get]
func (h *transferHandleApi) FindYearlyTransferAmounts(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindYearlyTransferAmounts(ctx, &pb.FindYearTransfer{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transfer amounts", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve yearly transfer amounts",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindMonthlyTransferAmountsBySenderCardNumber retrieves the monthly transfer amounts for a specific sender card number and year.
// @Summary Get monthly transfer amounts by sender card number
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the monthly transfer amounts for a specific sender card number and year.
// @Accept json
// @Produce json
// @Param card_number query string true "Sender Card Number"
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseTransferMonthAmount "Monthly transfer amounts by sender card number"
// @Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer amounts by sender card number"
// @Router /api/transfers/monthly-amounts-by-sender-card [get]
func (h *transferHandleApi) FindMonthlyTransferAmountsBySenderCardNumber(c echo.Context) error {
	cardNumber := c.QueryParam("card_number")
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindMonthlyTransferAmountsBySenderCardNumber(ctx, &pb.FindByCardNumberTransferRequest{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transfer amounts by sender card number", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve monthly transfer amounts by sender card number",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindMonthlyTransferAmountsByReceiverCardNumber retrieves the monthly transfer amounts for a specific receiver card number and year.
// @Summary Get monthly transfer amounts by receiver card number
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the monthly transfer amounts for a specific receiver card number and year.
// @Accept json
// @Produce json
// @Param card_number query string true "Receiver Card Number"
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseTransferMonthAmount "Monthly transfer amounts by receiver card number"
// @Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer amounts by receiver card number"
// @Router /api/transfers/monthly-amounts-by-receiver-card [get]
func (h *transferHandleApi) FindMonthlyTransferAmountsByReceiverCardNumber(c echo.Context) error {
	cardNumber := c.QueryParam("card_number")
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindMonthlyTransferAmountsByReceiverCardNumber(ctx, &pb.FindByCardNumberTransferRequest{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transfer amounts by receiver card number", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve monthly transfer amounts by receiver card number",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindYearlyTransferAmountsBySenderCardNumber retrieves the yearly transfer amounts for a specific sender card number and year.
// @Summary Get yearly transfer amounts by sender card number
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the yearly transfer amounts for a specific sender card number and year.
// @Accept json
// @Produce json
// @Param card_number query string true "Sender Card Number"
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseTransferYearAmount "Yearly transfer amounts by sender card number"
// @Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer amounts by sender card number"
// @Router /api/transfers/yearly-amounts-by-sender-card [get]
func (h *transferHandleApi) FindYearlyTransferAmountsBySenderCardNumber(c echo.Context) error {
	cardNumber := c.QueryParam("card_number")
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindYearlyTransferAmountsBySenderCardNumber(ctx, &pb.FindByCardNumberTransferRequest{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transfer amounts by sender card number", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve yearly transfer amounts by sender card number",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindYearlyTransferAmountsByReceiverCardNumber retrieves the yearly transfer amounts for a specific receiver card number and year.
// @Summary Get yearly transfer amounts by receiver card number
// @Tags Transfer
// @Security Bearer
// @Description Retrieve the yearly transfer amounts for a specific receiver card number and year.
// @Accept json
// @Produce json
// @Param card_number query string true "Receiver Card Number"
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseTransferYearAmount "Yearly transfer amounts by receiver card number"
// @Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer amounts by receiver card number"
// @Router /api/transfers/yearly-amounts-by-receiver-card [get]
func (h *transferHandleApi) FindYearlyTransferAmountsByReceiverCardNumber(c echo.Context) error {
	cardNumber := c.QueryParam("card_number")
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindYearlyTransferAmountsByReceiverCardNumber(ctx, &pb.FindByCardNumberTransferRequest{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transfer amounts by receiver card number", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve yearly transfer amounts by receiver card number",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Find transfers by transfer_from
// @Tags Transfer
// @Security Bearer
// @Description Retrieve a list of transfer records using the transfer_from parameter
// @Accept json
// @Produce json
// @Param transfer_from path string true "Transfer From"
// @Success 200 {object} pb.ApiResponseTransfers "Transfer data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
// @Router /api/transfer/transfer_from/{transfer_from} [get]
func (h *transferHandleApi) FindByTransferByTransferFrom(c echo.Context) error {
	transfer_from := c.Param("transfer_from")

	ctx := c.Request().Context()

	res, err := h.client.FindTransferByTransferFrom(ctx, &pb.FindTransferByTransferFromRequest{
		TransferFrom: transfer_from,
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve transfer data: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Find transfers by transfer_to
// @Tags Transfer
// @Security Bearer
// @Description Retrieve a list of transfer records using the transfer_to parameter
// @Accept json
// @Produce json
// @Param transfer_to path string true "Transfer To"
// @Success 200 {object} pb.ApiResponseTransfers "Transfer data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
// @Router /api/transfer/transfer_to/{transfer_to} [get]
func (h *transferHandleApi) FindByTransferByTransferTo(c echo.Context) error {
	transfer_to := c.Param("transfer_to")

	ctx := c.Request().Context()

	res, err := h.client.FindTransferByTransferTo(ctx, &pb.FindTransferByTransferToRequest{
		TransferTo: transfer_to,
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve transfer data: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Find active transfers
// @Tags Transfer
// @Security Bearer
// @Description Retrieve a list of active transfer records
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseTransfers "Active transfer data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
// @Router /api/transfer/active [get]

func (h *transferHandleApi) FindByActiveTransfer(c echo.Context) error {
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

	req := &pb.FindAllTransferRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByActiveTransfer(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve transfer data: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Retrieve trashed transfers
// @Tags Transfer
// @Security Bearer
// @Description Retrieve a list of trashed transfer records
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseTransfers "List of trashed transfer records"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
// @Router /api/transfer/trashed [get]
func (h *transferHandleApi) FindByTrashedTransfer(c echo.Context) error {
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

	req := &pb.FindAllTransferRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByTrashedTransfer(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve transfer data: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Create a transfer
// @Tags Transfer
// @Security Bearer
// @Description Create a new transfer record
// @Accept json
// @Produce json
// @Param body body requests.CreateTransferRequest true "Transfer request"
// @Success 200 {object} pb.ApiResponseTransfer "Transfer data"
// @Failure 400 {object} response.ErrorResponse "Validation Error"
// @Failure 500 {object} response.ErrorResponse "Failed to create transfer"
// @Router /api/transfer/create [post]
func (h *transferHandleApi) CreateTransfer(c echo.Context) error {
	var body requests.CreateTransferRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request body: ", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error: ", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.CreateTransfer(ctx, &pb.CreateTransferRequest{
		TransferFrom:   body.TransferFrom,
		TransferTo:     body.TransferTo,
		TransferAmount: int32(body.TransferAmount),
	})

	if err != nil {
		h.logger.Debug("Failed to create transfer: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transfer: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Update a transfer
// @Tags Transfer
// @Security Bearer
// @Description Update an existing transfer record
// @Accept json
// @Produce json
// @Param id path int true "Transfer ID"
// @Param body body requests.UpdateTransferRequest true "Transfer request"
// @Success 200 {object} pb.ApiResponseTransfer "Transfer data"
// @Failure 400 {object} response.ErrorResponse "Validation Error"
// @Failure 500 {object} response.ErrorResponse "Failed to update transfer"
// @Router /api/transfer/update/{id} [post]
func (h *transferHandleApi) UpdateTransfer(c echo.Context) error {
	var body requests.UpdateTransferRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request body: ", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error: ", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.UpdateTransfer(ctx, &pb.UpdateTransferRequest{
		TransferId:     int32(body.TransferID),
		TransferFrom:   body.TransferFrom,
		TransferTo:     body.TransferTo,
		TransferAmount: int32(body.TransferAmount),
	})

	if err != nil {
		h.logger.Debug("Failed to update transfer: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transfer: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Soft delete a transfer
// @Tags Transfer
// @Security Bearer
// @Description Soft delete a transfer record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transfer ID"
// @Success 200 {object} pb.ApiResponseTransfer "Successfully trashed transfer record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed transfer"
// @Router /api/transfer/trash/{id} [post]
func (h *transferHandleApi) TrashTransfer(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})

	}

	ctx := c.Request().Context()

	res, err := h.client.TrashedTransfer(ctx, &pb.FindByIdTransferRequest{
		TransferId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to trash transfer: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash transfer:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Restore a trashed transfer
// @Tags Transfer
// @Security Bearer
// @Description Restore a trashed transfer record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transfer ID"
// @Success 200 {object} pb.ApiResponseTransfer "Successfully restored transfer record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore transfer:"
// @Router /api/transfer/restore/{id} [post]
func (h *transferHandleApi) RestoreTransfer(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.RestoreTransfer(ctx, &pb.FindByIdTransferRequest{
		TransferId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to restore transfer: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore transfer:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Permanently delete a transfer
// @Tags Transfer
// @Security Bearer
// @Description Permanently delete a transfer record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transfer ID"
// @Success 200 {object} pb.ApiResponseTransferDelete "Successfully deleted transfer record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete transfer:"
// @Router /api/transfer/permanent/{id} [delete]
func (h *transferHandleApi) DeleteTransferPermanent(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.DeleteTransferPermanent(ctx, &pb.FindByIdTransferRequest{
		TransferId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transfer:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Restore a trashed transfer
// @Tags Transfer
// @Security Bearer
// @Description Restore a trashed transfer all
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseTransferAll "Successfully restored transfer record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore transfer:"
// @Router /api/transfer/restore/all [post]
func (h *transferHandleApi) RestoreAllTransfer(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.RestoreAllTransfer(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to restore all transfer", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently restore all transfer",
		})
	}

	h.logger.Debug("Successfully restored all transfer")

	return c.JSON(http.StatusOK, res)
}

// @Summary Permanently delete a transfer
// @Tags Transfer
// @Security Bearer
// @Description Permanently delete a transfer record all.
// @Accept json
// @Produce json
// @Param id path int true "Transfer ID"
// @Success 200 {object} pb.ApiResponseTransferAll "Successfully deleted transfer all"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete transfer:"
// @Router /api/transfer/permanent/{id} [delete]
func (h *transferHandleApi) DeleteAllTransferPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.DeleteAllTransferPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to permanently delete all transfer", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all transfer",
		})
	}

	h.logger.Debug("Successfully deleted all transfer permanently")

	return c.JSON(http.StatusOK, res)
}
