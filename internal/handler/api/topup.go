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

	routerTopup.GET("", topupHandler.FindAll)
	routerTopup.GET("/:id", topupHandler.FindById)
	routerTopup.GET("/active", topupHandler.FindByActive)
	routerTopup.GET("/trashed", topupHandler.FindByTrashed)
	routerTopup.GET("/card_number/:card_number", topupHandler.FindByCardNumber)

	routerTopup.POST("/create", topupHandler.Create)
	routerTopup.POST("/update/:id", topupHandler.Update)
	routerTopup.POST("/trashed/:id", topupHandler.TrashTopup)
	routerTopup.POST("/restore/:id", topupHandler.RestoreTopup)
	routerTopup.DELETE("/:id", topupHandler.DeleteTopupPermanent)

	return topupHandler

}

func (h topupHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllTopupRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAllTopup(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve topup data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h topupHandleApi) FindById(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.FindByIdTopup(ctx, &pb.FindByIdTopupRequest{
		TopupId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve topup data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *topupHandleApi) FindByCardNumber(c echo.Context) error {
	cardNumber := c.Param("card_number")

	ctx := c.Request().Context()

	req := &pb.FindByCardNumberTopupRequest{
		CardNumber: cardNumber,
	}

	topup, err := h.client.FindByCardNumberTopup(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve topup data: ",
		})
	}

	return c.JSON(http.StatusOK, topup)
}

func (h *topupHandleApi) FindByActive(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.FindByActive(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve topup data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *topupHandleApi) FindByTrashed(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.FindByTrashed(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve topup data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *topupHandleApi) Create(c echo.Context) error {
	var body requests.CreateTopupRequest

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

	res, err := h.client.CreateTopup(ctx, &pb.CreateTopupRequest{
		CardNumber:  body.CardNumber,
		TopupNo:     body.TopupNo,
		TopupAmount: int32(body.TopupAmount),
		TopupMethod: body.TopupMethod,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create topup: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *topupHandleApi) Update(c echo.Context) error {
	var body requests.UpdateTopupRequest

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

	res, err := h.client.UpdateTopup(ctx, &pb.UpdateTopupRequest{
		TopupId:     int32(body.TopupID),
		CardNumber:  body.CardNumber,
		TopupAmount: int32(body.TopupAmount),
		TopupMethod: body.TopupMethod,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update topup: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *topupHandleApi) TrashTopup(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.TrashedTopup(ctx, &pb.FindByIdTopupRequest{
		TopupId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed topup:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *topupHandleApi) RestoreTopup(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.RestoreTopup(ctx, &pb.FindByIdTopupRequest{
		TopupId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore topup:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *topupHandleApi) DeleteTopupPermanent(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.DeleteTopupPermanent(ctx, &pb.FindByIdTopupRequest{
		TopupId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete topup:",
		})
	}

	return c.JSON(http.StatusOK, res)
}
