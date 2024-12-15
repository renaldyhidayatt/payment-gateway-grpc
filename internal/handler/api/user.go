package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandleApi struct {
	client pb.UserServiceClient
}

func NewHandlerUser(client pb.UserServiceClient, router *echo.Echo) *userHandleApi {
	userHandler := &userHandleApi{
		client: client,
	}
	routerUser := router.Group("/api/user")

	routerUser.GET("/", userHandler.FindAllUser)
	routerUser.GET("/:id", userHandler.FindById)

	routerUser.POST("/create", userHandler.Create)
	routerUser.POST("/update/:id", userHandler.Update)
	
	routerUser.POST("/trashed/:id", userHandler.TrashedUser)
	routerUser.POST("/restore/:id", userHandler.RestoreUser)
	routerUser.DELETE("/:id", userHandler.DeleteUserPermanent)

	return userHandler
}

func (h *userHandleApi) FindAllUser(c echo.Context) error {
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

	req := &pb.FindAllUserRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAll(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve user data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *userHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid user ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdUserRequest{
		Id: int32(id),
	}

	user, err := h.client.FindById(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve user data: ",
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandleApi) Create(c echo.Context) error {
	var body requests.CreateUserRequest

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

	req := &pb.CreateUserRequest{
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	res, err := h.client.Create(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create user: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *userHandleApi) Update(c echo.Context) error {
	var body requests.UpdateUserRequest

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

	req := &pb.UpdateUserRequest{
		Id:              int32(body.UserID),
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	res, err := h.client.Update(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update user: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *userHandleApi) TrashedUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid user ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdUserRequest{
		Id: int32(id),
	}

	user, err := h.client.TrashedUser(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed user: ",
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandleApi) RestoreUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid user ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdUserRequest{
		Id: int32(id),
	}

	user, err := h.client.RestoreUser(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore user: ",
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandleApi) DeleteUserPermanent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid user ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdUserRequest{
		Id: int32(id),
	}

	user, err := h.client.DeleteUserPermanent(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete user: ",
		})
	}

	return c.JSON(http.StatusOK, user)
}
