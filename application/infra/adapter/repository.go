package adapter

import (
	"context"
	"poc-ddd/application/domain/auth/model"
)

type AuthRepo interface {
	GetAuthById(ctx context.Context, authId string) (*model.AuthEntity, error)
}
