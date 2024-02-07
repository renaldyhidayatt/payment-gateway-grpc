package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"

	"github.com/labstack/echo/v4"
)

type authHandleApi struct {
	client pb.AuthServiceClient
}

func NewHandlerAuth(client pb.AuthServiceClient, router *echo.Echo) *authHandleApi {
	authHandler := &authHandleApi{
		client: client,
	}
	routerAuth := router.Group("/api/auth")

	routerAuth.GET("/hello", authHandler.handleHello)
	routerAuth.POST("/register", authHandler.register)
	routerAuth.POST("/login", authHandler.login)

	return authHandler
}

func (h *authHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

func (h *authHandleApi) register(c echo.Context) error {
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

	data := &pb.RegisterRequest{
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	ctx := c.Request().Context()

	res, err := h.client.RegisterUser(ctx, data)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}

func (h *authHandleApi) login(c echo.Context) error {
	var body requests.AuthLoginRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Bad Request",
			Data:       nil,
		})
	}

	data := &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	}

	ctx := c.Request().Context()

	res, err := h.client.LoginUser(ctx, data)

	if err != nil {
		return c.JSON(400, response.ResponseMessage{
			StatusCode: 400,
			Message:    "Error",
			Data:       nil,
		})
	}

	return c.JSON(200, response.ResponseMessage{
		StatusCode: 200,
		Message:    "Success",
		Data:       res,
	})
}
