package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type roleService struct {
	roleRepository repository.RoleRepository
	logger         logger.LoggerInterface
	mapping        responsemapper.RoleResponseMapper
}

func NewRoleService(roleRepository repository.RoleRepository, logger logger.LoggerInterface, mapping responsemapper.RoleResponseMapper) *roleService {
	return &roleService{
		roleRepository: roleRepository,
		logger:         logger,
		mapping:        mapping,
	}
}

func (s *roleService) FindAll(page int, pageSize int, search string) ([]*response.RoleResponse, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.roleRepository.FindAllRoles(page, pageSize, search)
	if err != nil {
		s.logger.Error("Failed to fetch role records", zap.Error(err))

		s.logger.Error("Failed to fetch role records", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch role records",
		}
	}

	if len(res) == 0 {
		s.logger.Debug("No role records found", zap.String("search", search))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No role records found",
		}
	}

	so := s.mapping.ToRolesResponse(res)

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return so, totalPages, nil
}

func (s *roleService) FindById(id int) (*response.RoleResponse, *response.ErrorResponse) {
	res, err := s.roleRepository.FindById(id)
	if err != nil {
		s.logger.Error("Failed to fetch role record by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch role record by ID",
		}
	}

	so := s.mapping.ToRoleResponse(res)

	return so, nil
}

func (s *roleService) FindByUserId(id int) ([]*response.RoleResponse, *response.ErrorResponse) {
	res, err := s.roleRepository.FindByUserId(id)

	if err != nil {
		s.logger.Error("Failed to fetch role record by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch role record by ID",
		}
	}

	so := s.mapping.ToRolesResponse(res)

	return so, nil
}

func (s *roleService) FindByActiveRole(page int, pageSize int, search string) ([]*response.RoleResponse, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.roleRepository.FindByActiveRole(page, pageSize, search)
	if err != nil {
		s.logger.Error("Failed to fetch role records", zap.Error(err))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch role records",
		}
	}

	if len(res) == 0 {
		s.logger.Debug("No role records found", zap.String("search", search))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No role records found",
		}
	}

	so := s.mapping.ToRolesResponse(res)

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return so, totalPages, nil
}

func (s *roleService) FindByTrashedRole(page int, pageSize int, search string) ([]*response.RoleResponse, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.roleRepository.FindByTrashedRole(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to fetch role records", zap.Error(err))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch role records",
		}
	}

	if len(res) == 0 {
		s.logger.Debug("No role records found", zap.String("search", search))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No role records found",
		}
	}

	so := s.mapping.ToRolesResponse(res)

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return so, totalPages, nil
}

func (s *roleService) CreateRole(request *requests.CreateRoleRequest) (*response.RoleResponse, *response.ErrorResponse) {
	res, err := s.roleRepository.CreateRole(request)
	if err != nil {
		s.logger.Error("Failed to create role", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create role record",
		}
	}

	so := s.mapping.ToRoleResponse(res)

	return so, nil
}

func (s *roleService) UpdateRole(request *requests.UpdateRoleRequest) (*response.RoleResponse, *response.ErrorResponse) {
	res, err := s.roleRepository.UpdateRole(request)

	if err != nil {
		s.logger.Error("Failed to update role", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update role record",
		}
	}

	so := s.mapping.ToRoleResponse(res)

	return so, nil
}

func (s *roleService) TrashedRole(id int) (*response.RoleResponse, *response.ErrorResponse) {
	res, err := s.roleRepository.TrashedRole(id)

	if err != nil {
		s.logger.Error("Failed to trashed role", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed role record",
		}
	}

	so := s.mapping.ToRoleResponse(res)

	return so, nil
}

func (s *roleService) RestoreRole(id int) (*response.RoleResponse, *response.ErrorResponse) {
	res, err := s.roleRepository.RestoreRole(id)

	if err != nil {
		s.logger.Error("Failed to restore role", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore role record",
		}
	}

	so := s.mapping.ToRoleResponse(res)

	return so, nil
}

func (s *roleService) DeleteRolePermanent(id int) (interface{}, *response.ErrorResponse) {
	err := s.roleRepository.DeleteRolePermanent(id)

	if err != nil {
		s.logger.Error("Failed to delete role", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete role record",
		}
	}

	return nil, nil
}
