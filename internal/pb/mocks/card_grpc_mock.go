// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pb/card_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=internal/pb/card_grpc.pb.go -destination=internal/pb/mocks/card_grpc_mock.go
//

// Package mock_pb is a generated GoMock package.
package mock_pb

import (
	pb "MamangRust/paymentgatewaygrpc/internal/pb"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCardServiceClient is a mock of CardServiceClient interface.
type MockCardServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockCardServiceClientMockRecorder
	isgomock struct{}
}

// MockCardServiceClientMockRecorder is the mock recorder for MockCardServiceClient.
type MockCardServiceClientMockRecorder struct {
	mock *MockCardServiceClient
}

// NewMockCardServiceClient creates a new mock instance.
func NewMockCardServiceClient(ctrl *gomock.Controller) *MockCardServiceClient {
	mock := &MockCardServiceClient{ctrl: ctrl}
	mock.recorder = &MockCardServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCardServiceClient) EXPECT() *MockCardServiceClientMockRecorder {
	return m.recorder
}

// CreateCard mocks base method.
func (m *MockCardServiceClient) CreateCard(ctx context.Context, in *pb.CreateCardRequest, opts ...grpc.CallOption) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCard", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCard indicates an expected call of CreateCard.
func (mr *MockCardServiceClientMockRecorder) CreateCard(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCard", reflect.TypeOf((*MockCardServiceClient)(nil).CreateCard), varargs...)
}

// DeleteCardPermanent mocks base method.
func (m *MockCardServiceClient) DeleteCardPermanent(ctx context.Context, in *pb.FindByIdCardRequest, opts ...grpc.CallOption) (*pb.ApiResponseCardDelete, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCardPermanent", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseCardDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCardPermanent indicates an expected call of DeleteCardPermanent.
func (mr *MockCardServiceClientMockRecorder) DeleteCardPermanent(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCardPermanent", reflect.TypeOf((*MockCardServiceClient)(nil).DeleteCardPermanent), varargs...)
}

// FindAllCard mocks base method.
func (m *MockCardServiceClient) FindAllCard(ctx context.Context, in *pb.FindAllCardRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationCard, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAllCard", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllCard indicates an expected call of FindAllCard.
func (mr *MockCardServiceClientMockRecorder) FindAllCard(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllCard", reflect.TypeOf((*MockCardServiceClient)(nil).FindAllCard), varargs...)
}

// FindByActiveCard mocks base method.
func (m *MockCardServiceClient) FindByActiveCard(ctx context.Context, in *pb.FindAllCardRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationCardDeleteAt, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByActiveCard", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationCardDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActiveCard indicates an expected call of FindByActiveCard.
func (mr *MockCardServiceClientMockRecorder) FindByActiveCard(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActiveCard", reflect.TypeOf((*MockCardServiceClient)(nil).FindByActiveCard), varargs...)
}

// FindByCardNumber mocks base method.
func (m *MockCardServiceClient) FindByCardNumber(ctx context.Context, in *pb.FindByCardNumberRequest, opts ...grpc.CallOption) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByCardNumber", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCardNumber indicates an expected call of FindByCardNumber.
func (mr *MockCardServiceClientMockRecorder) FindByCardNumber(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCardNumber", reflect.TypeOf((*MockCardServiceClient)(nil).FindByCardNumber), varargs...)
}

// FindByIdCard mocks base method.
func (m *MockCardServiceClient) FindByIdCard(ctx context.Context, in *pb.FindByIdCardRequest, opts ...grpc.CallOption) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByIdCard", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdCard indicates an expected call of FindByIdCard.
func (mr *MockCardServiceClientMockRecorder) FindByIdCard(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdCard", reflect.TypeOf((*MockCardServiceClient)(nil).FindByIdCard), varargs...)
}

// FindByTrashedCard mocks base method.
func (m *MockCardServiceClient) FindByTrashedCard(ctx context.Context, in *pb.FindAllCardRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationCardDeleteAt, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByTrashedCard", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationCardDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTrashedCard indicates an expected call of FindByTrashedCard.
func (mr *MockCardServiceClientMockRecorder) FindByTrashedCard(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTrashedCard", reflect.TypeOf((*MockCardServiceClient)(nil).FindByTrashedCard), varargs...)
}

// FindByUserIdCard mocks base method.
func (m *MockCardServiceClient) FindByUserIdCard(ctx context.Context, in *pb.FindByUserIdCardRequest, opts ...grpc.CallOption) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByUserIdCard", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserIdCard indicates an expected call of FindByUserIdCard.
func (mr *MockCardServiceClientMockRecorder) FindByUserIdCard(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserIdCard", reflect.TypeOf((*MockCardServiceClient)(nil).FindByUserIdCard), varargs...)
}

// RestoreCard mocks base method.
func (m *MockCardServiceClient) RestoreCard(ctx context.Context, in *pb.FindByIdCardRequest, opts ...grpc.CallOption) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RestoreCard", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreCard indicates an expected call of RestoreCard.
func (mr *MockCardServiceClientMockRecorder) RestoreCard(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreCard", reflect.TypeOf((*MockCardServiceClient)(nil).RestoreCard), varargs...)
}

// TrashedCard mocks base method.
func (m *MockCardServiceClient) TrashedCard(ctx context.Context, in *pb.FindByIdCardRequest, opts ...grpc.CallOption) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TrashedCard", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrashedCard indicates an expected call of TrashedCard.
func (mr *MockCardServiceClientMockRecorder) TrashedCard(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrashedCard", reflect.TypeOf((*MockCardServiceClient)(nil).TrashedCard), varargs...)
}

// UpdateCard mocks base method.
func (m *MockCardServiceClient) UpdateCard(ctx context.Context, in *pb.UpdateCardRequest, opts ...grpc.CallOption) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCard", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCard indicates an expected call of UpdateCard.
func (mr *MockCardServiceClientMockRecorder) UpdateCard(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCard", reflect.TypeOf((*MockCardServiceClient)(nil).UpdateCard), varargs...)
}

// MockCardServiceServer is a mock of CardServiceServer interface.
type MockCardServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockCardServiceServerMockRecorder
	isgomock struct{}
}

// MockCardServiceServerMockRecorder is the mock recorder for MockCardServiceServer.
type MockCardServiceServerMockRecorder struct {
	mock *MockCardServiceServer
}

// NewMockCardServiceServer creates a new mock instance.
func NewMockCardServiceServer(ctrl *gomock.Controller) *MockCardServiceServer {
	mock := &MockCardServiceServer{ctrl: ctrl}
	mock.recorder = &MockCardServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCardServiceServer) EXPECT() *MockCardServiceServerMockRecorder {
	return m.recorder
}

// CreateCard mocks base method.
func (m *MockCardServiceServer) CreateCard(arg0 context.Context, arg1 *pb.CreateCardRequest) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCard", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCard indicates an expected call of CreateCard.
func (mr *MockCardServiceServerMockRecorder) CreateCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCard", reflect.TypeOf((*MockCardServiceServer)(nil).CreateCard), arg0, arg1)
}

// DeleteCardPermanent mocks base method.
func (m *MockCardServiceServer) DeleteCardPermanent(arg0 context.Context, arg1 *pb.FindByIdCardRequest) (*pb.ApiResponseCardDelete, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCardPermanent", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseCardDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCardPermanent indicates an expected call of DeleteCardPermanent.
func (mr *MockCardServiceServerMockRecorder) DeleteCardPermanent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCardPermanent", reflect.TypeOf((*MockCardServiceServer)(nil).DeleteCardPermanent), arg0, arg1)
}

// FindAllCard mocks base method.
func (m *MockCardServiceServer) FindAllCard(arg0 context.Context, arg1 *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllCard", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllCard indicates an expected call of FindAllCard.
func (mr *MockCardServiceServerMockRecorder) FindAllCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllCard", reflect.TypeOf((*MockCardServiceServer)(nil).FindAllCard), arg0, arg1)
}

// FindByActiveCard mocks base method.
func (m *MockCardServiceServer) FindByActiveCard(arg0 context.Context, arg1 *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCardDeleteAt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByActiveCard", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationCardDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActiveCard indicates an expected call of FindByActiveCard.
func (mr *MockCardServiceServerMockRecorder) FindByActiveCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActiveCard", reflect.TypeOf((*MockCardServiceServer)(nil).FindByActiveCard), arg0, arg1)
}

// FindByCardNumber mocks base method.
func (m *MockCardServiceServer) FindByCardNumber(arg0 context.Context, arg1 *pb.FindByCardNumberRequest) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCardNumber", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCardNumber indicates an expected call of FindByCardNumber.
func (mr *MockCardServiceServerMockRecorder) FindByCardNumber(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCardNumber", reflect.TypeOf((*MockCardServiceServer)(nil).FindByCardNumber), arg0, arg1)
}

// FindByIdCard mocks base method.
func (m *MockCardServiceServer) FindByIdCard(arg0 context.Context, arg1 *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdCard", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdCard indicates an expected call of FindByIdCard.
func (mr *MockCardServiceServerMockRecorder) FindByIdCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdCard", reflect.TypeOf((*MockCardServiceServer)(nil).FindByIdCard), arg0, arg1)
}

// FindByTrashedCard mocks base method.
func (m *MockCardServiceServer) FindByTrashedCard(arg0 context.Context, arg1 *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCardDeleteAt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTrashedCard", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationCardDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTrashedCard indicates an expected call of FindByTrashedCard.
func (mr *MockCardServiceServerMockRecorder) FindByTrashedCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTrashedCard", reflect.TypeOf((*MockCardServiceServer)(nil).FindByTrashedCard), arg0, arg1)
}

// FindByUserIdCard mocks base method.
func (m *MockCardServiceServer) FindByUserIdCard(arg0 context.Context, arg1 *pb.FindByUserIdCardRequest) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserIdCard", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserIdCard indicates an expected call of FindByUserIdCard.
func (mr *MockCardServiceServerMockRecorder) FindByUserIdCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserIdCard", reflect.TypeOf((*MockCardServiceServer)(nil).FindByUserIdCard), arg0, arg1)
}

// RestoreCard mocks base method.
func (m *MockCardServiceServer) RestoreCard(arg0 context.Context, arg1 *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreCard", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreCard indicates an expected call of RestoreCard.
func (mr *MockCardServiceServerMockRecorder) RestoreCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreCard", reflect.TypeOf((*MockCardServiceServer)(nil).RestoreCard), arg0, arg1)
}

// TrashedCard mocks base method.
func (m *MockCardServiceServer) TrashedCard(arg0 context.Context, arg1 *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrashedCard", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrashedCard indicates an expected call of TrashedCard.
func (mr *MockCardServiceServerMockRecorder) TrashedCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrashedCard", reflect.TypeOf((*MockCardServiceServer)(nil).TrashedCard), arg0, arg1)
}

// UpdateCard mocks base method.
func (m *MockCardServiceServer) UpdateCard(arg0 context.Context, arg1 *pb.UpdateCardRequest) (*pb.ApiResponseCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCard", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCard indicates an expected call of UpdateCard.
func (mr *MockCardServiceServerMockRecorder) UpdateCard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCard", reflect.TypeOf((*MockCardServiceServer)(nil).UpdateCard), arg0, arg1)
}

// mustEmbedUnimplementedCardServiceServer mocks base method.
func (m *MockCardServiceServer) mustEmbedUnimplementedCardServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedCardServiceServer")
}

// mustEmbedUnimplementedCardServiceServer indicates an expected call of mustEmbedUnimplementedCardServiceServer.
func (mr *MockCardServiceServerMockRecorder) mustEmbedUnimplementedCardServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedCardServiceServer", reflect.TypeOf((*MockCardServiceServer)(nil).mustEmbedUnimplementedCardServiceServer))
}

// MockUnsafeCardServiceServer is a mock of UnsafeCardServiceServer interface.
type MockUnsafeCardServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeCardServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeCardServiceServerMockRecorder is the mock recorder for MockUnsafeCardServiceServer.
type MockUnsafeCardServiceServerMockRecorder struct {
	mock *MockUnsafeCardServiceServer
}

// NewMockUnsafeCardServiceServer creates a new mock instance.
func NewMockUnsafeCardServiceServer(ctrl *gomock.Controller) *MockUnsafeCardServiceServer {
	mock := &MockUnsafeCardServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeCardServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeCardServiceServer) EXPECT() *MockUnsafeCardServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedCardServiceServer mocks base method.
func (m *MockUnsafeCardServiceServer) mustEmbedUnimplementedCardServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedCardServiceServer")
}

// mustEmbedUnimplementedCardServiceServer indicates an expected call of mustEmbedUnimplementedCardServiceServer.
func (mr *MockUnsafeCardServiceServerMockRecorder) mustEmbedUnimplementedCardServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedCardServiceServer", reflect.TypeOf((*MockUnsafeCardServiceServer)(nil).mustEmbedUnimplementedCardServiceServer))
}
