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

type saldoHandleApi struct {
	saldo pb.SaldoServiceClient
}

func NewHandlerSaldo(client pb.SaldoServiceClient, router *echo.Echo) *saldoHandleApi {
	saldoHandler := &saldoHandleApi{
		saldo: client,
	}
	routerSaldo := router.Group("/api/saldo")

	routerSaldo.GET("", saldoHandler.FindAll)
	routerSaldo.GET("/:id", saldoHandler.FindById)
	routerSaldo.GET("/active", saldoHandler.FindByActive)
	routerSaldo.GET("/trashed", saldoHandler.FindByTrashed)
	routerSaldo.GET("/card_number/:card_number", saldoHandler.FindByCardNumber)

	routerSaldo.POST("/create", saldoHandler.Create)
	routerSaldo.POST("/update/:id", saldoHandler.Update)
	routerSaldo.POST("/trashed/:id", saldoHandler.TrashSaldo)
	routerSaldo.POST("/restore/:id", saldoHandler.RestoreSaldo)
	routerSaldo.DELETE("/:id", saldoHandler.Delete)

	return saldoHandler

}

func (h *saldoHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllSaldoRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.saldo.FindAllSaldo(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *saldoHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid saldo ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdSaldoRequest{
		SaldoId: int32(id),
	}

	saldo, err := h.saldo.FindByIdSaldo(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	return c.JSON(http.StatusOK, saldo)
}

func (h *saldoHandleApi) FindByCardNumber(c echo.Context) error {
	cardNumber := c.Param("card_number")

	ctx := c.Request().Context()

	req := &pb.FindByCardNumberRequest{
		CardNumber: cardNumber,
	}

	saldo, err := h.saldo.FindByCardNumber(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	return c.JSON(http.StatusOK, saldo)
}

func (h *saldoHandleApi) FindByActive(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.saldo.FindByActive(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *saldoHandleApi) FindByTrashed(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.saldo.FindByTrashed(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *saldoHandleApi) Create(c echo.Context) error {
	var body requests.CreateSaldoRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.CreateSaldo(ctx, &pb.CreateSaldoRequest{
		CardNumber:   body.CardNumber,
		TotalBalance: int32(body.TotalBalance),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create saldo: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *saldoHandleApi) Update(c echo.Context) error {
	var body requests.UpdateSaldoRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.UpdateSaldo(ctx, &pb.UpdateSaldoRequest{
		SaldoId:      int32(body.SaldoID),
		CardNumber:   body.CardNumber,
		TotalBalance: int32(body.TotalBalance),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *saldoHandleApi) TrashSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.TrashSaldo(ctx, &pb.FindByIdSaldoRequest{
		SaldoId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed saldo:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *saldoHandleApi) RestoreSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.RestoreSaldo(ctx, &pb.FindByIdSaldoRequest{
		SaldoId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore saldo:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *saldoHandleApi) Delete(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.DeleteSaldoPermanent(ctx, &pb.FindByIdSaldoRequest{
		SaldoId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete saldo:",
		})
	}

	return c.JSON(http.StatusOK, res)
}
