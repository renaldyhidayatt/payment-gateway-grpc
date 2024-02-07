package api

import (
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/pkg/auth"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func NewHandler(conn *grpc.ClientConn, token auth.TokenManager, e *echo.Echo) {

	clientAuth := pb.NewAuthServiceClient(conn)
	clientUser := pb.NewUserServiceClient(conn)
	clientSaldo := pb.NewSaldoServiceClient(conn)
	clientTopup := pb.NewTopupServiceClient(conn)
	clientTransfer := pb.NewTransferServiceClient(conn)
	clientWithdraw := pb.NewWithdrawServiceClient(conn)

	NewHandlerAuth(clientAuth, e)
	NewHandlerUser(clientUser, e)
	NewHandlerSaldo(clientSaldo, e)
	NewHandlerTopup(clientTopup, e)
	NewHandlerTransfer(clientTransfer, e)
	NewHandlerWithdraw(clientWithdraw, e)
}
