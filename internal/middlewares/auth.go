package middlewares

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var whiteListPaths = []string{
	"/api/auth/login",
	"/api/auth/register",
	"/api/auth/hello",
}

func WebSecurityConfig(e *echo.Echo) {
	config := echojwt.Config{
		SigningKey: []byte(viper.GetString("SECRET_KEY")),
		Skipper:    skipAuth,
	}
	e.Use(echojwt.WithConfig(config))
}

func skipAuth(e echo.Context) bool {
	path := e.Path()
	for _, p := range whiteListPaths {
		if path == p {
			return true
		}
	}
	return false
}
