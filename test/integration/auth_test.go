//go:build integration
// +build integration

package integration

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ezex-io/ezex-gateway/internal/adapter/ezex_notification"
	"github.com/ezex-io/ezex-gateway/internal/adapter/redis"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/gopkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type authData struct {
	authInteractor *auth.Auth
	redisPort      port.RedisPort
	notifPort      port.NotificationPort
}

func setupAuth(t *testing.T) *authData {
	t.Helper()

	notifAddressEnv := os.Getenv("NOTIFICATION_ADDRESS")
	notifPortEnv := os.Getenv("NOTIFICATION_PORT")

	notifCfg := ezex_notification.DefaultConfig

	if notifAddressEnv != "" {
		notifCfg.Address = notifAddressEnv
	}

	if notifPortEnv != "" {
		p, err := strconv.Atoi(notifPortEnv)
		require.NoError(t, err)

		notifCfg.Port = p
	}

	notificationPort, err := ezex_notification.New(notifCfg)
	require.NoError(t, err)

	redisCfg := redis.DefaultConfig

	redisHostEnv := os.Getenv("REDIS_HOST")
	redisPortEnv := os.Getenv("REDIS_PORT")
	redisDBEnv := os.Getenv("REDIS_DB")
	redisPasswordEnv := os.Getenv("REDIS_PASSWORD")

	if redisHostEnv != "" {
		redisCfg.Host = redisHostEnv
	}

	if redisPortEnv != "" {
		p, err := strconv.Atoi(redisPortEnv)
		require.NoError(t, err)

		redisCfg.Port = p
	}

	if redisDBEnv != "" {
		db, err := strconv.Atoi(redisDBEnv)
		require.NoError(t, err)

		redisCfg.DB = db
	}

	if redisPasswordEnv != "" {
		redisCfg.Password = redisPasswordEnv
	}

	redisPort, err := redis.New(redisCfg)
	require.NoError(t, err)

	a := auth.NewAuth(auth.DefaultConfig, logger.DefaultSlog, notificationPort, redisPort)
	require.NotNil(t, a)

	return &authData{
		authInteractor: a,
		redisPort:      redisPort,
		notifPort:      notificationPort,
	}
}

func (a *authData) cleanup() func() {
	return func() {
		_ = a.redisPort.Close()
		_ = a.notifPort.Close()
	}
}

func TestAuth_SendConfirmationCode(t *testing.T) {
	a := setupAuth(t)

	t.Cleanup(a.cleanup())

	t.Run("TestSendConfirmationWithEmail", func(t *testing.T) {
		c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		recip := "staging@ezex.io"

		err := a.authInteractor.SendConfirmationCode(c, recip, gen.DeliveryMethodEmail)
		require.NoError(t, err)

		v, err := a.redisPort.Get(c, recip)
		require.NoError(t, err)
		assert.NotEmpty(t, v)

		err = a.redisPort.Del(c, recip)
		require.NoError(t, err)
	})

	t.Run("TestSendConfirmationWithAlreadySent", func(t *testing.T) {
		c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		recip := "staging@ezex.io"

		err := a.authInteractor.SendConfirmationCode(c, recip, gen.DeliveryMethodEmail)
		require.NoError(t, err)

		v, err := a.redisPort.Get(c, recip)
		require.NoError(t, err)
		assert.NotEmpty(t, v)

		err = a.authInteractor.SendConfirmationCode(c, recip, gen.DeliveryMethodEmail)
		require.ErrorIs(t, err, auth.ErrConfirmationCodeAlreadySent)
	})

	t.Run("TestSendConfirmationWithInvalidDeliveryMethod", func(t *testing.T) {
		c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		recip := "staging@ezex.io"
		method := gen.DeliveryMethod("unknown")

		err := a.authInteractor.SendConfirmationCode(c, recip, method)
		require.Error(t, err)
	})
}
