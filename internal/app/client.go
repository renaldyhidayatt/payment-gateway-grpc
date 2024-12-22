package app

import (
	"MamangRust/paymentgatewaygrpc/internal/handler/api"
	"MamangRust/paymentgatewaygrpc/internal/middlewares"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/auth"
	"MamangRust/paymentgatewaygrpc/pkg/dotenv"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	_ "MamangRust/paymentgatewaygrpc/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "server:50051", "the address to connect to")
)

// @title PaymentGateway gRPC
// @version 1.0
// @description gRPC based Payment Gateway service

// @host localhost:5000
// @BasePath /api/

// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization
func RunClient() {
	flag.Parse()

	logger, err := logger.NewLogger()

	if err != nil {
		logger.Fatal("Failed to create logger", zap.Error(err))
	}

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logger.Fatal("Failed to connect to server", zap.Error(err))
	}

	err = dotenv.Viper()

	if err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	middlewares.WebSecurityConfig(e)

	depsRepo := repository.Deps{
		DB:           DB,
		Ctx:          ctx,
		MapperRecord: mapperRecord,
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	token, err := auth.NewManager(viper.GetString("SECRET_KEY"))

	if err != nil {
		logger.Fatal("Failed to create token manager", zap.Error(err))
	}

	depsHandler := api.Deps{
		Conn:    conn,
		Token:   token,
		E:       e,
		Service: service,
		Logger:  logger,
	}

	api.NewHandler(depsHandler)

	go func() {
		if err := e.Start(":5000"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Server.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
