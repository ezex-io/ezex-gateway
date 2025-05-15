package auth

import (
	"errors"
	"fmt"

	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gateway"
)

var (
	ErrConfirmationCodeAlreadySent = errors.New("confirmation code already sent")
	ErrConfirmationCodeExpired     = errors.New("confirmation code has expired")
	ErrConfirmationCodeIsInvalid   = errors.New("confirmation code is invalid")
)

type UnknownDeliveryMethodError struct {
	Method gateway.DeliveryMethod
}

func (e UnknownDeliveryMethodError) Error() string {
	return fmt.Sprintf("unknown delivery method: %s", e.Method)
}
