package resolver

import (
	"context"

	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gen"
	"github.com/ezex-io/ezex-gateway/internal/entity"
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
	err := m.auth.SaveSecurityImage(ctx, &entity.SaveSecurityImageReq{
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
	resp, err := m.auth.GetSecurityImage(ctx, &entity.GetSecurityImageReq{
		Email: inp.Email,
	})
	if err != nil {
		return nil, err
	}

	return &gen.GetSecurityImagePayload{
		Image:  resp.Image,
		Phrase: resp.Phrase,
	}, nil
}

func (m *mutationResolver) ProcessFirebaseAuth(ctx context.Context,
	inp gen.ProcessFirebaseAuthInput,
) (*gen.ProcessFirebaseAuthPayload, error) {
	userID, err := m.auth.ProcessFirebaseLogin(ctx, inp.Token)
	if err != nil {
		return nil, err
	}

	return &gen.ProcessFirebaseAuthPayload{
		UserID: &userID,
	}, nil
}
