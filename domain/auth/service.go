package auth

import "context"

type authService struct {
}

func newService() authService {
	return authService{}
}

func (a authService) register(ctx context.Context, req register) (authItem Auth, err error) {
	return
}
