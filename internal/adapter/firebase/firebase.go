package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/ezex-io/ezex-gateway/internal/port"
	"google.golang.org/api/option"
)

type Firebase struct {
	app  *firebase.App
	auth *auth.Client
}

func New(ctx context.Context, cfg *Config) (port.FirebasePort, error) {
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: cfg.ProjectID,
	}, option.WithAPIKey(cfg.APIKey))
	if err != nil {
		return nil, err
	}

	authCli, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &Firebase{
		app:  app,
		auth: authCli,
	}, nil
}

func (f *Firebase) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := f.auth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}
