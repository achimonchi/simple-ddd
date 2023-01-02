package param

import (
	"poc-ddd/application/domain/auth/model"
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterRequest) ParseToModel() *model.AuthEntity {
	return &model.AuthEntity{
		Id:        uuid.New(),
		Email:     r.Email,
		CreatedAt: time.Now(),
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
