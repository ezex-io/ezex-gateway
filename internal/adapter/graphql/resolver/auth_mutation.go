package resolver

import (
	"context"

	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gen"
	"github.com/ezex-io/ezex-gateway/internal/port"
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

func (m *mutationResolver) SetSecurityImage(ctx context.Context,
	inp gen.SetSecurityImageInput,
) (*gen.VoidPayload, error) {
	_, err := m.auth.SaveSecurityImage(ctx, &port.SaveSecurityImageRequest{
		Email:  inp.Email,
		Image:  inp.Image,
		Phrase: inp.Phrase,
	})
	if err != nil {
		return nil, err
	}

	return &gen.VoidPayload{}, nil
}

func (m *mutationResolver) GetSecurityImage(ctx context.Context,
	inp gen.GetSecurityImageInput,
) (*gen.GetSecurityImagePayload, error) {
	res, err := m.auth.GetSecurityImage(ctx, &port.GetSecurityImageRequest{
		Email: inp.Email,
	})
	if err != nil {
		return nil, err
	}

	return &gen.GetSecurityImagePayload{
		Image:  res.Image,
		Phrase: res.Phrase,
	}, nil
}

func (m *mutationResolver) ProcessFirebaseAuth(ctx context.Context,
	inp gen.ProcessFirebaseAuthInput,
) (*gen.ProcessFirebaseAuthPayload, error) {
	res, err := m.auth.ProcessLogin(ctx, &port.VerifyIDTokenRequest{
		IDToken: inp.Token,
	})
	if err != nil {
		return nil, err
	}

	return &gen.ProcessFirebaseAuthPayload{
		UserID: &res.UserID,
	}, nil
}
