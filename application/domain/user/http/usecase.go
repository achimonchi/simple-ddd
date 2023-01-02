package infrahttp

import (
	"context"
	"poc-ddd/application/domain/user/param"
	"poc-ddd/application/infra/pkg/response"
)

type Usecase interface {
	CreateProfile(ctx context.Context, req *param.CreateProfileRequest) *response.CustomError
	GetProfile(ctx context.Context, authId string) (*param.GetProfileResponse, *response.CustomError)
}
