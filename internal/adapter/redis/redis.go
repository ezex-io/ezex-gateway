package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/redis/go-redis/v9"
)

var ErrNotFound = errors.New("key not found")

type Adapter struct {
	rdb *redis.Client
}

func New(cfg *Config) (port.RedisPort, error) {
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

func (a *Adapter) Close() error {
	return a.rdb.Close()
}

func (a *Adapter) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return a.rdb.Set(ctx, key, value, ttl).Err()
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

func (a *Adapter) SetJSON(ctx context.Context, key string, val any, ttl time.Duration) error {
	data, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	return a.rdb.Set(ctx, key, data, ttl).Err()
}

func (a *Adapter) GetJSON(ctx context.Context, key string, out any) error {
	data, err := a.rdb.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return ErrNotFound
	} else if err != nil {
		return fmt.Errorf("redis get failed: %w", err)
	}

	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("json unmarshal failed: %w", err)
	}

	return nil
}
