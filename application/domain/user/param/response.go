package param

import (
	"poc-ddd/application/domain/user/model"

	"github.com/google/uuid"
)

type GetProfileResponse struct {
	Id          uuid.UUID `json:"id"`
	AuthId      string    `json:"auth_id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phone_number"`
}

func (p *GetProfileResponse) MakeResponse(model *model.UserEntity) *GetProfileResponse {
	p = &GetProfileResponse{
		Id:          model.Id,
		AuthId:      model.AuthId,
		FullName:    model.FullName,
		Gender:      model.Gender,
		PhoneNumber: model.PhoneNumber,
	}

	return p
}
