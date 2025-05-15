package resolver

import (
	"context"

	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gateway"
)

func (m *mutationResolver) SendConfirmationCode(ctx context.Context,
	inp gateway.SendConfirmationCodeInput,
) (*gateway.SendConfirmationCodePayload, error) {
	return m.auth.SendConfirmationCode(ctx, &inp)
}

func (m *mutationResolver) VerifyConfirmationCode(ctx context.Context,
	inp gateway.VerifyConfirmationCodeInput,
) (*gateway.VerifyConfirmationCodePayload, error) {
	return m.auth.VerifyConfirmationCode(ctx, &inp)
}

func (m *mutationResolver) SaveSecurityImage(ctx context.Context,
	inp gateway.SaveSecurityImageInput,
) (*gateway.SaveSecurityImagePayload, error) {
	return m.auth.SaveSecurityImage(ctx, &inp)
}

func (m *mutationResolver) GetSecurityImage(ctx context.Context,
	inp gateway.GetSecurityImageInput,
) (*gateway.GetSecurityImagePayload, error) {
	return m.auth.GetSecurityImage(ctx, &inp)
}

func (m *mutationResolver) ProcessAuthToken(ctx context.Context,
	inp gateway.ProcessAuthTokenInput,
) (*gateway.ProcessAuthTokenPayload, error) {
	return m.auth.ProcessAuthToken(ctx, &inp)
}
