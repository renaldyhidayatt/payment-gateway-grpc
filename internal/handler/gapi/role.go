package gapi

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	protomapper "MamangRust/paymentgatewaygrpc/internal/mapper/proto"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"MamangRust/paymentgatewaygrpc/internal/service"
	"context"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type roleHandleGrpc struct {
	pb.UnimplementedRoleServiceServer
	roleService service.RoleService
	mapping     protomapper.RoleProtoMapper
}

func NewRoleHandleGrpc(role service.RoleService, mapping protomapper.RoleProtoMapper) *roleHandleGrpc {
	return &roleHandleGrpc{
		roleService: role,
		mapping:     mapping,
	}
}

func (s *roleHandleGrpc) FindAllRole(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRole, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	role, totalRecords, err := s.roleService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card records: " + err.Message,
		})
	}

	so := s.mapping.ToResponsesRole(role)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationRole{
		Status:     "success",
		Message:    "Successfully fetched card records",
		Data:       so,
		Pagination: paginationMeta,
	}, nil
}

func (s *roleHandleGrpc) FindByIdRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRole, error) {
	roleID := int(req.GetRoleId())

	role, err := s.roleService.FindById(roleID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch role: " + err.Message,
		})
	}

	roleResponse := s.mapping.ToResponseRole(role)

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully fetched role",
		Data:    roleResponse,
	}, nil
}

func (s *roleHandleGrpc) FindByUserId(ctx context.Context, req *pb.FindByIdUserRoleRequest) (*pb.ApiResponsesRole, error) {
	userID := int(req.GetUserId())

	role, err := s.roleService.FindByUserId(userID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch role by user ID: " + err.Message,
		})
	}

	roleResponse := s.mapping.ToResponsesRole(role)

	return &pb.ApiResponsesRole{
		Status:  "success",
		Message: "Successfully fetched role by user ID",
		Data:    roleResponse,
	}, nil
}

func (s *roleHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRoleDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	roles, totalRecords, err := s.roleService.FindByActiveRole(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active roles: " + err.Message,
		})
	}

	roleResponses := s.mapping.ToResponsesRoleDeleteAt(roles)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationRoleDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active roles",
		Data:       roleResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *roleHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRoleDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	roles, totalRecords, err := s.roleService.FindByTrashedRole(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed roles: " + err.Message,
		})
	}

	roleResponses := s.mapping.ToResponsesRoleDeleteAt(roles)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	return &pb.ApiResponsePaginationRoleDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed roles",
		Data:       roleResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *roleHandleGrpc) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.ApiResponseRole, error) {
	name := req.GetName()

	role, err := s.roleService.CreateRole(&requests.CreateRoleRequest{
		Name: name,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create role: t" + err.Message,
		})
	}

	roleResponse := s.mapping.ToResponseRole(role)

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully created role",
		Data:    roleResponse,
	}, nil
}

func (s *roleHandleGrpc) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.ApiResponseRole, error) {
	roleID := int(req.GetId())
	name := req.GetName()

	role, err := s.roleService.UpdateRole(&requests.UpdateRoleRequest{
		ID:   roleID,
		Name: name,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update role: " + err.Message,
		})
	}

	roleResponse := s.mapping.ToResponseRole(role)

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully updated role",
		Data:    roleResponse,
	}, nil
}

func (s *roleHandleGrpc) TrashedRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRole, error) {
	roleID := int(req.GetRoleId())

	role, err := s.roleService.TrashedRole(roleID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash role: " + err.Message,
		})
	}

	roleResponse := s.mapping.ToResponseRole(role)

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully trashed role",
		Data:    roleResponse,
	}, nil
}

func (s *roleHandleGrpc) RestoreRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRole, error) {
	roleID := int(req.GetRoleId())

	role, err := s.roleService.RestoreRole(roleID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore role: " + err.Message,
		})
	}

	roleResponse := s.mapping.ToResponseRole(role)

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully restored role",
		Data:    roleResponse,
	}, nil
}

func (s *roleHandleGrpc) DeleteRolePermanent(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRoleDelete, error) {
	roleID := int(req.GetRoleId())

	_, err := s.roleService.DeleteRolePermanent(roleID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete role permanently: " + err.Message,
		})
	}

	return &pb.ApiResponseRoleDelete{
		Status:  "success",
		Message: "Successfully deleted role permanently",
	}, nil
}

func (s *roleHandleGrpc) RestoreAllRole(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseRoleAll, error) {
	_, err := s.roleService.RestoreAllRole()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all roles: " + err.Message,
		})
	}

	return &pb.ApiResponseRoleAll{
		Status:  "success",
		Message: "Successfully restored all roles",
	}, nil
}

func (s *roleHandleGrpc) DeleteAllRolePermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseRoleAll, error) {
	_, err := s.roleService.DeleteAllRolePermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all roles permanently: " + err.Message,
		})
	}

	return &pb.ApiResponseRoleAll{
		Status:  "success",
		Message: "Successfully deleted all roles permanently",
	}, nil
}
