package model

import (
	"time"

	"github.com/google/uuid"
)

type AuthEntity struct {
	Id        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
}
