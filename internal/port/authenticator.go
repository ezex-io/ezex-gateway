package port

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

type VerifyIDTokenRequest struct {
	IDToken string
}

type VerifyIDTokenResponse struct {
	Token *auth.Token
}

type AuthenticatorPort interface {
	VerifyIDToken(ctx context.Context, req *VerifyIDTokenRequest) (*VerifyIDTokenResponse, error)
}
