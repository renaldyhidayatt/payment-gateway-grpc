package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type authHandleApi struct {
	client pb.AuthServiceClient
	logger logger.LoggerInterface
}

func NewHandlerAuth(client pb.AuthServiceClient, router *echo.Echo, logger logger.LoggerInterface) *authHandleApi {
	authHandler := &authHandleApi{
		client: client,
		logger: logger,
	}
	routerAuth := router.Group("/api/auth")

	routerAuth.GET("/hello", authHandler.HandleHello)
	routerAuth.POST("/register", authHandler.Register)
	routerAuth.POST("/login", authHandler.Login)
	routerAuth.POST("/refresh-token", authHandler.RefreshToken)
	routerAuth.GET("/me", authHandler.GetMe)

	return authHandler
}

// HandleHello godoc
// @Summary Returns a "Hello" message
// @Tags Auth
// @Description Returns a simple "Hello" message for testing purposes.
// @Produce json
// @Success 200 {string} string "Hello"
// @Router /auth/hello [get]
func (h *authHandleApi) HandleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

// Register godoc
// @Summary Register a new user
// @Tags Auth
// @Description Registers a new user with the provided details.
// @Accept json
// @Produce json
// @Param request body requests.CreateUserRequest true "User registration data"
// @Success 200 {object} pb.ApiResponseRegister "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/register [post]
func (h *authHandleApi) Register(c echo.Context) error {
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
			Message: "Internal Server Error: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Authenticate a user
// @Tags Auth
// @Description Authenticates a user using the provided email and password.
// @Accept json
// @Produce json
// @Param request body requests.AuthRequest true "User login credentials"
// @Success 200 {object} pb.ApiResponseLogin "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/login [post]
func (h *authHandleApi) Login(c echo.Context) error {
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

// RefreshToken godoc
// @Summary Refresh access token
// @Tags Auth
// @Security Bearer
// @Description Refreshes the access token using a valid refresh token.
// @Accept json
// @Produce json
// @Param request body requests.RefreshTokenRequest true "Refresh token data"
// @Success 200 {object} pb.ApiResponseRefreshToken "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/refresh-token [post]
func (h *authHandleApi) RefreshToken(c echo.Context) error {
	var body requests.RefreshTokenRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	res, err := h.client.RefreshToken(c.Request().Context(), &pb.RefreshTokenRequest{
		RefreshToken: body.RefreshToken,
	})

	if err != nil {
		h.logger.Debug("Failed to refresh token", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Internal Server Error: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// GetMe godoc
// @Summary Get current user information
// @Tags Auth
// @Security Bearer
// @Description Retrieves the current user's information using a valid access token from the Authorization header.
// @Produce json
// @Security BearerToken
// @Success 200 {object} pb.ApiResponseGetMe "Success"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/me [get]
func (h *authHandleApi) GetMe(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")

	h.logger.Debug("Authorization header: ", zap.String("authHeader", authHeader))

	if !strings.HasPrefix(authHeader, "Bearer ") {
		h.logger.Debug("Authorization header is missing or invalid format")
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized: Missing or invalid Authorization header",
		})
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	res, err := h.client.GetMe(c.Request().Context(), &pb.GetMeRequest{
		AccessToken: accessToken,
	})

	if err != nil {
		h.logger.Debug("Failed to get user information", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Internal Server Error: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
