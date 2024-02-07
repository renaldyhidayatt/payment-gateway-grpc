// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: topup.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TopupService_GetTopups_FullMethodName        = "/pb.TopupService/GetTopups"
	TopupService_GetTopup_FullMethodName         = "/pb.TopupService/GetTopup"
	TopupService_GetTopupByUsers_FullMethodName  = "/pb.TopupService/GetTopupByUsers"
	TopupService_GetTopupByUserId_FullMethodName = "/pb.TopupService/GetTopupByUserId"
	TopupService_CreateTopup_FullMethodName      = "/pb.TopupService/CreateTopup"
	TopupService_UpdateTopup_FullMethodName      = "/pb.TopupService/UpdateTopup"
	TopupService_DeleteTopup_FullMethodName      = "/pb.TopupService/DeleteTopup"
)

// TopupServiceClient is the client API for TopupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopupServiceClient interface {
	GetTopups(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TopupsResponse, error)
	GetTopup(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TopupResponse, error)
	GetTopupByUsers(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TopupsResponse, error)
	GetTopupByUserId(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TopupResponse, error)
	CreateTopup(ctx context.Context, in *CreateTopupRequest, opts ...grpc.CallOption) (*TopupResponse, error)
	UpdateTopup(ctx context.Context, in *UpdateTopupRequest, opts ...grpc.CallOption) (*TopupResponse, error)
	DeleteTopup(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*DeleteTopupResponse, error)
}

type topupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopupServiceClient(cc grpc.ClientConnInterface) TopupServiceClient {
	return &topupServiceClient{cc}
}

func (c *topupServiceClient) GetTopups(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TopupsResponse, error) {
	out := new(TopupsResponse)
	err := c.cc.Invoke(ctx, TopupService_GetTopups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topupServiceClient) GetTopup(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TopupResponse, error) {
	out := new(TopupResponse)
	err := c.cc.Invoke(ctx, TopupService_GetTopup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topupServiceClient) GetTopupByUsers(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TopupsResponse, error) {
	out := new(TopupsResponse)
	err := c.cc.Invoke(ctx, TopupService_GetTopupByUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topupServiceClient) GetTopupByUserId(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TopupResponse, error) {
	out := new(TopupResponse)
	err := c.cc.Invoke(ctx, TopupService_GetTopupByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topupServiceClient) CreateTopup(ctx context.Context, in *CreateTopupRequest, opts ...grpc.CallOption) (*TopupResponse, error) {
	out := new(TopupResponse)
	err := c.cc.Invoke(ctx, TopupService_CreateTopup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topupServiceClient) UpdateTopup(ctx context.Context, in *UpdateTopupRequest, opts ...grpc.CallOption) (*TopupResponse, error) {
	out := new(TopupResponse)
	err := c.cc.Invoke(ctx, TopupService_UpdateTopup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topupServiceClient) DeleteTopup(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*DeleteTopupResponse, error) {
	out := new(DeleteTopupResponse)
	err := c.cc.Invoke(ctx, TopupService_DeleteTopup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopupServiceServer is the server API for TopupService service.
// All implementations must embed UnimplementedTopupServiceServer
// for forward compatibility
type TopupServiceServer interface {
	GetTopups(context.Context, *empty.Empty) (*TopupsResponse, error)
	GetTopup(context.Context, *TopupRequest) (*TopupResponse, error)
	GetTopupByUsers(context.Context, *TopupRequest) (*TopupsResponse, error)
	GetTopupByUserId(context.Context, *TopupRequest) (*TopupResponse, error)
	CreateTopup(context.Context, *CreateTopupRequest) (*TopupResponse, error)
	UpdateTopup(context.Context, *UpdateTopupRequest) (*TopupResponse, error)
	DeleteTopup(context.Context, *TopupRequest) (*DeleteTopupResponse, error)
	mustEmbedUnimplementedTopupServiceServer()
}

// UnimplementedTopupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTopupServiceServer struct {
}

func (UnimplementedTopupServiceServer) GetTopups(context.Context, *empty.Empty) (*TopupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopups not implemented")
}
func (UnimplementedTopupServiceServer) GetTopup(context.Context, *TopupRequest) (*TopupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopup not implemented")
}
func (UnimplementedTopupServiceServer) GetTopupByUsers(context.Context, *TopupRequest) (*TopupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopupByUsers not implemented")
}
func (UnimplementedTopupServiceServer) GetTopupByUserId(context.Context, *TopupRequest) (*TopupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopupByUserId not implemented")
}
func (UnimplementedTopupServiceServer) CreateTopup(context.Context, *CreateTopupRequest) (*TopupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopup not implemented")
}
func (UnimplementedTopupServiceServer) UpdateTopup(context.Context, *UpdateTopupRequest) (*TopupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopup not implemented")
}
func (UnimplementedTopupServiceServer) DeleteTopup(context.Context, *TopupRequest) (*DeleteTopupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopup not implemented")
}
func (UnimplementedTopupServiceServer) mustEmbedUnimplementedTopupServiceServer() {}

// UnsafeTopupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopupServiceServer will
// result in compilation errors.
type UnsafeTopupServiceServer interface {
	mustEmbedUnimplementedTopupServiceServer()
}

func RegisterTopupServiceServer(s grpc.ServiceRegistrar, srv TopupServiceServer) {
	s.RegisterService(&TopupService_ServiceDesc, srv)
}

func _TopupService_GetTopups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopupServiceServer).GetTopups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopupService_GetTopups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopupServiceServer).GetTopups(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopupService_GetTopup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopupServiceServer).GetTopup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopupService_GetTopup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopupServiceServer).GetTopup(ctx, req.(*TopupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopupService_GetTopupByUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopupServiceServer).GetTopupByUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopupService_GetTopupByUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopupServiceServer).GetTopupByUsers(ctx, req.(*TopupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopupService_GetTopupByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopupServiceServer).GetTopupByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopupService_GetTopupByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopupServiceServer).GetTopupByUserId(ctx, req.(*TopupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopupService_CreateTopup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopupServiceServer).CreateTopup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopupService_CreateTopup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopupServiceServer).CreateTopup(ctx, req.(*CreateTopupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopupService_UpdateTopup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopupServiceServer).UpdateTopup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopupService_UpdateTopup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopupServiceServer).UpdateTopup(ctx, req.(*UpdateTopupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopupService_DeleteTopup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopupServiceServer).DeleteTopup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopupService_DeleteTopup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopupServiceServer).DeleteTopup(ctx, req.(*TopupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TopupService_ServiceDesc is the grpc.ServiceDesc for TopupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TopupService",
	HandlerType: (*TopupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopups",
			Handler:    _TopupService_GetTopups_Handler,
		},
		{
			MethodName: "GetTopup",
			Handler:    _TopupService_GetTopup_Handler,
		},
		{
			MethodName: "GetTopupByUsers",
			Handler:    _TopupService_GetTopupByUsers_Handler,
		},
		{
			MethodName: "GetTopupByUserId",
			Handler:    _TopupService_GetTopupByUserId_Handler,
		},
		{
			MethodName: "CreateTopup",
			Handler:    _TopupService_CreateTopup_Handler,
		},
		{
			MethodName: "UpdateTopup",
			Handler:    _TopupService_UpdateTopup_Handler,
		},
		{
			MethodName: "DeleteTopup",
			Handler:    _TopupService_DeleteTopup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topup.proto",
}
