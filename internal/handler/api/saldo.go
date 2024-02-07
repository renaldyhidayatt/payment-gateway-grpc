package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type saldoHandleApi struct {
	client pb.SaldoServiceClient
}

func NewHandlerSaldo(client pb.SaldoServiceClient, router *echo.Echo) *saldoHandleApi {
	saldoHandler := &saldoHandleApi{
		client: client,
	}
	routerSaldo := router.Group("/api/saldo")

	routerSaldo.GET("/hello", saldoHandler.handleHello)
	routerSaldo.GET("/", saldoHandler.handleGetSaldos)
	routerSaldo.GET("/:id", saldoHandler.handleGetSaldo)
	routerSaldo.GET("/user-all/:id", saldoHandler.handleGetSaldoByUsers)
	routerSaldo.GET("/user/:id", saldoHandler.handleGetSaldobyUserId)
	routerSaldo.POST("/create", saldoHandler.handleCreateSaldo)
	routerSaldo.PUT("/update/:id", saldoHandler.handleUpdateSaldo)
	routerSaldo.DELETE("/:id", saldoHandler.handleDeleteSaldo)

	return saldoHandler

}

func (h *saldoHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

func (h *saldoHandleApi) handleGetSaldos(c echo.Context) error {
	res, err := h.client.GetSaldos(c.Request().Context(), &emptypb.Empty{})

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

func (h *saldoHandleApi) handleGetSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetSaldo(c.Request().Context(), &pb.SaldoRequest{
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

func (h *saldoHandleApi) handleGetSaldoByUsers(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetSaldoByUsers(c.Request().Context(), &pb.SaldoRequest{
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

func (h *saldoHandleApi) handleGetSaldobyUserId(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetSaldoByUserId(c.Request().Context(), &pb.SaldoRequest{
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

func (h *saldoHandleApi) handleCreateSaldo(c echo.Context) error {
	var body requests.CreateSaldoRequest

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

	data := &pb.CreateSaldoRequest{
		UserId:       int32(body.UserID),
		TotalBalance: int32(body.TotalBalance),
	}

	res, err := h.client.CreateSaldo(c.Request().Context(), data)

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

func (h *saldoHandleApi) handleUpdateSaldo(c echo.Context) error {
	var body requests.UpdateSaldoRequest

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

	data := &pb.UpdateSaldoRequest{
		SaldoId:        int32(body.SaldoID),
		UserId:         int32(body.UserID),
		TotalBalance:   int32(body.TotalBalance),
		WithdrawAmount: int32(body.WithdrawAmount),
		WithdrawTime:   timestamppb.New(body.WithdrawTime),
	}

	res, err := h.client.UpdateSaldo(c.Request().Context(), data)

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

func (h *saldoHandleApi) handleDeleteSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.DeleteSaldo(c.Request().Context(), &pb.SaldoRequest{
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
