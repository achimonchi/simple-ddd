package adapter

import (
	"context"
)

type AuthAdapter struct {
	repo AuthRepo
}

// GetAuthById implements usecase.AuthRepositoryAdapter
func (a *AuthAdapter) GetAuthById(ctx context.Context, authId string) (*AuthAdapterEntity, error) {
	auth, err := a.repo.GetAuthById(ctx, authId)
	if err != nil {
		return nil, err
	}

	var authAdapterEntity = &AuthAdapterEntity{
		Id:        auth.Id,
		Email:     auth.Email,
		CreatedAt: auth.CreatedAt,
	}
	return authAdapterEntity, nil
}

func NewAuthAdapter(repo AuthRepo) *AuthAdapter {
	return &AuthAdapter{
		repo: repo,
	}
}
