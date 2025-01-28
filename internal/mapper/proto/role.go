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

func (s *roleProtoMapper) ToProtoResponseRoleAll(status string, message string) *pb.ApiResponseRoleAll {
	return &pb.ApiResponseRoleAll{
		Status:  status,
		Message: message,
	}
}

func (s *roleProtoMapper) ToProtoResponseRoleDelete(status string, message string) *pb.ApiResponseRoleDelete {
	return &pb.ApiResponseRoleDelete{
		Status:  status,
		Message: message,
	}
}

func (s *roleProtoMapper) ToProtoResponseRole(status string, message string, pbResponse *response.RoleResponse) *pb.ApiResponseRole {
	return &pb.ApiResponseRole{
		Status:  status,
		Message: message,
		Data:    s.mapResponseRole(pbResponse),
	}
}

func (s *roleProtoMapper) ToProtoResponsesRole(status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsesRole {
	return &pb.ApiResponsesRole{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesRole(pbResponse),
	}
}

func (s *roleProtoMapper) ToProtoResponsePaginationRole(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsePaginationRole {
	return &pb.ApiResponsePaginationRole{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesRole(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *roleProtoMapper) ToProtoResponsePaginationRoleDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponseDeleteAt) *pb.ApiResponsePaginationRoleDeleteAt {
	return &pb.ApiResponsePaginationRoleDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesRoleDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *roleProtoMapper) mapResponseRole(role *response.RoleResponse) *pb.RoleResponse {
	return &pb.RoleResponse{
		Id:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (s *roleProtoMapper) mapResponsesRole(roles []*response.RoleResponse) []*pb.RoleResponse {
	var responseRoles []*pb.RoleResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRole(role))
	}

	return responseRoles
}

func (s *roleProtoMapper) mapResponseRoleDeleteAt(role *response.RoleResponseDeleteAt) *pb.RoleResponseDeleteAt {
	return &pb.RoleResponseDeleteAt{
		Id:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt,
	}
}

func (s *roleProtoMapper) mapResponsesRoleDeleteAt(roles []*response.RoleResponseDeleteAt) []*pb.RoleResponseDeleteAt {
	var responseRoles []*pb.RoleResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRoleDeleteAt(role))
	}

	return responseRoles
}
