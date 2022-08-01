package faculty

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tellmeac/extended-schedule/userconfig/domain/values"
	"github.com/tellmeac/extended-schedule/userconfig/pkg/cache"
)

func NewCacheService(inner Service, facultyStore cache.Store[[]values.Faculty], groupsStore cache.Store[[]values.StudyGroup]) Service {
	return &cachedService{
		inner:        inner,
		facultyStore: facultyStore,
		groupsStore:  groupsStore,
	}
}

const (
	facultiesPrefix = "all-faculties"
	groupPrefix     = "faculty-groups-"
	expiration      = 30 * time.Minute
)

type cachedService struct {
	inner        Service
	facultyStore cache.Store[[]values.Faculty]
	groupsStore  cache.Store[[]values.StudyGroup]
}

func (s cachedService) GetByFaculty(ctx context.Context, facultyID string) ([]values.StudyGroup, error) {
	groups, err := s.groupsStore.Fetch(ctx, groupPrefix+facultyID)
	switch {
	case err != nil && !errors.Is(err, cache.ErrNotFound):
		log.Error().Err(err).Msg("failed to fetch faculties from cache")
	case err == nil:
		return groups, nil
	}

	groups, err = s.inner.GetByFaculty(ctx, facultyID)
	if err != nil {
		return nil, err
	}
	if err = s.groupsStore.Save(ctx, groupPrefix+facultyID, groups, cache.WithExpiration(expiration)); err != nil {
		log.Error().Err(err).Msg("failed to save faculty groups")
	}

	return groups, nil
}

func (s cachedService) GetAllFaculties(ctx context.Context) ([]values.Faculty, error) {
	faculties, err := s.facultyStore.Fetch(ctx, facultiesPrefix)
	switch {
	case err != nil && !errors.Is(err, cache.ErrNotFound):
		log.Error().Err(err).Msg("failed to fetch faculties from cache")
	case err == nil:
		return faculties, nil
	}

	faculties, err = s.inner.GetAllFaculties(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.facultyStore.Save(ctx, facultiesPrefix, faculties, cache.WithExpiration(expiration)); err != nil {
		log.Error().Err(err).Msg("failed to save faculty groups")
	}

	return faculties, nil
}
