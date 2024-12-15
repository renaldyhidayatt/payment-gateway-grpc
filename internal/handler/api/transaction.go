package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type transactionHandler struct {
	transaction pb.TransactionServiceClient
}

func NewTransactionHandler(transaction pb.TransactionServiceClient, router *echo.Echo) *transactionHandler {
	transactionHandler := transactionHandler{
		transaction: transaction,
	}

	return &transactionHandler
}

func (h *transactionHandler) FindByActive(c echo.Context) error {
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) FindById(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) FindByCardNumber(c echo.Context) error {
	cardNumber := c.QueryParam("card_number")

	ctx := c.Request().Context()

	req := &pb.FindByCardNumberTransactionRequest{
		CardNumber: cardNumber,
	}

	res, err := h.transaction.FindByCardNumberTransaction(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) FindByActiveTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.transaction.FindByActiveTransaction(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) FindByTrashedTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.transaction.FindByTrashedTransaction(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transaction data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) Create(c echo.Context) error {
	var body requests.CreateTransactionRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.transaction.CreateTransaction(ctx, &pb.CreateTransactionRequest{
		ApiKey:          "",
		CardNumber:      body.CardNumber,
		Amount:          int32(body.Amount),
		PaymentMethod:   body.PaymentMethod,
		MerchantId:      int32(*body.MerchantID),
		TransactionTime: timestamppb.New(body.TransactionTime),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transaction: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) Update(c echo.Context) error {
	var body requests.UpdateTransactionRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.transaction.UpdateTransaction(ctx, &pb.UpdateTransactionRequest{
		TransactionId:   int32(body.TransactionID),
		CardNumber:      body.CardNumber,
		Amount:          int32(body.Amount),
		PaymentMethod:   body.PaymentMethod,
		MerchantId:      int32(*body.MerchantID),
		TransactionTime: timestamppb.New(body.TransactionTime),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) TrashedTransaction(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed transaction:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) RestoreTransaction(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore transaction:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transactionHandler) DeletePermanent(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transaction:",
		})
	}

	return c.JSON(http.StatusOK, res)
}
