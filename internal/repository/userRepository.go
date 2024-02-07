package repository

import (
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
	"context"
	"errors"
	"fmt"
)

type userRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewUserRepository(db *db.Queries, ctx context.Context) *userRepository {
	return &userRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *userRepository) FindAll() ([]*db.User, error) {

	user, err := r.db.GetAllUsers(r.ctx)

	if err != nil {
		return nil, errors.New("failed get users")
	}

	return user, nil
}

func (r *userRepository) FindById(id int) (*db.User, error) {
	user, err := r.db.GetUserById(r.ctx, int32(id))

	if err != nil {
		return nil, errors.New("failed get user")
	}

	return user, nil
}

func (r *userRepository) Create(input *db.CreateUserParams) (*db.User, error) {
	var userRequest db.CreateUserParams

	userRequest.Firstname = input.Firstname
	userRequest.Lastname = input.Lastname
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	user, err := r.db.CreateUser(r.ctx, userRequest)

	if err != nil {
		return nil, errors.New("failed create user")
	}

	return user, nil
}

func (r *userRepository) Update(input *db.UpdateUserParams) (*db.User, error) {
	var userRequest db.UpdateUserParams

	userRequest.Firstname = input.Firstname
	userRequest.Lastname = input.Lastname
	userRequest.Email = input.Email
	userRequest.Password = input.Password
	userRequest.NocTransfer = "as"

	res, err := r.db.UpdateUser(r.ctx, userRequest)

	if err != nil {
		return nil, errors.New("failed update user")
	}

	return res, nil

}

func (r *userRepository) Delete(id int) error {
	resid, err := r.db.GetUserById(r.ctx, int32(id))

	if err != nil {
		return errors.New("failed get user")
	}

	err = r.db.DeleteUser(r.ctx, resid.UserID)

	if err != nil {
		return fmt.Errorf("failed error")
	}

	return nil

}

func (r *userRepository) FindByEmail(email string) (*db.User, error) {
	res, err := r.db.GetUserByEmail(r.ctx, email)

	if err != nil {
		return nil, errors.New("failed get user")
	}

	return res, nil
}
