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

	routerWithdraw.GET("/", withdrawHandler.FindAll)
	routerWithdraw.GET("/:id", withdrawHandler.FindById)
	routerWithdraw.GET("/card_number/:card_number", withdrawHandler.FindByCardNumber)
	routerWithdraw.GET("/active", withdrawHandler.FindByActive)
	routerWithdraw.GET("/trashed", withdrawHandler.FindByTrashed)
	routerWithdraw.POST("/create", withdrawHandler.Create)
	routerWithdraw.POST("/update/:id", withdrawHandler.Update)
	routerWithdraw.POST("/trash/:id", withdrawHandler.TrashWithdraw)
	routerWithdraw.POST("/restore/:id", withdrawHandler.RestoreWithdraw)
	routerWithdraw.DELETE("/:id", withdrawHandler.DeleteWithdrawPermanent)

	return withdrawHandler
}

func (h *withdrawHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllWithdrawRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAllWithdraw(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve withdraw data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *withdrawHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid withdraw ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdWithdrawRequest{
		WithdrawId: int32(id),
	}

	withdraw, err := h.client.FindByIdWithdraw(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve withdraw data: ",
		})
	}

	return c.JSON(http.StatusOK, withdraw)
}

func (h *withdrawHandleApi) FindByCardNumber(c echo.Context) error {
	cardNumber := c.QueryParam("card_number")

	ctx := c.Request().Context()

	req := &pb.FindByCardNumberRequest{
		CardNumber: cardNumber,
	}

	withdraw, err := h.client.FindByCardNumber(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve withdraw data: ",
		})
	}

	return c.JSON(http.StatusOK, withdraw)
}

func (h *withdrawHandleApi) FindByActive(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.FindByActive(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve withdraw data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *withdrawHandleApi) FindByTrashed(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.FindByTrashed(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve withdraw data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *withdrawHandleApi) Create(c echo.Context) error {
	var body requests.CreateWithdrawRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.CreateWithdraw(ctx, &pb.CreateWithdrawRequest{
		CardNumber:     body.CardNumber,
		WithdrawAmount: int32(body.WithdrawAmount),
		WithdrawTime:   timestamppb.New(body.WithdrawTime),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create withdraw: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *withdrawHandleApi) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid withdraw ID",
		})
	}

	var body requests.UpdateWithdrawRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.UpdateWithdraw(ctx, &pb.UpdateWithdrawRequest{
		WithdrawId:     int32(id),
		CardNumber:     body.CardNumber,
		WithdrawAmount: int32(body.WithdrawAmount),
		WithdrawTime:   timestamppb.New(body.WithdrawTime),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update withdraw: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *withdrawHandleApi) TrashWithdraw(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid withdraw ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.TrashedWithdraw(ctx, &pb.FindByIdWithdrawRequest{
		WithdrawId: int32(id),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash withdraw: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *withdrawHandleApi) RestoreWithdraw(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid withdraw ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.RestoreWithdraw(ctx, &pb.FindByIdWithdrawRequest{
		WithdrawId: int32(id),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore withdraw: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *withdrawHandleApi) DeleteWithdrawPermanent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid withdraw ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.DeleteWithdrawPermanent(ctx, &pb.FindByIdWithdrawRequest{
		WithdrawId: int32(id),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete withdraw permanently: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
