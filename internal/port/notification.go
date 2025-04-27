package port

import "context"

type NotificationPort interface {
	SendEmail(ctx context.Context, recipient, subject, template string, fields map[string]string) error
}
