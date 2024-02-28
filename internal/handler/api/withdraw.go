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
	routerWithdraw.GET("/:id", withdrawHandler.handleGetWithdraw)
	routerWithdraw.GET("/user-all/:id", withdrawHandler.handleGetWithdrawByUsers)
	routerWithdraw.GET("/user/:id", withdrawHandler.handleGetWithdrawByUserId)
	routerWithdraw.POST("/create", withdrawHandler.handleCreateWithdraw)
	routerWithdraw.PUT("/update/:id", withdrawHandler.handleUpdateWithdraw)
	routerWithdraw.DELETE("/:id", withdrawHandler.handleDeleteWithdraw)

	return withdrawHandler
}

// handleHello godoc
// @Summary Menampilkan pesan hello
// @Description Menampilkan pesan hello
// @Tags Saldo
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hello"
// @Router /withdraw/hello [get]
func (h *withdrawHandleApi) handleHello(c echo.Context) error {
	return c.JSON(200, "Hello World")
}

// @Summary Get all withdraws
// @Description Get all withdraws
// @Tags Withdraw
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /withdraw/ [get]
func (h *withdrawHandleApi) handleGetWithdraws(c echo.Context) error {
	res, err := h.client.GetWithdraws(c.Request().Context(), &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to retrieve withdraws: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get a withdraw by ID
// @Description Get a withdraw by ID
// @Tags Withdraw
// @Accept json
// @Produce json
// @Param id path int true "Withdraw ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /withdraw/{id} [get]
func (h *withdrawHandleApi) handleGetWithdraw(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.GetWithdraw(c.Request().Context(), &pb.WithdrawRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to retrieve withdraw: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get all withdraws by user
// @Description Get all withdraws by user
// @Tags Withdraw
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /withdraw/user-all/{id} [get]
func (h *withdrawHandleApi) handleGetWithdrawByUsers(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.GetWithdrawByUsers(c.Request().Context(), &pb.WithdrawRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to retrieve withdraw by user: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get all withdraws by user ID
// @Description Get all withdraws by user ID
// @Tags Withdraw
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /withdraw/user/{id} [get]
func (h *withdrawHandleApi) handleGetWithdrawByUserId(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.GetWithdrawByUserId(c.Request().Context(), &pb.WithdrawRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to retrieve withdraw by user ID: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Create a withdraw
// @Description Create a new withdraw
// @Tags Withdraw
// @Accept json
// @Produce json
// @Param body body requests.CreateWithdrawRequest true "Withdraw data"
// @Success 200 {object} response.ResponseMessage "Withdraw created successfully"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /withdraw/create [post]
func (h *withdrawHandleApi) handleCreateWithdraw(c echo.Context) error {
	var body requests.CreateWithdrawRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create withdraw: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Validation failed: " + err.Error(),
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
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create withdraw: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Withdraw created successfully",
		Data:       res,
	})
}

// @Summary Update a withdraw
// @Description Update an existing withdraw
// @Tags Withdraw
// @Accept json
// @Produce json
// @Param body body requests.UpdateWithdrawRequest true "Withdraw data"
// @Success 200 {object} response.ResponseMessage "Withdraw updated successfully"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /withdraw/update/{id} [put]
func (h *withdrawHandleApi) handleUpdateWithdraw(c echo.Context) error {
	var body requests.UpdateWithdrawRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update withdraw: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Validation failed: " + err.Error(),
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
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update withdraw: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Withdraw updated successfully",
		Data:       res,
	})
}

// @Summary Delete a withdraw
// @Description Delete a withdraw by ID
// @Tags Withdraw
// @Accept json
// @Produce json
// @Param id path int true "Withdraw ID"
// @Success 200 {object} response.ResponseMessage "Withdraw deleted successfully"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /withdraw/{id} [delete]
func (h *withdrawHandleApi) handleDeleteWithdraw(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.DeleteWithdraw(c.Request().Context(), &pb.WithdrawRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to delete withdraw: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Withdraw deleted successfully",
		Data:       res,
	})
}
