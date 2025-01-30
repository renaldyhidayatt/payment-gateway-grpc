package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	apimapper "MamangRust/paymentgatewaygrpc/internal/mapper/response/api"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type saldoHandleApi struct {
	saldo   pb.SaldoServiceClient
	logger  logger.LoggerInterface
	mapping apimapper.SaldoResponseMapper
}

func NewHandlerSaldo(client pb.SaldoServiceClient, router *echo.Echo, logger logger.LoggerInterface, mapping apimapper.SaldoResponseMapper) *saldoHandleApi {
	saldoHandler := &saldoHandleApi{
		saldo:   client,
		logger:  logger,
		mapping: mapping,
	}
	routerSaldo := router.Group("/api/saldos")

	routerSaldo.GET("", saldoHandler.FindAll)
	routerSaldo.GET("/:id", saldoHandler.FindById)

	routerSaldo.GET("/monthly-total-balance", saldoHandler.FindMonthlyTotalSaldoBalance)
	routerSaldo.GET("/yearly-total-balance", saldoHandler.FindYearTotalSaldoBalance)
	routerSaldo.GET("/monthly-balances", saldoHandler.FindMonthlySaldoBalances)
	routerSaldo.GET("/yearly-balances", saldoHandler.FindYearlySaldoBalances)

	routerSaldo.GET("/active", saldoHandler.FindByActive)
	routerSaldo.GET("/trashed", saldoHandler.FindByTrashed)
	routerSaldo.GET("/card_number/:card_number", saldoHandler.FindByCardNumber)

	routerSaldo.POST("/create", saldoHandler.Create)
	routerSaldo.POST("/update/:id", saldoHandler.Update)
	routerSaldo.POST("/trashed/:id", saldoHandler.TrashSaldo)
	routerSaldo.POST("/restore/:id", saldoHandler.RestoreSaldo)
	routerSaldo.DELETE("/permanent/:id", saldoHandler.Delete)

	routerSaldo.POST("/restore/all", saldoHandler.RestoreAllSaldo)
	routerSaldo.POST("/permanent/all", saldoHandler.DeleteAllSaldoPermanent)

	return saldoHandler

}

// @Summary Find all saldo data
// @Tags Saldo
// @Security Bearer
// @Description Retrieve a list of all saldo data with pagination and search
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} pb.ApiResponsePaginationSaldo "List of saldo data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldos [get]
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
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	so := h.mapping.ToApiResponsePaginationSaldo(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Find a saldo by ID
// @Tags Saldo
// @Security Bearer
// @Description Retrieve a saldo by its ID
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Success 200 {object} pb.ApiResponseSaldo "Saldo data"
// @Failure 400 {object} response.ErrorResponse "Invalid saldo ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldos/{id} [get]
func (h *saldoHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid saldo ID", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid saldo ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdSaldoRequest{
		SaldoId: int32(id),
	}

	res, err := h.saldo.FindByIdSaldo(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	so := h.mapping.ToApiResponseSaldo(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTotalSaldoBalance retrieves the total saldo balance for a specific month and year.
// @Summary Get monthly total saldo balance
// @Tags Saldo
// @Security Bearer
// @Description Retrieve the total saldo balance for a specific month and year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} pb.ApiResponseMonthTotalSaldo "Monthly total saldo balance"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly total saldo balance"
// @Router /api/saldos/monthly-total-balance [get]
func (h *saldoHandleApi) FindMonthlyTotalSaldoBalance(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		h.logger.Debug("Invalid month parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid month parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.FindMonthlyTotalSaldoBalance(ctx, &pb.FindMonthlySaldoTotalBalance{
		Year:  int32(year),
		Month: int32(month),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly total saldo balance", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve monthly total saldo balance",
		})
	}

	so := h.mapping.ToApiResponseMonthTotalSaldo(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearTotalSaldoBalance retrieves the total saldo balance for a specific year.
// @Summary Get yearly total saldo balance
// @Tags Saldo
// @Security Bearer
// @Description Retrieve the total saldo balance for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseYearTotalSaldo "Yearly total saldo balance"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly total saldo balance"
// @Router /api/saldos/yearly-total-balance [get]
func (h *saldoHandleApi) FindYearTotalSaldoBalance(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.FindYearTotalSaldoBalance(ctx, &pb.FindYearlySaldo{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve year total saldo balance", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve year total saldo balance",
		})
	}

	so := h.mapping.ToApiResponseYearTotalSaldo(res)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlySaldoBalances retrieves monthly saldo balances for a specific year.
// @Summary Get monthly saldo balances
// @Tags Saldo
// @Security Bearer
// @Description Retrieve monthly saldo balances for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseMonthSaldoBalances "Monthly saldo balances"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly saldo balances"
// @Router /api/saldos/monthly-balances [get]
func (h *saldoHandleApi) FindMonthlySaldoBalances(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.FindMonthlySaldoBalances(ctx, &pb.FindYearlySaldo{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly saldo balances", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve monthly saldo balances",
		})
	}

	so := h.mapping.ToApiResponseMonthSaldoBalances(res)

	return c.JSON(http.StatusOK, so)
}

// FindYearlySaldoBalances retrieves yearly saldo balances for a specific year.
// @Summary Get yearly saldo balances
// @Tags Saldo
// @Security Bearer
// @Description Retrieve yearly saldo balances for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} pb.ApiResponseYearSaldoBalances "Yearly saldo balances"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly saldo balances"
// @Router /api/saldo/yearly-balances [get]
func (h *saldoHandleApi) FindYearlySaldoBalances(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid year parameter",
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.FindYearlySaldoBalances(ctx, &pb.FindYearlySaldo{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly saldo balances", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve yearly saldo balances",
		})
	}

	so := h.mapping.ToApiResponseYearSaldoBalances(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Find a saldo by card number
// @Tags Saldo
// @Security Bearer
// @Description Retrieve a saldo by its card number
// @Accept json
// @Produce json
// @Param card_number path string true "Card number"
// @Success 200 {object} pb.ApiResponseSaldo "Saldo data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldos/card_number/{card_number} [get]
func (h *saldoHandleApi) FindByCardNumber(c echo.Context) error {
	cardNumber := c.Param("card_number")

	ctx := c.Request().Context()

	req := &pb.FindByCardNumberRequest{
		CardNumber: cardNumber,
	}

	res, err := h.saldo.FindByCardNumber(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	so := h.mapping.ToApiResponseSaldo(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Retrieve all active saldo data
// @Tags Saldo
// @Security Bearer
// @Description Retrieve a list of all active saldo data
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsesSaldo "List of saldo data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldos/active [get]
func (h *saldoHandleApi) FindByActive(c echo.Context) error {
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

	res, err := h.saldo.FindByActive(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	so := h.mapping.ToApiResponsePaginationSaldoDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Retrieve trashed saldo data
// @Tags Saldo
// @Security Bearer
// @Description Retrieve a list of all trashed saldo data
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsesSaldo "List of trashed saldo data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldos/trashed [get]
func (h *saldoHandleApi) FindByTrashed(c echo.Context) error {
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

	res, err := h.saldo.FindByTrashed(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo data: ",
		})
	}

	so := h.mapping.ToApiResponsePaginationSaldoDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Create a new saldo
// @Tags Saldo
// @Security Bearer
// @Description Create a new saldo record with the provided card number and total balance.
// @Accept json
// @Produce json
// @Param CreateSaldoRequest body requests.CreateSaldoRequest true "Create Saldo Request"
// @Success 200 {object} pb.ApiResponseSaldo "Successfully created saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create saldo"
// @Router /api/saldos/create [post]
func (h *saldoHandleApi) Create(c echo.Context) error {
	var body requests.CreateSaldoRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
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
		h.logger.Debug("Failed to create saldo", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create saldo: ",
		})
	}

	so := h.mapping.ToApiResponseSaldo(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Update an existing saldo
// @Tags Saldo
// @Security Bearer
// @Description Update an existing saldo record with the provided card number and total balance.
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Param UpdateSaldoRequest body requests.UpdateSaldoRequest true "Update Saldo Request"
// @Success 200 {object} pb.ApiResponseSaldo "Successfully updated saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update saldo"
// @Router /api/saldos/update/{id} [post]
func (h *saldoHandleApi) Update(c echo.Context) error {
	idint, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))

		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	var body requests.UpdateSaldoRequest

	body.SaldoID = idint

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: ",
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.UpdateSaldo(ctx, &pb.UpdateSaldoRequest{
		SaldoId:      int32(body.SaldoID),
		CardNumber:   body.CardNumber,
		TotalBalance: int32(body.TotalBalance),
	})

	if err != nil {
		h.logger.Debug("Failed to update saldo", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo: ",
		})
	}

	so := h.mapping.ToApiResponseSaldo(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Soft delete a saldo
// @Tags Saldo
// @Security Bearer
// @Description Soft delete an existing saldo record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Success 200 {object} pb.ApiResponseSaldo "Successfully trashed saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed saldo"
// @Router /api/saldos/trashed/{id} [post]
func (h *saldoHandleApi) TrashSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.saldo.TrashedSaldo(ctx, &pb.FindByIdSaldoRequest{
		SaldoId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to trashed saldo", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed saldo:",
		})
	}

	so := h.mapping.ToApiResponseSaldo(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Restore a trashed saldo
// @Tags Saldo
// @Security Bearer
// @Description Restore an existing saldo record from the trash by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Success 200 {object} pb.ApiResponseSaldo "Successfully restored saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore saldo"
// @Router /api/saldos/restore/{id} [post]
func (h *saldoHandleApi) RestoreSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))
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
		h.logger.Debug("Failed to restore saldo", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore saldo:",
		})
	}

	so := h.mapping.ToApiResponseSaldo(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Permanently delete a saldo
// @Tags Saldo
// @Security Bearer
// @Description Permanently delete an existing saldo record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Success 200 {object} pb.ApiResponseSaldoDelete "Successfully deleted saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete saldo"
// @Router /api/saldos/permanent/{id} [delete]
func (h *saldoHandleApi) Delete(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))
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
		h.logger.Debug("Failed to delete saldo", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete saldo:",
		})
	}

	so := h.mapping.ToApiResponseSaldoDelete(res)

	return c.JSON(http.StatusOK, so)
}

// RestoreAllSaldo restores all saldo records.
// @Summary Restore all saldo records
// @Tags Saldo
// @Security Bearer
// @Description Restore all saldo records that were previously deleted.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseSaldoAll "Successfully restored all saldo records"
// @Failure 500 {object} response.ErrorResponse "Failed to restore all saldo records"
// @Router /api/saldos/restore/all [post]
func (h *saldoHandleApi) RestoreAllSaldo(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.saldo.RestoreAllSaldo(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to restore all saldo", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently restore all saldo",
		})
	}

	h.logger.Debug("Successfully restored all saldo")

	so := h.mapping.ToApiResponseSaldoAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Permanently delete all saldo records
// @Tags Saldo
// @Security Bearer
// @Description Permanently delete all saldo records from the database.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseSaldoAll "Successfully deleted all saldo records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to permanently delete all saldo records"
// @Router /api/saldos/permanent/all [post]
func (h *saldoHandleApi) DeleteAllSaldoPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.saldo.DeleteAllSaldoPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to permanently delete all saldo", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all saldo",
		})
	}

	h.logger.Debug("Successfully deleted all merchant permanently")

	so := h.mapping.ToApiResponseSaldoAll(res)

	return c.JSON(http.StatusOK, so)
}
