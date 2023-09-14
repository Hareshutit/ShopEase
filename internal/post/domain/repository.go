package domain

import (
	"context"

	"github.com/google/uuid"
)

type CUDRepository interface {
	Create(ctx context.Context, post Post) (int, error)
	Update(ctx context.Context, post Post) (int, error)
	IncrementViews(ctx context.Context, PostId uuid.UUID, UserId uuid.UUID) (int, error)
	Delete(ctx context.Context, id uuid.UUID) (int, error)
}

type RRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*Post, int, error)
	GetMiniObject(ctx context.Context, par Parameters) ([]Post, int, error)
}
