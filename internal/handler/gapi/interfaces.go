package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthHandleGrpc interface {
	pb.AuthServiceServer
	LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.ApiResponseLogin, error)
	RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.ApiResponseRegister, error)
}

type UserHandleGrpc interface {
	pb.UserServiceServer

	FindAll(ctx context.Context, request *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUser, error)
	FindById(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error)
	FindByActive(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesUser, error)
	FindByTrashed(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesUser, error)
	Create(ctx context.Context, request *pb.CreateUserRequest) (*pb.ApiResponseUser, error)
	Update(ctx context.Context, request *pb.UpdateUserRequest) (*pb.ApiResponseUser, error)
	TrashedUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error)
	RestoreUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error)
	DeleteUserPermanent(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUserDelete, error)
}

type CardHandleGrpc interface {
	pb.CardServiceServer

	FindAllCard(ctx context.Context, req *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCard, error)
	FindByIdCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error)
	FindByUserIdCard(ctx context.Context, req *pb.FindByUserIdCardRequest) (*pb.ApiResponseCard, error)
	FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponseCard, error)
	FindByActiveCard(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCards, error)
	FindByTrashedCard(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCards, error)
	CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.ApiResponseCard, error)
	UpdateCard(ctx context.Context, req *pb.UpdateCardRequest) (*pb.ApiResponseCard, error)
	TrashedCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error)
	RestoreCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error)
	DeleteCardPermanent(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCardDelete, error)
}

type MerchantHandleGrpc interface {
	pb.MerchantServiceServer

	FindAll(ctx context.Context, req *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchant, error)
	FindById(ctx context.Context, req *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error)
	FindByApiKey(ctx context.Context, req *pb.FindByApiKeyRequest) (*pb.ApiResponseMerchant, error)
	FindByMerchantUserId(ctx context.Context, req *pb.FindByMerchantUserIdRequest) (*pb.ApiResponsesMerchant, error)
	FindByActive(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesMerchant, error)
	FindByTrashed(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesMerchant, error)
	CreateMerchant(ctx context.Context, req *pb.CreateMerchantRequest) (*pb.ApiResponseMerchant, error)
	UpdateMerchant(ctx context.Context, req *pb.UpdateMerchantRequest) (*pb.ApiResponseMerchant, error)
	TrashedMerchant(ctx context.Context, req *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error)
	RestoreMerchant(ctx context.Context, req *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error)
	DeleteMerchant(ctx context.Context, req *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchatDelete, error)
}

type SaldoHandleGrpc interface {
	pb.SaldoServiceServer

	FindAllSaldo(ctx context.Context, req *pb.FindAllSaldoRequest) (*pb.ApiResponsePaginationSaldo, error)
	FindByIdSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error)
	FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponseSaldo, error)
	FindByActive(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesSaldo, error)
	FindByTrashed(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesSaldo, error)
	CreateSaldo(ctx context.Context, req *pb.CreateSaldoRequest) (*pb.ApiResponseSaldo, error)
	UpdateSaldo(ctx context.Context, req *pb.UpdateSaldoRequest) (*pb.ApiResponseSaldo, error)
	TrashedSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error)
	RestoreSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error)
	DeleteSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldoDelete, error)
}

type TopupHandleGrpc interface {
	pb.TopupServiceServer

	FindAllTopups(ctx context.Context, req *pb.FindAllTopupRequest) (*pb.ApiResponsePaginationTopup, error)
	FindByIdTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopup, error)
	FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponsesTopup, error)
	FindByActive(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesTopup, error)
	FindByTrashed(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesTopup, error)
	CreateTopup(ctx context.Context, req *pb.CreateTopupRequest) (*pb.ApiResponseTopup, error)
	UpdateTopup(ctx context.Context, req *pb.UpdateTopupRequest) (*pb.ApiResponseTopup, error)
	TrashedTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopup, error)
	RestoreTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopup, error)
	DeleteTopupPermanent(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopupDelete, error)
}

type TransactionHandleGrpc interface {
	pb.TransactionServiceServer

	FindAllTransactions(ctx context.Context, req *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransaction, error)
	FindTransactionById(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.TransactionResponse, error)
	FindByCardNumberTransaction(ctx context.Context, request *pb.FindByCardNumberTransactionRequest) (*pb.ApiResponseTransactions, error)
	FindTransactionByMerchantIdRequest(ctx context.Context, request *pb.FindTransactionByMerchantIdRequest) (*pb.ApiResponseTransactions, error)
	FindByActiveTransaction(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactions, error)
	FindByTrashedTransaction(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactions, error)
	CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.ApiResponseTransaction, error)
	UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.ApiResponseTransaction, error)
	TrashedTransaction(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransaction, error)
	RestoreTransaction(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransaction, error)
	DeleteTransaction(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDelete, error)
}

type TransferHandleGrpc interface {
	pb.TransferServiceServer

	FindAllTransfer(ctx context.Context, req *pb.FindAllTransferRequest) (*pb.ApiResponsePaginationTransfer, error)
	FindByIdTransfer(ctx context.Context, req *pb.FindByIdTransferRequest) (*pb.ApiResponseTransfer, error)
	FindByTransferByTransferFrom(ctx context.Context, request *pb.FindTransferByTransferFromRequest) (*pb.ApiResponseTransfers, error)
	FindByTransferByTransferTo(ctx context.Context, request *pb.FindTransferByTransferToRequest) (*pb.ApiResponseTransfers, error)
	FindByActiveTransfer(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransfers, error)
	FindByTrashedTransfer(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransfers, error)
	CreateTransfer(ctx context.Context, req *pb.CreateTransferRequest) (*pb.ApiResponseTransfer, error)
	UpdateTransfer(ctx context.Context, req *pb.UpdateTransferRequest) (*pb.ApiResponseTransfer, error)
	TrashedTransfer(ctx context.Context, req *pb.FindByIdTransferRequest) (*pb.ApiResponseTransfer, error)
	RestoreTransfer(ctx context.Context, req *pb.FindByIdTransferRequest) (*pb.ApiResponseTransfer, error)
	DeleteTransferPermanent(ctx context.Context, req *pb.FindByIdTransferRequest) (*pb.ApiResponseTransferDelete, error)
}

type WithdrawHandleGrpc interface {
	pb.WithdrawServiceServer

	FindAllWithdraw(ctx context.Context, req *pb.FindAllWithdrawRequest) (*pb.ApiResponsePaginationWithdraw, error)
	FindByIdWithdraw(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdraw, error)
	FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponsesWithdraw, error)
	FindByActive(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesWithdraw, error)
	FindByTrashed(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponsesWithdraw, error)
	CreateWithdraw(ctx context.Context, req *pb.CreateWithdrawRequest) (*pb.ApiResponseWithdraw, error)
	UpdateWithdraw(ctx context.Context, req *pb.UpdateWithdrawRequest) (*pb.ApiResponseWithdraw, error)
	TrashedWithdraw(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdraw, error)
	RestoreWithdraw(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdraw, error)
	DeleteWithdrawPermanent(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdrawDelete, error)
}
