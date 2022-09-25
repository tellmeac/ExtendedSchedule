package adapters

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"tellmeac/extended-schedule/common/tsu"
	"tellmeac/extended-schedule/common/utils"
	"tellmeac/extended-schedule/schedule"
	"tellmeac/extended-schedule/userconfig"
)

// NewTargetProvider возвращает стандартный провайдер списков групп и преподавателей.
func NewTargetProvider(client tsu.ClientWithResponsesInterface) TargetProvider {
	return TargetProvider{client: client}
}

type TargetProvider struct {
	client tsu.ClientWithResponsesInterface
}

func (fp TargetProvider) Teachers(ctx context.Context) ([]schedule.Teacher, error) {
	resp, err := fp.client.GetProfessorsWithResponse(ctx, utils.ApplyFakeUserAgent)
	switch {
	case err != nil:
		return nil, err
	case resp.StatusCode() != 200:
		return nil, fmt.Errorf("failed with status code = %d", resp.StatusCode())
	}

	return lo.Map(*resp.JSON200, func(t tsu.Teacher, _ int) schedule.Teacher {
		return schedule.Teacher{
			ID:   t.ID,
			Name: t.FullName,
		}
	}), nil
}

func (fp TargetProvider) Faculties(ctx context.Context) ([]schedule.Faculty, error) {
	resp, err := fp.client.GetFacultiesWithResponse(ctx, utils.ApplyFakeUserAgent)
	switch {
	case err != nil:
		return nil, err
	case resp.StatusCode() != 200:
		return nil, fmt.Errorf("failed with status code = %d", resp.StatusCode())
	}

	return lo.Map(*resp.JSON200, func(f tsu.Faculty, _ int) schedule.Faculty {
		return schedule.Faculty{
			ID:   f.ID,
			Name: f.Name,
		}
	}), nil
}

func (fp TargetProvider) GroupsByFaculty(ctx context.Context, id string) ([]userconfig.StudyGroup, error) {
	resp, err := fp.client.GetFacultiesIdGroupsWithResponse(ctx, id, utils.ApplyFakeUserAgent)
	switch {
	case err != nil:
		return nil, err
	case resp.StatusCode() != 200:
		return nil, fmt.Errorf("failed with status code = %d", resp.StatusCode())
	}

	return lo.Map(*resp.JSON200, func(g tsu.StudyGroup, _ int) userconfig.StudyGroup {
		return userconfig.StudyGroup{
			ID:   g.ID,
			Name: g.Name,
		}
	}), nil
}
