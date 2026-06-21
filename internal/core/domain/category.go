package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID     uuid.UUID
	UserID uuid.UUID

	Name      string
	CreatedAt time.Time
}
