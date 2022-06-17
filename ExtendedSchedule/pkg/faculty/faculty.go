package faculty

import (
	"context"
	"errors"
	"github.com/samber/lo"
	"go.uber.org/fx"
	commonmodels "tellmeac/extended-schedule/common/models"
	"tellmeac/extended-schedule/pkg/tsuschedule"
)

var Module = fx.Options(fx.Provide(New))

// Manager provides methods to get faculties and groups.
type Manager interface {
	GetByFaculty(ctx context.Context, facultyID string) ([]commonmodels.GroupInfo, error)
	GetAllFaculties(ctx context.Context) ([]commonmodels.FacultyInfo, error)
}

func New(client *tsuschedule.Client) Manager {
	return &manager{client: client}
}

type manager struct {
	client *tsuschedule.Client
}

func (m manager) GetByFaculty(ctx context.Context, facultyID string) ([]commonmodels.GroupInfo, error) {
	response, err := m.client.GetFacultiesFacultyIDGroups(ctx, facultyID)
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

	var result = make([]commonmodels.GroupInfo, len(*dto.JSON200))
	lo.ForEach(*dto.JSON200, func(g tsuschedule.GroupInfo, i int) {
		result[i] = commonmodels.GroupInfo{
			ID:   g.Id,
			Name: g.Name,
		}
	})
	return result, nil
}

func (m manager) GetAllFaculties(ctx context.Context) ([]commonmodels.FacultyInfo, error) {
	response, err := m.client.GetFaculties(ctx)
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

	var result = make([]commonmodels.FacultyInfo, len(*dto.JSON200))
	lo.ForEach(*dto.JSON200, func(faculty tsuschedule.FacultyInfo, i int) {
		result[i] = commonmodels.FacultyInfo{
			ID:   faculty.Id,
			Name: faculty.Name,
		}
	})
	return result, nil
}
