package ezex_users

import (
	"context"

	"github.com/ezex-io/ezex-gateway/internal/port"
	client "github.com/ezex-io/ezex-proto/go/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ port.UserPort = &Users{}

type Users struct {
	conn        *grpc.ClientConn
	usersClient client.UsersServiceClient
}

func New(cfg *Config) (*Users, error) {
	conn, err := grpc.NewClient(cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Users{
		conn:        conn,
		usersClient: client.NewUsersServiceClient(conn),
	}, nil
}

func (u *Users) Close() error {
	return u.conn.Close()
}

func (u *Users) ProcessLogin(ctx context.Context, req *port.ProcessLoginRequest) (
	*port.ProcessLoginResponse, error,
) {
	res, err := u.usersClient.ProcessFirebaseLogin(ctx, &client.ProcessFirebaseLoginRequest{
		Email:          req.Email,
		FirebaseUserId: req.FirebaseUID,
	})
	if err != nil {
		return nil, err
	}

	return &port.ProcessLoginResponse{
		UserID: res.UserId,
	}, nil
}

func (u *Users) SaveSecurityImage(ctx context.Context, req *port.SaveSecurityImageRequest) (
	*port.SaveSecurityImageResponse, error,
) {
	_, err := u.usersClient.SaveSecurityImage(ctx, &client.SaveSecurityImageRequest{
		Email:          req.Email,
		SecurityImage:  req.Image,
		SecurityPhrase: req.Phrase,
	})
	if err != nil {
		return nil, err
	}

	return &port.SaveSecurityImageResponse{
		Email: req.Email,
	}, nil
}

func (u *Users) GetSecurityImage(ctx context.Context, req *port.GetSecurityImageRequest) (
	*port.GetSecurityImageResponse, error,
) {
	res, err := u.usersClient.GetSecurityImage(ctx, &client.GetSecurityImageRequest{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &port.GetSecurityImageResponse{
		Image:  res.SecurityImage,
		Phrase: res.SecurityPhrase,
	}, nil
}
