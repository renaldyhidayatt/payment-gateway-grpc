package service

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/repository"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"MamangRust/paymentgatewaygrpc/pkg/hash"
	"errors"
)

type userService struct {
	user repository.UserRepository
	hash hash.Hashing
}

func NewUserService(user repository.UserRepository, hash hash.Hashing) *userService {
	return &userService{
		user: user,
		hash: hash,
	}
}

func (s *userService) FindAll() ([]*db.User, error) {
	res, err := s.user.FindAll()

	if err != nil {
		return nil, errors.New("failed get user")
	}

	return res, nil
}

func (s *userService) FindById(id int) (*db.User, error) {
	res, err := s.user.FindById(id)

	if err != nil {
		return nil, errors.New("failed get user")
	}

	return res, nil
}

func (s *userService) Create(input *requests.CreateUserRequest) (*db.User, error) {
	_, err := s.user.FindByEmail(input.Email)

	if err == nil {
		return nil, errors.New("failed email already exist")
	}

	passwordHash, err := s.hash.HashPassword(input.Password)

	if err != nil {
		return nil, errors.New("failed hash password")
	}

	user := db.CreateUserParams{
		Firstname: input.FirstName,
		Lastname:  input.LastName,
		Email:     input.Email,
		Password:  passwordHash,
	}

	res, err := s.user.Create(&user)

	if err != nil {
		return nil, errors.New("failed create user")
	}

	return res, nil
}

func (s *userService) Update(input *requests.UpdateUserRequest) (*db.User, error) {
	_, err := s.user.FindById(input.ID)

	if err != nil {
		return nil, errors.New("user not found")
	}

	user := db.UpdateUserParams{
		Firstname: input.FirstName,
		Lastname:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}

	res, err := s.user.Update(&user)

	if err != nil {
		return nil, errors.New("failed update user")
	}

	return res, nil
}

func (s *userService) Delete(id int) error {
	res, err := s.user.FindById(id)

	if err != nil {
		return errors.New("user not found")
	}

	err = s.user.Delete(int(res.UserID))

	if err != nil {
		return errors.New("failed delete user")
	}

	return nil
}
