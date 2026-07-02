package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID
	Version   int
	UserID    uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type CategoryPatch struct {
	Name *string
}
