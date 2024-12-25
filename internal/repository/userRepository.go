package repository

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	recordmapper "MamangRust/paymentgatewaygrpc/internal/mapper/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"
	"database/sql"
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

	req := db.SearchUsersParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.SearchUsers(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find users: %w", err)
	}
	totalRecords := len(res)

	return r.mapping.ToUsersRecord(res), totalRecords, nil
}

func (r *userRepository) FindById(user_id int) (*record.UserRecord, error) {
	res, err := r.db.GetUserByID(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	return r.mapping.ToUserRecord(res), nil
}

func (r *userRepository) FindByActive() ([]*record.UserRecord, error) {
	res, err := r.db.GetActiveUsers(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	return r.mapping.ToUsersRecord(res), nil
}

func (r *userRepository) FindByTrashed() ([]*record.UserRecord, error) {
	res, err := r.db.GetTrashedUsers(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	return r.mapping.ToUsersRecord(res), nil
}

func (r *userRepository) SearchUsersByEmail(email string) ([]*record.UserRecord, error) {
	nullEmail := sql.NullString{
		String: email,
		Valid:  email != "",
	}

	res, err := r.db.SearchUsersByEmail(r.ctx, nullEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to search users by email '%s': %w", email, err)
	}

	users := r.mapping.ToUsersRecord(res)
	return users, nil
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

func (r *userRepository) DeleteUserPermanent(user_id int) error {
	err := r.db.DeleteUserPermanently(r.ctx, int32(user_id))

	if err != nil {
		return nil
	}

	return fmt.Errorf("failed to delete user: %w", err)
}
