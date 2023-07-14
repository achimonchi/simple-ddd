package auth

import (
	"context"
	"sync"
)

type memmory struct {
}

type repo struct {
	items []Auth
	mu    sync.Mutex
}

var items = repo{}

func newMemmory() memmory {
	return memmory{}
}

func (m memmory) save(ctx context.Context, item Auth) (err error) {
	items.mu.Lock()
	items.items = append(items.items, item)
	items.mu.Unlock()

	return nil
}

func (m memmory) findByEmail(ctx context.Context, email string) (auth Auth, err error) {
	for _, item := range items.items {
		if item.Email == email {
			return item, nil
		}
	}

	err = ErrEmailNotFound
	return
}
