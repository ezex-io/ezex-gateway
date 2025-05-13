package port

import "context"

type UserPort interface {
	ProcessFirebaseLogin(ctx context.Context, email, firebaseUID string) (userID string, err error)
	SaveSecurityImage(ctx context.Context, email, securityImage, securityPhrase string) error
	GetSecurityImage(ctx context.Context, email string) (securityImage string, securityPhrase string, err error)
}
