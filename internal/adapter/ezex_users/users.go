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

func (u *Users) ProcessFirebaseLogin(ctx context.Context, email, firebaseUID string) (string, error) {
	resp, err := u.usersClient.ProcessFirebaseLogin(ctx, &client.ProcessFirebaseLoginRequest{
		Email:          email,
		FirebaseUserId: firebaseUID,
	})
	if err != nil {
		return "", err
	}

	return resp.UserId, nil
}

func (u *Users) SaveSecurityImage(ctx context.Context, email, securityImage, securityPhrase string) error {
	_, err := u.usersClient.SaveSecurityImage(ctx, &client.SaveSecurityImageRequest{
		Email:          email,
		SecurityImage:  securityImage,
		SecurityPhrase: securityPhrase,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *Users) GetSecurityImage(ctx context.Context, email string) (image string, phrase string, err error) {
	resp, err := u.usersClient.GetSecurityImage(ctx, &client.GetSecurityImageRequest{
		Email: email,
	})
	if err != nil {
		return "", "", err
	}

	return resp.SecurityImage, resp.SecurityPhrase, nil
}
