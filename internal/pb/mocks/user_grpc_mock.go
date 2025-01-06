// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pb/user_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=internal/pb/user_grpc.pb.go -destination=internal/pb/mocks/user_grpc_mock.go
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

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
	isgomock struct{}
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserServiceClient) Create(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserServiceClientMockRecorder) Create(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserServiceClient)(nil).Create), varargs...)
}

// DeleteUserPermanent mocks base method.
func (m *MockUserServiceClient) DeleteUserPermanent(ctx context.Context, in *pb.FindByIdUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUserDelete, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteUserPermanent", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseUserDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserPermanent indicates an expected call of DeleteUserPermanent.
func (mr *MockUserServiceClientMockRecorder) DeleteUserPermanent(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserPermanent", reflect.TypeOf((*MockUserServiceClient)(nil).DeleteUserPermanent), varargs...)
}

// FindAll mocks base method.
func (m *MockUserServiceClient) FindAll(ctx context.Context, in *pb.FindAllUserRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationUser, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAll", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserServiceClientMockRecorder) FindAll(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserServiceClient)(nil).FindAll), varargs...)
}

// FindByActive mocks base method.
func (m *MockUserServiceClient) FindByActive(ctx context.Context, in *pb.FindAllUserRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationUserDeleteAt, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByActive", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationUserDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActive indicates an expected call of FindByActive.
func (mr *MockUserServiceClientMockRecorder) FindByActive(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActive", reflect.TypeOf((*MockUserServiceClient)(nil).FindByActive), varargs...)
}

// FindById mocks base method.
func (m *MockUserServiceClient) FindById(ctx context.Context, in *pb.FindByIdUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindById", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUserServiceClientMockRecorder) FindById(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserServiceClient)(nil).FindById), varargs...)
}

// FindByTrashed mocks base method.
func (m *MockUserServiceClient) FindByTrashed(ctx context.Context, in *pb.FindAllUserRequest, opts ...grpc.CallOption) (*pb.ApiResponsePaginationUserDeleteAt, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByTrashed", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationUserDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTrashed indicates an expected call of FindByTrashed.
func (mr *MockUserServiceClientMockRecorder) FindByTrashed(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTrashed", reflect.TypeOf((*MockUserServiceClient)(nil).FindByTrashed), varargs...)
}

// RestoreUser mocks base method.
func (m *MockUserServiceClient) RestoreUser(ctx context.Context, in *pb.FindByIdUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RestoreUser", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreUser indicates an expected call of RestoreUser.
func (mr *MockUserServiceClientMockRecorder) RestoreUser(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreUser", reflect.TypeOf((*MockUserServiceClient)(nil).RestoreUser), varargs...)
}

// TrashedUser mocks base method.
func (m *MockUserServiceClient) TrashedUser(ctx context.Context, in *pb.FindByIdUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TrashedUser", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrashedUser indicates an expected call of TrashedUser.
func (mr *MockUserServiceClientMockRecorder) TrashedUser(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrashedUser", reflect.TypeOf((*MockUserServiceClient)(nil).TrashedUser), varargs...)
}

// Update mocks base method.
func (m *MockUserServiceClient) Update(ctx context.Context, in *pb.UpdateUserRequest, opts ...grpc.CallOption) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserServiceClientMockRecorder) Update(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserServiceClient)(nil).Update), varargs...)
}

// MockUserServiceServer is a mock of UserServiceServer interface.
type MockUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceServerMockRecorder
	isgomock struct{}
}

// MockUserServiceServerMockRecorder is the mock recorder for MockUserServiceServer.
type MockUserServiceServerMockRecorder struct {
	mock *MockUserServiceServer
}

// NewMockUserServiceServer creates a new mock instance.
func NewMockUserServiceServer(ctrl *gomock.Controller) *MockUserServiceServer {
	mock := &MockUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceServer) EXPECT() *MockUserServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserServiceServer) Create(arg0 context.Context, arg1 *pb.CreateUserRequest) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserServiceServerMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserServiceServer)(nil).Create), arg0, arg1)
}

// DeleteUserPermanent mocks base method.
func (m *MockUserServiceServer) DeleteUserPermanent(arg0 context.Context, arg1 *pb.FindByIdUserRequest) (*pb.ApiResponseUserDelete, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserPermanent", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseUserDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserPermanent indicates an expected call of DeleteUserPermanent.
func (mr *MockUserServiceServerMockRecorder) DeleteUserPermanent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserPermanent", reflect.TypeOf((*MockUserServiceServer)(nil).DeleteUserPermanent), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockUserServiceServer) FindAll(arg0 context.Context, arg1 *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserServiceServerMockRecorder) FindAll(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserServiceServer)(nil).FindAll), arg0, arg1)
}

// FindByActive mocks base method.
func (m *MockUserServiceServer) FindByActive(arg0 context.Context, arg1 *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUserDeleteAt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByActive", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationUserDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActive indicates an expected call of FindByActive.
func (mr *MockUserServiceServerMockRecorder) FindByActive(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActive", reflect.TypeOf((*MockUserServiceServer)(nil).FindByActive), arg0, arg1)
}

// FindById mocks base method.
func (m *MockUserServiceServer) FindById(arg0 context.Context, arg1 *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUserServiceServerMockRecorder) FindById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserServiceServer)(nil).FindById), arg0, arg1)
}

// FindByTrashed mocks base method.
func (m *MockUserServiceServer) FindByTrashed(arg0 context.Context, arg1 *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUserDeleteAt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTrashed", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponsePaginationUserDeleteAt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTrashed indicates an expected call of FindByTrashed.
func (mr *MockUserServiceServerMockRecorder) FindByTrashed(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTrashed", reflect.TypeOf((*MockUserServiceServer)(nil).FindByTrashed), arg0, arg1)
}

// RestoreUser mocks base method.
func (m *MockUserServiceServer) RestoreUser(arg0 context.Context, arg1 *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreUser", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreUser indicates an expected call of RestoreUser.
func (mr *MockUserServiceServerMockRecorder) RestoreUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreUser", reflect.TypeOf((*MockUserServiceServer)(nil).RestoreUser), arg0, arg1)
}

// TrashedUser mocks base method.
func (m *MockUserServiceServer) TrashedUser(arg0 context.Context, arg1 *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrashedUser", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrashedUser indicates an expected call of TrashedUser.
func (mr *MockUserServiceServerMockRecorder) TrashedUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrashedUser", reflect.TypeOf((*MockUserServiceServer)(nil).TrashedUser), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserServiceServer) Update(arg0 context.Context, arg1 *pb.UpdateUserRequest) (*pb.ApiResponseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*pb.ApiResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserServiceServerMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserServiceServer)(nil).Update), arg0, arg1)
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUnsafeUserServiceServer is a mock of UnsafeUserServiceServer interface.
type MockUnsafeUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUserServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeUserServiceServerMockRecorder is the mock recorder for MockUnsafeUserServiceServer.
type MockUnsafeUserServiceServerMockRecorder struct {
	mock *MockUnsafeUserServiceServer
}

// NewMockUnsafeUserServiceServer creates a new mock instance.
func NewMockUnsafeUserServiceServer(ctrl *gomock.Controller) *MockUnsafeUserServiceServer {
	mock := &MockUnsafeUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUserServiceServer) EXPECT() *MockUnsafeUserServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUnsafeUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUnsafeUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}
