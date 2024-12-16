package middlewares

import (
	"MamangRust/paymentgatewaygrpc/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ApiKeyMiddleware(merchantService service.MerchantService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiKey := c.Request().Header.Get("X-Api-Key")
			if apiKey == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "API Key is required")
			}

			_, err := merchantService.FindByApiKey(apiKey)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API Key")
			}

			c.Set("apiKey", apiKey)

			return next(c)
		}
	}
}
