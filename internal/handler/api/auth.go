package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"net/http"

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

// handleHello menangani permintaan GET "/hello" dan mengembalikan pesan "Hello".
// @Summary Mengembalikan pesan "Hello"
// @Tags Auth
// @Description Mengembalikan pesan "Hello"
// @Produce json
// @Success 200 {string} string "Hello"
// @Router /auth/hello [get]
func (h *authHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

// register menangani permintaan POST "/register" untuk mendaftarkan pengguna baru.
// @Summary Mendaftarkan pengguna baru
// @Tags Auth
// @Description Mendaftarkan pengguna baru dengan data yang diberikan.
// @Accept json
// @Produce json
// @Param request body requests.CreateUserRequest true "Data pengguna yang ingin didaftarkan"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /auth/register [post]
func (h *authHandleApi) register(c echo.Context) error {
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
			Message:    "Validation Error: " + err.Error(),
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
		return c.JSON(http.StatusInternalServerError, response.ResponseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}

// login menangani permintaan POST "/login" untuk melakukan login pengguna.
// @Summary Melakukan login pengguna
// @Tags Auth
// @Description Melakukan login pengguna dengan data yang diberikan.
// @Accept json
// @Produce json
// @Param request body requests.AuthLoginRequest true "Data login pengguna"
// @Success 200 {object} response.ResponseMessage "Success"
// @Failure 400 {object} response.ResponseMessage "Bad Request"
// @Failure 500 {object} response.ResponseMessage "Internal Server Error"
// @Router /api/login [post]
func (h *authHandleApi) login(c echo.Context) error {
	var body requests.AuthLoginRequest

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
			Message:    "Validation Error: " + err.Error(),
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
		return c.JSON(http.StatusInternalServerError, response.ResponseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error: " + err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, response.ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	})
}
