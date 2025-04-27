package ezex_notification

import (
	"github.com/ezex-io/ezex-gateway/internal/port"
	client "github.com/ezex-io/ezex-notification/pkg/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ port.NotificationPort = &Adapter{}

type Adapter struct {
	conn               *grpc.ClientConn
	notificationClient client.NotificationServiceClient
}

func New(cfg *Config) (*Adapter, error) {
	conn, err := grpc.NewClient(cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Adapter{
		conn:               conn,
		notificationClient: client.NewNotificationServiceClient(conn),
	}, nil
}

func (a *Adapter) Close() error {
	return a.conn.Close()
}

func (a *Adapter) SendEmail(ctx context.Context, recipient, subject, template string, fields map[string]string) error {
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
