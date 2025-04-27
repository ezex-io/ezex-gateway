package resolver

import (
	"context"

	gen "github.com/ezex-io/ezex-gateway/pkg/graphql"
)

func (m *mutationResolver) SendConfirmationCode(ctx context.Context,
	inp gen.SendConfirmationCodeInput,
) (*gen.VoidPayload, error) {
	err := m.auth.SendConfirmationCode(ctx, inp.Recipient, inp.Method)
	if err != nil {
		return nil, err
	}

	return &gen.VoidPayload{}, nil
}

func (m *mutationResolver) VerifyConfirmationCode(ctx context.Context,
	inp gen.VerifyConfirmationCodeInput,
) (*gen.VoidPayload, error) {
	err := m.auth.VerifyConfirmationCode(ctx, inp.Recipient, inp.Code)
	if err != nil {
		return nil, err
	}

	return &gen.VoidPayload{}, nil
}

func (*mutationResolver) SetSecurityImage(_ context.Context,
	inp gen.SetSecurityImageInput,
) (*gen.SetSecurityImagePayload, error) {
	return &gen.SetSecurityImagePayload{
		Email: inp.Email,
	}, nil
}

func (*mutationResolver) GetSecurityImage(_ context.Context,
	inp gen.GetSecurityImageInput,
) (*gen.GetSecurityImagePayload, error) {
	return &gen.GetSecurityImagePayload{
		Email:  inp.Email,
		Image:  "moon.png",
		Phrase: "Security phrase",
	}, nil
}

func (*mutationResolver) ProcessFirebaseAuth(_ context.Context,
	_ gen.ProcessFirebaseAuthInput,
) (*gen.VoidPayload, error) {
	return &gen.VoidPayload{}, nil
}
