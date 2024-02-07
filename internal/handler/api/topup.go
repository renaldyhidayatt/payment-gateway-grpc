package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type topupHandleApi struct {
	client pb.TopupServiceClient
}

func NewHandlerTopup(client pb.TopupServiceClient, router *echo.Echo) *topupHandleApi {
	topupHandler := &topupHandleApi{
		client: client,
	}
	routerTopup := router.Group("/api/topup")

	routerTopup.GET("/hello", topupHandler.handleHello)
	routerTopup.GET("/", topupHandler.handleGetTopups)
	routerTopup.GET("/:id", topupHandler.handleGetTopup)
	routerTopup.GET("/user-all/:id", topupHandler.handleGetTopupByUsers)
	routerTopup.GET("/user/:id", topupHandler.GetTopupByUserId)
	routerTopup.POST("/create", topupHandler.handleCreateTopup)
	routerTopup.PUT("/update/:id", topupHandler.handleUpdateTopup)
	routerTopup.DELETE("/:id", topupHandler.handleDeleteTopup)

	return topupHandler

}

func (h *topupHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

func (h *topupHandleApi) handleGetTopups(c echo.Context) error {
	res, err := h.client.GetTopups(c.Request().Context(), &emptypb.Empty{})

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

func (h *topupHandleApi) handleGetTopup(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetTopup(c.Request().Context(), &pb.TopupRequest{
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

func (h *topupHandleApi) handleGetTopupByUsers(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetTopupByUsers(c.Request().Context(), &pb.TopupRequest{
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

func (h *topupHandleApi) GetTopupByUserId(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetTopupByUserId(c.Request().Context(), &pb.TopupRequest{
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

func (h *topupHandleApi) handleCreateTopup(c echo.Context) error {
	var body requests.CreateTopupRequest

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

	data := &pb.CreateTopupRequest{
		UserId:      int32(body.UserID),
		TopupNo:     body.TopupNo,
		TopupAmount: int32(body.TopupAmount),
		TopupMethod: body.TopupMethod,
	}

	res, err := h.client.CreateTopup(c.Request().Context(), data)

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

func (h *topupHandleApi) handleUpdateTopup(c echo.Context) error {
	var body requests.UpdateTopupRequest

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

	data := &pb.UpdateTopupRequest{
		UserId:      int32(body.UserID),
		TopupId:     int32(body.TopupID),
		TopupAmount: int32(body.TopupAmount),
		TopupMethod: body.TopupMethod,
	}

	res, err := h.client.UpdateTopup(c.Request().Context(), data)

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

func (h *topupHandleApi) handleDeleteTopup(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.DeleteTopup(c.Request().Context(), &pb.TopupRequest{
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
