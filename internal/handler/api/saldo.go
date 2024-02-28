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

// handleHello godoc
// @Summary Menampilkan pesan hello
// @Description Menampilkan pesan hello
// @Tags Saldo
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hello"
// @Router /hello [get]
func (h *saldoHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

// handleGetSaldos godoc
// @Summary Mengambil semua saldo
// @Description Mengambil semua saldo
// @Tags Saldo
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /saldo [get]
func (h *saldoHandleApi) handleGetSaldos(c echo.Context) error {
	res, err := h.client.GetSaldos(c.Request().Context(), &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// handleGetSaldo godoc
// @Summary Mengambil saldo berdasarkan ID
// @Description Mengambil saldo berdasarkan ID
// @Tags Saldo
// @Accept  json
// @Produce  json
// @Param id path int true "ID saldo"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request Invalid ID"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /saldo/{id} [get]
func (h *saldoHandleApi) handleGetSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID",
			Data:       nil,
		})
	}

	res, err := h.client.GetSaldo(c.Request().Context(), &pb.SaldoRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// handleGetSaldoByUsers godoc
// @Summary Mengambil saldo berdasarkan pengguna
// @Description Mengambil saldo berdasarkan pengguna
// @Tags Saldo
// @Accept  json
// @Produce  json
// @Param id path int true "ID pengguna"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request Invalid ID"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /saldo/user-all/{id} [get]
func (h *saldoHandleApi) handleGetSaldoByUsers(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID",
			Data:       nil,
		})
	}

	res, err := h.client.GetSaldoByUsers(c.Request().Context(), &pb.SaldoRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// handleGetSaldobyUserId godoc
// @Summary Mengambil saldo berdasarkan ID pengguna
// @Description Mengambil saldo berdasarkan ID pengguna
// @Tags Saldo
// @Accept  json
// @Produce  json
// @Param id path int true "ID pengguna"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request Invalid ID"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /saldo/user/{id} [get]
func (h *saldoHandleApi) handleGetSaldobyUserId(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID",
			Data:       nil,
		})
	}

	res, err := h.client.GetSaldoByUserId(c.Request().Context(), &pb.SaldoRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// handleCreateSaldo godoc
// @Summary Membuat saldo baru
// @Description Membuat saldo baru
// @Tags Saldo
// @Accept  json
// @Produce  json
// @Param request body requests.CreateSaldoRequest true "Data saldo baru"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request Validate"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /saldo/create [post]
func (h *saldoHandleApi) handleCreateSaldo(c echo.Context) error {
	var body requests.CreateSaldoRequest

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

	data := &pb.CreateSaldoRequest{
		UserId:       int32(body.UserID),
		TotalBalance: int32(body.TotalBalance),
	}

	res, err := h.client.CreateSaldo(c.Request().Context(), data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// handleUpdateSaldo godoc
// @Summary Memperbarui saldo
// @Description Memperbarui saldo
// @Tags Saldo
// @Accept  json
// @Produce  json
// @Param id path int true "ID saldo"
// @Param request body requests.UpdateSaldoRequest true "Data perubahan saldo"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request Validate"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /saldo/update/{id} [put]
func (h *saldoHandleApi) handleUpdateSaldo(c echo.Context) error {
	var body requests.UpdateSaldoRequest

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

	data := &pb.UpdateSaldoRequest{
		SaldoId:        int32(body.SaldoID),
		UserId:         int32(body.UserID),
		TotalBalance:   int32(body.TotalBalance),
		WithdrawAmount: int32(body.WithdrawAmount),
		WithdrawTime:   timestamppb.New(body.WithdrawTime),
	}

	res, err := h.client.UpdateSaldo(c.Request().Context(), data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Delete saldo by ID
// @Description Delete saldo by ID
// @Tags Saldo
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Success 200 {object} response.ResponseMessage	"Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Invalid ID"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /saldo/{id} [delete]
func (h *saldoHandleApi) handleDeleteSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID",
			Data:       nil,
		})
	}

	res, err := h.client.DeleteSaldo(c.Request().Context(), &pb.SaldoRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}
