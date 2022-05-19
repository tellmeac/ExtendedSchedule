package repository

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/adapters/ent"
	"tellmeac/extended-schedule/adapters/ent/excludedlesson"
	"tellmeac/extended-schedule/adapters/ent/joinedgroups"
	"tellmeac/extended-schedule/adapters/ent/userinfo"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/domain/entity"
	"tellmeac/extended-schedule/domain/repository"
)

// NewEntUserConfigRepository создает репозиторий для пользовательской конфигурации.
func NewEntUserConfigRepository(client *ent.Client) repository.IUserConfigRepository {
	return &entUserConfigRepository{client: client}
}

type entUserConfigRepository struct {
	client *ent.Client
}

func (r entUserConfigRepository) Put(ctx context.Context, userIdentifier string) (aggregate.UserConfig, error) {
	userInfo, err := r.client.UserInfo.Create().SetEmail(userIdentifier).Save(ctx)
	if err != nil {
		return aggregate.UserConfig{}, err
	}

	return aggregate.UserConfig{
		UserIdentifier:  userInfo.Email,
		JoinedGroups:    nil,
		ExcludedLessons: nil,
	}, nil
}

func (r entUserConfigRepository) Get(ctx context.Context, userIdentifier string) (aggregate.UserConfig, error) {
	dbo, err := r.client.UserInfo.Query().Where(userinfo.EmailEqualFold(userIdentifier)).
		WithExcludedLessons().
		WithJoinedGroups().
		Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return aggregate.UserConfig{}, fmt.Errorf("user = %s: %w", userIdentifier, repository.ErrConfigNotFound)
	case err != nil:
		return aggregate.UserConfig{}, err
	default:
		return aggregate.UserConfig{
			UserIdentifier:  dbo.Email,
			JoinedGroups:    mapJoinedGroups(dbo.Edges.JoinedGroups),
			ExcludedLessons: mapExcludedLessons(dbo.Edges.ExcludedLessons),
		}, nil
	}
}

func mapJoinedGroups(groups []*ent.JoinedGroups) []entity.GroupInfo {
	if len(groups) == 0 {
		return nil
	}
	return groups[0].JoinedGroups
}

func mapExcludedLessons(excluded []*ent.ExcludedLesson) []entity.ExcludedLesson {
	var result = make([]entity.ExcludedLesson, 0, len(excluded))
	for _, e := range excluded {
		result = append(result, entity.ExcludedLesson{
			ID:       e.UserID,
			LessonID: e.LessonID,
			Position: e.Position,
			WeekDay:  e.Weekday,
		})
	}
	return result
}

func (r entUserConfigRepository) Update(ctx context.Context, userIdentifier string, desired aggregate.UserConfig) error {
	userInfo, err := r.client.UserInfo.Query().Where(userinfo.EmailEqualFold(userIdentifier)).Only(ctx)
	if err != nil {
		return err
	}

	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}

	if _, err := tx.ExcludedLesson.Delete().Where(excludedlesson.UserIDEQ(userInfo.ID)).Exec(ctx); err != nil {
		return rollback(tx, err)
	}

	for _, excluded := range desired.ExcludedLessons {
		_, err := tx.ExcludedLesson.Create().
			SetLessonID(excluded.LessonID).
			SetPosition(excluded.Position).
			SetWeekday(excluded.WeekDay).
			SetGroups(excluded.Groups).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	if err := tx.JoinedGroups.Update().
		Where(joinedgroups.UserIDEQ(userInfo.ID)).
		SetJoinedGroups(desired.JoinedGroups).
		Exec(ctx); err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func rollback(tx *ent.Tx, err error) error {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		err = fmt.Errorf("%w: %v", err, rollbackErr)
	}
	return err
}
