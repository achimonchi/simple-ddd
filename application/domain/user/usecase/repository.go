package usecase

import (
	"context"
	"poc-ddd/application/domain/user/model"
	"poc-ddd/application/infra/adapter"
)

type Repository interface {
	CreateProfile(ctx context.Context, profile *model.UserEntity) error
	GetProfile(ctx context.Context, authId string) (*model.UserEntity, error)
}

type AuthRepositoryAdapter interface {
	GetAuthById(ctx context.Context, authId string) (*adapter.AuthAdapterEntity, error)
}
