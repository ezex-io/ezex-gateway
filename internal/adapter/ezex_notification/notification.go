package ezex_notification

import (
	"github.com/ezex-io/ezex-gateway/internal/port"
	client "github.com/ezex-io/ezex-proto/go/notification"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ port.NotificationPort = &Notification{}

type Notification struct {
	conn               *grpc.ClientConn
	notificationClient client.NotificationServiceClient
}

func New(cfg *Config) (*Notification, error) {
	conn, err := grpc.NewClient(cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Notification{
		conn:               conn,
		notificationClient: client.NewNotificationServiceClient(conn),
	}, nil
}

func (a *Notification) Close() error {
	return a.conn.Close()
}

func (a *Notification) SendEmail(ctx context.Context, recipient, subject, template string,
	fields map[string]string,
) error {
	_, err := a.notificationClient.SendEmail(ctx, &client.SendEmailRequest{
		Recipient:      recipient,
		Subject:        subject,
		TemplateName:   template,
		TemplateFields: fields,
	})
	if err != nil {
		return err
	}

	return nil
}
