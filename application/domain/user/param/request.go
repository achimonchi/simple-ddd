package param

import (
	"poc-ddd/application/domain/user/model"

	"github.com/google/uuid"
)

type CreateProfileRequest struct {
	AuthId      string
	FullName    string `json:"full_name"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
}

func (c *CreateProfileRequest) ParseToModel() *model.UserEntity {
	return &model.UserEntity{
		Id:          uuid.New(),
		AuthId:      c.AuthId,
		FullName:    c.FullName,
		Gender:      c.Gender,
		PhoneNumber: c.PhoneNumber,
	}
}
