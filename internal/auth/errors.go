package auth

import "errors"

var (
	ErrAuthFailed   = errors.New("authorization failed")
	ErrUserNotFound = errors.New("user not found")
)
