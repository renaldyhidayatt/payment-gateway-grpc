package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"context"
	"errors"
	"fmt"
)

type userRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.UserRecordMapping
}

func NewUserRepository(db *db.Queries, ctx context.Context, mapping recordmapper.UserRecordMapping) *userRepository {
	return &userRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *userRepository) FindAllUsers(search string, page, pageSize int) ([]*record.UserRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetUsersWithPaginationParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetUsersWithPagination(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find users: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToUsersRecordPagination(res), totalCount, nil
}

func (r *userRepository) FindById(user_id int) (*record.UserRecord, error) {
	fmt.Printf("Searching for user with ID: %d\n", user_id)
	res, err := r.db.GetUserByID(r.ctx, int32(user_id))

	if err != nil {
		fmt.Printf("Error fetching user: %v\n", err)

		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	return r.mapping.ToUserRecord(res), nil
}

func (r *userRepository) FindByActive(search string, page, pageSize int) ([]*record.UserRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetActiveUsersWithPaginationParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetActiveUsersWithPagination(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find users: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToUsersRecordActivePagination(res), totalCount, nil
}

func (r *userRepository) FindByTrashed(search string, page, pageSize int) ([]*record.UserRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTrashedUsersWithPaginationParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTrashedUsersWithPagination(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find users: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToUsersRecordTrashedPagination(res), totalCount, nil
}

func (r *userRepository) FindByEmail(email string) (*record.UserRecord, error) {
	res, err := r.db.GetUserByEmail(r.ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	return r.mapping.ToUserRecord(res), nil
}

func (r *userRepository) CreateUser(request *requests.CreateUserRequest) (*record.UserRecord, error) {
	req := db.CreateUserParams{
		Firstname: request.FirstName,
		Lastname:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}

	user, err := r.db.CreateUser(r.ctx, req)

	if err != nil {
		return nil, errors.New("failed create user")
	}

	return r.mapping.ToUserRecord(user), nil
}

func (r *userRepository) UpdateUser(request *requests.UpdateUserRequest) (*record.UserRecord, error) {
	req := db.UpdateUserParams{
		UserID:    int32(request.UserID),
		Firstname: request.FirstName,
		Lastname:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}

	err := r.db.UpdateUser(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	res, err := r.db.GetUserByID(r.ctx, int32(request.UserID))

	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return r.mapping.ToUserRecord(res), nil
}

func (r *userRepository) TrashedUser(user_id int) (*record.UserRecord, error) {
	err := r.db.TrashUser(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash user: %w", err)
	}

	merchant, err := r.db.GetTrashedUserByID(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find trashed by id user: %w", err)
	}

	return r.mapping.ToUserRecord(merchant), nil
}

func (r *userRepository) RestoreUser(user_id int) (*record.UserRecord, error) {
	err := r.db.RestoreUser(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed to restore topup: %w", err)
	}

	user, err := r.db.GetUserByID(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed not found user :%w", err)
	}

	return r.mapping.ToUserRecord(user), nil
}

func (r *userRepository) DeleteUserPermanent(user_id int) (bool, error) {
	err := r.db.DeleteUserPermanently(r.ctx, int32(user_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}

	return true, nil
}

func (r *userRepository) RestoreAllUser() (bool, error) {
	err := r.db.RestoreAllUsers(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all users: %w", err)
	}
	return true, nil
}

func (r *userRepository) DeleteAllUserPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentUsers(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all users permanently: %w", err)
	}
	return true, nil
}
