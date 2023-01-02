package memory

import (
	"context"
	"database/sql"
	"fmt"
	"poc-ddd/application/domain/user/model"
	"strings"
)

type Memory struct{}

var profilesData = []*model.UserEntity{}

// CreateProfile implements usecase.Repository
func (*Memory) CreateProfile(ctx context.Context, profile *model.UserEntity) error {
	profilesData = append(profilesData, profile)
	for _, profile := range profilesData {
		divide := strings.Repeat("=", 10)
		fmt.Printf("%s [%s] %s\n", divide, profile.FullName, divide)
		fmt.Printf("%+v\n", profile)
		fmt.Printf("%s%s\n", divide, divide)
	}
	return nil
}

// GetProfile implements usecase.Repository
func (*Memory) GetProfile(ctx context.Context, authId string) (*model.UserEntity, error) {
	for _, profile := range profilesData {
		fmt.Println(authId, profile.AuthId, authId == profile.AuthId)
		if profile.AuthId == authId {
			return profile, nil
		}
	}
	return nil, sql.ErrNoRows
}

func NewRepoMemory() *Memory {
	return &Memory{}
}
