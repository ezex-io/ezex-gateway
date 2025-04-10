package resolver

import (
	"context"

	"github.com/ezex-io/ezex-gateway/api/graphql/gen"
)

func (m *mutationResolver) SendConfirmationCode(ctx context.Context,
	in gen.SendConfirmationCodeInput,
) (*gen.VoidPayload, error) {
	err := m.auth.SendConfirmationCode(ctx, in.Recipient, in.Method)
	if err != nil {
		return nil, err
	}

	return &gen.VoidPayload{}, nil
}

func (m *mutationResolver) VerifyConfirmationCode(ctx context.Context,
	in gen.VerifyConfirmationCodeInput,
) (*gen.VoidPayload, error) {
	err := m.auth.VerifyConfirmationCode(ctx, in.Recipient, in.Code)
	if err != nil {
		return nil, err
	}

	return &gen.VoidPayload{}, nil
}

func (m *mutationResolver) SecurityImage(ctx context.Context,
	input gen.SecurityImageInput) (*gen.SecurityImagePayload, error) {
	return &gen.SecurityImagePayload{
		Image: "mock.png",
	}, nil
}

func (m *mutationResolver) ProcessFirebaseAuth(ctx context.Context,
	input gen.ProcessFirebaseAuthInput) (*gen.VoidPayload, error) {
	return &gen.VoidPayload{}, nil
}
