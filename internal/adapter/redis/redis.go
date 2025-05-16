package redis

import (
	"context"
	"errors"

	"github.com/ezex-io/ezex-gateway/internal/port"
	redis "github.com/redis/go-redis/v9"
)

var _ port.CachePort = &Adapter{}

type Adapter struct {
	rdb *redis.Client
}

func New(cfg *Config) (*Adapter, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Address,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.PoolSize,
		Protocol:     cfg.Protocol,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &Adapter{rdb: rdb}, nil
}

func (a *Adapter) Close() {
	_ = a.rdb.Close()
}

func (a *Adapter) Set(ctx context.Context, key string, value string, opts ...port.CacheOption) error {
	options := &port.CacheOptions{}
	for _, opt := range opts {
		opt(options)
	}

	return a.rdb.Set(ctx, key, value, options.TTL).Err()
}

func (a *Adapter) Get(ctx context.Context, key string) (string, error) {
	val, err := a.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", ErrNotFound
	}

	return val, err
}

func (a *Adapter) Del(ctx context.Context, keys ...string) error {
	return a.rdb.Del(ctx, keys...).Err()
}

func (a *Adapter) Exists(ctx context.Context, key string) (bool, error) {
	n, err := a.rdb.Exists(ctx, key).Result()

	return n > 0, err
}
