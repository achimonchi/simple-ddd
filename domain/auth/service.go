package auth

import "context"

type repository interface {
	writeRepository
	readRepository
}

type writeRepository interface {
	save(ctx context.Context, item Auth) (err error)
}

type readRepository interface {
}

type authService struct {
	repo repository
}

func newService(repo repository) authService {
	return authService{
		repo: repo,
	}
}

func (a authService) register(ctx context.Context, req Auth) (authItem Auth, err error) {
	if err = req.EncryptPassword(); err != nil {
		return
	}

	if err = a.repo.save(ctx, req); err != nil {
		return
	}

	authItem = req
	return
}
