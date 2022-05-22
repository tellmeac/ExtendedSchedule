package groups

import (
	"context"
	"tellmeac/extended-schedule/adapters/provider"
	"tellmeac/extended-schedule/domain/entity"
)

type IService interface {
	GetAllFaculties(ctx context.Context) ([]entity.FacultyInfo, error)
	GetFacultyGroups(ctx context.Context, facultyID string) ([]entity.GroupInfo, error)
}

func NewService(provider provider.IGroupsProvider) IService {
	return &service{provider: provider}
}

type service struct {
	provider provider.IGroupsProvider
}

func (s service) GetAllFaculties(ctx context.Context) ([]entity.FacultyInfo, error) {
	return s.provider.GetAllFaculties(ctx)
}

func (s service) GetFacultyGroups(ctx context.Context, facultyID string) ([]entity.GroupInfo, error) {
	return s.provider.GetByFaculty(ctx, facultyID)
}
