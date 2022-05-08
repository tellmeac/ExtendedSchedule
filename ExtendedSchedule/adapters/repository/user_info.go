package repository

import (
	"context"
	"tellmeac/extended-schedule/adapters/ent"
	"tellmeac/extended-schedule/domain/entity"
	"tellmeac/extended-schedule/domain/repository"
)

func NewEntUserInfoRepository(client *ent.Client) repository.IUserInfoRepository {
	return &userInfoRepository{client: client}
}

type userInfoRepository struct {
	client *ent.Client
}

func (repository userInfoRepository) GetByEmail(ctx context.Context, email string) (entity.UserInfo, error) {
	dbo, err := repository.client.UserInfo.Query().Where().Only(ctx)
	if err != nil {
		return entity.UserInfo{}, err
	}

	return entity.UserInfo{
		UserID: dbo.ID,
		Email:  dbo.Email,
	}, nil
}
