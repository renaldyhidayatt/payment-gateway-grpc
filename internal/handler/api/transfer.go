package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
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

func (h *transferHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

func (h *transferHandleApi) handleGetTransfers(c echo.Context) error {
	res, err := h.client.GetTransfers(c.Request().Context(), &emptypb.Empty{})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *transferHandleApi) handleGetTransfer(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetTransfer(c.Request().Context(), &pb.TransferRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *transferHandleApi) handleGetTransferByUsers(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetTransferByUsers(c.Request().Context(), &pb.TransferRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *transferHandleApi) GetTransferByUserId(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetTransferByUserId(c.Request().Context(), &pb.TransferRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *transferHandleApi) handleCreateTransfer(c echo.Context) error {
	var body requests.CreateTransferRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
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
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *transferHandleApi) handleUpdateTransfer(c echo.Context) error {
	var body requests.UpdateTransferRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
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
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *transferHandleApi) handleDeleteTransfer(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.DeleteTransfer(c.Request().Context(), &pb.TransferRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}
