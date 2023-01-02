package usecase

import (
	"context"
	"poc-ddd/application/domain/auth/param"
	"poc-ddd/application/infra/helper"
	"poc-ddd/application/infra/pkg/crypto"
	"poc-ddd/application/infra/pkg/response"
	"poc-ddd/application/infra/pkg/token"
)

type Usecase struct {
	repo Repository
}

// Login implements domainhttp.Usecase
func (u *Usecase) Login(ctx context.Context, req *param.LoginRequest) (*param.LoginResponse, *response.CustomError) {
	auth, err := u.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		if err == helper.ERR_NotFound {
			return nil, response.LoginError().NotFound().WithMessage("error when try to find auth by email").WithInfo("Login", err.Error())
		}
		return nil, response.LoginError().BadRequest().WithMessage("error when try to find auth by email").WithInfo("Login", err.Error())
	}

	tok, err := token.GenerateToken(auth.Id.String())
	if err != nil {
		return nil, response.LoginError().BadRequest().WithMessage("error when try to find auth by email").WithInfo("Login", err.Error())

	}

	return &param.LoginResponse{
		Token: tok,
		Email: auth.Email,
	}, nil
}

// RegisterAccount implements domainhttp.Usecase
func (u *Usecase) RegisterAccount(ctx context.Context, req *param.RegisterRequest) *response.CustomError {
	passHash, err := crypto.GeneratePassword(req.Password)
	if err != nil {
		return response.RegisterError().BadRequest().WithMessage("error when try to generate password").WithInfo("RegisterAccount", err.Error())
	}
	auth := req.ParseToModel()
	auth.Password = passHash

	err = u.repo.Create(ctx, auth)
	if err != nil {
		return response.RegisterError().BadRequest().WithMessage("error when try to store to datastore").WithInfo("RegisterAccount", err.Error())
	}
	return nil
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}
