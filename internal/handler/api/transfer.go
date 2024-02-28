package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transferHandleApi struct {
	client pb.TransferServiceClient
}

func NewHandlerTransfer(client pb.TransferServiceClient, router *echo.Echo) *transferHandleApi {
	transferHandler := &transferHandleApi{
		client: client,
	}
	routerTransfer := router.Group("/api/transfer")

	routerTransfer.GET("/hello", transferHandler.handleHello)
	routerTransfer.GET("/", transferHandler.handleGetTransfers)
	routerTransfer.GET("/:id", transferHandler.handleGetTransfer)
	routerTransfer.GET("/user-all/:id", transferHandler.handleGetTransferByUsers)
	routerTransfer.GET("/user/:id", transferHandler.GetTransferByUserId)
	routerTransfer.POST("/create", transferHandler.handleCreateTransfer)
	routerTransfer.PUT("/update/:id", transferHandler.handleUpdateTransfer)
	routerTransfer.DELETE("/:id", transferHandler.handleDeleteTransfer)

	return transferHandler

}

// @Summary Get a greeting message
// @Description Get a greeting message
// @Tags Transfer
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello"
// @Router /transfer/hello [get]
func (h *transferHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

// @Summary Get all transfers
// @Description Get all transfers
// @Tags Transfer
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Failed to retrieve transfers: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /transfer [get]
func (h *transferHandleApi) handleGetTransfers(c echo.Context) error {
	res, err := h.client.GetTransfers(c.Request().Context(), &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to retrieve transfers: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get a transfer by ID
// @Description Get a transfer by ID
// @Tags Transfer
// @Accept json
// @Produce json
// @Param id path int true "Transfer ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /transfer/{id} [get]
func (h *transferHandleApi) handleGetTransfer(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.GetTransfer(c.Request().Context(), &pb.TransferRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to retrieve transfer: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get transfers by user ID
// @Description Get transfers by user ID
// @Tags Transfer
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /transfer/user-all/{id} [get]
func (h *transferHandleApi) handleGetTransferByUsers(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.GetTransferByUsers(c.Request().Context(), &pb.TransferRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to retrieve transfers by user: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get a transfer by user ID
// @Description Get a transfer by user ID
// @Tags Transfer
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /transfer/user/{id} [get]
func (h *transferHandleApi) GetTransferByUserId(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.GetTransferByUserId(c.Request().Context(), &pb.TransferRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to retrieve transfer by user: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Create a new transfer
// @Description Create a new transfer
// @Tags Transfer
// @Accept json
// @Produce json
// @Param body body requests.CreateTransferRequest true "Transfer details"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /transfer/create [post]
func (h *transferHandleApi) handleCreateTransfer(c echo.Context) error {
	var body requests.CreateTransferRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request Validate: " + err.Error(),
			Data:       nil,
		})
	}

	data := &pb.CreateTransferRequest{
		TransferFrom:   int32(body.TransferFrom),
		TransferTo:     int32(body.TransferTo),
		TransferAmount: int32(body.TransferAmount),
	}

	res, err := h.client.CreateTransfer(c.Request().Context(), data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create transfer: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Update a transfer
// @Description Update a transfer
// @Tags Transfer
// @Accept json
// @Produce json
// @Param body body requests.UpdateTransferRequest true "Transfer details"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /transfer/update/{id} [put]
func (h *transferHandleApi) handleUpdateTransfer(c echo.Context) error {
	var body requests.UpdateTransferRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request Validate: " + err.Error(),
			Data:       nil,
		})
	}

	data := &pb.UpdateTransferRequest{
		Id:             int32(body.TransferID),
		TransferFrom:   int32(body.TransferFrom),
		TransferTo:     int32(body.TransferTo),
		TransferAmount: int32(body.TransferAmount),
	}

	res, err := h.client.UpdateTransfer(c.Request().Context(), data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update transfer: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Delete a transfer by ID
// @Description Delete a transfer by ID
// @Tags Transfer
// @Accept json
// @Produce json
// @Param id path int true "Transfer ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /transfer/delete/{id} [delete]
func (h *transferHandleApi) handleDeleteTransfer(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.DeleteTransfer(c.Request().Context(), &pb.TransferRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to delete transfer: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}
