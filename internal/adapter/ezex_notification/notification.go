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

func (a *Notification) SendEmail(ctx context.Context, req *port.SendEmailRequest) (*port.SendEmailResponse, error) {
	_, err := a.notificationClient.SendEmail(ctx, &client.SendEmailRequest{
		Recipient:      req.Recipient,
		Subject:        req.Subject,
		TemplateName:   req.Template,
		TemplateFields: req.Fields,
	})
	if err != nil {
		return nil, err
	}

	return &port.SendEmailResponse{
		Recipient: req.Recipient,
	}, nil
}
