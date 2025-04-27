package auth

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/ezex-gateway/internal/utils"
	gen "github.com/ezex-io/ezex-gateway/pkg/graphql"
	"github.com/ezex-io/gopkg/logger"
)

type Auth struct {
	notificationPort port.NotificationPort
	redisPort        port.CachePort

	cfg     *Config
	logging logger.Logger
}

func NewAuth(cfg *Config, logging logger.Logger,
	notificationPort port.NotificationPort, redisPort port.CachePort,
) *Auth {
	return &Auth{
		notificationPort: notificationPort,
		redisPort:        redisPort,
		cfg:              cfg,
		logging:          logging,
	}
}

func (a *Auth) SendConfirmationCode(ctx context.Context, recipient string, method gen.DeliveryMethod) error {
	ok, err := a.redisPort.Exists(ctx, recipient)
	if ok && err == nil {
		return ErrConfirmationCodeAlreadySent
	}

	code := utils.GenerateRandomCode(6)

	switch method {
	case gen.DeliveryMethodEmail:
		if err := a.notificationPort.SendEmail(ctx, recipient,
			fmt.Sprintf(a.cfg.ConfirmationCodeSubject, code),
			a.cfg.ConfirmationTemplateName,
			map[string]string{
				"Code": code,
			},
		); err != nil {
			return err
		}

		return a.redisPort.Set(ctx, recipient, code, port.CacheWithTTL(a.cfg.ConfirmationCodeTTL))
	default:
		return fmt.Errorf("unknown delivery method: %s", method)
	}
}

func (a *Auth) VerifyConfirmationCode(ctx context.Context, recipient, code string) error {
	v, err := a.redisPort.Get(ctx, recipient)
	if err != nil {
		return ErrConfirmationCodeExpired
	}

	if v != code {
		return ErrConfirmationCodeIsInvalid
	}

	go func() {
		if err := a.redisPort.Del(ctx, recipient); err != nil {
			a.logging.Error("failed to delete recipient confirmation code from redis",
				"recipient", recipient, "err", err)
		}
	}()

	return nil
}
