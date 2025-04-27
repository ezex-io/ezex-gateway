package resolver

import "context"

func (*queryResolver) Dummy(_ context.Context) (*string, error) {
	msg := "dummy"

	return &msg, nil
}
