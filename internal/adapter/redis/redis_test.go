//go:build integration
// +build integration

package redis

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func setupRedis(t *testing.T) *Adapter {
	t.Helper()

	host := os.Getenv("EZEX_GATEWAY_REDIS_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("EZEX_GATEWAY_REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	cfg := &Config{
		Host:         host,
		Port:         mustAtoi(port),
		DB:           0,
		DialTimeout:  2 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolSize:     5,
		Protocol:     3,
	}

	r, err := New(cfg)
	assert.NoError(t, err)

	return r.(*Adapter)
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func TestRedis_SetGetDel(t *testing.T) {
	ctx := context.Background()
	r := setupRedis(t)

	t.Cleanup(func() {
		_ = r.Close()
	})

	key := "test:set"
	val := "hello"

	err := r.Set(ctx, key, val, 10*time.Second)
	assert.NoError(t, err)

	got, err := r.Get(ctx, key)
	assert.NoError(t, err)
	assert.Equal(t, val, got)

	err = r.Del(ctx, key)
	assert.NoError(t, err)

	_, err = r.Get(ctx, key)
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestRedis_Exists(t *testing.T) {
	ctx := context.Background()
	r := setupRedis(t)

	t.Cleanup(func() {
		_ = r.Close()
	})

	key := "test:exists"
	_ = r.Set(ctx, key, "yes", 10*time.Second)

	ok, err := r.Exists(ctx, key)
	assert.NoError(t, err)
	assert.True(t, ok)

	_ = r.Del(ctx, key)

	ok, err = r.Exists(ctx, key)
	assert.NoError(t, err)
	assert.False(t, ok)
}

func TestRedis_SetGetJSON(t *testing.T) {
	ctx := context.Background()
	r := setupRedis(t)

	t.Cleanup(func() {
		_ = r.Close()
	})

	key := "test:json"
	data := testStruct{Name: "ezex", Email: "foo@example.com"}

	err := r.SetJSON(ctx, key, data, 10*time.Second)
	assert.NoError(t, err)

	var out testStruct
	err = r.GetJSON(ctx, key, &out)
	assert.NoError(t, err)
	assert.Equal(t, data, out)

	_ = r.Del(ctx, key)

	err = r.GetJSON(ctx, key, &out)
	assert.ErrorIs(t, err, ErrNotFound)
}
