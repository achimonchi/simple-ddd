package repository

import (
	"poc-ddd/application/domain/auth/repository/memory"
	"poc-ddd/application/domain/auth/usecase"
)

type RepoFactory struct{}

func NewRepoFactory() *RepoFactory {
	return &RepoFactory{}
}

const (
	REPO_Memory   = "memory"
	REPO_Postgres = "postgres"
)

func (r *RepoFactory) Create(repoType string) usecase.Repository {
	switch repoType {
	case REPO_Memory:
		return memory.NewRepoMemory()
	default:
		return nil
	}
}
