package repository

import (
	"context"
	"poc-ddd/application/infra/adapter"
)

type AuthRepo struct{}

// GetAuthById implements adapter.AuthRepo
// this repo will be used if we migrate it to microservices
func (*AuthRepo) GetAuthById(ctx context.Context, authId string) *adapter.AuthAdapterEntity {
	panic("unimplemented")
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}
