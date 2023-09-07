package domain

import (
	"context"

	"github.com/google/uuid"
)

type CUDRepository interface {
	Create(ctx context.Context, post Post) (int, error)
	Update(ctx context.Context, post Post) (int, error)
	Delete(ctx context.Context, id uuid.UUID) (int, error)
	AddFavorite(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (int, error)
	RemoveFavorite(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (int, error)
}

type RRepository interface {
	GetIdPost(ctx context.Context, id uuid.UUID) (*Post, int, error)
	GetMiniPostSortNew(ctx context.Context, par Parameters) ([]Post, int, error)
	GetFavorite(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, int, error)
	GetCart(ctx context.Context, postId []uuid.UUID) ([]Post, int, error)
	SearchPost(ctx context.Context, search string) ([]uuid.UUID, int, error)
}

type CUDCartRepository interface {
	Add(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (int, error)
	Remove(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (int, error)
}

type RCartRepository interface {
	GetUUID(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, int, error)
}
