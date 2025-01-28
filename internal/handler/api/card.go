package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	apimapper "MamangRust/paymentgatewaygrpc/internal/mapper/response/api"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type cardHandleApi struct {
	card    pb.CardServiceClient
	logger  logger.LoggerInterface
	mapping apimapper.CardResponseMapper
}

func NewHandlerCard(card pb.CardServiceClient, router *echo.Echo, logger logger.LoggerInterface, mapper apimapper.CardResponseMapper) *cardHandleApi {
	cardHandler := &cardHandleApi{
		card:    card,
		logger:  logger,
		mapping: mapper,
	}
	routerCard := router.Group("/api/card")

	routerCard.GET("", cardHandler.FindAll)
	routerCard.GET("/:id", cardHandler.FindById)

	routerCard.GET("/dashboard", cardHandler.DashboardCard)
	routerCard.GET("/dashboard/:cardNumber", cardHandler.DashboardCardCardNumber)

	routerCard.GET("/monthly-balance", cardHandler.FindMonthlyBalance)
	routerCard.GET("/yearly-balance", cardHandler.FindYearlyBalance)

	routerCard.GET("/monthly-topup-amount", cardHandler.FindMonthlyTopupAmount)
	routerCard.GET("/yearly-topup-amount", cardHandler.FindYearlyTopupAmount)
	routerCard.GET("/monthly-withdraw-amount", cardHandler.FindMonthlyWithdrawAmount)
	routerCard.GET("/yearly-withdraw-amount", cardHandler.FindYearlyWithdrawAmount)

	routerCard.GET("/monthly-transaction-amount", cardHandler.FindMonthlyTransactionAmount)
	routerCard.GET("/yearly-transaction-amount", cardHandler.FindYearlyTransactionAmount)

	routerCard.GET("/monthly-transfer-sender-amount", cardHandler.FindMonthlyTransferSenderAmount)
	routerCard.GET("/yearly-transfer-sender-amount", cardHandler.FindYearlyTransferSenderAmount)
	routerCard.GET("/monthly-transfer-receiver-amount", cardHandler.FindMonthlyTransferReceiverAmount)
	routerCard.GET("/yearly-transfer-receiver-amount", cardHandler.FindYearlyTransferReceiverAmount)

	routerCard.GET("/monthly-balance-by-card", cardHandler.FindMonthlyBalanceByCardNumber)
	routerCard.GET("/yearly-balance-by-card", cardHandler.FindYearlyBalanceByCardNumber)
	routerCard.GET("/monthly-topup-amount-by-card", cardHandler.FindMonthlyTopupAmountByCardNumber)
	routerCard.GET("/yearly-topup-amount-by-card", cardHandler.FindYearlyTopupAmountByCardNumber)

	routerCard.GET("/monthly-withdraw-amount-by-card", cardHandler.FindMonthlyWithdrawAmountByCardNumber)
	routerCard.GET("/yearly-withdraw-amount-by-card", cardHandler.FindYearlyWithdrawAmountByCardNumber)
	routerCard.GET("/monthly-transaction-amount-by-card", cardHandler.FindMonthlyTransactionAmountByCardNumber)
	routerCard.GET("/yearly-transaction-amount-by-card", cardHandler.FindYearlyTransactionAmountByCardNumber)

	routerCard.GET("/monthly-transfer-sender-amount-by-card", cardHandler.FindMonthlyTransferSenderAmountByCardNumber)
	routerCard.GET("/yearly-transfer-sender-amount-by-card", cardHandler.FindYearlyTransferSenderAmountByCardNumber)
	routerCard.GET("/monthly-transfer-receiver-amount-by-card", cardHandler.FindMonthlyTransferReceiverAmountByCardNumber)
	routerCard.GET("/yearly-transfer-receiver-amount-by-card", cardHandler.FindYearlyTransferReceiverAmountByCardNumber)

	routerCard.GET("/user", cardHandler.FindByUserID)
	routerCard.GET("/active", cardHandler.FindByActive)
	routerCard.GET("/trashed", cardHandler.FindByTrashed)
	routerCard.GET("/card_number/:card_number", cardHandler.FindByCardNumber)

	routerCard.POST("/create", cardHandler.CreateCard)
	routerCard.POST("/update/:id", cardHandler.UpdateCard)
	routerCard.POST("/trashed/:id", cardHandler.TrashedCard)
	routerCard.POST("/restore/:id", cardHandler.RestoreCard)
	routerCard.DELETE("/permanent/:id", cardHandler.DeleteCardPermanent)

	routerCard.POST("/restore/all", cardHandler.RestoreAllCard)
	routerCard.POST("/permanent/all", cardHandler.DeleteAllCardPermanent)

	return cardHandler
}

// FindAll godoc
// @Summary Retrieve all cards
// @Tags Card
// @Security Bearer
// @Description Retrieve all cards with pagination
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Number of data per page"
// @Param search query string false "Search keyword"
// @Success 200 {object} pb.ApiResponsePaginationCard "Card data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card data"
// @Router /api/card [get]
func (h *cardHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllCardRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	cards, err := h.card.FindAllCard(ctx, req)
	if err != nil {
		if errors.Is(err, echo.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status:  "error",
				Message: "Unauthorized",
			})
		}

		h.logger.Debug("Failed to fetch card records", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card records: ",
		})
	}

	so := h.mapping.ToApiResponsesCard(cards)

	return c.JSON(http.StatusOK, so)
}

// FindById godoc
// @Summary Retrieve card by ID
// @Tags Card
// @Security Bearer
// @Description Retrieve a card by its ID
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} pb.ApiResponseCard "Card data"
// @Failure 400 {object} response.ErrorResponse "Invalid card ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/{id} [get]
func (h *cardHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid card ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid card ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCardRequest{
		CardId: int32(id),
	}

	card, err := h.card.FindByIdCard(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to fetch card record", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseCard(card)

	return c.JSON(http.StatusOK, so)
}

// FindByUserID godoc
// @Summary Retrieve cards by user ID
// @Tags Card
// @Security Bearer
// @Description Retrieve a list of cards associated with a user by their ID
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseCard "Card data"
// @Failure 400 {object} response.ErrorResponse "Invalid user ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/user [get]
func (h *cardHandleApi) FindByUserID(c echo.Context) error {
	userID, ok := c.Get("user_id").(int32)
	if !ok {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to parse UserID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByUserIdCardRequest{
		UserId: userID,
	}

	card, err := h.card.FindByUserIdCard(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to fetch card record", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseCard(card)

	return c.JSON(http.StatusOK, so)
}

// DashboardCard godoc
// @Summary Get dashboard card data
// @Description Retrieve dashboard card data
// @Tags Card
// @Security Bearer
// @Produce json
// @Success 200 {object} pb.ApiResponseDashboardCard
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/dashboard [get]
func (h *cardHandleApi) DashboardCard(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.card.DashboardCard(ctx, &emptypb.Empty{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve dashboard card: %v", err),
		})
	}

	so := h.mapping.ToApiResponseDashboardCard(res)

	return c.JSON(http.StatusOK, so)
}

// DashboardCardCardNumber godoc
// @Summary Get dashboard card data by card number
// @Description Retrieve dashboard card data for a specific card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param cardNumber path string true "Card Number"
// @Success 200 {object} pb.ApiResponseDashboardCardNumber
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/dashboard/{cardNumber} [get]
func (h *cardHandleApi) DashboardCardCardNumber(c echo.Context) error {
	ctx := c.Request().Context()

	cardNumber := c.Param("cardNumber")

	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	req := &pb.FindByCardNumberRequest{
		CardNumber: cardNumber,
	}

	res, err := h.card.DashboardCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve dashboard card for card number %s: %v", cardNumber, err),
		})
	}

	so := h.mapping.ToApiResponseDashboardCardCardNumber(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyBalance godoc
// @Summary Get monthly balance data
// @Description Retrieve monthly balance data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMonthlyBalance
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-balance [get]
func (h *cardHandleApi) FindMonthlyBalance(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearBalance{
		Year: int32(year),
	}

	res, err := h.card.FindMonthlyBalance(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly balance: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyBalances(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyBalance godoc
// @Summary Get yearly balance data
// @Description Retrieve yearly balance data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseYearlyBalance
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-balance [get]
func (h *cardHandleApi) FindYearlyBalance(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearBalance{
		Year: int32(year),
	}

	res, err := h.card.FindYearlyBalance(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly balance: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyBalances(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTopupAmount godoc
// @Summary Get monthly topup amount data
// @Description Retrieve monthly topup amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-topup-amount [get]
func (h *cardHandleApi) FindMonthlyTopupAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindMonthlyTopupAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly topup amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyTopupAmount godoc
// @Summary Get yearly topup amount data
// @Description Retrieve yearly topup amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/topup/yearly-topup-amount [get]
func (h *cardHandleApi) FindYearlyTopupAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindYearlyTopupAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly topup amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyWithdrawAmount godoc
// @Summary Get monthly withdraw amount data
// @Description Retrieve monthly withdraw amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-withdraw-amount [get]
func (h *cardHandleApi) FindMonthlyWithdrawAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindMonthlyWithdrawAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly withdraw amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyWithdrawAmount godoc
// @Summary Get yearly withdraw amount data
// @Description Retrieve yearly withdraw amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-withdraw-amount [get]
func (h *cardHandleApi) FindYearlyWithdrawAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindYearlyWithdrawAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly withdraw amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTransactionAmount godoc
// @Summary Get monthly transaction amount data
// @Description Retrieve monthly transaction amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-transaction-amount [get]
func (h *cardHandleApi) FindMonthlyTransactionAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindMonthlyTransactionAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transaction amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyTransactionAmount godoc
// @Summary Get yearly transaction amount data
// @Description Retrieve yearly transaction amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-transaction-amount [get]
func (h *cardHandleApi) FindYearlyTransactionAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindYearlyTransactionAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transaction amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTransferSenderAmount godoc
// @Summary Get monthly transfer sender amount data
// @Description Retrieve monthly transfer sender amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-transfer-sender-amount [get]
func (h *cardHandleApi) FindMonthlyTransferSenderAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindMonthlyTransferSenderAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transfer sender amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyTransferSenderAmount godoc
// @Summary Get yearly transfer sender amount data
// @Description Retrieve yearly transfer sender amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transfer/yearly-transfer-sender-amount [get]
func (h *cardHandleApi) FindYearlyTransferSenderAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindYearlyTransferSenderAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transfer sender amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTransferReceiverAmount godoc
// @Summary Get monthly transfer receiver amount data
// @Description Retrieve monthly transfer receiver amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-transfer-receiver-amount [get]
func (h *cardHandleApi) FindMonthlyTransferReceiverAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindMonthlyTransferReceiverAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transfer receiver amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyTransferReceiverAmount godoc
// @Summary Get yearly transfer receiver amount data
// @Description Retrieve yearly transfer receiver amount data for a specific year
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-transfer-receiver-amount [get]
func (h *cardHandleApi) FindYearlyTransferReceiverAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindYearlyTransferReceiverAmount(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transfer receiver amount: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyBalanceByCardNumber godoc
// @Summary Get monthly balance data by card number
// @Description Retrieve monthly balance data for a specific year and card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} pb.ApiResponseMonthlyBalance
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-balance-by-card [get]
func (h *cardHandleApi) FindMonthlyBalanceByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")
	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearBalanceCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindMonthlyBalanceByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly balance: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyBalances(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyBalanceByCardNumber godoc
// @Summary Get yearly balance data by card number
// @Description Retrieve yearly balance data for a specific year and card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} pb.ApiResponseYearlyBalance
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-balance-by-card [get]
func (h *cardHandleApi) FindYearlyBalanceByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")
	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearBalanceCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindYearlyBalanceByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly balance: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyBalances(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTopupAmountByCardNumber godoc
// @Summary Get monthly topup amount data by card number
// @Description Retrieve monthly topup amount data for a specific year and card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} pb.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-topup-amount-by-card [get]
func (h *cardHandleApi) FindMonthlyTopupAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")
	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindMonthlyTopupAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly topup amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyTopupAmountByCardNumber godoc
// @Summary Get yearly topup amount data by card number
// @Description Retrieve yearly topup amount data for a specific year and card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-topup-amount-by-card [get]
func (h *cardHandleApi) FindYearlyTopupAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindYearlyTopupAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly topup amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyWithdrawAmountByCardNumber godoc
// @Summary Get monthly withdraw amount data by card number
// @Description Retrieve monthly withdraw amount data for a specific year and card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} pb.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-withdraw-amount-by-card [get]
func (h *cardHandleApi) FindMonthlyWithdrawAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")
	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindMonthlyWithdrawAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly withdraw amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyWithdrawAmountByCardNumber godoc
// @Summary Get yearly withdraw amount data by card number
// @Description Retrieve yearly withdraw amount data for a specific year and card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-withdraw-amount-by-card [get]
func (h *cardHandleApi) FindYearlyWithdrawAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindYearlyWithdrawAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly withdraw amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTransactionAmountByCardNumber godoc
// @Summary Get monthly transaction amount data by card number
// @Description Retrieve monthly transaction amount data for a specific year and card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} pb.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-transaction-amount-by-card [get]
func (h *cardHandleApi) FindMonthlyTransactionAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindMonthlyTransactionAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transaction amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyTransactionAmountByCardNumber godoc
// @Summary Get yearly transaction amount data by card number
// @Description Retrieve yearly transaction amount data for a specific year and card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-transaction-amount-by-card [get]
func (h *cardHandleApi) FindYearlyTransactionAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindYearlyTransactionAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transaction amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTransferSenderAmountByCardNumber godoc
// @Summary Get monthly transfer sender amount data by card number
// @Description Retrieve monthly transfer sender amount data for a specific year and card number
// @Tags Card
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} pb.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-transfer-sender-amount-by-card [get]
func (h *cardHandleApi) FindMonthlyTransferSenderAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")
	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindMonthlyTransferSenderAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transfer sender amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyTransferSenderAmountByCardNumber godoc
// @Summary Get yearly transfer sender amount by card number
// @Description Retrieve the total amount sent by a specific card number in a given year
// @Tags Card
// @Security Bearer
// @Accept json
// @Produce json
// @Param year query int true "Year for which the data is requested"
// @Param card_number query string true "Card number for which the data is requested"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-transfer-sender-amount-by-card [get]
func (h *cardHandleApi) FindYearlyTransferSenderAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")
	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindYearlyTransferSenderAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transfer sender amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTransferReceiverAmountByCardNumber godoc
// @Summary Get monthly transfer receiver amount by card number
// @Description Retrieve the total amount received by a specific card number in a given year, broken down by month
// @Tags Card
// @Security Bearer
// @Accept json
// @Produce json
// @Param year query int true "Year for which the data is requested"
// @Param card_number query string true "Card number for which the data is requested"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/monthly-transfer-receiver-amount-by-card [get]
func (h *cardHandleApi) FindMonthlyTransferReceiverAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindMonthlyTransferReceiverAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve monthly transfer receiver amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseMonthlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlyTransferReceiverAmountByCardNumber godoc
// @Summary Get yearly transfer receiver amount by card number
// @Description Retrieve the total amount received by a specific card number in a given year
// @Tags Card
// @Security Bearer
// @Accept json
// @Produce json
// @Param year query int true "Year for which the data is requested"
// @Param card_number query string true "Card number for which the data is requested"
// @Success 200 {object} pb.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card/yearly-transfer-receiver-amount-by-card [get]
func (h *cardHandleApi) FindYearlyTransferReceiverAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year format",
		})
	}

	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Card number is required",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindYearlyTransferReceiverAmountByCardNumber(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve yearly transfer receiver amount by card number: %v", err),
		})
	}

	so := h.mapping.ToApiResponseYearlyAmounts(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active card by Saldo ID
// @Tags Card
// @Description Retrieve an active card associated with a Saldo ID
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsePaginationCardDeleteAt "Card data"
// @Failure 400 {object} response.ErrorResponse "Invalid Saldo ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/active [get]
func (h *cardHandleApi) FindByActive(c echo.Context) error {
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

	req := &pb.FindAllCardRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.card.FindByActiveCard(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to fetch card record", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponsesCardDeletedAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Retrieve trashed cards
// @Tags Card
// @Security Bearer
// @Description Retrieve a list of trashed cards
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsePaginationCardDeleteAt "Card data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/trashed [get]
func (h *cardHandleApi) FindByTrashed(c echo.Context) error {
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

	req := &pb.FindAllCardRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.card.FindByTrashedCard(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to fetch card record", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponsesCardDeletedAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve card by card number
// @Tags Card
// @Description Retrieve a card by its card number
// @Accept json
// @Produce json
// @Param card_number path string true "Card number"
// @Success 200 {object} pb.ApiResponseCard "Card data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/{card_number} [get]
func (h *cardHandleApi) FindByCardNumber(c echo.Context) error {
	cardNumber := c.Param("card_number")

	ctx := c.Request().Context()

	req := &pb.FindByCardNumberRequest{
		CardNumber: cardNumber,
	}

	res, err := h.card.FindByCardNumber(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to fetch card record", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: ",
		})
	}

	so := h.mapping.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Create a new card
// @Tags Card
// @Description Create a new card for a user
// @Accept json
// @Produce json
// @Param CreateCardRequest body requests.CreateCardRequest true "Create card request"
// @Success 200 {object} pb.ApiResponseCard "Created card"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create card"
// @Router /api/card/create [post]
func (h *cardHandleApi) CreateCard(c echo.Context) error {
	var body requests.CreateCardRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Bad Request: ", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error: ", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: ",
		})
	}

	ctx := c.Request().Context()

	req := &pb.CreateCardRequest{
		UserId:       int32(body.UserID),
		CardType:     body.CardType,
		ExpireDate:   timestamppb.New(body.ExpireDate),
		Cvv:          body.CVV,
		CardProvider: body.CardProvider,
	}

	res, err := h.card.CreateCard(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to create card", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create card: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update a card
// @Tags Card
// @Description Update a card for a user
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Param UpdateCardRequest body requests.UpdateCardRequest true "Update card request"
// @Success 200 {object} pb.ApiResponseCard "Updated card"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update card"
// @Router /api/card/update/{id} [post]
func (h *cardHandleApi) UpdateCard(c echo.Context) error {
	var body requests.UpdateCardRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Bad Request: ", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
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

	req := &pb.UpdateCardRequest{
		CardId:       int32(body.CardID),
		UserId:       int32(body.UserID),
		CardType:     body.CardType,
		ExpireDate:   timestamppb.New(body.ExpireDate),
		Cvv:          body.CVV,
		CardProvider: body.CardProvider,
	}

	res, err := h.card.UpdateCard(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to update card", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update card: ",
		})
	}

	so := h.mapping.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Trashed a card
// @Tags Card
// @Description Trashed a card by its ID
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} pb.ApiResponseCard "Trashed card"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed card"
// @Router /api/card/trashed/{id} [post]
func (h *cardHandleApi) TrashedCard(c echo.Context) error {
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

	req := &pb.FindByIdCardRequest{
		CardId: int32(idInt),
	}

	res, err := h.card.TrashedCard(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to trashed card", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed card: ",
		})
	}

	so := h.mapping.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore a card
// @Tags Card
// @Description Restore a card by its ID
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} pb.ApiResponseCard "Restored card"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore card"
// @Router /api/card/restore/{id} [post]
func (h *cardHandleApi) RestoreCard(c echo.Context) error {

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

	req := &pb.FindByIdCardRequest{
		CardId: int32(idInt),
	}

	res, err := h.card.RestoreCard(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to restore card", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore card: ",
		})
	}

	so := h.mapping.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Delete a card permanently
// @Tags Card
// @Description Delete a card by its ID permanently
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} pb.ApiResponseCardDelete "Deleted card"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete card"
// @Router /api/card/permanent/{id} [delete]
func (h *cardHandleApi) DeleteCardPermanent(c echo.Context) error {
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

	req := &pb.FindByIdCardRequest{
		CardId: int32(idInt),
	}

	res, err := h.card.DeleteCardPermanent(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to delete card", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete card: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseCardDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore all card records
// @Tags Card
// @Description Restore all card records that were previously deleted.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseCardAll "Successfully restored all card records"
// @Failure 500 {object} response.ErrorResponse "Failed to restore all card records"
// @Router /api/card/restore/all [post]
func (h *cardHandleApi) RestoreAllCard(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.card.RestoreAllCard(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Error("Failed to restore all cards", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently restore all cards",
		})
	}

	h.logger.Debug("Successfully restored all cards")

	so := h.mapping.ToApiResponseCardAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer.
// @Summary Permanently delete all card records
// @Tags Card
// @Description Permanently delete all card records from the database.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseCardAll "Successfully deleted all card records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to permanently delete all card records"
// @Router /api/card/permanent/all [post]
func (h *cardHandleApi) DeleteAllCardPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.card.DeleteAllCardPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to permanently delete all cards", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all cards",
		})
	}

	h.logger.Debug("Successfully deleted all cards permanently")

	so := h.mapping.ToApiResponseCardAll(res)

	return c.JSON(http.StatusOK, so)
}
