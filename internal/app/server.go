package app

import (
	"MamangRust/paymentgatewaygrpc/internal/handler/gapi"
	protomapper "MamangRust/paymentgatewaygrpc/internal/mapper/proto"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"MamangRust/paymentgatewaygrpc/pkg/auth"
	"MamangRust/paymentgatewaygrpc/pkg/database/postgres"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"MamangRust/paymentgatewaygrpc/pkg/dotenv"
	"MamangRust/paymentgatewaygrpc/pkg/hash"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func RunServer() {
	logger, err := logger.NewLogger()

	if err != nil {
		logger.Fatal("Failed to create logger", zap.Error(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	err = dotenv.Viper()

	if err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
	}

	token, err := auth.NewManager(viper.GetString("SECRET_KEY"))

	if err != nil {
		logger.Fatal("Failed to create token manager", zap.Error(err))
	}

	conn, err := postgres.NewClient(logger)

	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	hash := hash.NewHashingPassword()

	DB := db.New(conn)

	ctx := context.Background()

	mapperRecord := recordmapper.NewRecordMapper()
	mapperResponse := responsemapper.NewResponseMapper()
	mapperProto := protomapper.NewProtoMapper()

	depsRepo := repository.Deps{
		DB:           DB,
		Ctx:          ctx,
		MapperRecord: mapperRecord,
	}

	repository := repository.NewRepositories(depsRepo)

	service := service.NewService(service.Deps{
		Repositories: repository,
		Hash:         hash,
		Token:        token,
		Logger:       logger,
		Mapper:       *mapperResponse,
	})

	if err != nil {
		logger.Fatal("Failed to create service", zap.Error(err))
	}

	handlerAuth := gapi.NewAuthHandleGrpc(service.Auth, mapperProto.AuthProtoMapper)
	handlerCard := gapi.NewCardHandleGrpc(service.Card, mapperProto.CardProtoMapper)
	handleMerchant := gapi.NewMerchantHandleGrpc(service.MerchantService, mapperProto.MerchantProtoMapper)
	handlerUser := gapi.NewUserHandleGrpc(service.User, mapperProto.UserProtoMapper)
	handlerSaldo := gapi.NewSaldoHandleGrpc(service.Saldo, mapperProto.SaldoProtoMapper)
	handlerTopup := gapi.NewTopupHandleGrpc(service.Topup, mapperProto.TopupProtoMapper)
	handlerTransaction := gapi.NewTransactionHandleGrpc(service.Transaction, mapperProto.TransactionProtoMapper)
	handlerTransfer := gapi.NewTransferHandleGrpc(service.Transfer, mapperProto.TransferProtoMapper)
	handlerWithraw := gapi.NewWithdrawHandleGrpc(service.Withdraw, mapperProto.WithdrawalProtoMapper)

	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, handlerAuth)
	pb.RegisterUserServiceServer(s, handlerUser)
	pb.RegisterCardServiceServer(s, handlerCard)
	pb.RegisterMerchantServiceServer(s, handleMerchant)
	pb.RegisterSaldoServiceServer(s, handlerSaldo)
	pb.RegisterTopupServiceServer(s, handlerTopup)
	pb.RegisterTransactionServiceServer(s, handlerTransaction)
	pb.RegisterTransferServiceServer(s, handlerTransfer)
	pb.RegisterWithdrawServiceServer(s, handlerWithraw)

	logger.Info(fmt.Sprintf("Server running on port %d", *port))

	if err := s.Serve(lis); err != nil {
		logger.Fatal("Failed to serve", zap.Error(err))
	}

}
