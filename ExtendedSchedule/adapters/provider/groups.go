package provider

import (
	"context"
	"errors"
	"github.com/samber/lo"
	"tellmeac/extended-schedule/adapters/client/tsuschedule"
	"tellmeac/extended-schedule/domain/entity"
)

// IGroupsProvider представляет провайдер для получения групп по факультетам.
type IGroupsProvider interface {
	GetByFaculty(ctx context.Context, facultyID string) ([]entity.GroupInfo, error)
	GetAllFaculties(ctx context.Context) ([]entity.FacultyInfo, error)
}

func NewGroupsProvider(client *tsuschedule.Client) IGroupsProvider {
	return &groupsProvider{client: client}
}

type groupsProvider struct {
	client *tsuschedule.Client
}

func (g groupsProvider) GetByFaculty(ctx context.Context, facultyID string) ([]entity.GroupInfo, error) {
	response, err := g.client.GetFacultiesFacultyIDGroups(ctx, facultyID)
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

	var result = make([]entity.GroupInfo, len(*dto.JSON200))
	lo.ForEach(*dto.JSON200, func(g tsuschedule.GroupInfo, i int) {
		result[i] = entity.GroupInfo{
			ID:   g.Id,
			Name: g.Name,
		}
	})
	return result, nil
}

func (g groupsProvider) GetAllFaculties(ctx context.Context) ([]entity.FacultyInfo, error) {
	response, err := g.client.GetFaculties(ctx)
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

	var result = make([]entity.FacultyInfo, len(*dto.JSON200))
	lo.ForEach(*dto.JSON200, func(faculty tsuschedule.FacultyInfo, i int) {
		result[i] = entity.FacultyInfo{
			ID:   faculty.Id,
			Name: faculty.Name,
		}
	})
	return result, nil
}

type cacheGroupsProvider struct {
	provider IGroupsProvider

	// TODO: cache cannot be invalidated
	cachedGroups    map[string][]entity.GroupInfo
	cachedFaculties []entity.FacultyInfo
}

func (c cacheGroupsProvider) GetByFaculty(ctx context.Context, facultyID string) ([]entity.GroupInfo, error) {
	_, ok := c.cachedGroups[facultyID]
	if !ok {
		groups, err := c.provider.GetByFaculty(ctx, facultyID)
		if err != nil {
			return nil, err
		}
		c.cachedGroups[facultyID] = groups
	}

	return c.cachedGroups[facultyID], nil
}

func (c cacheGroupsProvider) GetAllFaculties(ctx context.Context) ([]entity.FacultyInfo, error) {
	if c.cachedFaculties == nil {
		faculties, err := c.provider.GetAllFaculties(ctx)
		if err != nil {
			return nil, err
		}
		c.cachedFaculties = faculties
	}

	return c.cachedFaculties, nil
}

func MakeNewCachingGroupsProvider(client *tsuschedule.Client) IGroupsProvider {
	return &cacheGroupsProvider{
		provider: NewGroupsProvider(client),

		cachedGroups:    map[string][]entity.GroupInfo{},
		cachedFaculties: make([]entity.FacultyInfo, 0),
	}
}
