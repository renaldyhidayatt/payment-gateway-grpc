package recordmapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/record"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
)

type roleRecordMapper struct {
}

func NewRoleRecordMapper() *roleRecordMapper {
	return &roleRecordMapper{}
}

func (s *roleRecordMapper) ToRoleRecord(role *db.Role) *record.RoleRecord {
	deletedAt := role.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.RoleRecord{
		ID:        int(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *roleRecordMapper) ToRolesRecord(roles []*db.Role) []*record.RoleRecord {
	var result []*record.RoleRecord

	for _, role := range roles {
		result = append(result, s.ToRoleRecord(role))
	}

	return result
}
