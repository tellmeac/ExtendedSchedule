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
	GetByFaculty(ctx context.Context, facultyID string) ([]values.StudyGroup, error)
	GetAllFaculties(ctx context.Context) ([]values.Faculty, error)
}

func New(client *tsuschedule.Client) Service {
	return &service{client: client}
}

type service struct {
	client *tsuschedule.Client
}

func (s service) GetByFaculty(ctx context.Context, facultyID string) ([]values.StudyGroup, error) {
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

	var result = make([]values.StudyGroup, len(*dto.JSON200))
	lo.ForEach(*dto.JSON200, func(g tsuschedule.GroupInfo, i int) {
		result[i] = values.StudyGroup{
			ExternalID: g.Id,
			Name:       g.Name,
		}
	})
	return result, nil
}

func (s service) GetAllFaculties(ctx context.Context) ([]values.Faculty, error) {
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

	var result = make([]values.Faculty, len(*dto.JSON200))
	lo.ForEach(*dto.JSON200, func(faculty tsuschedule.FacultyInfo, i int) {
		result[i] = values.Faculty{
			ExternalID: faculty.Id,
			Name:       faculty.Name,
		}
	})
	return result, nil
}
