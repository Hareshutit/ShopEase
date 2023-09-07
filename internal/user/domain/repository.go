package domain

import (
	"context"

	"github.com/google/uuid"
)

type CUDRepository interface {
	Create(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type RRepository interface {
	Get(ctx context.Context, id uuid.UUID) (User, error)
	FindById(ctx context.Context, id uuid.UUID) (User, error)
	CheckUser(ctx context.Context, login string, password string) string
}
