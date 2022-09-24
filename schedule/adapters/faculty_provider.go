package adapters

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"tellmeac/extended-schedule/common/tsu"
	"tellmeac/extended-schedule/schedule"
)

func NewFacultyProvider(client tsu.ClientWithResponsesInterface) FacultyProvider {
	return FacultyProvider{client: client}
}

type FacultyProvider struct {
	client tsu.ClientWithResponsesInterface
}

func (fp FacultyProvider) Faculties(ctx context.Context) ([]schedule.Faculty, error) {
	resp, err := fp.client.GetFacultiesWithResponse(ctx)
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
