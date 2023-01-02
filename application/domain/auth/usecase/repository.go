package usecase

import (
	"context"
	"poc-ddd/application/domain/auth/model"
)

// interface should be declared
// in package that will used the interface
type Repository interface {
	Create(ctx context.Context, auth *model.AuthEntity) error
	FindByEmail(ctx context.Context, email string) (*model.AuthEntity, error)
	GetAuthById(ctx context.Context, authId string) (*model.AuthEntity, error)
}
