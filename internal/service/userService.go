package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	responsemapper "MamangRust/paymentgatewaygrpc/internal/mapper/response"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	"MamangRust/paymentgatewaygrpc/pkg/hash"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type userService struct {
	userRepository repository.UserRepository
	logger         *logger.Logger
	mapping        responsemapper.UserResponseMapper
	hashing        *hash.Hashing
}

func NewUserService(
	userRepository repository.UserRepository,
	logger *logger.Logger,
	mapper responsemapper.UserResponseMapper,
	hashing *hash.Hashing,
) *userService {
	return &userService{
		userRepository: userRepository,
		logger:         logger,
		mapping:        mapper,
		hashing:        hashing,
	}
}

func (ds *userService) FindAll(page int, pageSize int, search string) ([]*response.UserResponse, int, *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := ds.userRepository.FindAllUsers(search, page, pageSize)

	if err != nil {
		ds.logger.Error("failed to fetch users", zap.Error(err))
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch users",
		}
	}

	if len(users) == 0 {
		ds.logger.Error("no users found")
		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "No users found",
		}
	}

	userResponses := ds.mapping.ToUsersResponse(users)

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return userResponses, totalPages, nil
}

func (ds *userService) FindByID(id int) (*response.UserResponse, *response.ErrorResponse) {
	user, err := ds.userRepository.FindById(id)
	if err != nil {
		ds.logger.Error("failed to find user by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	so := ds.mapping.ToUserResponse(*user)

	return so, nil
}

func (s *userService) FindByActive() ([]*response.UserResponse, *response.ErrorResponse) {
	res, err := s.userRepository.FindByActive()

	if err != nil {
		s.logger.Error("Failed to find active users", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find active users",
		}
	}

	return s.mapping.ToUsersResponse(res), nil
}

func (s *userService) FindByTrashed() ([]*response.UserResponse, *response.ErrorResponse) {
	res, err := s.userRepository.FindByTrashed()

	if err != nil {
		s.logger.Error("Failed to find trashed users", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find trashed users",
		}
	}

	return s.mapping.ToUsersResponse(res), nil
}

func (s *userService) CreateUser(request requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse) {
	existingUser, _ := s.userRepository.FindByEmail(request.Email)

	if existingUser != nil {
		s.logger.Error("Email is already in use", zap.String("email", request.Email))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Email is already in use",
		}
	}

	hash, err := s.hashing.HashPassword(request.Password)
	if err != nil {
		s.logger.Error("Failed to hash password", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to hash password",
		}
	}

	request.Password = hash

	res, err := s.userRepository.CreateUser(request)
	if err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create user",
		}
	}

	so := s.mapping.ToUserResponse(*res)

	return so, nil
}

func (s *userService) UpdateUser(request requests.UpdateUserRequest) (*response.UserResponse, *response.ErrorResponse) {
	existingUser, err := s.userRepository.FindById(request.UserID)
	if err != nil {
		s.logger.Error("Failed to find user by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	if request.Email != "" && request.Email != existingUser.Email {
		duplicateUser, _ := s.userRepository.FindByEmail(request.Email)
		

		if duplicateUser != nil {
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Email is already in use",
			}
		}

		existingUser.Email = request.Email
	}

	if request.Password != "" {
		hash, err := s.hashing.HashPassword(request.Password)
		if err != nil {
			s.logger.Error("Failed to hash password", zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Failed to hash password",
			}
		}
		existingUser.Password = hash
	}

	res, err := s.userRepository.UpdateUser(request)
	if err != nil {
		s.logger.Error("Failed to update user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update user",
		}
	}

	so := s.mapping.ToUserResponse(*res)

	return so, nil
}

func (s *userService) TrashedUser(user_id int) (*response.UserResponse, *response.ErrorResponse) {
	s.logger.Debug("Trashing user", zap.Int("user_id", user_id))

	res, err := s.userRepository.TrashedUser(user_id)

	if err != nil {
		s.logger.Error("Failed to trash user", zap.Error(err), zap.Int("user_id", user_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash user",
		}
	}

	so := s.mapping.ToUserResponse(*res)

	s.logger.Debug("Successfully trashed user", zap.Int("user_id", user_id))

	return so, nil
}

func (s *userService) RestoreUser(user_id int) (*response.UserResponse, *response.ErrorResponse) {
	s.logger.Debug("Restoring user", zap.Int("user_id", user_id))

	res, err := s.userRepository.RestoreUser(user_id)
	if err != nil {
		s.logger.Error("Failed to restore user", zap.Error(err), zap.Int("user_id", user_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore user",
		}
	}

	so := s.mapping.ToUserResponse(*res)

	s.logger.Debug("Successfully restored user", zap.Int("user_id", user_id))

	return so, nil
}

func (s *userService) DeleteUserPermanent(user_id int) (interface{}, *response.ErrorResponse) {
	s.logger.Debug("Deleting user permanently", zap.Int("user_id", user_id))

	err := s.userRepository.DeleteUserPermanent(user_id)
	if err != nil {
		s.logger.Error("Failed to delete user permanently", zap.Error(err), zap.Int("user_id", user_id))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete user permanently",
		}
	}

	s.logger.Debug("Successfully deleted user permanently", zap.Int("user_id", user_id))

	return nil, nil
}
