package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gateway"
	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/ezex-gateway/internal/utils"
	"github.com/ezex-io/ezex-proto/go/notification"
	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/gopkg/logger"
)

type Auth struct {
	notificationPort  port.NotificationPort
	usersPort         port.UsersPort
	redisPort         port.CachePort
	authenticatorPort port.AuthenticatorPort

	cfg     *Config
	logging logger.Logger
}

func NewAuth(cfg *Config, logging logger.Logger,
	notificationPort port.NotificationPort, redisPort port.CachePort,
	authenticatorPort port.AuthenticatorPort,
	usersPort port.UsersPort,
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

func (a *Auth) SendConfirmationCode(ctx context.Context, inp *gateway.SendConfirmationCodeInput) (
	*gateway.SendConfirmationCodePayload, error,
) {
	ok, err := a.redisPort.Exists(ctx, inp.Recipient)
	if ok && err == nil {
		return nil, ErrConfirmationCodeAlreadySent
	}

	code := utils.GenerateRandomCode(6)

	switch inp.Method {
	case gateway.DeliveryMethodEmail:
		req := &notification.SendTemplatedEmailRequest{
			Recipient:    inp.Recipient,
			Subject:      fmt.Sprintf(a.cfg.ConfirmationCodeSubject, code),
			TemplateName: a.cfg.ConfirmationTemplateName,
			TemplateFields: map[string]string{
				"Code": code,
			},
		}

		res, err := a.notificationPort.SendTemplatedEmail(ctx, req)
		if err != nil {
			return nil, err
		}

		err = a.redisPort.Set(ctx, inp.Recipient, code, port.CacheWithTTL(a.cfg.ConfirmationCodeTTL))
		if err != nil {
			return nil, err
		}

		return &gateway.SendConfirmationCodePayload{
			Recipient: res.Recipient,
		}, nil

	default:
		return nil, UnknownDeliveryMethodError{Method: inp.Method}
	}
}

func (a *Auth) VerifyConfirmationCode(ctx context.Context, inp *gateway.VerifyConfirmationCodeInput) (
	*gateway.VerifyConfirmationCodePayload, error,
) {
	v, err := a.redisPort.Get(ctx, inp.Recipient)
	if err != nil {
		return nil, ErrConfirmationCodeExpired
	}

	if v != inp.Code {
		return nil, ErrConfirmationCodeIsInvalid
	}

	if err := a.redisPort.Del(ctx, inp.Recipient); err != nil {
		a.logging.Error("failed to delete recipient confirmation code from redis",
			"recipient", inp.Recipient, "err", err)
	}

	return &gateway.VerifyConfirmationCodePayload{
		Recipient: inp.Recipient,
	}, nil
}

func (a *Auth) ProcessAuthToken(ctx context.Context, inp *gateway.ProcessAuthTokenInput) (
	*gateway.ProcessAuthTokenPayload, error,
) {
	verifyRes, err := a.authenticatorPort.VerifyIDToken(ctx, &port.VerifyIDTokenRequest{
		IDToken: inp.IDToken,
	})
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

	res, err := a.usersPort.CreateUser(ctx, &users.CreateUserRequest{
		Email:       emailStr,
		FirebaseUid: firebaseUID,
	})
	if err != nil {
		return nil, err
	}

	return &gateway.ProcessAuthTokenPayload{
		UserID: res.UserId,
	}, nil
}

func (a *Auth) SaveSecurityImage(ctx context.Context, inp *gateway.SaveSecurityImageInput) (
	*gateway.SaveSecurityImagePayload, error,
) {
	req := &users.SaveSecurityImageRequest{
		Email:          inp.Email,
		SecurityImage:  inp.SecurityImage,
		SecurityPhrase: inp.SecurityPhrase,
	}
	res, err := a.usersPort.SaveSecurityImage(ctx, req)
	if err != nil {
		return nil, err
	}

	return &gateway.SaveSecurityImagePayload{
		Email: res.Email,
	}, nil
}

func (a *Auth) GetSecurityImage(ctx context.Context,
	inp *gateway.GetSecurityImageInput,
) (*gateway.GetSecurityImagePayload, error) {
	req := &users.GetSecurityImageRequest{
		Email: inp.Email,
	}
	res, err := a.usersPort.GetSecurityImage(ctx, req)
	if err != nil {
		return nil, err
	}

	return &gateway.GetSecurityImagePayload{
		Email:          res.Email,
		SecurityImage:  res.SecurityImage,
		SecurityPhrase: res.SecurityPhrase,
	}, nil
}
