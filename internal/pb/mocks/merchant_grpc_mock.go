// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pb/merchant_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=internal/pb/merchant_grpc.pb.go -destination=internal/pb/mocks/merchant_grpc_mock.go
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

// MockMerchantServiceClient is a mock of MerchantServiceClient interface.
type MockMerchantServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockMerchantServiceClientMockRecorder
	isgomock struct{}
}

// MockMerchantServiceClientMockRecorder is the mock recorder for MockMerchantServiceClient.
type MockMerchantServiceClientMockRecorder struct {
	mock *MockMerchantServiceClient
}

// NewMockMerchantServiceClient creates a new mock instance.
func NewMockMerchantServiceClient(ctrl *gomock.Controller) *MockMerchantServiceClient {
	mock := &MockMerchantServiceClient{ctrl: ctrl}
	mock.recorder = &MockMerchantServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMerchantServiceClient) EXPECT() *MockMerchantServiceClientMockRecorder {
	return m.recorder
}

// CreateMerchant mocks base method.
func (m *MockMerchantServiceClient) CreateMerchant(ctx context.Context, in *pb.CreateMerchantRequest, opts ...grpc.CallOption) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMerchant", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMerchant indicates an expected call of CreateMerchant.
func (mr *MockMerchantServiceClientMockRecorder) CreateMerchant(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMerchant", reflect.TypeOf((*MockMerchantServiceClient)(nil).CreateMerchant), varargs...)
}

// DeleteMerchantPermanent mocks base method.
func (m *MockMerchantServiceClient) DeleteMerchantPermanent(ctx context.Context, in *pb.FindByIdMerchantRequest, opts ...grpc.CallOption) (*pb.ApiResponseMerchatDelete, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMerchantPermanent", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseMerchatDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMerchantPermanent indicates an expected call of DeleteMerchantPermanent.
func (mr *MockMerchantServiceClientMockRecorder) DeleteMerchantPermanent(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMerchantPermanent", reflect.TypeOf((*MockMerchantServiceClient)(nil).DeleteMerchantPermanent), varargs...)
}

// FindAllMerchant mocks base method.
func (m *MockMerchantServiceClient) FindAllMerchant(ctx context.Context, in *pb.FindAllMerchantRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationMerchant, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAllMerchant", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllMerchant indicates an expected call of FindAllMerchant.
func (mr *MockMerchantServiceClientMockRecorder) FindAllMerchant(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllMerchant", reflect.TypeOf((*MockMerchantServiceClient)(nil).FindAllMerchant), varargs...)
}

// FindByActive mocks base method.
func (m *MockMerchantServiceClient) FindByActive(ctx context.Context, in *pb.FindAllMerchantRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByActive", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationMerchantDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActive indicates an expected call of FindByActive.
func (mr *MockMerchantServiceClientMockRecorder) FindByActive(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActive", reflect.TypeOf((*MockMerchantServiceClient)(nil).FindByActive), varargs...)
}

// FindByApiKey mocks base method.
func (m *MockMerchantServiceClient) FindByApiKey(ctx context.Context, in *pb.FindByApiKeyRequest, opts ...grpc.CallOption) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByApiKey", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByApiKey indicates an expected call of FindByApiKey.
func (mr *MockMerchantServiceClientMockRecorder) FindByApiKey(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByApiKey", reflect.TypeOf((*MockMerchantServiceClient)(nil).FindByApiKey), varargs...)
}

// FindByIdMerchant mocks base method.
func (m *MockMerchantServiceClient) FindByIdMerchant(ctx context.Context, in *pb.FindByIdMerchantRequest, opts ...grpc.CallOption) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByIdMerchant", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdMerchant indicates an expected call of FindByIdMerchant.
func (mr *MockMerchantServiceClientMockRecorder) FindByIdMerchant(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdMerchant", reflect.TypeOf((*MockMerchantServiceClient)(nil).FindByIdMerchant), varargs...)
}

// FindByMerchantUserId mocks base method.
func (m *MockMerchantServiceClient) FindByMerchantUserId(ctx context.Context, in *pb.FindByMerchantUserIdRequest, opts ...grpc.CallOption) (*pb.ApiResponsesMerchant, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByMerchantUserId", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsesMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMerchantUserId indicates an expected call of FindByMerchantUserId.
func (mr *MockMerchantServiceClientMockRecorder) FindByMerchantUserId(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMerchantUserId", reflect.TypeOf((*MockMerchantServiceClient)(nil).FindByMerchantUserId), varargs...)
}

// FindByTrashed mocks base method.
func (m *MockMerchantServiceClient) FindByTrashed(ctx context.Context, in *pb.FindAllMerchantRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByTrashed", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationMerchantDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTrashed indicates an expected call of FindByTrashed.
func (mr *MockMerchantServiceClientMockRecorder) FindByTrashed(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTrashed", reflect.TypeOf((*MockMerchantServiceClient)(nil).FindByTrashed), varargs...)
}

// RestoreMerchant mocks base method.
func (m *MockMerchantServiceClient) RestoreMerchant(ctx context.Context, in *pb.FindByIdMerchantRequest, opts ...grpc.CallOption) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RestoreMerchant", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreMerchant indicates an expected call of RestoreMerchant.
func (mr *MockMerchantServiceClientMockRecorder) RestoreMerchant(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreMerchant", reflect.TypeOf((*MockMerchantServiceClient)(nil).RestoreMerchant), varargs...)
}

// TrashedMerchant mocks base method.
func (m *MockMerchantServiceClient) TrashedMerchant(ctx context.Context, in *pb.FindByIdMerchantRequest, opts ...grpc.CallOption) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TrashedMerchant", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrashedMerchant indicates an expected call of TrashedMerchant.
func (mr *MockMerchantServiceClientMockRecorder) TrashedMerchant(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrashedMerchant", reflect.TypeOf((*MockMerchantServiceClient)(nil).TrashedMerchant), varargs...)
}

// UpdateMerchant mocks base method.
func (m *MockMerchantServiceClient) UpdateMerchant(ctx context.Context, in *pb.UpdateMerchantRequest, opts ...grpc.CallOption) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMerchant", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMerchant indicates an expected call of UpdateMerchant.
func (mr *MockMerchantServiceClientMockRecorder) UpdateMerchant(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMerchant", reflect.TypeOf((*MockMerchantServiceClient)(nil).UpdateMerchant), varargs...)
}

// MockMerchantServiceServer is a mock of MerchantServiceServer interface.
type MockMerchantServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockMerchantServiceServerMockRecorder
	isgomock struct{}
}

// MockMerchantServiceServerMockRecorder is the mock recorder for MockMerchantServiceServer.
type MockMerchantServiceServerMockRecorder struct {
	mock *MockMerchantServiceServer
}

// NewMockMerchantServiceServer creates a new mock instance.
func NewMockMerchantServiceServer(ctrl *gomock.Controller) *MockMerchantServiceServer {
	mock := &MockMerchantServiceServer{ctrl: ctrl}
	mock.recorder = &MockMerchantServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMerchantServiceServer) EXPECT() *MockMerchantServiceServerMockRecorder {
	return m.recorder
}

// CreateMerchant mocks base method.
func (m *MockMerchantServiceServer) CreateMerchant(arg0 context.Context, arg1 *pb.CreateMerchantRequest) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMerchant", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMerchant indicates an expected call of CreateMerchant.
func (mr *MockMerchantServiceServerMockRecorder) CreateMerchant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMerchant", reflect.TypeOf((*MockMerchantServiceServer)(nil).CreateMerchant), arg0, arg1)
}

// DeleteMerchantPermanent mocks base method.
func (m *MockMerchantServiceServer) DeleteMerchantPermanent(arg0 context.Context, arg1 *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchatDelete, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMerchantPermanent", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseMerchatDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMerchantPermanent indicates an expected call of DeleteMerchantPermanent.
func (mr *MockMerchantServiceServerMockRecorder) DeleteMerchantPermanent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMerchantPermanent", reflect.TypeOf((*MockMerchantServiceServer)(nil).DeleteMerchantPermanent), arg0, arg1)
}

// FindAllMerchant mocks base method.
func (m *MockMerchantServiceServer) FindAllMerchant(arg0 context.Context, arg1 *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllMerchant", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllMerchant indicates an expected call of FindAllMerchant.
func (mr *MockMerchantServiceServerMockRecorder) FindAllMerchant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllMerchant", reflect.TypeOf((*MockMerchantServiceServer)(nil).FindAllMerchant), arg0, arg1)
}

// FindByActive mocks base method.
func (m *MockMerchantServiceServer) FindByActive(arg0 context.Context, arg1 *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByActive", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationMerchantDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActive indicates an expected call of FindByActive.
func (mr *MockMerchantServiceServerMockRecorder) FindByActive(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActive", reflect.TypeOf((*MockMerchantServiceServer)(nil).FindByActive), arg0, arg1)
}

// FindByApiKey mocks base method.
func (m *MockMerchantServiceServer) FindByApiKey(arg0 context.Context, arg1 *pb.FindByApiKeyRequest) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByApiKey", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByApiKey indicates an expected call of FindByApiKey.
func (mr *MockMerchantServiceServerMockRecorder) FindByApiKey(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByApiKey", reflect.TypeOf((*MockMerchantServiceServer)(nil).FindByApiKey), arg0, arg1)
}

// FindByIdMerchant mocks base method.
func (m *MockMerchantServiceServer) FindByIdMerchant(arg0 context.Context, arg1 *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdMerchant", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdMerchant indicates an expected call of FindByIdMerchant.
func (mr *MockMerchantServiceServerMockRecorder) FindByIdMerchant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdMerchant", reflect.TypeOf((*MockMerchantServiceServer)(nil).FindByIdMerchant), arg0, arg1)
}

// FindByMerchantUserId mocks base method.
func (m *MockMerchantServiceServer) FindByMerchantUserId(arg0 context.Context, arg1 *pb.FindByMerchantUserIdRequest) (*pb.ApiResponsesMerchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMerchantUserId", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsesMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMerchantUserId indicates an expected call of FindByMerchantUserId.
func (mr *MockMerchantServiceServerMockRecorder) FindByMerchantUserId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMerchantUserId", reflect.TypeOf((*MockMerchantServiceServer)(nil).FindByMerchantUserId), arg0, arg1)
}

// FindByTrashed mocks base method.
func (m *MockMerchantServiceServer) FindByTrashed(arg0 context.Context, arg1 *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTrashed", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationMerchantDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTrashed indicates an expected call of FindByTrashed.
func (mr *MockMerchantServiceServerMockRecorder) FindByTrashed(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTrashed", reflect.TypeOf((*MockMerchantServiceServer)(nil).FindByTrashed), arg0, arg1)
}

// RestoreMerchant mocks base method.
func (m *MockMerchantServiceServer) RestoreMerchant(arg0 context.Context, arg1 *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreMerchant", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreMerchant indicates an expected call of RestoreMerchant.
func (mr *MockMerchantServiceServerMockRecorder) RestoreMerchant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreMerchant", reflect.TypeOf((*MockMerchantServiceServer)(nil).RestoreMerchant), arg0, arg1)
}

// TrashedMerchant mocks base method.
func (m *MockMerchantServiceServer) TrashedMerchant(arg0 context.Context, arg1 *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrashedMerchant", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrashedMerchant indicates an expected call of TrashedMerchant.
func (mr *MockMerchantServiceServerMockRecorder) TrashedMerchant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrashedMerchant", reflect.TypeOf((*MockMerchantServiceServer)(nil).TrashedMerchant), arg0, arg1)
}

// UpdateMerchant mocks base method.
func (m *MockMerchantServiceServer) UpdateMerchant(arg0 context.Context, arg1 *pb.UpdateMerchantRequest) (*pb.ApiResponseMerchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMerchant", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseMerchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMerchant indicates an expected call of UpdateMerchant.
func (mr *MockMerchantServiceServerMockRecorder) UpdateMerchant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMerchant", reflect.TypeOf((*MockMerchantServiceServer)(nil).UpdateMerchant), arg0, arg1)
}

// mustEmbedUnimplementedMerchantServiceServer mocks base method.
func (m *MockMerchantServiceServer) mustEmbedUnimplementedMerchantServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMerchantServiceServer")
}

// mustEmbedUnimplementedMerchantServiceServer indicates an expected call of mustEmbedUnimplementedMerchantServiceServer.
func (mr *MockMerchantServiceServerMockRecorder) mustEmbedUnimplementedMerchantServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMerchantServiceServer", reflect.TypeOf((*MockMerchantServiceServer)(nil).mustEmbedUnimplementedMerchantServiceServer))
}

// MockUnsafeMerchantServiceServer is a mock of UnsafeMerchantServiceServer interface.
type MockUnsafeMerchantServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeMerchantServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeMerchantServiceServerMockRecorder is the mock recorder for MockUnsafeMerchantServiceServer.
type MockUnsafeMerchantServiceServerMockRecorder struct {
	mock *MockUnsafeMerchantServiceServer
}

// NewMockUnsafeMerchantServiceServer creates a new mock instance.
func NewMockUnsafeMerchantServiceServer(ctrl *gomock.Controller) *MockUnsafeMerchantServiceServer {
	mock := &MockUnsafeMerchantServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeMerchantServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeMerchantServiceServer) EXPECT() *MockUnsafeMerchantServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedMerchantServiceServer mocks base method.
func (m *MockUnsafeMerchantServiceServer) mustEmbedUnimplementedMerchantServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMerchantServiceServer")
}

// mustEmbedUnimplementedMerchantServiceServer indicates an expected call of mustEmbedUnimplementedMerchantServiceServer.
func (mr *MockUnsafeMerchantServiceServerMockRecorder) mustEmbedUnimplementedMerchantServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMerchantServiceServer", reflect.TypeOf((*MockUnsafeMerchantServiceServer)(nil).mustEmbedUnimplementedMerchantServiceServer))
}
