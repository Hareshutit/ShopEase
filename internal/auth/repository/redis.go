package repository

import (
	"context"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	auth *redis.Client
}

func (r *RedisRepository) Create(ctx context.Context,
	userId string, tokenId string, token []byte) (int, error) {

	err := r.auth.Set(ctx, userId+":"+tokenId, token, 30*24*time.Hour).Err() // Добавить конфигурацию времени
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (r *RedisRepository) Update(ctx context.Context,
	userId string, tokenId string, token []byte) (int, error) {

	var args redis.SetArgs
	args.KeepTTL = true
	args.Mode = "XX"
	err := r.auth.SetArgs(ctx, userId+":"+tokenId, token, args).Err()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (r *RedisRepository) Delete(ctx context.Context,
	userId string, tokenId string) (int, error) {

	err := r.auth.Del(ctx, userId+":"+tokenId).Err()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}