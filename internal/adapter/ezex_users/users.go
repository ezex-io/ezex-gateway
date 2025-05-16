package ezex_users

import (
	"context"

	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/ezex-proto/go/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ port.UsersPort = &Users{}

type Users struct {
	conn        *grpc.ClientConn
	usersClient users.UsersServiceClient
}

func New(cfg *Config) (*Users, error) {
	conn, err := grpc.NewClient(cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Users{
		conn:        conn,
		usersClient: users.NewUsersServiceClient(conn),
	}, nil
}

func (u *Users) Close() {
	_ = u.conn.Close()
}

func (u *Users) CreateUser(ctx context.Context, req *users.CreateUserRequest) (
	*users.CreateUserResponse, error,
) {
	return u.usersClient.CreateUser(ctx, req)
}

func (u *Users) GetUserByEmail(ctx context.Context, req *users.GetUserByEmailRequest) (
	*users.GetUserByEmailResponse, error,
) {
	return u.usersClient.GetUserByEmail(ctx, req)
}

func (u *Users) SaveSecurityImage(ctx context.Context, req *users.SaveSecurityImageRequest) (
	*users.SaveSecurityImageResponse, error,
) {
	return u.usersClient.SaveSecurityImage(ctx, req)
}

func (u *Users) GetSecurityImage(ctx context.Context, req *users.GetSecurityImageRequest) (
	*users.GetSecurityImageResponse, error,
) {
	return u.usersClient.GetSecurityImage(ctx, req)
}
