// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: saldo.proto

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
	SaldoService_GetSaldos_FullMethodName        = "/pb.SaldoService/GetSaldos"
	SaldoService_GetSaldo_FullMethodName         = "/pb.SaldoService/GetSaldo"
	SaldoService_GetSaldoByUsers_FullMethodName  = "/pb.SaldoService/GetSaldoByUsers"
	SaldoService_GetSaldoByUserId_FullMethodName = "/pb.SaldoService/GetSaldoByUserId"
	SaldoService_CreateSaldo_FullMethodName      = "/pb.SaldoService/CreateSaldo"
	SaldoService_UpdateSaldo_FullMethodName      = "/pb.SaldoService/UpdateSaldo"
	SaldoService_DeleteSaldo_FullMethodName      = "/pb.SaldoService/DeleteSaldo"
)

// SaldoServiceClient is the client API for SaldoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SaldoServiceClient interface {
	GetSaldos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SaldoResponses, error)
	GetSaldo(ctx context.Context, in *SaldoRequest, opts ...grpc.CallOption) (*SaldoResponse, error)
	GetSaldoByUsers(ctx context.Context, in *SaldoRequest, opts ...grpc.CallOption) (*SaldoResponses, error)
	GetSaldoByUserId(ctx context.Context, in *SaldoRequest, opts ...grpc.CallOption) (*SaldoResponse, error)
	CreateSaldo(ctx context.Context, in *CreateSaldoRequest, opts ...grpc.CallOption) (*SaldoResponse, error)
	UpdateSaldo(ctx context.Context, in *UpdateSaldoRequest, opts ...grpc.CallOption) (*SaldoResponse, error)
	DeleteSaldo(ctx context.Context, in *SaldoRequest, opts ...grpc.CallOption) (*DeleteSaldoResponse, error)
}

type saldoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSaldoServiceClient(cc grpc.ClientConnInterface) SaldoServiceClient {
	return &saldoServiceClient{cc}
}

func (c *saldoServiceClient) GetSaldos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SaldoResponses, error) {
	out := new(SaldoResponses)
	err := c.cc.Invoke(ctx, SaldoService_GetSaldos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saldoServiceClient) GetSaldo(ctx context.Context, in *SaldoRequest, opts ...grpc.CallOption) (*SaldoResponse, error) {
	out := new(SaldoResponse)
	err := c.cc.Invoke(ctx, SaldoService_GetSaldo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saldoServiceClient) GetSaldoByUsers(ctx context.Context, in *SaldoRequest, opts ...grpc.CallOption) (*SaldoResponses, error) {
	out := new(SaldoResponses)
	err := c.cc.Invoke(ctx, SaldoService_GetSaldoByUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saldoServiceClient) GetSaldoByUserId(ctx context.Context, in *SaldoRequest, opts ...grpc.CallOption) (*SaldoResponse, error) {
	out := new(SaldoResponse)
	err := c.cc.Invoke(ctx, SaldoService_GetSaldoByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saldoServiceClient) CreateSaldo(ctx context.Context, in *CreateSaldoRequest, opts ...grpc.CallOption) (*SaldoResponse, error) {
	out := new(SaldoResponse)
	err := c.cc.Invoke(ctx, SaldoService_CreateSaldo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saldoServiceClient) UpdateSaldo(ctx context.Context, in *UpdateSaldoRequest, opts ...grpc.CallOption) (*SaldoResponse, error) {
	out := new(SaldoResponse)
	err := c.cc.Invoke(ctx, SaldoService_UpdateSaldo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saldoServiceClient) DeleteSaldo(ctx context.Context, in *SaldoRequest, opts ...grpc.CallOption) (*DeleteSaldoResponse, error) {
	out := new(DeleteSaldoResponse)
	err := c.cc.Invoke(ctx, SaldoService_DeleteSaldo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SaldoServiceServer is the server API for SaldoService service.
// All implementations must embed UnimplementedSaldoServiceServer
// for forward compatibility
type SaldoServiceServer interface {
	GetSaldos(context.Context, *empty.Empty) (*SaldoResponses, error)
	GetSaldo(context.Context, *SaldoRequest) (*SaldoResponse, error)
	GetSaldoByUsers(context.Context, *SaldoRequest) (*SaldoResponses, error)
	GetSaldoByUserId(context.Context, *SaldoRequest) (*SaldoResponse, error)
	CreateSaldo(context.Context, *CreateSaldoRequest) (*SaldoResponse, error)
	UpdateSaldo(context.Context, *UpdateSaldoRequest) (*SaldoResponse, error)
	DeleteSaldo(context.Context, *SaldoRequest) (*DeleteSaldoResponse, error)
	mustEmbedUnimplementedSaldoServiceServer()
}

// UnimplementedSaldoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSaldoServiceServer struct {
}

func (UnimplementedSaldoServiceServer) GetSaldos(context.Context, *empty.Empty) (*SaldoResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSaldos not implemented")
}
func (UnimplementedSaldoServiceServer) GetSaldo(context.Context, *SaldoRequest) (*SaldoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSaldo not implemented")
}
func (UnimplementedSaldoServiceServer) GetSaldoByUsers(context.Context, *SaldoRequest) (*SaldoResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSaldoByUsers not implemented")
}
func (UnimplementedSaldoServiceServer) GetSaldoByUserId(context.Context, *SaldoRequest) (*SaldoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSaldoByUserId not implemented")
}
func (UnimplementedSaldoServiceServer) CreateSaldo(context.Context, *CreateSaldoRequest) (*SaldoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSaldo not implemented")
}
func (UnimplementedSaldoServiceServer) UpdateSaldo(context.Context, *UpdateSaldoRequest) (*SaldoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSaldo not implemented")
}
func (UnimplementedSaldoServiceServer) DeleteSaldo(context.Context, *SaldoRequest) (*DeleteSaldoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSaldo not implemented")
}
func (UnimplementedSaldoServiceServer) mustEmbedUnimplementedSaldoServiceServer() {}

// UnsafeSaldoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SaldoServiceServer will
// result in compilation errors.
type UnsafeSaldoServiceServer interface {
	mustEmbedUnimplementedSaldoServiceServer()
}

func RegisterSaldoServiceServer(s grpc.ServiceRegistrar, srv SaldoServiceServer) {
	s.RegisterService(&SaldoService_ServiceDesc, srv)
}

func _SaldoService_GetSaldos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaldoServiceServer).GetSaldos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaldoService_GetSaldos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaldoServiceServer).GetSaldos(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaldoService_GetSaldo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaldoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaldoServiceServer).GetSaldo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaldoService_GetSaldo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaldoServiceServer).GetSaldo(ctx, req.(*SaldoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaldoService_GetSaldoByUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaldoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaldoServiceServer).GetSaldoByUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaldoService_GetSaldoByUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaldoServiceServer).GetSaldoByUsers(ctx, req.(*SaldoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaldoService_GetSaldoByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaldoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaldoServiceServer).GetSaldoByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaldoService_GetSaldoByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaldoServiceServer).GetSaldoByUserId(ctx, req.(*SaldoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaldoService_CreateSaldo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSaldoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaldoServiceServer).CreateSaldo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaldoService_CreateSaldo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaldoServiceServer).CreateSaldo(ctx, req.(*CreateSaldoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaldoService_UpdateSaldo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSaldoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaldoServiceServer).UpdateSaldo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaldoService_UpdateSaldo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaldoServiceServer).UpdateSaldo(ctx, req.(*UpdateSaldoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaldoService_DeleteSaldo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaldoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaldoServiceServer).DeleteSaldo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaldoService_DeleteSaldo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaldoServiceServer).DeleteSaldo(ctx, req.(*SaldoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SaldoService_ServiceDesc is the grpc.ServiceDesc for SaldoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SaldoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SaldoService",
	HandlerType: (*SaldoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSaldos",
			Handler:    _SaldoService_GetSaldos_Handler,
		},
		{
			MethodName: "GetSaldo",
			Handler:    _SaldoService_GetSaldo_Handler,
		},
		{
			MethodName: "GetSaldoByUsers",
			Handler:    _SaldoService_GetSaldoByUsers_Handler,
		},
		{
			MethodName: "GetSaldoByUserId",
			Handler:    _SaldoService_GetSaldoByUserId_Handler,
		},
		{
			MethodName: "CreateSaldo",
			Handler:    _SaldoService_CreateSaldo_Handler,
		},
		{
			MethodName: "UpdateSaldo",
			Handler:    _SaldoService_UpdateSaldo_Handler,
		},
		{
			MethodName: "DeleteSaldo",
			Handler:    _SaldoService_DeleteSaldo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "saldo.proto",
}
