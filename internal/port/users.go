package port

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
)

type UsersPort interface {
	CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error)
	GetUserByEmail(ctx context.Context, req *users.GetUserByEmailRequest) (*users.GetUserByEmailResponse, error)
	SaveSecurityImage(ctx context.Context, req *users.SaveSecurityImageRequest) (*users.SaveSecurityImageResponse, error)
	GetSecurityImage(ctx context.Context, req *users.GetSecurityImageRequest) (*users.GetSecurityImageResponse, error)
}
