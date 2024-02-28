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

type userHandleApi struct {
	client pb.UserServiceClient
}

func NewHandlerUser(client pb.UserServiceClient, router *echo.Echo) *userHandleApi {
	userHandler := &userHandleApi{
		client: client,
	}
	routerUser := router.Group("/api/user")

	routerUser.GET("/hello", userHandler.handleHello)
	routerUser.GET("/", userHandler.handleGetUsers)
	routerUser.GET("/:id", userHandler.handleGetUser)
	routerUser.POST("/create", userHandler.handleCreateUser)
	routerUser.PUT("/update/:id", userHandler.handleUpdateUser)
	routerUser.DELETE("/delete/:id", userHandler.handleDeleteUser)

	return userHandler
}

// handleHello godoc
// @Summary Menampilkan pesan hello
// @Description Menampilkan pesan hello
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hello"
// @Router /user/hello [get]
func (h *userHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /user/ [get]
func (h *userHandleApi) handleGetUsers(c echo.Context) error {
	res, err := h.client.GetUsers(c.Request().Context(), &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to get users: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /user/{id} [get]
func (h *userHandleApi) handleGetUser(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID",
			Data:       nil,
		})
	}

	res, err := h.client.GetUser(c.Request().Context(), &pb.UserRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to get user: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param body body requests.CreateUserRequest true "User details"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /user/create [post]
func (h *userHandleApi) handleCreateUser(c echo.Context) error {
	var body requests.CreateUserRequest

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

	data := &pb.CreateUserRequest{
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	res, err := h.client.CreateUser(c.Request().Context(), data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create user: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Update a user
// @Description Update a user
// @Tags User
// @Accept json
// @Produce json
// @Param body body requests.UpdateUserRequest true "User details"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /user/update/{id} [put]
func (h *userHandleApi) handleUpdateUser(c echo.Context) error {
	var body requests.UpdateUserRequest

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

	data := &pb.UpdateUserRequest{
		Id:              int32(body.ID),
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	res, err := h.client.UpdateUser(c.Request().Context(), data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update user: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request: Error message"
// @Router /user/delete/{id} [delete]
func (h *userHandleApi) handleDeleteUser(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request: Invalid ID: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.DeleteUser(c.Request().Context(), &pb.UserRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to delete user: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}
