package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gen"
	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/ezex-gateway/internal/utils"
	"github.com/ezex-io/gopkg/logger"
)

type Auth struct {
	notificationPort  port.NotificationPort
	usersPort         port.UserPort
	redisPort         port.CachePort
	authenticatorPort port.AuthenticatorPort

	cfg     *Config
	logging logger.Logger
}

func NewAuth(cfg *Config, logging logger.Logger,
	notificationPort port.NotificationPort, redisPort port.CachePort,
	authenticatorPort port.AuthenticatorPort,
	usersPort port.UserPort,
) *Auth {
	return &Auth{
		notificationPort:  notificationPort,
		usersPort:         usersPort,
		redisPort:         redisPort,
		authenticatorPort: authenticatorPort,
		cfg:               cfg,
		logging:           logging,
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
		req := &port.SendEmailRequest{
			Recipient: recipient,
			Subject:   fmt.Sprintf(a.cfg.ConfirmationCodeSubject, code),
			Template:  a.cfg.ConfirmationTemplateName,
			Fields: map[string]string{
				"Code": code,
			},
		}

		_, err := a.notificationPort.SendEmail(ctx, req)
		if err != nil {
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

	if err := a.redisPort.Del(ctx, recipient); err != nil {
		a.logging.Error("failed to delete recipient confirmation code from redis",
			"recipient", recipient, "err", err)
	}

	return nil
}

func (a *Auth) ProcessLogin(ctx context.Context, req *port.VerifyIDTokenRequest) (
	*port.ProcessLoginResponse, error,
) {
	// TODO: A little bit Strange to me. @Javad Please double check.

	verifyRes, err := a.authenticatorPort.VerifyIDToken(ctx, req)
	if err != nil {
		return nil, err
	}

	firebaseUID := verifyRes.Token.UID
	email, ok := verifyRes.Token.Claims["email"]
	if !ok {
		return nil, errors.New("no email claim found from firebase")
	}

	emailStr, ok := email.(string)
	if !ok {
		return nil, errors.New("invalid email claim found from firebase")
	}

	return a.usersPort.ProcessLogin(ctx, &port.ProcessLoginRequest{
		Email:       emailStr,
		FirebaseUID: firebaseUID,
	})
}

func (a *Auth) SaveSecurityImage(ctx context.Context, req *port.SaveSecurityImageRequest) (
	*port.SaveSecurityImageResponse, error,
) {
	return a.usersPort.SaveSecurityImage(ctx, req)
}

func (a *Auth) GetSecurityImage(ctx context.Context,
	req *port.GetSecurityImageRequest,
) (*port.GetSecurityImageResponse, error) {
	res, err := a.usersPort.GetSecurityImage(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
