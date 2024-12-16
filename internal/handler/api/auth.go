package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type authHandleApi struct {
	client pb.AuthServiceClient
	logger *logger.Logger
}

func NewHandlerAuth(client pb.AuthServiceClient, router *echo.Echo, logger *logger.Logger) *authHandleApi {
	authHandler := &authHandleApi{
		client: client,
		logger: logger,
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
// @Success 200 {object} pb.ApiResponseRegister "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /auth/register [post]
func (h *authHandleApi) register(c echo.Context) error {
	var body requests.CreateUserRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Bad Request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: ",
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: ",
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
		h.logger.Debug("Internal Server Error", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Internal Server Error: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

// login menangani permintaan POST "/login" untuk melakukan login pengguna.
// @Summary Melakukan login pengguna
// @Tags Auth
// @Description Melakukan login pengguna dengan data yang diberikan.
// @Accept json
// @Produce json
// @Param request body requests.AuthRequest true "Data login pengguna"
// @Success 200 {object} pb.ApiResponseLogin "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/login [post]
func (h *authHandleApi) login(c echo.Context) error {
	var body requests.AuthRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: ",
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: ",
		})
	}

	data := &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	}

	ctx := c.Request().Context()

	res, err := h.client.LoginUser(ctx, data)

	if err != nil {
		h.logger.Debug("Failed to login user", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Internal Server Error: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}
