package utils

import "errors"

var (
	ErrorInvalidpassword        = errors.New("invalid password")
	ErrorUserAlreadyExists      = errors.New("user already exists")
	ErrorUserForbidden          = errors.New("login required")
	ErrorOccurredAuthentication = errors.New("error occurred during authentication")
)
