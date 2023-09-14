package domain

import (
	"context"
)

type CUDRepository interface {
	Create(ctx context.Context, userId string, tokenId string, token []byte) (int, error)
	Update(ctx context.Context, userId string, tokenId string, token []byte) (int, error)
	Delete(ctx context.Context, userId string, tokenId string) (int, error)
}
