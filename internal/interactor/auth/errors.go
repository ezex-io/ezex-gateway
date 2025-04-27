package auth

import "errors"

var (
	ErrConfirmationCodeAlreadySent = errors.New("confirmation code already sent")
	ErrConfirmationCodeExpired     = errors.New("confirmation code has expired")
	ErrConfirmationCodeIsInvalid   = errors.New("confirmation code is invalid")
)
