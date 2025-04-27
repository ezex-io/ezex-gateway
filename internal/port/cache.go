package port

import (
	"context"
	"encoding/json"
	"time"
)

type CacheOption func(*CacheOptions)

type CacheOptions struct {
	// Time-To-Live for cached item.
	TTL time.Duration
}

func CacheWithTTL(ttl time.Duration) CacheOption {
	return func(opt *CacheOptions) { opt.TTL = ttl }
}

type CachePort interface {
	Set(ctx context.Context, key string, value string, opts ...CacheOption) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, key string) (bool, error)
}

type CacheWithJSON struct {
	CachePort
}

func (c CacheWithJSON) SetJSON(ctx context.Context, key string, val any, opts ...CacheOption) error {
	bz, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return c.Set(ctx, key, string(bz), opts...)
}

func (c CacheWithJSON) GetJSON(ctx context.Context, key string, out any) error {
	str, err := c.Get(ctx, key)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), out)
}
