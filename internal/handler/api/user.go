package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
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

func (h *userHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

func (h *userHandleApi) handleGetUsers(c echo.Context) error {
	res, err := h.client.GetUsers(c.Request().Context(), &emptypb.Empty{})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *userHandleApi) handleGetUser(c echo.Context) error {
	id := c.Param("id")

	idErr, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	res, err := h.client.GetUser(c.Request().Context(), &pb.UserRequest{
		Id: int32(idErr),
	})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *userHandleApi) handleCreateUser(c echo.Context) error {
	var body requests.CreateUserRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
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
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *userHandleApi) handleUpdateUser(c echo.Context) error {
	var body requests.UpdateUserRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
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
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})

}

func (h *userHandleApi) handleDeleteUser(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.client.DeleteUser(c.Request().Context(), &pb.UserRequest{
		Id: int32(idInt),
	})

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})

}
