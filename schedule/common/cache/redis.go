package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

type redisStore[T any] struct {
	client *redis.Client
	codec  Codec
}

func NewStoreRedis[T any](client *redis.Client) Store[T] {
	return &redisStore[T]{
		client: client,
		codec:  DefaultCodec(),
	}
}

func (s redisStore[T]) Save(ctx context.Context, key string, val T, options ...Option) error {
	var opt Options
	for _, setter := range options {
		setter(&opt)
	}

	bytes, err := s.codec.Encode(val)
	if err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	var d time.Duration
	switch {
	case opt.Expiration == nil:
		d = redis.KeepTTL
	default:
		d = *opt.Expiration
	}

	statusCmd := s.client.Set(ctx, key, bytes, d)
	if statusCmd.Err() != nil {
		return fmt.Errorf("failed to save: %w", statusCmd.Err())
	}

	return nil
}

func (s redisStore[T]) Fetch(ctx context.Context, key string) (T, error) {
	var result T

	stringCmd := s.client.Get(ctx, key)
	switch {
	case errors.Is(stringCmd.Err(), redis.Nil):
		return result, ErrNotFound
	case stringCmd.Err() != nil:
		return result, stringCmd.Err()
	default:
		bytes, err := stringCmd.Bytes()
		if err != nil {
			return result, fmt.Errorf("failed to get value bytes: %w", err)
		}

		if err = s.codec.Decode(bytes, &result); err != nil {
			return result, fmt.Errorf("failed to decode bytes to value: %w", err)
		}

		return result, err
	}
}
