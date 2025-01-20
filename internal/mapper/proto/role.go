package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type roleProtoMapper struct {
}

func NewRoleProtoMapper() *roleProtoMapper {
	return &roleProtoMapper{}
}

func (s *roleProtoMapper) ToResponseRole(role *response.RoleResponse) *pb.RoleResponse {
	return &pb.RoleResponse{
		Id:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (s *roleProtoMapper) ToResponsesRole(roles []*response.RoleResponse) []*pb.RoleResponse {
	var responseRoles []*pb.RoleResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.ToResponseRole(role))
	}

	return responseRoles
}

func (s *roleProtoMapper) ToResponseRoleDeleteAt(role *response.RoleResponseDeleteAt) *pb.RoleResponseDeleteAt {
	return &pb.RoleResponseDeleteAt{
		Id:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt,
	}
}

func (s *roleProtoMapper) ToResponsesRoleDeleteAt(roles []*response.RoleResponseDeleteAt) []*pb.RoleResponseDeleteAt {
	var responseRoles []*pb.RoleResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.ToResponseRoleDeleteAt(role))
	}

	return responseRoles
}
