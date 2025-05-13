package port

import "context"

type SendEmailRequest struct {
	Recipient string
	Subject   string
	Template  string
	Fields    map[string]string
}

type SendEmailResponse struct {
	Recipient string
}

type NotificationPort interface {
	SendEmail(ctx context.Context, req *SendEmailRequest) (*SendEmailResponse, error)
}
