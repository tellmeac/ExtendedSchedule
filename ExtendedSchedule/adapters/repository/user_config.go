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

func NewEntUserConfigRepository(client *ent.Client) repository.IUserConfigRepository {
	return &entUserConfigRepository{client: client}
}

// entUserConfigRepository реализует репозиторий для пользовательской конфигурации.
type entUserConfigRepository struct {
	client *ent.Client
}

func (repository entUserConfigRepository) Get(ctx context.Context, userIdentifier string) (aggregate.UserConfig, error) {
	dbo, err := repository.client.UserInfo.Query().Where(userinfo.EmailEqualFold(userIdentifier)).
		WithExcludedLessons().
		WithJoinedGroups().
		Only(ctx)
	if err != nil {
		return aggregate.UserConfig{}, err
	}

	return aggregate.UserConfig{
		UserIdentifier:  dbo.Email,
		JoinedGroups:    mapJoinedGroups(dbo.Edges.JoinedGroups),
		ExcludedLessons: mapExcluded(dbo.Edges.ExcludedLessons),
	}, nil
}

func mapJoinedGroups(groups []*ent.JoinedGroups) []entity.GroupInfo {
	if len(groups) == 0 {
		return nil
	}
	return groups[0].JoinedGroups
}

func mapExcluded(excluded []*ent.ExcludedLesson) []entity.ExcludedLesson {
	var result = make([]entity.ExcludedLesson, 0, len(excluded))
	for _, e := range excluded {
		result = append(result, entity.ExcludedLesson{
			ID:       e.UserID,
			LessonID: e.LessonID,
			Teacher:  e.Teacher,
			Position: e.Position,
			WeekDay:  e.Weekday,
		})
	}
	return result
}

func (repository entUserConfigRepository) Update(ctx context.Context, userIdentifier string, desired aggregate.UserConfig) error {
	userInfo, err := repository.client.UserInfo.Query().Where(userinfo.EmailEqualFold(userIdentifier)).Only(ctx)
	if err != nil {
		return err
	}

	tx, err := repository.client.Tx(ctx)
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
			SetTeacher(excluded.Teacher).
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
