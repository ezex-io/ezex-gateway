package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gen"
	"github.com/ezex-io/ezex-gateway/internal/entity"
	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/ezex-gateway/internal/utils"
	"github.com/ezex-io/gopkg/logger"
)

type Auth struct {
	notificationPort port.NotificationPort
	usersPort        port.UserPort
	redisPort        port.CachePort
	firebasePort     port.FirebasePort

	cfg     *Config
	logging logger.Logger
}

func NewAuth(cfg *Config, logging logger.Logger,
	notificationPort port.NotificationPort, redisPort port.CachePort,
	firebasePort port.FirebasePort,
	usersPort port.UserPort,
) *Auth {
	return &Auth{
		notificationPort: notificationPort,
		usersPort:        usersPort,
		redisPort:        redisPort,
		firebasePort:     firebasePort,
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

func (a *Auth) ProcessFirebaseLogin(ctx context.Context, idToken string) (string, error) {
	token, err := a.firebasePort.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}

	firebaseUID := token.UID
	email, ok := token.Claims["email"]
	if !ok {
		return "", errors.New("no email claim found from firebase")
	}

	emailStr, ok := email.(string)
	if !ok {
		return "", errors.New("invalid email claim found from firebase")
	}

	return a.usersPort.ProcessFirebaseLogin(ctx, emailStr, firebaseUID)
}

func (a *Auth) SaveSecurityImage(ctx context.Context, req *entity.SaveSecurityImageReq) error {
	return a.usersPort.SaveSecurityImage(ctx, req.Email, req.Image, req.Phrase)
}

func (a *Auth) GetSecurityImage(ctx context.Context,
	req *entity.GetSecurityImageReq,
) (*entity.GetSecurityImageResp, error) {
	image, phrase, err := a.usersPort.GetSecurityImage(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &entity.GetSecurityImageResp{
		Image:  image,
		Phrase: phrase,
	}, nil
}
