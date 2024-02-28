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

// @Summary Get hello message
// @Description Get hello message
// @Tags Topup
// @Produce plain
// @Success 200 {string} string	"Hello"
// @Router /topup/hello [get]
func (h *topupHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

// @Summary Get list of Topups
// @Description Get list of Topups
// @Tags Topup
// @Produce json
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /topup/ [get]
func (h *topupHandleApi) handleGetTopups(c echo.Context) error {
	res, err := h.client.GetTopups(c.Request().Context(), &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get a Topup by ID
// @Description Get a Topup by ID
// @Tags Topup
// @Produce json
// @Param id path int true "Topup ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Invalid ID"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /topup/{id} [get]
func (h *topupHandleApi) handleGetTopup(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID",
			Data:       nil,
		})
	}

	res, err := h.client.GetTopup(c.Request().Context(), &pb.TopupRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get list of Topups by user ID
// @Description Get list of Topups by user ID
// @Tags Topup
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Invalid ID"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /topup/user-all/{id} [get]
func (h *topupHandleApi) handleGetTopupByUsers(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID",
			Data:       nil,
		})
	}

	res, err := h.client.GetTopupByUsers(c.Request().Context(), &pb.TopupRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get a Topup by user ID
// @Description Get a Topup by user ID
// @Tags Topup
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Invalid ID"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /topup/user/{id} [get]
func (h *topupHandleApi) GetTopupByUserId(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID",
			Data:       nil,
		})
	}

	res, err := h.client.GetTopupByUserId(c.Request().Context(), &pb.TopupRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Create a new topup
// @Description Create a new topup
// @Tags Topup
// @Accept json
// @Produce json
// @Param body body requests.CreateTopupRequest true "Topup data"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /topup/create [post]
func (h *topupHandleApi) handleCreateTopup(c echo.Context) error {
	var body requests.CreateTopupRequest

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

	data := &pb.CreateTopupRequest{
		UserId:      int32(body.UserID),
		TopupNo:     body.TopupNo,
		TopupAmount: int32(body.TopupAmount),
		TopupMethod: body.TopupMethod,
	}

	res, err := h.client.CreateTopup(c.Request().Context(), data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Error creating topup: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Update an existing topup
// @Description Update an existing topup
// @Tags Topup
// @Accept json
// @Produce json
// @Param body body requests.UpdateTopupRequest true "Topup data"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /topup/update/{id} [put]
func (h *topupHandleApi) handleUpdateTopup(c echo.Context) error {
	var body requests.UpdateTopupRequest

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

	data := &pb.UpdateTopupRequest{
		UserId:      int32(body.UserID),
		TopupId:     int32(body.TopupID),
		TopupAmount: int32(body.TopupAmount),
		TopupMethod: body.TopupMethod,
	}

	res, err := h.client.UpdateTopup(c.Request().Context(), data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Error updating topup: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Summary Delete a topup by ID
// @Description Delete a topup by ID
// @Tags Topup
// @Accept json
// @Produce json
// @Param id path int true "Topup ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Invalid ID"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /topup/{id} [delete]
func (h *topupHandleApi) handleDeleteTopup(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID",
			Data:       nil,
		})
	}

	res, err := h.client.DeleteTopup(c.Request().Context(), &pb.TopupRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Error deleting topup: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}
