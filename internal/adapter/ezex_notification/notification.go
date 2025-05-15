package ezex_notification

import (
	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/ezex-proto/go/notification"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ port.NotificationPort = &Notification{}

type Notification struct {
	conn               *grpc.ClientConn
	notificationClient notification.NotificationServiceClient
}

func New(cfg *Config) (*Notification, error) {
	conn, err := grpc.NewClient(cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Notification{
		conn:               conn,
		notificationClient: notification.NewNotificationServiceClient(conn),
	}, nil
}

func (a *Notification) Close() error {
	return a.conn.Close()
}

func (a *Notification) SendTemplatedEmail(ctx context.Context, req *notification.SendTemplatedEmailRequest) (
	*notification.SendTemplatedEmailResponse, error,
) {
	return a.notificationClient.SendTemplatedEmail(ctx, req)
}
