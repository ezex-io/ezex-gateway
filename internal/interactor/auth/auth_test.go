package auth_test

import (
	"errors"
	"testing"
	"time"

	gauth "firebase.google.com/go/v4/auth"
	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gateway"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
	"github.com/ezex-io/ezex-gateway/internal/mock"
	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/ezex-proto/go/notification"
	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/gopkg/logger"
	"github.com/ezex-io/gopkg/testsuite"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSendConfirmationCode(t *testing.T) {
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockNotification := mock.NewMockNotificationPort(ctrl)
	mockRedis := mock.NewMockCachePort(ctrl)
	mockUsers := mock.NewMockUsersPort(ctrl)
	mockAuth := mock.NewMockAuthenticatorPort(ctrl)

	cfg := &auth.Config{
		ConfirmationCodeSubject:  "Your code is %s",
		ConfirmationTemplateName: "confirmation",
		ConfirmationCodeTTL:      time.Minute * 5,
	}

	authInteractor := auth.NewAuth(
		cfg,
		logger.NewSlog(nil),
		mockNotification,
		mockRedis,
		mockAuth,
		mockUsers,
	)

	t.Run("successful email confirmation code", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(false, nil)
		mockNotification.EXPECT().SendTemplatedEmail(ctx, gomock.Any()).Return(
			&notification.SendTemplatedEmailResponse{
				Recipient: recipient,
			}, nil)
		mockRedis.EXPECT().Set(ctx, recipient, gomock.Any(), gomock.Any()).Return(nil)

		pld, err := authInteractor.SendConfirmationCode(ctx, &gateway.SendConfirmationCodeInput{
			Method:    gateway.DeliveryMethodEmail,
			Recipient: recipient,
		})
		assert.NoError(t, err)
		assert.Equal(t, recipient, pld.Recipient)
	})

	t.Run("code already sent", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(true, nil)

		_, err := authInteractor.SendConfirmationCode(ctx, &gateway.SendConfirmationCodeInput{
			Method:    gateway.DeliveryMethodEmail,
			Recipient: recipient,
		})
		assert.Equal(t, auth.ErrConfirmationCodeAlreadySent, err)
	})

	t.Run("unknown delivery method", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(false, nil)

		_, err := authInteractor.SendConfirmationCode(ctx, &gateway.SendConfirmationCodeInput{
			Method:    "invalid-method",
			Recipient: recipient,
		})
		assert.ErrorIs(t, err, auth.UnknownDeliveryMethodError{Method: "invalid-method"})
	})

	t.Run("email send failure", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(false, nil)
		mockNotification.EXPECT().SendTemplatedEmail(ctx, gomock.Any()).Return(nil,
			errors.New("send failed"))

		_, err := authInteractor.SendConfirmationCode(ctx, &gateway.SendConfirmationCodeInput{
			Method:    gateway.DeliveryMethodEmail,
			Recipient: recipient,
		})
		assert.Error(t, err)
	})

	t.Run("redis set failure", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Exists(ctx, recipient).Return(false, nil)
		mockNotification.EXPECT().SendTemplatedEmail(ctx, gomock.Any()).Return(
			&notification.SendTemplatedEmailResponse{}, nil)
		mockRedis.EXPECT().Set(ctx, recipient, gomock.Any(), gomock.Any()).Return(errors.New("redis error"))

		_, err := authInteractor.SendConfirmationCode(ctx, &gateway.SendConfirmationCodeInput{
			Method:    gateway.DeliveryMethodEmail,
			Recipient: recipient,
		})
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
		logger.NewSlog(nil),
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

		pld, err := authInteractor.VerifyConfirmationCode(ctx, &gateway.VerifyConfirmationCodeInput{
			Recipient: recipient,
			Code:      code,
		})
		assert.NoError(t, err)
		assert.Equal(t, recipient, pld.Recipient)
	})

	t.Run("code expired", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Get(ctx, recipient).Return("", errors.New("not found"))

		_, err := authInteractor.VerifyConfirmationCode(ctx, &gateway.VerifyConfirmationCodeInput{
			Recipient: recipient,
			Code:      ts.RandString(6),
		})
		assert.ErrorIs(t, err, auth.ErrConfirmationCodeExpired)
	})

	t.Run("invalid code", func(t *testing.T) {
		recipient := "test@example.com"

		mockRedis.EXPECT().Get(ctx, recipient).Return("654321", nil)

		_, err := authInteractor.VerifyConfirmationCode(ctx,
			&gateway.VerifyConfirmationCodeInput{
				Recipient: recipient,
				Code:      "123456",
			})
		assert.ErrorIs(t, err, auth.ErrConfirmationCodeIsInvalid)
	})

	t.Run("delete failure logs error", func(t *testing.T) {
		recipient := "test@example.com"
		code := "123456"

		mockRedis.EXPECT().Get(ctx, recipient).Return(code, nil)
		mockRedis.EXPECT().Del(ctx, recipient).Return(errors.New("delete failed"))

		_, err := authInteractor.VerifyConfirmationCode(ctx, &gateway.VerifyConfirmationCodeInput{
			Recipient: recipient,
			Code:      code,
		})
		assert.NoError(t, err) // deletion failure shouldn't fail the operation
	})
}

func TestProcessAuthToken(t *testing.T) {
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockAuth := mock.NewMockAuthenticatorPort(ctrl)
	mockUsers := mock.NewMockUsersPort(ctrl)

	auth := auth.NewAuth(
		&auth.Config{},
		logger.NewSlog(nil),
		nil, // notification not needed
		nil, // redis not needed
		mockAuth,
		mockUsers,
	)

	t.Run("successful", func(t *testing.T) {
		inp := &gateway.ProcessAuthTokenInput{IDToken: "test-token"}

		mockAuth.EXPECT().VerifyIDToken(ctx, &port.VerifyIDTokenRequest{
			IDToken: inp.IDToken,
		}).Return(&port.VerifyIDTokenResponse{
			Token: &gauth.Token{
				UID: "firebase-uid",
				Claims: map[string]any{
					"email": "test@example.com",
				},
			},
		}, nil)
		mockUsers.EXPECT().CreateUser(ctx, &users.CreateUserRequest{
			Email:       "test@example.com",
			FirebaseUid: "firebase-uid",
		}).Return(&users.CreateUserResponse{
			UserId: "user-id",
		}, nil)

		pld, err := auth.ProcessAuthToken(ctx, inp)
		assert.NoError(t, err)
		assert.Equal(t, "user-id", pld.UserID)
	})

	t.Run("token verification failure", func(t *testing.T) {
		inp := &gateway.ProcessAuthTokenInput{IDToken: "test-token"}

		mockAuth.EXPECT().VerifyIDToken(ctx, gomock.Any()).Return(nil, errors.New("invalid token"))

		res, err := auth.ProcessAuthToken(ctx, inp)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("missing email claim", func(t *testing.T) {
		inp := &gateway.ProcessAuthTokenInput{IDToken: "test-token"}

		mockAuth.EXPECT().VerifyIDToken(ctx, gomock.Any()).Return(&port.VerifyIDTokenResponse{
			Token: &gauth.Token{
				UID:    "firebase-uid",
				Claims: map[string]any{},
			},
		}, nil)

		res, err := auth.ProcessAuthToken(ctx, inp)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("invalid email claim type", func(t *testing.T) {
		inp := &gateway.ProcessAuthTokenInput{IDToken: "test-token"}

		mockAuth.EXPECT().VerifyIDToken(ctx, gomock.Any()).Return(&port.VerifyIDTokenResponse{
			Token: &gauth.Token{
				UID: "firebase-uid",
				Claims: map[string]any{
					"email": 123, // not a string
				},
			},
		}, nil)

		res, err := auth.ProcessAuthToken(ctx, inp)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Create user failure", func(t *testing.T) {
		inp := &gateway.ProcessAuthTokenInput{IDToken: "test-token"}

		mockAuth.EXPECT().VerifyIDToken(ctx, gomock.Any()).Return(&port.VerifyIDTokenResponse{
			Token: &gauth.Token{
				UID: "firebase-uid",
				Claims: map[string]any{
					"email": "test@example.com",
				},
			},
		}, nil)
		mockUsers.EXPECT().CreateUser(ctx, gomock.Any()).Return(nil, errors.New("user error"))

		res, err := auth.ProcessAuthToken(ctx, inp)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestSecurityImageOperations(t *testing.T) {
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockUsers := mock.NewMockUsersPort(ctrl)

	auth := auth.NewAuth(
		&auth.Config{},
		logger.NewSlog(nil),
		nil, // notification not needed
		nil, // redis not needed
		nil, // authenticator not needed
		mockUsers,
	)

	t.Run("save security image", func(t *testing.T) {
		inp := &gateway.SaveSecurityImageInput{Email: "test@example.com", SecurityImage: "image.jpg"}

		mockUsers.EXPECT().SaveSecurityImage(ctx, &users.SaveSecurityImageRequest{
			Email:         inp.Email,
			SecurityImage: inp.SecurityImage,
		}).Return(&users.SaveSecurityImageResponse{
			Email: inp.Email,
		}, nil)

		pld, err := auth.SaveSecurityImage(ctx, inp)
		assert.NoError(t, err)
		assert.Equal(t, inp.Email, pld.Email)
	})

	t.Run("get security image", func(t *testing.T) {
		inp := &gateway.GetSecurityImageInput{Email: "test@example.com"}

		mockUsers.EXPECT().GetSecurityImage(ctx, &users.GetSecurityImageRequest{
			Email: inp.Email,
		}).Return(&users.GetSecurityImageResponse{
			Email:         inp.Email,
			SecurityImage: "image.jpg",
		}, nil)

		pld, err := auth.GetSecurityImage(ctx, inp)
		assert.NoError(t, err)
		assert.Equal(t, "image.jpg", pld.SecurityImage)
	})
}
