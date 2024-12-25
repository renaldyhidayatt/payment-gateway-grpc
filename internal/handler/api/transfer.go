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
	}
	routerTransfer := router.Group("/api/transfer")

	routerTransfer.GET("", transferHandler.FindAll)
	routerTransfer.GET("/:id", transferHandler.FindById)
	routerTransfer.GET("/transfer_from/:transfer_from", transferHandler.FindByTransferByTransferFrom)
	routerTransfer.GET("/transfer_to/:transfer_to", transferHandler.FindByTransferByTransferTo)

	routerTransfer.GET("/active", transferHandler.FindByActiveTransfer)
	routerTransfer.GET("/trashed", transferHandler.FindByTrashedTransfer)

	routerTransfer.POST("/create", transferHandler.CreateTransfer)
	routerTransfer.POST("/update/:id", transferHandler.UpdateTransfer)
	routerTransfer.POST("/trashed/:id", transferHandler.TrashTransfer)
	routerTransfer.POST("/restore/:id", transferHandler.RestoreTransfer)
	routerTransfer.DELETE("/permanent/:id", transferHandler.DeleteTransferPermanent)

	return transferHandler

}

// @Summary Find all transfer records
// @Tags Transfer
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

// FindById retrieves a transfer record by its ID.
// @Summary Find a transfer by ID
// @Tags Transfer
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

// FindByTransferByTransferFrom retrieves transfer records based on the transfer_from parameter.
// @Summary Find transfers by transfer_from
// @Tags Transfer
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

// FindByActiveTransfer retrieves a list of active transfer records.
// @Summary Find active transfers
// @Tags Transfer
// @Description Retrieve a list of active transfer records
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseTransfers "Active transfer data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
// @Router /api/transfer/active [get]

func (h *transferHandleApi) FindByActiveTransfer(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.FindByActiveTransfer(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Debug("Failed to retrieve transfer data: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// FindByTrashedTransfer retrieves a list of trashed transfer records.
// @Summary Retrieve trashed transfers
// @Tags Transfer
// @Description Retrieve a list of trashed transfer records
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseTransfers "List of trashed transfer records"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
// @Router /api/transfer/trashed [get]
func (h *transferHandleApi) FindByTrashedTransfer(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.FindByTrashedTransfer(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Debug("Failed to retrieve transfer data: ", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// CreateTransfer creates a new transfer record.
// @Summary Create a transfer
// @Tags Transfer
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

// UpdateTransfer updates an existing transfer record.
// @Summary Update a transfer
// @Tags Transfer
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

// TrashTransfer soft deletes a transfer record by its ID.
// @Summary Soft delete a transfer
// @Tags Transfer
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

// RestoreTransfer restores a trashed transfer record by its ID.
// @Summary Restore a trashed transfer
// @Tags Transfer
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

// DeleteTransferPermanent permanently deletes a transfer record by its ID.
// @Summary Permanently delete a transfer
// @Tags Transfer
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
