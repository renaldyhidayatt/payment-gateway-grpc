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

type merchantHandleApi struct {
	merchant pb.MerchantServiceClient
}

func NewMerchantHandleApi(merchant pb.MerchantServiceClient, router *echo.Echo) *merchantHandleApi {
	merchantHandler := &merchantHandleApi{
		merchant: merchant,
	}

	routerMerchant := router.Group("/api/merchant")

	routerMerchant.GET("/", merchantHandler.FindAll)
	routerMerchant.GET("/:id", merchantHandler.FindById)
	routerMerchant.GET("/api-key", merchantHandler.FindByApiKey)
	routerMerchant.GET("/merchant-user/:id", merchantHandler.FindByMerchantUserId)
	routerMerchant.GET("/active", merchantHandler.FindByActive)

	routerMerchant.GET("/trashed", merchantHandler.FindByTrashed)

	routerMerchant.POST("/create", merchantHandler.Create)
	routerMerchant.POST("/update/:id", merchantHandler.Update)
	routerMerchant.POST("/trashed/:id", merchantHandler.TrashedMerchant)
	routerMerchant.POST("/restore/:id", merchantHandler.RestoreMerchant)
	routerMerchant.DELETE("/permanent/:id", merchantHandler.Delete)

	return merchantHandler
}

// @Summary Find all merchants
// @Tags Merchant
// @Description Retrieve a list of all merchants
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} pb.ApiResponsePaginationMerchant "List of merchants"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant [get]
func (s *merchantHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllMerchantRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := s.merchant.FindAllMerchant(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	return c.JSON(http.StatusOK, res)

}

// @Summary Find a merchant by ID
// @Tags Merchant
// @Description Retrieve a merchant by its ID
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} pb.ApiResponseMerchant "Merchant data"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/{id} [get]
func (s *merchantHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdMerchantRequest{
		MerchantId: int32(id),
	}

	merchant, err := s.merchant.FindByIdMerchant(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	return c.JSON(http.StatusOK, merchant)
}

// @Summary Find a merchant by API key
// @Tags Merchant
// @Description Retrieve a merchant by its API key
// @Accept json
// @Produce json
// @Param api_key query string true "API key"
// @Success 200 {object} pb.ApiResponseMerchant "Merchant data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/api-key [get]
func (s *merchantHandleApi) FindByApiKey(c echo.Context) error {
	apiKey := c.QueryParam("api_key")

	ctx := c.Request().Context()

	req := &pb.FindByApiKeyRequest{
		ApiKey: apiKey,
	}

	merchant, err := s.merchant.FindByApiKey(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	return c.JSON(http.StatusOK, merchant)
}

// @Summary Find a merchant by user ID
// @Tags Merchant
// @Description Retrieve a merchant by its user ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} pb.ApiResponsesMerchant "Merchant data"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/user/{id} [get]
func (s *merchantHandleApi) FindByMerchantUserId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByMerchantUserIdRequest{
		UserId: int32(id),
	}

	merchant, err := s.merchant.FindByMerchantUserId(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	return c.JSON(http.StatusOK, merchant)
}

// @Summary Find active merchants
// @Tags Merchant
// @Description Retrieve a list of active merchants
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsesMerchant "List of active merchants"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/active [get]
func (s *merchantHandleApi) FindByActive(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := s.merchant.FindByActive(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Find trashed merchants
// @Tags Merchant
// @Description Retrieve a list of trashed merchants
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponsesMerchant "List of trashed merchants"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant data"
// @Router /api/merchant/trashed [get]
func (s *merchantHandleApi) FindByTrashed(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := s.merchant.FindByTrashed(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Create a new merchant
// @Tags Merchant
// @Description Create a new merchant with the given name and user ID
// @Accept json
// @Produce json
// @Param body body requests.CreateMerchantRequest true "Create merchant request"
// @Success 200 {object} pb.ApiResponseMerchant "Created merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create merchant"
// @Router /api/merchant [post]
func (s *merchantHandleApi) Create(c echo.Context) error {
	var body requests.CreateMerchantRequest

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

	req := &pb.CreateMerchantRequest{
		Name:   body.Name,
		UserId: int32(body.UserID),
	}

	res, err := s.merchant.CreateMerchant(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create merchant:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Update a merchant
// @Tags Merchant
// @Description Update a merchant with the given ID
// @Accept json
// @Produce json
// @Param body body requests.UpdateMerchantRequest true "Update merchant request"
// @Success 200 {object} pb.ApiResponseMerchant "Updated merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update merchant"
// @Router /api/merchant/{id} [post]
func (s *merchantHandleApi) Update(c echo.Context) error {
	var body requests.UpdateMerchantRequest

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

	req := &pb.UpdateMerchantRequest{
		MerchantId: int32(body.MerchantID),
		Name:       body.Name,
		UserId:     int32(body.UserID),
		Status:     body.Status,
	}

	res, err := s.merchant.UpdateMerchant(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update merchant:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Trashed a merchant
// @Tags Merchant
// @Description Trashed a merchant by its ID
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} pb.ApiResponseMerchant "Trashed merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed merchant"
// @Router /api/merchant/trashed/{id} [post]
func (s *merchantHandleApi) TrashedMerchant(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := s.merchant.TrashedMerchant(ctx, &pb.FindByIdMerchantRequest{
		MerchantId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed merchant:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Restore a merchant
// @Tags Merchant
// @Description Restore a merchant by its ID
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} pb.ApiResponseMerchant "Restored merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore merchant"
// @Router /api/merchant/restore/{id} [post]
func (s *merchantHandleApi) RestoreMerchant(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := s.merchant.RestoreMerchant(ctx, &pb.FindByIdMerchantRequest{
		MerchantId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore merchant:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Delete a merchant permanently
// @Tags Merchant
// @Description Delete a merchant by its ID permanently
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} pb.ApiResponseMerchatDelete "Deleted merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete merchant"
// @Router /api/merchant/{id} [delete]
func (s *merchantHandleApi) Delete(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := s.merchant.DeleteMerchantPermanent(ctx, &pb.FindByIdMerchantRequest{
		MerchantId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete merchant:",
		})
	}

	return c.JSON(http.StatusOK, res)
}
