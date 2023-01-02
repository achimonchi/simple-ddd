package memory

import (
	"context"
	"database/sql"
	"poc-ddd/application/domain/auth/model"
	"poc-ddd/application/infra/helper"
)

var authData = []*model.AuthEntity{}

type RepoMemory struct{}

// GetAuthById implements usecase.Repository
func (*RepoMemory) GetAuthById(ctx context.Context, authId string) (*model.AuthEntity, error) {
	for _, data := range authData {
		if data.Id.String() == authId {
			return data, nil
		}
	}

	return nil, sql.ErrNoRows
}

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
