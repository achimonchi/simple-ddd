package memory

import (
	"context"
	"poc-ddd/application/domain/auth/model"
	"poc-ddd/application/infra/helper"
)

var authData = []*model.AuthEntity{}

type RepoMemory struct{}

// FindByEmail implements usecase.Repository
func (*RepoMemory) FindByEmail(ctx context.Context, email string) (*model.AuthEntity, error) {
	for _, data := range authData {
		if data.Email == email {
			return data, nil
		}
	}

	return nil, helper.ERR_NotFound
}

// Create implements usecase.Repository
func (*RepoMemory) Create(ctx context.Context, auth *model.AuthEntity) error {
	authData = append(authData, auth)
	return nil
}

func NewRepoMemory() *RepoMemory {
	return &RepoMemory{}
}
