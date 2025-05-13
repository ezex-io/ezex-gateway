package port

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

type FirebasePort interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}
