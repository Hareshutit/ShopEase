package repository

import (
	"context"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type CartRedisRepository struct {
	cart redis.Conn
}

func (r *CartRedisRepository) Add(ctx context.Context,
	userId uuid.UUID, postId uuid.UUID) (int, error) {

	_, err := r.cart.Do("sadd", userId.String(), postId.String())

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (r *CartRedisRepository) Remove(ctx context.Context,
	userId uuid.UUID, postId uuid.UUID) (int, error) {

	_, err := r.cart.Do("srem", userId.String(), postId.String())

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (r *CartRedisRepository) GetUUID(ctx context.Context,
	userId uuid.UUID) ([]uuid.UUID, int, error) {

	values, err := redis.Strings(r.cart.Do("smembers", userId.String()))

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var result []uuid.UUID
	for _, val := range values {
		postId, err := uuid.Parse(val)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		result = append(result, postId)
	}

	return result, http.StatusOK, nil
}
