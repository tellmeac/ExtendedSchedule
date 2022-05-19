package provider

import (
	"context"
	"tellmeac/extended-schedule/domain/entity"
)

// IGroupsProvider представляет провайдер для получения групп по факультетам.
type IGroupsProvider interface {
	GetByFaculty(ctx context.Context, facultyID string) ([]entity.GroupInfo, error)
	GetAllFaculties(ctx context.Context) ([]entity.FacultyInfo, error)
}
