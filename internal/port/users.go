package port

import (
	"context"
)

type ProcessLoginRequest struct {
	Email       string
	FirebaseUID string
}

type ProcessLoginResponse struct {
	UserID string
}

type SaveSecurityImageRequest struct {
	Email  string
	Image  string
	Phrase string
}

type SaveSecurityImageResponse struct {
	Email string
}

type GetSecurityImageRequest struct {
	Email string
}

type GetSecurityImageResponse struct {
	Image  string
	Phrase string
}

type UsersPort interface {
	ProcessLogin(ctx context.Context, req *ProcessLoginRequest) (*ProcessLoginResponse, error)
	SaveSecurityImage(ctx context.Context, req *SaveSecurityImageRequest) (*SaveSecurityImageResponse, error)
	GetSecurityImage(ctx context.Context, req *GetSecurityImageRequest) (*GetSecurityImageResponse, error)
}
