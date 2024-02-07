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

type withdrawHandleApi struct {
	client pb.WithdrawServiceClient
}

func NewHandlerWithdraw(client pb.WithdrawServiceClient, router *echo.Echo) *withdrawHandleApi {
	withdrawHandler := &withdrawHandleApi{
		client: client,
	}
	routerWithdraw := router.Group("/api/withdraw")

	routerWithdraw.GET("/hello", withdrawHandler.handleHello)
	routerWithdraw.GET("/", withdrawHandler.handleGetWithdraws)
	routerWithdraw.GET("/user-all/:id", withdrawHandler.handleGetWithdrawByUsers)
	routerWithdraw.GET("/user/:id", withdrawHandler.handleGetWithdrawByUserId)
	routerWithdraw.POST("/create", withdrawHandler.handleCreateWithdraw)
	routerWithdraw.PUT("/update/:id", withdrawHandler.handleUpdateWithdraw)
	routerWithdraw.DELETE("/:id", withdrawHandler.handleDeleteWithdraw)

	return withdrawHandler
}

func (h *withdrawHandleApi) handleHello(c echo.Context) error {
	return c.JSON(200, "Hello World")
}

func (h *withdrawHandleApi) handleGetWithdraws(c echo.Context) error {
	res, err := h.client.GetWithdraws(c.Request().Context(), &emptypb.Empty{})

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

func (h *withdrawHandleApi) handleGetWithdrawByUsers(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetWithdrawByUsers(c.Request().Context(), &pb.WithdrawRequest{
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

func (h *withdrawHandleApi) handleGetWithdrawByUserId(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetWithdrawByUserId(c.Request().Context(), &pb.WithdrawRequest{
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

func (h *withdrawHandleApi) handleCreateWithdraw(c echo.Context) error {
	var body requests.CreateWithdrawRequest

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

	data := &pb.CreateWithdrawRequest{
		UserId:         int32(body.UserID),
		WithdrawAmount: int32(body.WithdrawAmount),
		WithdrawTime:   timestamppb.New(body.WithdrawTime),
	}

	res, err := h.client.CreateWithdraw(c.Request().Context(), data)

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

func (h *withdrawHandleApi) handleUpdateWithdraw(c echo.Context) error {
	var body requests.UpdateWithdrawRequest

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

	data := &pb.UpdateWithdrawRequest{
		WithdrawId:     int32(body.WithdrawID),
		UserId:         int32(body.UserID),
		WithdrawAmount: int32(body.WithdrawAmount),
		WithdrawTime:   timestamppb.New(body.WithdrawTime),
	}

	res, err := h.client.UpdateWithdraw(c.Request().Context(), data)

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

func (h *withdrawHandleApi) handleDeleteWithdraw(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.DeleteWithdraw(c.Request().Context(), &pb.WithdrawRequest{
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
