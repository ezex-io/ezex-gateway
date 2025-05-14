package auth

import (
	"errors"
	"fmt"
)

var (
	ErrConfirmationCodeAlreadySent = errors.New("confirmation code already sent")
	ErrConfirmationCodeExpired     = errors.New("confirmation code has expired")
	ErrConfirmationCodeIsInvalid   = errors.New("confirmation code is invalid")
)

type UnknownDeliveryMethodError struct {
	Method string // TODO: should be same enum as GraphQL definition
}

func (e UnknownDeliveryMethodError) Error() string {
	return fmt.Sprintf("unknown delivery method: %s", e.Method)
}
