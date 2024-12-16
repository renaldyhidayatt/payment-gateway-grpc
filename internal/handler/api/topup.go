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
	routerTopup.DELETE("/permanent/:id", topupHandler.DeleteTopupPermanent)

	return topupHandler

}

// @Summary Find all topup data
// @Tags Topup
// @Description Retrieve a list of all topup data with pagination and search
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} pb.ApiResponsePaginationTopup "List of topup data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
// @Router /api/topup [get]
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

// FindById retrieves a topup record by its ID.
// @Summary Find a topup by ID
// @Tags Topup
// @Description Retrieve a topup record using its ID
// @Accept json
// @Produce json
// @Param id path string true "Topup ID"
// @Success 200 {object} pb.ApiResponseTopup "Topup data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
// @Router /api/topup/{id} [get]
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

// FindByCardNumber retrieves a topup record by its card number.
// @Summary Find a topup by its card number
// @Tags Topup
// @Description Retrieve a topup record using its card number
// @Accept json
// @Produce json
// @Param card_number path string true "Card number"
// @Success 200 {object} pb.ApiResponsesTopup "Topup data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
// @Router /api/topup/card_number/{card_number} [get]
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

// FindByActive retrieves a list of active topup records.
// @Summary Find active topups
// @Tags Topup
// @Description Retrieve a list of active topup records
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsesTopup "Active topup data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
// @Router /api/topup/active [get]
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

// FindByTrashed retrieves a list of trashed topup records.
// @Summary Retrieve trashed topups
// @Tags Topup
// @Description Retrieve a list of trashed topup records
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsesTopup "List of trashed topup data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
// @Router /api/topup/trashed [get]
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

// Create creates a new topup record.
// @Summary Create topup
// @Tags Topup
// @Description Create a new topup record
// @Accept json
// @Produce json
// @Param CreateTopupRequest body requests.CreateTopupRequest true "Create topup request"
// @Success 200 {object} pb.ApiResponseTopup "Created topup data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: "
// @Failure 500 {object} response.ErrorResponse "Failed to create topup: "
// @Router /api/topup/create [post]
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

// Update updates an existing topup record.
// @Summary Update topup
// @Tags Topup
// @Description Update an existing topup record with the provided details
// @Accept json
// @Produce json
// @Param id path int true "Topup ID"
// @Param UpdateTopupRequest body requests.UpdateTopupRequest true "Update topup request"
// @Success 200 {object} pb.ApiResponseTopup "Updated topup data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid input data"
// @Failure 500 {object} response.ErrorResponse "Failed to update topup: "
// @Router /api/topup/update/{id} [post]
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

// TrashTopup trashes a topup record by its ID.
//
// @Summary Trash a topup
// @Tags Topup
// @Description Trash a topup record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Topup ID"
// @Success 200 {object} pb.ApiResponseTopup "Successfully trashed topup record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trash topup:"
// @Router /api/topup/trash/{id} [post]
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

// RestoreTopup restores a trashed topup record by its ID.
//
// @Summary Restore a trashed topup
// @Tags Topup
// @Description Restore a trashed topup record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Topup ID"
// @Success 200 {object} pb.ApiResponseTopup "Successfully restored topup record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore topup:"
// @Router /api/topup/restore/{id} [post]
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

// DeleteTopupPermanent deletes a topup record permanently by its ID.
//
// @Summary Permanently delete a topup
// @Tags Topup
// @Description Permanently delete a topup record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Topup ID"
// @Success 200 {object} pb.ApiResponseTopupDelete "Successfully deleted topup record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete topup:"
// @Router /api/topup/permanent/{id} [delete]
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
