package usecase

import (
	"context"
	"fmt"
	"log"
	"poc-ddd/application/domain/user/param"
	"poc-ddd/application/infra/helper"
	"poc-ddd/application/infra/pkg/response"
)

type Usecase struct {
	repo        Repository
	authAdapter AuthRepositoryAdapter
}

// CreateProfile implements infrahttp.Usecase
func (u *Usecase) CreateProfile(ctx context.Context, req *param.CreateProfileRequest) *response.CustomError {
	profile := req.ParseToModel()
	err := u.repo.CreateProfile(ctx, profile)
	if err != nil {
		return response.CreateProfileError()
	}
	return nil
}

// GetProfile implements infrahttp.Usecase
func (u *Usecase) GetProfile(ctx context.Context, authId string) (*param.GetProfileResponse, *response.CustomError) {
	profile, err := u.repo.GetProfile(ctx, authId)
	if err != nil {
		if err == helper.ERR_NotFound {
			return nil, response.GetProfileError().NotFound().WithMessage("error when try to get profile").WithInfo("GetProfile", err.Error())
		}
		return nil, response.GetProfileError().InternalServerError().WithMessage("error when try to get profile").WithInfo("GetProfile", err.Error())
	}

	auth, err := u.authAdapter.GetAuthById(ctx, profile.AuthId)
	if err != nil {
		if err == helper.ERR_NotFound {
			return nil, response.GetProfileError().NotFound().WithMessage("error when try to get auth with auth id").WithInfo("GetProfile", err.Error())
		}
		return nil, response.GetProfileError().InternalServerError().WithMessage("error when try to get auth with auth id").WithInfo("GetProfile", err.Error())
	}

	resp := &param.GetProfileResponse{}
	resp = resp.MakeResponse(profile)
	resp.Email = auth.Email

	fmt.Printf("%+v\n%+v\n", resp, profile)

	return resp, nil
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) SetRepo(repo Repository) *Usecase {
	u.repo = repo
	return u
}
func (u *Usecase) SetAuthAdapter(authAdapter AuthRepositoryAdapter) *Usecase {
	u.authAdapter = authAdapter
	return u
}

func (u *Usecase) Build() *Usecase {
	if u.authAdapter == nil {
		log.Fatal("auth adapter is not set in user usecase")
	}

	if u.repo == nil {
		log.Fatal("repository is not set in user usecase")
	}

	return u
}
