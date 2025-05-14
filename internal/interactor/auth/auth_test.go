package auth_test

import (
	"errors"
	"testing"
	"time"

	gauth "firebase.google.com/go/v4/auth"
	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gen"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
	"github.com/ezex-io/ezex-gateway/internal/mock"
	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/gopkg/testsuite"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSendConfirmationCode(t *testing.T) {
	ctx := t.Context()
	ts := testsuite.NewTestSuite(t)
	ctrl := gomock.NewController(t)

	mockNotification := mock.NewMockNotificationPort(ctrl)
	mockRedis := mock.NewMockCachePort(ctrl)
	mockUsers := mock.NewMockUserPort(ctrl)
	mockAuth := mock.NewMockAuthenticatorPort(ctrl)

	cfg := &auth.Config{
		ConfirmationCodeSubject:  "Your code is %s",
		ConfirmationTemplateName: "confirmation",
		ConfirmationCodeTTL:      time.Minute * 5,
	}

	authInteractor := auth.NewAuth(
		cfg,
		ts.TestLogger(),
		mockNotification,
		mockRedis,
		mockAuth,
		mockUsers,
	)

	t.Run("successful email confirmation code", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(false, nil)
		mockNotification.EXPECT().SendEmail(ctx, gomock.Any()).Return(&port.SendEmailResponse{}, nil)
		mockRedis.EXPECT().Set(ctx, recipient, gomock.Any(), gomock.Any()).Return(nil)

		err := authInteractor.SendConfirmationCode(ctx, recipient, gen.DeliveryMethodEmail)
		assert.NoError(t, err)
	})

	t.Run("code already sent", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(true, nil)

		err := authInteractor.SendConfirmationCode(ctx, recipient, gen.DeliveryMethodEmail)
		assert.Equal(t, auth.ErrConfirmationCodeAlreadySent, err)
	})

	t.Run("unknown delivery method", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(false, nil)

		err := authInteractor.SendConfirmationCode(ctx, recipient, "invalid-method")
		assert.ErrorIs(t, err, auth.UnknownDeliveryMethodError{Method: "invalid-method"})
	})

	t.Run("email send failure", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(false, nil)
		mockNotification.EXPECT().SendEmail(ctx, gomock.Any()).Return(nil, errors.New("send failed"))

		err := authInteractor.SendConfirmationCode(ctx, recipient, gen.DeliveryMethodEmail)
		assert.Error(t, err)
	})

	t.Run("redis set failure", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(false, nil)
		mockNotification.EXPECT().SendEmail(ctx, gomock.Any()).Return(&port.SendEmailResponse{}, nil)
		mockRedis.EXPECT().Set(ctx, recipient, gomock.Any(), gomock.Any()).Return(errors.New("redis error"))

		err := authInteractor.SendConfirmationCode(ctx, recipient, gen.DeliveryMethodEmail)
		assert.Error(t, err)
	})
}

func TestVerifyConfirmationCode(t *testing.T) {
	ctx := t.Context()
	ts := testsuite.NewTestSuite(t)
	ctrl := gomock.NewController(t)

	mockRedis := mock.NewMockCachePort(ctrl)

	authInteractor := auth.NewAuth(
		&auth.Config{},
		ts.TestLogger(),
		nil, // notification not needed
		mockRedis,
		nil, // authenticator not needed
		nil, // users not needed
	)

	t.Run("successful verification", func(t *testing.T) {
		recipient := "test@example.com"
		code := "123456"

		mockRedis.EXPECT().Get(ctx, recipient).Return(code, nil)
		mockRedis.EXPECT().Del(ctx, recipient).Return(nil)

		err := authInteractor.VerifyConfirmationCode(ctx, recipient, code)
		assert.NoError(t, err)
	})

	t.Run("code expired", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Get(ctx, recipient).Return("", errors.New("not found"))

		err := authInteractor.VerifyConfirmationCode(ctx, recipient, "123456")
		assert.ErrorIs(t, err, auth.ErrConfirmationCodeExpired)
	})

	t.Run("invalid code", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Get(ctx, recipient).Return("654321", nil)

		err := authInteractor.VerifyConfirmationCode(ctx, recipient, "123456")
		assert.ErrorIs(t, err, auth.ErrConfirmationCodeIsInvalid)
	})

	t.Run("delete failure logs error", func(t *testing.T) {
		recipient := "test@example.com"
		code := "123456"

		mockRedis.EXPECT().Get(ctx, recipient).Return(code, nil)
		mockRedis.EXPECT().Del(ctx, recipient).Return(errors.New("delete failed"))

		err := authInteractor.VerifyConfirmationCode(ctx, recipient, code)
		assert.NoError(t, err) // deletion failure shouldn't fail the operation
	})
}

func TestProcessLogin(t *testing.T) {
	ctx := t.Context()
	ts := testsuite.NewTestSuite(t)
	ctrl := gomock.NewController(t)

	mockAuth := mock.NewMockAuthenticatorPort(ctrl)
	mockUsers := mock.NewMockUserPort(ctrl)

	authInteractor := auth.NewAuth(
		&auth.Config{},
		ts.TestLogger(),
		nil, // notification not needed
		nil, // redis not needed
		mockAuth,
		mockUsers,
	)

	t.Run("successful login", func(t *testing.T) {
		req := &port.VerifyIDTokenRequest{IDToken: "test-token"}
		expectedResponse := &port.ProcessLoginResponse{UserID: "123"}

		mockAuth.EXPECT().VerifyIDToken(ctx, req).Return(&port.VerifyIDTokenResponse{
			Token: &gauth.Token{
				UID: "firebase-uid",
				Claims: map[string]any{
					"email": "test@example.com",
				},
			},
		}, nil)
		mockUsers.EXPECT().ProcessLogin(ctx, &port.ProcessLoginRequest{
			Email:       "test@example.com",
			FirebaseUID: "firebase-uid",
		}).Return(expectedResponse, nil)

		res, err := authInteractor.ProcessLogin(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, res)
	})

	t.Run("token verification failure", func(t *testing.T) {
		req := &port.VerifyIDTokenRequest{IDToken: "test-token"}

		mockAuth.EXPECT().VerifyIDToken(ctx, req).Return(nil, errors.New("invalid token"))

		res, err := authInteractor.ProcessLogin(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("missing email claim", func(t *testing.T) {
		req := &port.VerifyIDTokenRequest{IDToken: "test-token"}

		mockAuth.EXPECT().VerifyIDToken(ctx, req).Return(&port.VerifyIDTokenResponse{
			Token: &gauth.Token{
				UID:    "firebase-uid",
				Claims: map[string]any{},
			},
		}, nil)

		res, err := authInteractor.ProcessLogin(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("invalid email claim type", func(t *testing.T) {
		req := &port.VerifyIDTokenRequest{IDToken: "test-token"}

		mockAuth.EXPECT().VerifyIDToken(ctx, req).Return(&port.VerifyIDTokenResponse{
			Token: &gauth.Token{
				UID: "firebase-uid",
				Claims: map[string]any{
					"email": 123, // not a string
				},
			},
		}, nil)

		res, err := authInteractor.ProcessLogin(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("user process login failure", func(t *testing.T) {
		req := &port.VerifyIDTokenRequest{IDToken: "test-token"}

		mockAuth.EXPECT().VerifyIDToken(ctx, req).Return(&port.VerifyIDTokenResponse{
			Token: &gauth.Token{
				UID: "firebase-uid",
				Claims: map[string]any{
					"email": "test@example.com",
				},
			},
		}, nil)
		mockUsers.EXPECT().ProcessLogin(ctx, gomock.Any()).Return(nil, errors.New("user error"))

		res, err := authInteractor.ProcessLogin(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestSecurityImageOperations(t *testing.T) {
	ctx := t.Context()
	ts := testsuite.NewTestSuite(t)
	ctrl := gomock.NewController(t)

	mockUsers := mock.NewMockUserPort(ctrl)

	authInteractor := auth.NewAuth(
		&auth.Config{},
		ts.TestLogger(),
		nil, // notification not needed
		nil, // redis not needed
		nil, // authenticator not needed
		mockUsers,
	)

	t.Run("save security image", func(t *testing.T) {
		req := &port.SaveSecurityImageRequest{Email: "test@example.com", Image: "image.jpg"}
		expectedResponse := &port.SaveSecurityImageResponse{Email: "test@example.com"}

		mockUsers.EXPECT().SaveSecurityImage(ctx, req).Return(expectedResponse, nil)

		res, err := authInteractor.SaveSecurityImage(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, res)
	})

	t.Run("get security image", func(t *testing.T) {
		req := &port.GetSecurityImageRequest{Email: "test@example.com"}
		expectedResponse := &port.GetSecurityImageResponse{Image: "image.jpg"}

		mockUsers.EXPECT().GetSecurityImage(ctx, req).Return(expectedResponse, nil)

		res, err := authInteractor.GetSecurityImage(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, res)
	})
}
