package faculty

import (
	"context"
	"errors"
	"github.com/samber/lo"
	"tellmeac/extended-schedule/domain/values"
	"tellmeac/extended-schedule/services/tsuschedule"
)

// Service provides methods to get faculties and groups.
type Service interface {
	GetByFaculty(ctx context.Context, facultyID string) ([]values.GroupInfo, error)
	GetAllFaculties(ctx context.Context) ([]values.FacultyInfo, error)
}

func New(client *tsuschedule.Client) Service {
	return &service{client: client}
}

type service struct {
	client *tsuschedule.Client
}

func (s service) GetByFaculty(ctx context.Context, facultyID string) ([]values.GroupInfo, error) {
	response, err := s.client.GetFacultiesFacultyIDGroups(ctx, facultyID)
	if err != nil {
		return nil, err
	}

	dto, err := tsuschedule.ParseGetFacultiesFacultyIDGroupsResponse(response)
	if err != nil {
		return nil, err
	}

	if dto == nil {
		return nil, errors.New("returned body json is nil")
	}

	var result = make([]values.GroupInfo, len(*dto.JSON200))
	lo.ForEach(*dto.JSON200, func(g tsuschedule.GroupInfo, i int) {
		result[i] = values.GroupInfo{
			ID:   g.Id,
			Name: g.Name,
		}
	})
	return result, nil
}

func (s service) GetAllFaculties(ctx context.Context) ([]values.FacultyInfo, error) {
	response, err := s.client.GetFaculties(ctx)
	if err != nil {
		return nil, err
	}

	dto, err := tsuschedule.ParseGetFacultiesResponse(response)
	if err != nil {
		return nil, err
	}

	if dto == nil {
		return nil, errors.New("returned body json is nil")
	}

	var result = make([]values.FacultyInfo, len(*dto.JSON200))
	lo.ForEach(*dto.JSON200, func(faculty tsuschedule.FacultyInfo, i int) {
		result[i] = values.FacultyInfo{
			ID:   faculty.Id,
			Name: faculty.Name,
		}
	})
	return result, nil
}
