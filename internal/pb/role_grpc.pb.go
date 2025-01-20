// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: role.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RoleService_FindAllRole_FullMethodName            = "/pb.RoleService/FindAllRole"
	RoleService_FindByIdRole_FullMethodName           = "/pb.RoleService/FindByIdRole"
	RoleService_FindByActive_FullMethodName           = "/pb.RoleService/FindByActive"
	RoleService_FindByTrashed_FullMethodName          = "/pb.RoleService/FindByTrashed"
	RoleService_FindByUserId_FullMethodName           = "/pb.RoleService/FindByUserId"
	RoleService_CreateRole_FullMethodName             = "/pb.RoleService/CreateRole"
	RoleService_UpdateRole_FullMethodName             = "/pb.RoleService/UpdateRole"
	RoleService_TrashedRole_FullMethodName            = "/pb.RoleService/TrashedRole"
	RoleService_RestoreRole_FullMethodName            = "/pb.RoleService/RestoreRole"
	RoleService_DeleteRolePermanent_FullMethodName    = "/pb.RoleService/DeleteRolePermanent"
	RoleService_RestoreAllRole_FullMethodName         = "/pb.RoleService/RestoreAllRole"
	RoleService_DeleteAllRolePermanent_FullMethodName = "/pb.RoleService/DeleteAllRolePermanent"
)

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	FindAllRole(ctx context.Context, in *FindAllRoleRequest, opts ...grpc.CallOption) (*ApiResponsePaginationRole, error)
	FindByIdRole(ctx context.Context, in *FindByIdRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error)
	FindByActive(ctx context.Context, in *FindAllRoleRequest, opts ...grpc.CallOption) (*ApiResponsePaginationRoleDeleteAt, error)
	FindByTrashed(ctx context.Context, in *FindAllRoleRequest, opts ...grpc.CallOption) (*ApiResponsePaginationRoleDeleteAt, error)
	FindByUserId(ctx context.Context, in *FindByIdUserRoleRequest, opts ...grpc.CallOption) (*ApiResponsesRole, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error)
	TrashedRole(ctx context.Context, in *FindByIdRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error)
	RestoreRole(ctx context.Context, in *FindByIdRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error)
	DeleteRolePermanent(ctx context.Context, in *FindByIdRoleRequest, opts ...grpc.CallOption) (*ApiResponseRoleDelete, error)
	RestoreAllRole(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseRoleAll, error)
	DeleteAllRolePermanent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseRoleAll, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) FindAllRole(ctx context.Context, in *FindAllRoleRequest, opts ...grpc.CallOption) (*ApiResponsePaginationRole, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationRole)
	err := c.cc.Invoke(ctx, RoleService_FindAllRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) FindByIdRole(ctx context.Context, in *FindByIdRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseRole)
	err := c.cc.Invoke(ctx, RoleService_FindByIdRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) FindByActive(ctx context.Context, in *FindAllRoleRequest, opts ...grpc.CallOption) (*ApiResponsePaginationRoleDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationRoleDeleteAt)
	err := c.cc.Invoke(ctx, RoleService_FindByActive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) FindByTrashed(ctx context.Context, in *FindAllRoleRequest, opts ...grpc.CallOption) (*ApiResponsePaginationRoleDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationRoleDeleteAt)
	err := c.cc.Invoke(ctx, RoleService_FindByTrashed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) FindByUserId(ctx context.Context, in *FindByIdUserRoleRequest, opts ...grpc.CallOption) (*ApiResponsesRole, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsesRole)
	err := c.cc.Invoke(ctx, RoleService_FindByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseRole)
	err := c.cc.Invoke(ctx, RoleService_CreateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseRole)
	err := c.cc.Invoke(ctx, RoleService_UpdateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) TrashedRole(ctx context.Context, in *FindByIdRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseRole)
	err := c.cc.Invoke(ctx, RoleService_TrashedRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RestoreRole(ctx context.Context, in *FindByIdRoleRequest, opts ...grpc.CallOption) (*ApiResponseRole, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseRole)
	err := c.cc.Invoke(ctx, RoleService_RestoreRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteRolePermanent(ctx context.Context, in *FindByIdRoleRequest, opts ...grpc.CallOption) (*ApiResponseRoleDelete, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseRoleDelete)
	err := c.cc.Invoke(ctx, RoleService_DeleteRolePermanent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RestoreAllRole(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseRoleAll, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseRoleAll)
	err := c.cc.Invoke(ctx, RoleService_RestoreAllRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteAllRolePermanent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseRoleAll, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseRoleAll)
	err := c.cc.Invoke(ctx, RoleService_DeleteAllRolePermanent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations must embed UnimplementedRoleServiceServer
// for forward compatibility.
type RoleServiceServer interface {
	FindAllRole(context.Context, *FindAllRoleRequest) (*ApiResponsePaginationRole, error)
	FindByIdRole(context.Context, *FindByIdRoleRequest) (*ApiResponseRole, error)
	FindByActive(context.Context, *FindAllRoleRequest) (*ApiResponsePaginationRoleDeleteAt, error)
	FindByTrashed(context.Context, *FindAllRoleRequest) (*ApiResponsePaginationRoleDeleteAt, error)
	FindByUserId(context.Context, *FindByIdUserRoleRequest) (*ApiResponsesRole, error)
	CreateRole(context.Context, *CreateRoleRequest) (*ApiResponseRole, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*ApiResponseRole, error)
	TrashedRole(context.Context, *FindByIdRoleRequest) (*ApiResponseRole, error)
	RestoreRole(context.Context, *FindByIdRoleRequest) (*ApiResponseRole, error)
	DeleteRolePermanent(context.Context, *FindByIdRoleRequest) (*ApiResponseRoleDelete, error)
	RestoreAllRole(context.Context, *emptypb.Empty) (*ApiResponseRoleAll, error)
	DeleteAllRolePermanent(context.Context, *emptypb.Empty) (*ApiResponseRoleAll, error)
	mustEmbedUnimplementedRoleServiceServer()
}

// UnimplementedRoleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoleServiceServer struct{}

func (UnimplementedRoleServiceServer) FindAllRole(context.Context, *FindAllRoleRequest) (*ApiResponsePaginationRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllRole not implemented")
}
func (UnimplementedRoleServiceServer) FindByIdRole(context.Context, *FindByIdRoleRequest) (*ApiResponseRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByIdRole not implemented")
}
func (UnimplementedRoleServiceServer) FindByActive(context.Context, *FindAllRoleRequest) (*ApiResponsePaginationRoleDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByActive not implemented")
}
func (UnimplementedRoleServiceServer) FindByTrashed(context.Context, *FindAllRoleRequest) (*ApiResponsePaginationRoleDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByTrashed not implemented")
}
func (UnimplementedRoleServiceServer) FindByUserId(context.Context, *FindByIdUserRoleRequest) (*ApiResponsesRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUserId not implemented")
}
func (UnimplementedRoleServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*ApiResponseRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRoleServiceServer) UpdateRole(context.Context, *UpdateRoleRequest) (*ApiResponseRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleServiceServer) TrashedRole(context.Context, *FindByIdRoleRequest) (*ApiResponseRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrashedRole not implemented")
}
func (UnimplementedRoleServiceServer) RestoreRole(context.Context, *FindByIdRoleRequest) (*ApiResponseRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreRole not implemented")
}
func (UnimplementedRoleServiceServer) DeleteRolePermanent(context.Context, *FindByIdRoleRequest) (*ApiResponseRoleDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRolePermanent not implemented")
}
func (UnimplementedRoleServiceServer) RestoreAllRole(context.Context, *emptypb.Empty) (*ApiResponseRoleAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreAllRole not implemented")
}
func (UnimplementedRoleServiceServer) DeleteAllRolePermanent(context.Context, *emptypb.Empty) (*ApiResponseRoleAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllRolePermanent not implemented")
}
func (UnimplementedRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}
func (UnimplementedRoleServiceServer) testEmbeddedByValue()                     {}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	// If the following call pancis, it indicates UnimplementedRoleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_FindAllRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindAllRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_FindAllRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindAllRole(ctx, req.(*FindAllRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_FindByIdRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindByIdRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_FindByIdRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindByIdRole(ctx, req.(*FindByIdRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_FindByActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindByActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_FindByActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindByActive(ctx, req.(*FindAllRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_FindByTrashed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindByTrashed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_FindByTrashed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindByTrashed(ctx, req.(*FindAllRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_FindByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_FindByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindByUserId(ctx, req.(*FindByIdUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_TrashedRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).TrashedRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_TrashedRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).TrashedRole(ctx, req.(*FindByIdRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RestoreRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RestoreRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_RestoreRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RestoreRole(ctx, req.(*FindByIdRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteRolePermanent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteRolePermanent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_DeleteRolePermanent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteRolePermanent(ctx, req.(*FindByIdRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RestoreAllRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RestoreAllRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_RestoreAllRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RestoreAllRole(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteAllRolePermanent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteAllRolePermanent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_DeleteAllRolePermanent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteAllRolePermanent(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllRole",
			Handler:    _RoleService_FindAllRole_Handler,
		},
		{
			MethodName: "FindByIdRole",
			Handler:    _RoleService_FindByIdRole_Handler,
		},
		{
			MethodName: "FindByActive",
			Handler:    _RoleService_FindByActive_Handler,
		},
		{
			MethodName: "FindByTrashed",
			Handler:    _RoleService_FindByTrashed_Handler,
		},
		{
			MethodName: "FindByUserId",
			Handler:    _RoleService_FindByUserId_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _RoleService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RoleService_UpdateRole_Handler,
		},
		{
			MethodName: "TrashedRole",
			Handler:    _RoleService_TrashedRole_Handler,
		},
		{
			MethodName: "RestoreRole",
			Handler:    _RoleService_RestoreRole_Handler,
		},
		{
			MethodName: "DeleteRolePermanent",
			Handler:    _RoleService_DeleteRolePermanent_Handler,
		},
		{
			MethodName: "RestoreAllRole",
			Handler:    _RoleService_RestoreAllRole_Handler,
		},
		{
			MethodName: "DeleteAllRolePermanent",
			Handler:    _RoleService_DeleteAllRolePermanent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}
