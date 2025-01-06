// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pb/saldo_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=internal/pb/saldo_grpc.pb.go -destination=internal/pb/mocks/saldo_grpc_mock.go
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

// MockSaldoServiceClient is a mock of SaldoServiceClient interface.
type MockSaldoServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockSaldoServiceClientMockRecorder
	isgomock struct{}
}

// MockSaldoServiceClientMockRecorder is the mock recorder for MockSaldoServiceClient.
type MockSaldoServiceClientMockRecorder struct {
	mock *MockSaldoServiceClient
}

// NewMockSaldoServiceClient creates a new mock instance.
func NewMockSaldoServiceClient(ctrl *gomock.Controller) *MockSaldoServiceClient {
	mock := &MockSaldoServiceClient{ctrl: ctrl}
	mock.recorder = &MockSaldoServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaldoServiceClient) EXPECT() *MockSaldoServiceClientMockRecorder {
	return m.recorder
}

// CreateSaldo mocks base method.
func (m *MockSaldoServiceClient) CreateSaldo(ctx context.Context, in *pb.CreateSaldoRequest, opts ...grpc.CallOption) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSaldo", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSaldo indicates an expected call of CreateSaldo.
func (mr *MockSaldoServiceClientMockRecorder) CreateSaldo(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSaldo", reflect.TypeOf((*MockSaldoServiceClient)(nil).CreateSaldo), varargs...)
}

// DeleteSaldoPermanent mocks base method.
func (m *MockSaldoServiceClient) DeleteSaldoPermanent(ctx context.Context, in *pb.FindByIdSaldoRequest, opts ...grpc.CallOption) (*pb.ApiResponseSaldoDelete, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSaldoPermanent", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseSaldoDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSaldoPermanent indicates an expected call of DeleteSaldoPermanent.
func (mr *MockSaldoServiceClientMockRecorder) DeleteSaldoPermanent(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSaldoPermanent", reflect.TypeOf((*MockSaldoServiceClient)(nil).DeleteSaldoPermanent), varargs...)
}

// FindAllSaldo mocks base method.
func (m *MockSaldoServiceClient) FindAllSaldo(ctx context.Context, in *pb.FindAllSaldoRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationSaldo, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAllSaldo", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllSaldo indicates an expected call of FindAllSaldo.
func (mr *MockSaldoServiceClientMockRecorder) FindAllSaldo(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllSaldo", reflect.TypeOf((*MockSaldoServiceClient)(nil).FindAllSaldo), varargs...)
}

// FindByActive mocks base method.
func (m *MockSaldoServiceClient) FindByActive(ctx context.Context, in *pb.FindAllSaldoRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationSaldoDeleteAt, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByActive", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationSaldoDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActive indicates an expected call of FindByActive.
func (mr *MockSaldoServiceClientMockRecorder) FindByActive(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActive", reflect.TypeOf((*MockSaldoServiceClient)(nil).FindByActive), varargs...)
}

// FindByCardNumber mocks base method.
func (m *MockSaldoServiceClient) FindByCardNumber(ctx context.Context, in *pb.FindByCardNumberRequest, opts ...grpc.CallOption) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByCardNumber", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCardNumber indicates an expected call of FindByCardNumber.
func (mr *MockSaldoServiceClientMockRecorder) FindByCardNumber(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCardNumber", reflect.TypeOf((*MockSaldoServiceClient)(nil).FindByCardNumber), varargs...)
}

// FindByIdSaldo mocks base method.
func (m *MockSaldoServiceClient) FindByIdSaldo(ctx context.Context, in *pb.FindByIdSaldoRequest, opts ...grpc.CallOption) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByIdSaldo", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdSaldo indicates an expected call of FindByIdSaldo.
func (mr *MockSaldoServiceClientMockRecorder) FindByIdSaldo(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdSaldo", reflect.TypeOf((*MockSaldoServiceClient)(nil).FindByIdSaldo), varargs...)
}

// FindByTrashed mocks base method.
func (m *MockSaldoServiceClient) FindByTrashed(ctx context.Context, in *pb.FindAllSaldoRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationSaldoDeleteAt, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByTrashed", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationSaldoDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTrashed indicates an expected call of FindByTrashed.
func (mr *MockSaldoServiceClientMockRecorder) FindByTrashed(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTrashed", reflect.TypeOf((*MockSaldoServiceClient)(nil).FindByTrashed), varargs...)
}

// RestoreSaldo mocks base method.
func (m *MockSaldoServiceClient) RestoreSaldo(ctx context.Context, in *pb.FindByIdSaldoRequest, opts ...grpc.CallOption) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RestoreSaldo", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreSaldo indicates an expected call of RestoreSaldo.
func (mr *MockSaldoServiceClientMockRecorder) RestoreSaldo(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreSaldo", reflect.TypeOf((*MockSaldoServiceClient)(nil).RestoreSaldo), varargs...)
}

// TrashedSaldo mocks base method.
func (m *MockSaldoServiceClient) TrashedSaldo(ctx context.Context, in *pb.FindByIdSaldoRequest, opts ...grpc.CallOption) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TrashedSaldo", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrashedSaldo indicates an expected call of TrashedSaldo.
func (mr *MockSaldoServiceClientMockRecorder) TrashedSaldo(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrashedSaldo", reflect.TypeOf((*MockSaldoServiceClient)(nil).TrashedSaldo), varargs...)
}

// UpdateSaldo mocks base method.
func (m *MockSaldoServiceClient) UpdateSaldo(ctx context.Context, in *pb.UpdateSaldoRequest, opts ...grpc.CallOption) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSaldo", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSaldo indicates an expected call of UpdateSaldo.
func (mr *MockSaldoServiceClientMockRecorder) UpdateSaldo(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSaldo", reflect.TypeOf((*MockSaldoServiceClient)(nil).UpdateSaldo), varargs...)
}

// MockSaldoServiceServer is a mock of SaldoServiceServer interface.
type MockSaldoServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockSaldoServiceServerMockRecorder
	isgomock struct{}
}

// MockSaldoServiceServerMockRecorder is the mock recorder for MockSaldoServiceServer.
type MockSaldoServiceServerMockRecorder struct {
	mock *MockSaldoServiceServer
}

// NewMockSaldoServiceServer creates a new mock instance.
func NewMockSaldoServiceServer(ctrl *gomock.Controller) *MockSaldoServiceServer {
	mock := &MockSaldoServiceServer{ctrl: ctrl}
	mock.recorder = &MockSaldoServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaldoServiceServer) EXPECT() *MockSaldoServiceServerMockRecorder {
	return m.recorder
}

// CreateSaldo mocks base method.
func (m *MockSaldoServiceServer) CreateSaldo(arg0 context.Context, arg1 *pb.CreateSaldoRequest) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSaldo", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSaldo indicates an expected call of CreateSaldo.
func (mr *MockSaldoServiceServerMockRecorder) CreateSaldo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSaldo", reflect.TypeOf((*MockSaldoServiceServer)(nil).CreateSaldo), arg0, arg1)
}

// DeleteSaldoPermanent mocks base method.
func (m *MockSaldoServiceServer) DeleteSaldoPermanent(arg0 context.Context, arg1 *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldoDelete, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSaldoPermanent", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseSaldoDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSaldoPermanent indicates an expected call of DeleteSaldoPermanent.
func (mr *MockSaldoServiceServerMockRecorder) DeleteSaldoPermanent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSaldoPermanent", reflect.TypeOf((*MockSaldoServiceServer)(nil).DeleteSaldoPermanent), arg0, arg1)
}

// FindAllSaldo mocks base method.
func (m *MockSaldoServiceServer) FindAllSaldo(arg0 context.Context, arg1 *pb.FindAllSaldoRequest) (*pb.ApiResponsePaginationSaldo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllSaldo", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllSaldo indicates an expected call of FindAllSaldo.
func (mr *MockSaldoServiceServerMockRecorder) FindAllSaldo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllSaldo", reflect.TypeOf((*MockSaldoServiceServer)(nil).FindAllSaldo), arg0, arg1)
}

// FindByActive mocks base method.
func (m *MockSaldoServiceServer) FindByActive(arg0 context.Context, arg1 *pb.FindAllSaldoRequest) (*pb.ApiResponsePaginationSaldoDeleteAt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByActive", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationSaldoDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActive indicates an expected call of FindByActive.
func (mr *MockSaldoServiceServerMockRecorder) FindByActive(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActive", reflect.TypeOf((*MockSaldoServiceServer)(nil).FindByActive), arg0, arg1)
}

// FindByCardNumber mocks base method.
func (m *MockSaldoServiceServer) FindByCardNumber(arg0 context.Context, arg1 *pb.FindByCardNumberRequest) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCardNumber", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCardNumber indicates an expected call of FindByCardNumber.
func (mr *MockSaldoServiceServerMockRecorder) FindByCardNumber(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCardNumber", reflect.TypeOf((*MockSaldoServiceServer)(nil).FindByCardNumber), arg0, arg1)
}

// FindByIdSaldo mocks base method.
func (m *MockSaldoServiceServer) FindByIdSaldo(arg0 context.Context, arg1 *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdSaldo", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdSaldo indicates an expected call of FindByIdSaldo.
func (mr *MockSaldoServiceServerMockRecorder) FindByIdSaldo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdSaldo", reflect.TypeOf((*MockSaldoServiceServer)(nil).FindByIdSaldo), arg0, arg1)
}

// FindByTrashed mocks base method.
func (m *MockSaldoServiceServer) FindByTrashed(arg0 context.Context, arg1 *pb.FindAllSaldoRequest) (*pb.ApiResponsePaginationSaldoDeleteAt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTrashed", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationSaldoDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTrashed indicates an expected call of FindByTrashed.
func (mr *MockSaldoServiceServerMockRecorder) FindByTrashed(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTrashed", reflect.TypeOf((*MockSaldoServiceServer)(nil).FindByTrashed), arg0, arg1)
}

// RestoreSaldo mocks base method.
func (m *MockSaldoServiceServer) RestoreSaldo(arg0 context.Context, arg1 *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreSaldo", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreSaldo indicates an expected call of RestoreSaldo.
func (mr *MockSaldoServiceServerMockRecorder) RestoreSaldo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreSaldo", reflect.TypeOf((*MockSaldoServiceServer)(nil).RestoreSaldo), arg0, arg1)
}

// TrashedSaldo mocks base method.
func (m *MockSaldoServiceServer) TrashedSaldo(arg0 context.Context, arg1 *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrashedSaldo", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrashedSaldo indicates an expected call of TrashedSaldo.
func (mr *MockSaldoServiceServerMockRecorder) TrashedSaldo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrashedSaldo", reflect.TypeOf((*MockSaldoServiceServer)(nil).TrashedSaldo), arg0, arg1)
}

// UpdateSaldo mocks base method.
func (m *MockSaldoServiceServer) UpdateSaldo(arg0 context.Context, arg1 *pb.UpdateSaldoRequest) (*pb.ApiResponseSaldo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSaldo", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseSaldo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSaldo indicates an expected call of UpdateSaldo.
func (mr *MockSaldoServiceServerMockRecorder) UpdateSaldo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSaldo", reflect.TypeOf((*MockSaldoServiceServer)(nil).UpdateSaldo), arg0, arg1)
}

// mustEmbedUnimplementedSaldoServiceServer mocks base method.
func (m *MockSaldoServiceServer) mustEmbedUnimplementedSaldoServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSaldoServiceServer")
}

// mustEmbedUnimplementedSaldoServiceServer indicates an expected call of mustEmbedUnimplementedSaldoServiceServer.
func (mr *MockSaldoServiceServerMockRecorder) mustEmbedUnimplementedSaldoServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSaldoServiceServer", reflect.TypeOf((*MockSaldoServiceServer)(nil).mustEmbedUnimplementedSaldoServiceServer))
}

// MockUnsafeSaldoServiceServer is a mock of UnsafeSaldoServiceServer interface.
type MockUnsafeSaldoServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSaldoServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeSaldoServiceServerMockRecorder is the mock recorder for MockUnsafeSaldoServiceServer.
type MockUnsafeSaldoServiceServerMockRecorder struct {
	mock *MockUnsafeSaldoServiceServer
}

// NewMockUnsafeSaldoServiceServer creates a new mock instance.
func NewMockUnsafeSaldoServiceServer(ctrl *gomock.Controller) *MockUnsafeSaldoServiceServer {
	mock := &MockUnsafeSaldoServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSaldoServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSaldoServiceServer) EXPECT() *MockUnsafeSaldoServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSaldoServiceServer mocks base method.
func (m *MockUnsafeSaldoServiceServer) mustEmbedUnimplementedSaldoServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSaldoServiceServer")
}

// mustEmbedUnimplementedSaldoServiceServer indicates an expected call of mustEmbedUnimplementedSaldoServiceServer.
func (mr *MockUnsafeSaldoServiceServerMockRecorder) mustEmbedUnimplementedSaldoServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSaldoServiceServer", reflect.TypeOf((*MockUnsafeSaldoServiceServer)(nil).mustEmbedUnimplementedSaldoServiceServer))
}
