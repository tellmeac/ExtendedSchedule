package cache

import (
	"context"
	"time"
)

type Options struct {
	Expiration *time.Duration
}

type Option func(*Options)

func WithExpiration(d time.Duration) Option {
	return func(o *Options) {
		o.Expiration = &d
	}
}

type Store[T any] interface {
	Fetch(ctx context.Context, key string) (T, error)
	Save(ctx context.Context, key string, val T, options ...Option) error
}
