// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: transfer.proto

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
	TransferService_GetTransfers_FullMethodName        = "/pb.TransferService/GetTransfers"
	TransferService_GetTransfer_FullMethodName         = "/pb.TransferService/GetTransfer"
	TransferService_GetTransferByUsers_FullMethodName  = "/pb.TransferService/GetTransferByUsers"
	TransferService_GetTransferByUserId_FullMethodName = "/pb.TransferService/GetTransferByUserId"
	TransferService_CreateTransfer_FullMethodName      = "/pb.TransferService/CreateTransfer"
	TransferService_UpdateTransfer_FullMethodName      = "/pb.TransferService/UpdateTransfer"
	TransferService_DeleteTransfer_FullMethodName      = "/pb.TransferService/DeleteTransfer"
)

// TransferServiceClient is the client API for TransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferServiceClient interface {
	GetTransfers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TransfersResponse, error)
	GetTransfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	GetTransferByUsers(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransfersResponse, error)
	GetTransferByUserId(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	UpdateTransfer(ctx context.Context, in *UpdateTransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	DeleteTransfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*DeleteTransferResponse, error)
}

type transferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferServiceClient(cc grpc.ClientConnInterface) TransferServiceClient {
	return &transferServiceClient{cc}
}

func (c *transferServiceClient) GetTransfers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TransfersResponse, error) {
	out := new(TransfersResponse)
	err := c.cc.Invoke(ctx, TransferService_GetTransfers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) GetTransfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, TransferService_GetTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) GetTransferByUsers(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransfersResponse, error) {
	out := new(TransfersResponse)
	err := c.cc.Invoke(ctx, TransferService_GetTransferByUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) GetTransferByUserId(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, TransferService_GetTransferByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, TransferService_CreateTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) UpdateTransfer(ctx context.Context, in *UpdateTransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, TransferService_UpdateTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) DeleteTransfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*DeleteTransferResponse, error) {
	out := new(DeleteTransferResponse)
	err := c.cc.Invoke(ctx, TransferService_DeleteTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferServiceServer is the server API for TransferService service.
// All implementations must embed UnimplementedTransferServiceServer
// for forward compatibility
type TransferServiceServer interface {
	GetTransfers(context.Context, *empty.Empty) (*TransfersResponse, error)
	GetTransfer(context.Context, *TransferRequest) (*TransferResponse, error)
	GetTransferByUsers(context.Context, *TransferRequest) (*TransfersResponse, error)
	GetTransferByUserId(context.Context, *TransferRequest) (*TransferResponse, error)
	CreateTransfer(context.Context, *CreateTransferRequest) (*TransferResponse, error)
	UpdateTransfer(context.Context, *UpdateTransferRequest) (*TransferResponse, error)
	DeleteTransfer(context.Context, *TransferRequest) (*DeleteTransferResponse, error)
	mustEmbedUnimplementedTransferServiceServer()
}

// UnimplementedTransferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransferServiceServer struct {
}

func (UnimplementedTransferServiceServer) GetTransfers(context.Context, *empty.Empty) (*TransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransfers not implemented")
}
func (UnimplementedTransferServiceServer) GetTransfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransfer not implemented")
}
func (UnimplementedTransferServiceServer) GetTransferByUsers(context.Context, *TransferRequest) (*TransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransferByUsers not implemented")
}
func (UnimplementedTransferServiceServer) GetTransferByUserId(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransferByUserId not implemented")
}
func (UnimplementedTransferServiceServer) CreateTransfer(context.Context, *CreateTransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedTransferServiceServer) UpdateTransfer(context.Context, *UpdateTransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransfer not implemented")
}
func (UnimplementedTransferServiceServer) DeleteTransfer(context.Context, *TransferRequest) (*DeleteTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransfer not implemented")
}
func (UnimplementedTransferServiceServer) mustEmbedUnimplementedTransferServiceServer() {}

// UnsafeTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServiceServer will
// result in compilation errors.
type UnsafeTransferServiceServer interface {
	mustEmbedUnimplementedTransferServiceServer()
}

func RegisterTransferServiceServer(s grpc.ServiceRegistrar, srv TransferServiceServer) {
	s.RegisterService(&TransferService_ServiceDesc, srv)
}

func _TransferService_GetTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).GetTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_GetTransfers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).GetTransfers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_GetTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).GetTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_GetTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).GetTransfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_GetTransferByUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).GetTransferByUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_GetTransferByUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).GetTransferByUsers(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_GetTransferByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).GetTransferByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_GetTransferByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).GetTransferByUserId(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_CreateTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).CreateTransfer(ctx, req.(*CreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_UpdateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).UpdateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_UpdateTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).UpdateTransfer(ctx, req.(*UpdateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_DeleteTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).DeleteTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_DeleteTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).DeleteTransfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransferService_ServiceDesc is the grpc.ServiceDesc for TransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TransferService",
	HandlerType: (*TransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransfers",
			Handler:    _TransferService_GetTransfers_Handler,
		},
		{
			MethodName: "GetTransfer",
			Handler:    _TransferService_GetTransfer_Handler,
		},
		{
			MethodName: "GetTransferByUsers",
			Handler:    _TransferService_GetTransferByUsers_Handler,
		},
		{
			MethodName: "GetTransferByUserId",
			Handler:    _TransferService_GetTransferByUserId_Handler,
		},
		{
			MethodName: "CreateTransfer",
			Handler:    _TransferService_CreateTransfer_Handler,
		},
		{
			MethodName: "UpdateTransfer",
			Handler:    _TransferService_UpdateTransfer_Handler,
		},
		{
			MethodName: "DeleteTransfer",
			Handler:    _TransferService_DeleteTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer.proto",
}
