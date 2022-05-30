package additions

import (
	"context"
	"tellmeac/extended-schedule/adapters/provider"
	"tellmeac/extended-schedule/domain/entity"
	"time"
)

type IService interface {
	GetAllFaculties(ctx context.Context) ([]entity.FacultyInfo, error)
	GetFacultyGroups(ctx context.Context, facultyID string) ([]entity.GroupInfo, error)
	GetGroupLessons(ctx context.Context, groupID string) ([]provider.LessonInfo, error)
}

func NewService(groups provider.IGroupsProvider, lessons provider.ILessonInfoProvider) IService {
	return &service{
		groups:  groups,
		lessons: lessons,
	}
}

type service struct {
	groups  provider.IGroupsProvider
	lessons provider.ILessonInfoProvider
}

func (s service) GetGroupLessons(ctx context.Context, groupID string) ([]provider.LessonInfo, error) {
	start := time.Now().Add(-1 * 7 * 24 * time.Hour)
	end := time.Now().Add(7 * 24 * time.Hour)

	return s.lessons.GetLessons(ctx, groupID, start, end)
}

func (s service) GetAllFaculties(ctx context.Context) ([]entity.FacultyInfo, error) {
	return s.groups.GetAllFaculties(ctx)
}

func (s service) GetFacultyGroups(ctx context.Context, facultyID string) ([]entity.GroupInfo, error) {
	return s.groups.GetByFaculty(ctx, facultyID)
}
