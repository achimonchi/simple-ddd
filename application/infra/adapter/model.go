package adapter

import (
	"time"

	"github.com/google/uuid"
)

type AuthAdapterEntity struct {
	Id        uuid.UUID
	Email     string
	CreatedAt time.Time
}
