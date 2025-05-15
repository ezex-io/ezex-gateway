package port

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/notification"
)

type NotificationPort interface {
	SendTemplatedEmail(ctx context.Context, req *notification.SendTemplatedEmailRequest) (
		*notification.SendTemplatedEmailResponse, error)
}
