package model

import (
	"github.com/google/uuid"
)

type UserEntity struct {
	Id          uuid.UUID
	AuthId      string
	FullName    string
	Gender      string
	PhoneNumber string
}
