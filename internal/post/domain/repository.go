package domain

import (
	"context"

	"github.com/google/uuid"
)

type CUDRepository interface {
	Create(ctx context.Context, post Post) (int, error)
	Update(ctx context.Context, post Post) (int, error)
	Delete(ctx context.Context, id uuid.UUID) (int, error)
}

type RRepository interface {
	GetIdPost(ctx context.Context, id uuid.UUID) (*Post, int, error)
	GetMiniPostSortNew(ctx context.Context, par Parameters) ([]Post, int, error)
}
