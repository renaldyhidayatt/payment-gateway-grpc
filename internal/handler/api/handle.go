package api

import (
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/pkg/auth"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"github.com/labstack/echo/v4"

	"google.golang.org/grpc"
)

type Deps struct {
	Conn   *grpc.ClientConn
	Token  auth.TokenManager
	E      *echo.Echo
	Logger logger.LoggerInterface
}

func NewHandler(deps Deps) {

	clientAuth := pb.NewAuthServiceClient(deps.Conn)
	clientRole := pb.NewRoleServiceClient(deps.Conn)
	clientCard := pb.NewCardServiceClient(deps.Conn)
	clientMerchant := pb.NewMerchantServiceClient(deps.Conn)
	clientUser := pb.NewUserServiceClient(deps.Conn)
	clientSaldo := pb.NewSaldoServiceClient(deps.Conn)
	clientTopup := pb.NewTopupServiceClient(deps.Conn)
	clientTransaction := pb.NewTransactionServiceClient(deps.Conn)
	clientTransfer := pb.NewTransferServiceClient(deps.Conn)
	clientWithdraw := pb.NewWithdrawServiceClient(deps.Conn)

	NewHandlerAuth(clientAuth, deps.E, deps.Logger)
	NewHandlerRole(clientRole, deps.E, deps.Logger)
	NewHandlerUser(clientUser, deps.E, deps.Logger)
	NewHandlerCard(clientCard, deps.E, deps.Logger)
	NewHandlerMerchant(clientMerchant, deps.E, deps.Logger)
	NewHandlerTransaction(clientTransaction, clientMerchant, deps.E, deps.Logger)
	NewHandlerSaldo(clientSaldo, deps.E, deps.Logger)
	NewHandlerTopup(clientTopup, deps.E, deps.Logger)
	NewHandlerTransfer(clientTransfer, deps.E, deps.Logger)
	NewHandlerWithdraw(clientWithdraw, deps.E, deps.Logger)
}
