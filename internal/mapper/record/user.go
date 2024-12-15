package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/postgres/schema"
)

type userRecordMapper struct {
}

func NewUserRecordMapper() *userRecordMapper {
	return &userRecordMapper{}
}

func (s *userRecordMapper) ToUserRecord(user *db.User) *record.UserRecord {
	var deletedAt *string

	if user.DeletedAt.Valid {
		formatedDeletedAt := user.DeletedAt.Time.Format("2006-01-02")

		deletedAt = &formatedDeletedAt
	}

	return &record.UserRecord{
		ID:        int(user.UserID),
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}

func (s *userRecordMapper) ToUsersRecord(users []*db.User) []*record.UserRecord {
	var userRecords []*record.UserRecord

	for _, user := range users {
		userRecords = append(userRecords, s.ToUserRecord(user))
	}

	return userRecords
}
