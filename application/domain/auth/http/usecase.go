package domainhttp

import (
	"context"
	"poc-ddd/application/domain/auth/param"
	"poc-ddd/application/infra/pkg/response"
)

// interface should be declared
// in package that will used the interface
type Usecase interface {
	RegisterAccount(ctx context.Context, req *param.RegisterRequest) *response.CustomError
	Login(ctx context.Context, req *param.LoginRequest) (*param.LoginResponse, *response.CustomError)
}
